package sms

import (
	"fmt"
)

type Sender interface {
	SendSMS(phoneNumber, message string) error
}

type SMSSender struct {
	Sender Sender
}

func (s *SMSSender) SendSMS(phoneNumber, message string) error {
	return s.Sender.SendSMS(phoneNumber, message)
}

type TwilioSMSSender struct{}

func (t *TwilioSMSSender) SendSMS(phoneNumber, message string) error {
	return nil
}

type MockSMSSender struct{}

func (m *MockSMSSender) SendSMS(phoneNumber, message string) error {
	fmt.Printf("sent message: %s to phone number : %s", message, phoneNumber)
	return nil
}
