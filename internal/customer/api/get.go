package api

import (
	"context"
	customerpb "tmprldemo/internal/customer/pb/customer/v1"
)

func (s *customerServiceGRPCServer) GetCustomer(ctx context.Context, request *customerpb.GetCustomerRequest) (*customerpb.GetCustomerResponse, error) {
	if err := validateUUIDs(request.CustomerId); err != nil {
		return nil, err
	}

	customer, err := s.customerGetter.Get(ctx, request.CustomerId)
	if err != nil {
		return nil, err
	}

	return &customerpb.GetCustomerResponse{
		Customer: ConvertFromCustomerToPb(*customer),
	}, nil
}

func (s *customerServiceGRPCServer) BatchGetCustomers(ctx context.Context, request *customerpb.BatchGetCustomersRequest) (*customerpb.BatchGetCustomersResponse, error) {
	if err := validateUUIDs(request.CustomerIds...); err != nil {
		return nil, err
	}

	customers, err := s.customerGetter.BatchGet(ctx, request.CustomerIds)
	if err != nil {
		return nil, err
	}

	return &customerpb.BatchGetCustomersResponse{
		Customers: ConvertFromCustomersToPb(customers),
	}, nil
}
