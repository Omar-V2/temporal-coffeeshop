package api

import (
	"fmt"
	"tmprldemo/internal/customer/domain"
	customerpb "tmprldemo/internal/customer/pb/customer/v1"

	"github.com/google/uuid"
)

// TOOD: Add unit tests

func ConvertFromPbToCustomer(customer *customerpb.Customer) (*domain.Customer, error) {
	var customerID uuid.UUID

	if customer.Id == "" {
		customerID = uuid.New()
	} else {
		var err error
		customerID, err = uuid.Parse(customer.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to convert provided customer ID into UUID:  %s is not a valid UUID", customer.Id)
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

func ConvertFromCustomersToPb(customers domain.Customers) []*customerpb.Customer {
	var customersPb []*customerpb.Customer
	for _, customer := range customers {
		customerPb := ConvertFromCustomerToPb(*customer)
		customersPb = append(customersPb, customerPb)
	}
	return customersPb
}
