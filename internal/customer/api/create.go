package api

import (
	"context"

	customerpb "tmprldemo/internal/pb/customer/v1"
)

// CreateCustomer creates a new customer in the coffee shop system.
func (s *customerServiceGRPCServer) CreateCustomer(ctx context.Context, request *customerpb.CreateCustomerRequest) (*customerpb.CreateCustomerResponse, error) {
	domainCustomer, err := ConvertFromPbToCustomer(request.Customer)
	if err != nil {
		return nil, err
	}

	// should place the request id on the context
	createdCustomer, err := s.customerCreator.Create(ctx, *domainCustomer)
	if err != nil {
		return nil, err
	}

	// TODO: execute temporal Verify Phone Workflow
	// customer ID should be used as the workflow ID.

	return &customerpb.CreateCustomerResponse{
		Customer: ConvertFromCustomerToPb(*createdCustomer),
	}, nil
}
