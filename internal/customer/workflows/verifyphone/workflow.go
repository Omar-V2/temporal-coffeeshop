package sms

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

type SMSSender interface {
	SendSMS(ctx context.Context, message, phoneNumber string) error
}

type VerificationState int

const (
	StateUnknown VerificationState = iota
	CodeExpired
	TooManyAttempts
	IncorrectCode
	CorrectCode
)

type VerifyPhoneWorkflow struct {
	ctx                  workflow.Context
	phoneNumber          string
	maximumAttempts      uint
	codeValidityDuration time.Duration
}

func NewVerifyPhoneWorkflow(ctx workflow.Context, phoneNumber string, maximumAttempts uint, codeValidityDuration time.Duration) error {
	wf := &VerifyPhoneWorkflow{
		ctx:                  ctx,
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

	for attempts < wf.maximumAttempts {
		var userCode string
		workflow.GetSignalChannel(wf.ctx, "receive_user_code").Receive(wf.ctx, &userCode)

	}


	wf.ctx = workflow.WithActivityOptions(wf.ctx, options)
	err := workflow.ExecuteActivity(wf.ctx, SMSSender.SendSMS, wf.phoneNumber).Get(wf.ctx, nil)
	if err != nil {
		return fmt.Errorf("unable to send sms to phone number: %s. err: %w", wf.phoneNumber, err)
	}

	return nil
}
