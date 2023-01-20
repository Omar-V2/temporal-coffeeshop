package api

import (
	"errors"
	"log"
	"tmprldemo/internal/customer/domain"
	customerpb "tmprldemo/internal/customer/pb/customer/v1"

	"github.com/google/uuid"
)

// TOOD: Add unit tests

var ErrInvalidUUID = errors.New("failed to convert provided customer ID into UUID: not a valid UUID")

func ConvertFromPbToCustomer(customer *customerpb.Customer) (*domain.Customer, error) {
	var customerID uuid.UUID

	log.Println("received customer ", customer)

	if customer.Id == "" {
		customerID = uuid.New()
	} else {
		var err error
		customerID, err = uuid.Parse(customer.Id)
		if err != nil {
			return nil, ErrInvalidUUID
		}
	}

	return domain.NewCustomer(
		customerID,
		customer.FirstName,
		customer.LastName,
		customer.Email,
		customer.PhoneNumber,
		customer.PhoneVerified,
	), nil
}

func ConvertFromCustomerToPb(customer domain.Customer) *customerpb.Customer {
	return &customerpb.Customer{
		Id:            customer.ID.String(), // consider handling nil case here?
		FirstName:     customer.FirstName,
		LastName:      customer.LastName,
		Email:         customer.Email,
		PhoneNumber:   customer.PhoneNumber,
		PhoneVerified: customer.PhoneVerified,
	}
}
