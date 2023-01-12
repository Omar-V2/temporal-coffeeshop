package sms

import (
	crypto "crypto/subtle"
	"time"
)

type OneTimeCode struct {
	code       string
	validUntil time.Time
}

func NewOneTimeCode(validFor time.Duration) *OneTimeCode {
	code := generateCode()
	return &OneTimeCode{
		code:       code,
		validUntil: time.Now().Add(validFor),
	}
}

func (o *OneTimeCode) IsExpired(currentTime time.Time) bool {
	return currentTime.After(o.validUntil)
}

func (o *OneTimeCode) Matches(codeToCompare string) bool {
	match := crypto.ConstantTimeCompare([]byte(o.code), []byte(codeToCompare))
	if match == 1 {
		return true
	}
	return false
}

func generateCode() string {
	return "acbde"
}
