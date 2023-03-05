package verifyphone

import "fmt"

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
