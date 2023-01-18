package api

import (
	"context"

	customerpb "tmprldemo/internal/customer/pb/customer/v1"
)

func (s *customerServiceGRPCServer) CreateCustomer(ctx context.Context, customer *customerpb.CreateCustomerRequest) (*customerpb.CreateCustomerResponse, error) {
	return nil, nil
	// domainCustomer := ConvertPbToCustomer(customer)
	// createdCustomer, err := s.CustomerCreator.Create(domainCustomer)
	// if err != nil {
	// 	return nil, err
	// }

	// s.temporalClient.ExecuteWorkflow(ctx, "VerifyPhoneWorkflow", customer.Phone)

	// return &customerpb.CreateCustomerResponse{
	// 	Customer: ConvertToPbCustomer(domainCustomer),
	// }
}
