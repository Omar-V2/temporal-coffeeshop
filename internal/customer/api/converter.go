package api

import (
	"fmt"
	"tmprldemo/internal/customer/domain"
	"tmprldemo/internal/customer/workflows/verifyphone"
	customerpb "tmprldemo/internal/pb/customer/v1"

	"github.com/google/uuid"
)

// ConvertFromPbToCustomer converts a protobuf Customer into a domain Customer.
// It also generates and sets a resource UUID for the customer if one is not provided.
// If an ID is provided but it is not a valid UUID an error will be returned.
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

// ConvertFromCustomerToPb converts a domain Customer into a protobuf Customer.
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

// ConvertFromCustomersToPb converts a slice of domain Customers to a slice of protobuf Customers.
func ConvertFromCustomersToPb(customers domain.Customers) []*customerpb.Customer {
	var customersPb []*customerpb.Customer
	for _, customer := range customers {
		customerPb := ConvertFromCustomerToPb(*customer)
		customersPb = append(customersPb, customerPb)
	}
	return customersPb
}

// convertWorkflowResultToPb converts the verify phone workflow state to proto verification result.
func convertWorkflowResultToPb(wfResult verifyphone.VerificationState) customerpb.VerificationResult {
	wfStateToPbResult := map[verifyphone.VerificationState]customerpb.VerificationResult{
		verifyphone.StateUnknown:       customerpb.VerificationResult_VERIFICATION_RESULT_UNSPECIFIED,
		verifyphone.InProgress:         customerpb.VerificationResult_VERIFICATION_RESULT_IN_PROGRESS,
		verifyphone.CodeExpired:        customerpb.VerificationResult_VERIFICATION_RESULT_CODE_EXPIRED,
		verifyphone.MaxAttemptsReached: customerpb.VerificationResult_VERIFICATION_RESULT_MAX_ATTEMPTS_REACHED,
		verifyphone.IncorrectCode:      customerpb.VerificationResult_VERIFICATION_RESULT_INCORRECT_CODE,
		verifyphone.CorrectCode:        customerpb.VerificationResult_VERIFICATION_RESULT_SUCCESS,
	}

	return wfStateToPbResult[wfResult]
}
