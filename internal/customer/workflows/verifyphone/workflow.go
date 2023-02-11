package verifyphone

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
	UserCodeSignal             = "verify_phone_user_code_signal"
	VerificationStateQueryType = "verify_phone_workflow_state"
)

type WorkflowParams struct {
	PhoneNumber          string
	MaximumAttempts      uint
	CodeValidityDuration time.Duration
}

func NewWorkflow(
	ctx workflow.Context,
	params WorkflowParams,
) error {
	// TODO: Add explicit retry policy
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	activityCtx := workflow.WithActivityOptions(ctx, options)

	var attempts uint
	var state VerificationState
	var smsSender *SMSSender

	err := workflow.SetQueryHandler(ctx, VerificationStateQueryType, func() (VerificationState, error) { return state, nil })
	if err != nil {
		return err
	}

	userCodeChannel := workflow.GetSignalChannel(ctx, UserCodeSignal)
	for attempts < params.MaximumAttempts {
		state = InProgress

		oneTimeCode := NewOneTimeCode(params.CodeValidityDuration)

		message := fmt.Sprintf(
			"Thanks for signing up to GoCoffee. Please enter the following code in our app to verify your phone number: %s",
			oneTimeCode.code,
		)
		err = workflow.ExecuteActivity(activityCtx, smsSender.SendSMS, params.PhoneNumber, message).Get(ctx, nil)
		if err != nil {
			return fmt.Errorf("unable to send sms to phone number: %s. err: %w", params.PhoneNumber, err)
		}

		var userCode string
		userCodeChannel.Receive(ctx, &userCode)
		attempts++

		// note: states CodeExpired and IncorrectCode are somewhat redundant since we transition to the next iteration
		// of the for loop immediately, so state moves directly to InProgress.
		if oneTimeCode.IsExpired(workflow.Now(ctx)) {
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
