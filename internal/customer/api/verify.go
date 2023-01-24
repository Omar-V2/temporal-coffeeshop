package api

import (
	"context"

	customerpb "tmprldemo/internal/pb/customer/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s customerServiceGRPCServer) VerifyCustomer(context.Context, *customerpb.VerifyCustomerRequest) (*emptypb.Empty, error) {
	return nil, nil
}
