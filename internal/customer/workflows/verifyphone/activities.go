package verifyphone

import (
	"context"
	"time"
)

// CustomerVerifier defines a method to mark the customer as verified in the database
type CustomerVerifier interface {
	Verify(ctx context.Context, customerID string) error
}

// Sender defines a method to send a text message via SMS to given phone number.
type Sender interface {
	SendMessage(phoneNumber, message string) error
}

type CodeGenerator interface {
	GenerateCode(length int) string
}

// activities defines the activity methods required for the verify phone workflow.
type activities struct {
	Sender           Sender
	CustomerVerifier CustomerVerifier
	CodeGenerator    CodeGenerator
}

// NewActivities creates and returns a new Activities struct.
func NewActivities(
	sender Sender,
	customerVerifier CustomerVerifier,
	codeGenerator CodeGenerator,
) *activities {
	return &activities{Sender: sender, CustomerVerifier: customerVerifier, CodeGenerator: codeGenerator}
}

// SendSMS sends an SMS message to the provided phone number.
func (a *activities) SendSMS(phoneNumber, message string) error {
	return a.Sender.SendMessage(phoneNumber, message)
}

// VerifyCustomer marks the customer as verified in the database.
func (a *activities) VerifyCustomer(ctx context.Context, customerID string) error {
	return a.CustomerVerifier.Verify(ctx, customerID)
}

// NewOneTimeCode generates a new one time code object containing a code and an expiry time
func (a *activities) NewOneTimeCode(length int, validityDuration time.Duration) (*OneTimeCode, error) {
	return &OneTimeCode{
		Code:       a.CodeGenerator.GenerateCode(length),
		ValidUntil: time.Now().Add(validityDuration),
	}, nil
}
