package api

import (
	"context"
	customerpb "tmprldemo/internal/customer/pb/customer/v1"

	"github.com/google/uuid"
)

func (s *customerServiceGRPCServer) GetCustomer(ctx context.Context, request *customerpb.GetCustomerRequest) (*customerpb.GetCustomerResponse, error) {
	if _, err := uuid.Parse(request.CustomerId); err != nil {
		return nil, ErrInvalidUUID
	}

	customer, err := s.customerGetter.Get(ctx, request.CustomerId)
	if err != nil {
		return nil, err
	}

	return &customerpb.GetCustomerResponse{
		Customer: ConvertFromCustomerToPb(*customer),
	}, nil
}
