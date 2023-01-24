package api

import (
	"testing"
	"tmprldemo/internal/customer/domain"
	customerpb "tmprldemo/internal/pb/customer/v1"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestConvertFromPbToCustomer(t *testing.T) {
	customerID := uuid.New()
	firstName := "Naruto"
	lastName := "Uzumaki"
	email := "naruto@leafvillage.com"
	phoneNumber := "07500512345"

	t.Run("converts customer", func(t *testing.T) {
		customerPb := &customerpb.Customer{
			Id:          customerID.String(),
			FirstName:   firstName,
			LastName:    lastName,
			Email:       email,
			PhoneNumber: phoneNumber,
		}

		expectedCustomer := &domain.Customer{
			ID:            customerID,
			FirstName:     firstName,
			LastName:      lastName,
			Email:         email,
			PhoneNumber:   phoneNumber,
			PhoneVerified: false,
		}

		customer, err := ConvertFromPbToCustomer(customerPb)
		assert.NoError(t, err)

		assert.Equal(t, expectedCustomer, customer)
	})

	t.Run("generates resource ID if not provided", func(t *testing.T) {
		customerPb := &customerpb.Customer{
			FirstName:   firstName,
			LastName:    lastName,
			Email:       email,
			PhoneNumber: phoneNumber,
		}

		customer, err := ConvertFromPbToCustomer(customerPb)
		assert.NoError(t, err)
		assert.NotEmpty(t, customer.ID)
	})

	t.Run("returns error if invalid UUID provided", func(t *testing.T) {
		customerPb := &customerpb.Customer{
			Id:          "invalid-uuid",
			FirstName:   firstName,
			LastName:    lastName,
			Email:       email,
			PhoneNumber: phoneNumber,
		}

		customer, err := ConvertFromPbToCustomer(customerPb)
		assert.Error(t, err)
		assert.EqualError(t, err, "failed to convert provided customer ID into UUID:  invalid-uuid is not a valid UUID")
		assert.Nil(t, customer)
	})
}

func TestConvertFromCustomerToPb(t *testing.T) {
	customerID := uuid.New()
	firstName := "Naruto"
	lastName := "Uzumaki"
	email := "naruto@leafvillage.com"
	phoneNumber := "07500512345"

	customer := domain.Customer{
		ID:            customerID,
		FirstName:     firstName,
		LastName:      lastName,
		Email:         email,
		PhoneNumber:   phoneNumber,
		PhoneVerified: true,
	}

	expectedCustomerPb := &customerpb.Customer{
		Id:            customerID.String(),
		FirstName:     firstName,
		LastName:      lastName,
		Email:         email,
		PhoneNumber:   phoneNumber,
		PhoneVerified: true,
	}

	customberPb := ConvertFromCustomerToPb(customer)
	assert.Equal(t, expectedCustomerPb, customberPb)
}
