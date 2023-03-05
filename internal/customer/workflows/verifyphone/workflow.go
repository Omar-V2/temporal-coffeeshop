package verifyphone

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// TODO: Add comments to all structs/funcs etc.

type VerificationResult int

const (
	NotStarted VerificationResult = iota
	InProgress
	CodeExpired
	MaxAttemptsReached
	IncorrectCode
	CorrectCode
)

const (
	UserCodeSignal              = "verify_phone_user_code_signal"
	VerificationResultQueryType = "verify_phone_workflow_verification_result"
)

type WorkflowParams struct {
	PhoneNumber          string
	MaximumAttempts      uint
	CodeValidityDuration time.Duration
}

func NewWorkflow(ctx workflow.Context, params WorkflowParams) error {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts:    5,
			InitialInterval:    time.Second * 5,
			MaximumInterval:    time.Minute * 2,
			BackoffCoefficient: 2.0,
		},
	}
	activityCtx := workflow.WithActivityOptions(ctx, options)

	var attempts uint
	var mostRecentAttempt VerificationResult
	var activities *activities

	err := workflow.SetQueryHandler(ctx, VerificationResultQueryType, func() (VerificationResult, error) { return mostRecentAttempt, nil })
	if err != nil {
		return err
	}

	userCodeChannel := workflow.GetSignalChannel(ctx, UserCodeSignal)

	for attempts < params.MaximumAttempts {
		encodedCodeExpiryTime := workflow.SideEffect(ctx, func(ctx workflow.Context) interface{} {
			return workflow.Now(ctx).Add(params.CodeValidityDuration)
		})

		var codeExpiryTime time.Time
		err = encodedCodeExpiryTime.Get(&codeExpiryTime)
		if err != nil {
			return fmt.Errorf("failed to decode one time code from side effect function: %w", err)
		}

		oneTimeCode := NewOneTimeCode(codeExpiryTime)

		message := fmt.Sprintf(
			"Thanks for signing up to GoCoffee. Please enter the following code in our app to verify your phone number: %s",
			oneTimeCode.code,
		)
		err = workflow.ExecuteActivity(activityCtx, activities.SendSMS, params.PhoneNumber, message).Get(ctx, nil)
		if err != nil {
			return fmt.Errorf("unable to send sms to phone number: %s. err: %w", params.PhoneNumber, err)
		}

		var userCode string
		userCodeChannel.Receive(ctx, &userCode)
		attempts++

		if oneTimeCode.IsExpired(workflow.Now(ctx)) {
			mostRecentAttempt = CodeExpired
			continue
		}

		if oneTimeCode.Matches(userCode) {
			mostRecentAttempt = CorrectCode

			err := workflow.ExecuteActivity(
				activityCtx,
				activities.VerifyCustomer,
				workflow.GetInfo(ctx).WorkflowExecution.ID, // workflow ID is the customer ID
			).Get(ctx, nil)
			if err != nil {
				return fmt.Errorf("unable to mark customer as verified: %w", err)
			}

			return nil
		}

		mostRecentAttempt = IncorrectCode
	}

	mostRecentAttempt = MaxAttemptsReached
	return nil
}
