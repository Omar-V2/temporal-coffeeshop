package sms

import (
	"errors"
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

type VerificationState int

const (
	StateUnknown VerificationState = iota
	CodeExpired
	MaxAttemptsReached
	IncorrectCode
	CorrectCode
)

const UserCodeChannel = "verify_phone_user_code_channel"

type VerifyPhoneWorkflow struct {
	ctx                  workflow.Context
	smsSender            SMSSender
	phoneNumber          string
	maximumAttempts      uint
	codeValidityDuration time.Duration
}

func NewVerifyPhoneWorkflow(
	ctx workflow.Context,
	smsSender SMSSender,
	phoneNumber string,
	maximumAttempts uint,
	codeValidityDuration time.Duration,
) error {
	wf := &VerifyPhoneWorkflow{
		ctx:                  ctx,
		smsSender:            smsSender,
		phoneNumber:          phoneNumber,
		maximumAttempts:      maximumAttempts,
		codeValidityDuration: codeValidityDuration,
	}

	return wf.run()
}

func (wf *VerifyPhoneWorkflow) run() error {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	var attempts uint
	var state VerificationState
	userCodeChannel := workflow.GetSignalChannel(wf.ctx, UserCodeChannel)

	for attempts < wf.maximumAttempts {
		wf.ctx = workflow.WithActivityOptions(wf.ctx, options)

		var oneTimeCode *OneTimeCode
		err := workflow.ExecuteActivity(wf.ctx, NewOneTimeCode, wf.codeValidityDuration).Get(wf.ctx, &oneTimeCode)
		if err != nil {
			return err
		}

		message := fmt.Sprintf("Thanks for signing up to GoCoffee. Please enter the following code in our app to verify your phone number: %s", oneTimeCode.code)
		err = workflow.ExecuteActivity(wf.ctx, wf.smsSender.SendSMS, wf.phoneNumber, message).Get(wf.ctx, nil)
		if err != nil {
			return fmt.Errorf("unable to send sms to phone number: %s. err: %w", wf.phoneNumber, err)
		}

		var userCode string
		userCodeChannel.Receive(wf.ctx, userCode)
		attempts++

		if oneTimeCode.IsExpired(workflow.Now(wf.ctx)) {
			state = CodeExpired
			continue
		}

		if oneTimeCode.Matches(userCode) {
			state = CorrectCode
			return nil
		}
		state = IncorrectCode
	}

	state = MaxAttemptsReached
	return errors.New("too many attempts")
}
