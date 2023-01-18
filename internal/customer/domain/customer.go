package domain

type Customer struct {
	ID            string
	FirstName     string
	LastName      string
	Email         string
	PhoneNumber   string
	PhoneVerified bool
}

func NewCustomer(
	firstName,
	lastName,
	email,
	phoneNumber string,
	phoneVerified bool,
) *Customer {
	return &Customer{
		FirstName:     firstName,
		LastName:      lastName,
		Email:         email,
		PhoneNumber:   phoneNumber,
		PhoneVerified: phoneVerified,
	}
}
