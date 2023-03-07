package verifyphone

import (
	"context"
)

// CustomerVerifier defines a method to mark the customer as verified in the database
type CustomerVerifier interface {
	Verify(ctx context.Context, customerID string) error
}

// Sender defines a method to send a text message via SMS to given phone number.
type Sender interface {
	SendMessage(phoneNumber, message string) error
}

// activities defines the activity methods required for the verify phone workflow.
type activities struct {
	Sender           Sender
	CustomerVerifier CustomerVerifier
}

// NewActivities creates and returns a new Activities struct.
func NewActivities(sender Sender, verifier CustomerVerifier) *activities {
	return &activities{Sender: sender, CustomerVerifier: verifier}
}

// SendSMS sends an SMS message to the provided phone number.
func (a *activities) SendSMS(phoneNumber, message string) error {
	return a.Sender.SendMessage(phoneNumber, message)
}

// VerifyCustomer marks the customer as verified in the database.
func (a *activities) VerifyCustomer(ctx context.Context, customerID string) error {
	return a.CustomerVerifier.Verify(ctx, customerID)
}
