package sms

import (
	"context"
	"fmt"
)

type SMSSender interface {
	SendSMS(ctx context.Context, phoneNumber, message string) error
}

type TwilioSMSSender struct {
	ctx context.Context
}

func (t *TwilioSMSSender) SendSMS(ctx context.Context, message, phoneNumber string) error {
	return nil
}

type MockSMSSender struct {
	ctx context.Context
}

func (m *MockSMSSender) SendSMS(ctx context.Context, message, phoneNumber string) error {
	fmt.Printf("sent message: %s to phone number : %s", message, phoneNumber)
	return nil
}
