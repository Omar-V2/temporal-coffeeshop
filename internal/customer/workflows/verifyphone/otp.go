package verifyphone

import (
	crypto "crypto/subtle"
	"time"
)

type OneTimeCode struct {
	code       string
	validUntil time.Time
}

func NewOneTimeCode(validUntil time.Time) *OneTimeCode {
	code := generateCode()
	return &OneTimeCode{
		code:       code,
		validUntil: validUntil,
	}
}

func (o *OneTimeCode) IsExpired(currentTime time.Time) bool {
	return currentTime.After(o.validUntil)
}

func (o *OneTimeCode) Matches(codeToCompare string) bool {
	match := crypto.ConstantTimeCompare([]byte(o.code), []byte(codeToCompare))
	return match == 1
}

// TODO: generate random four digit code
func generateCode() string {
	return "1234"
}
