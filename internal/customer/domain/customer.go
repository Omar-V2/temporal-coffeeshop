package domain

import "github.com/google/uuid"

type Customer struct {
	ID            uuid.UUID
	FirstName     string
	LastName      string
	Email         string
	PhoneNumber   string
	PhoneVerified bool
}

func NewCustomer(
	Id uuid.UUID,
	firstName,
	lastName,
	email,
	phoneNumber string,
	phoneVerified bool,
) *Customer {
	return &Customer{
		ID:            Id,
		FirstName:     firstName,
		LastName:      lastName,
		Email:         email,
		PhoneNumber:   phoneNumber,
		PhoneVerified: phoneVerified,
	}
}
