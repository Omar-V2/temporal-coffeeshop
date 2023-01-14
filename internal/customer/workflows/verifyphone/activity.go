package sms

import (
	"fmt"
)

type Sender interface {
	SendMessage(phoneNumber, message string) error
}

type SMSSender struct {
	Sender Sender
}

func (s *SMSSender) SendSMS(phoneNumber, message string) error {
	return s.Sender.SendMessage(phoneNumber, message)
}

type TwilioSMSSender struct{}

func (t *TwilioSMSSender) SendMessage(phoneNumber, message string) error {
	return nil
}

type MockSMSSender struct{}

func (m *MockSMSSender) SendMessage(phoneNumber, message string) error {
	fmt.Printf("sent message: %s to phone number : %s", message, phoneNumber)
	return nil
}
