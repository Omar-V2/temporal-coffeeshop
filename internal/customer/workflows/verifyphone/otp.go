package verifyphone

import (
	crypto "crypto/subtle"
	"math/rand"
	"time"
)

type OneTimeCode struct {
	Code       string
	ValidUntil time.Time
}

func (o *OneTimeCode) IsExpired(currentTime time.Time) bool {
	return currentTime.After(o.ValidUntil)
}

func (o *OneTimeCode) Matches(codeToCompare string) bool {
	match := crypto.ConstantTimeCompare([]byte(o.Code), []byte(codeToCompare))
	return match == 1
}

type RandomCodeGenerator struct{}

func (g RandomCodeGenerator) GenerateCode(length int) string {
	const characters = "0123456789"
	code := make([]byte, length)
	for i := range code {
		code[i] = characters[rand.Intn(len(characters))]
	}
	return string(code)
}

type StaticCodeGenerator struct{}

func (g StaticCodeGenerator) GenerateCode(length int) string {
	return "1234"
}
