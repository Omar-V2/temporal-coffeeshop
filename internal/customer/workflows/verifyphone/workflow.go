package sms

import (
	"errors"
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

// TODO: Add comments to all structs/funcs etc.

type VerificationState int

const (
	StateUnknown VerificationState = iota
	InProgress
	CodeExpired
	MaxAttemptsReached
	IncorrectCode
	CorrectCode
)

const (
	UserCodeChannel            = "verify_phone_user_code_channel"
	VerificationStateQueryType = "verify_phone_workflow_state"
)

type VerifyPhoneWorkflowParams struct {
	PhoneNumber          string
	MaximumAttempts      uint
	CodeValidityDuration time.Duration
}

type VerifyPhoneWorkflow struct {
	ctx                  workflow.Context
	phoneNumber          string
	maximumAttempts      uint
	codeValidityDuration time.Duration
}

func NewVerifyPhoneWorkflow(
	ctx workflow.Context,
	params VerifyPhoneWorkflowParams,
) error {
	wf := &VerifyPhoneWorkflow{
		ctx:                  ctx,
		phoneNumber:          params.PhoneNumber,
		maximumAttempts:      params.MaximumAttempts,
		codeValidityDuration: params.CodeValidityDuration,
	}

	return wf.run()
}

func (wf *VerifyPhoneWorkflow) run() error {

	// TODO: Add explicit retry policy
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	var attempts uint
	var state VerificationState
	var smsSender *SMSSender

	err := workflow.SetQueryHandler(wf.ctx, VerificationStateQueryType, func() (VerificationState, error) { return state, nil })
	if err != nil {
		return err
	}

	userCodeChannel := workflow.GetSignalChannel(wf.ctx, UserCodeChannel)
	for attempts < wf.maximumAttempts {
		state = InProgress
		wf.ctx = workflow.WithActivityOptions(wf.ctx, options)

		oneTimeCode := NewOneTimeCode(wf.codeValidityDuration)

		message := fmt.Sprintf("Thanks for signing up to GoCoffee. Please enter the following code in our app to verify your phone number: %s", oneTimeCode.code)
		err = workflow.ExecuteActivity(wf.ctx, smsSender.SendSMS, wf.phoneNumber, message).Get(wf.ctx, nil)
		if err != nil {
			return fmt.Errorf("unable to send sms to phone number: %s. err: %w", wf.phoneNumber, err)
		}

		var userCode string
		userCodeChannel.Receive(wf.ctx, &userCode)
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
