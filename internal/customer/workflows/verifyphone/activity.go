package verifyphone

import (
	"context"
	"fmt"
)

// activities defines the activity methods required for the verify phone workflow.
type activities struct {
	Sender           Sender
	CustomerVerifier CustomerVerifier
}

// NewActivities creates and returns a new Activities struct.
func NewActivities(sender Sender, verifier CustomerVerifier) *activities {
	return &activities{Sender: sender, CustomerVerifier: verifier}
}

func (a *activities) SendSMS(phoneNumber, message string) error {
	return a.Sender.SendMessage(phoneNumber, message)
}

func (a *activities) VerifyCustomer(ctx context.Context, customerID string) error {
	return a.CustomerVerifier.Verify(ctx, customerID)
}

// CustomerVerifier defines a method to mark the customer as verified in the database
type CustomerVerifier interface {
	Verify(ctx context.Context, customerID string) error
}

// Sender defines a method to send a text message via SMS to given phone number.
type Sender interface {
	SendMessage(phoneNumber, message string) error
}

type SMSSender struct {
	Sender Sender
}

func (s *SMSSender) SendSMS(phoneNumber, message string) error {
	return s.Sender.SendMessage(phoneNumber, message)
}

// TwilioSMSSender implements the SendMessage interface using the Twilio API
// to send an SMS message to the provided phone number.
type TwilioSMSSender struct{}

func (t *TwilioSMSSender) SendMessage(phoneNumber, message string) error {
	return nil
}

// MockSMSSender implements the SendMessage interface but simply prints the message to stdout.
type MockSMSSender struct{}

func (m *MockSMSSender) SendMessage(phoneNumber, message string) error {
	fmt.Printf("sent message: %s to phone number : %s", message, phoneNumber)
	return nil
}

// FaultySMSSender implements the SendMessage interface but always returns an error
// when attempting to send a message to simulate a failure scenario.
type FaultySMSSender struct{}

func (m *FaultySMSSender) SendMessage(phoneNumber, message string) error {
	return fmt.Errorf(
		"failed sendint text message: %s to phone number %s",
		message, phoneNumber,
	)
}
