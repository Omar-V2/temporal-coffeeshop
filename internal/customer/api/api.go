package api

import (
	customerdata "tmprldemo/internal/customer/data/customer"
	customerpb "tmprldemo/internal/customer/pb/customer/v1"
)

type customerServiceGRPCServer struct {
	customerpb.UnimplementedCustomerServiceServer
	customerCreator customerdata.CustomerCreator
}

func NewCustomerServiceGRPCServer(customerCreator customerdata.CustomerCreator) *customerServiceGRPCServer {
	return &customerServiceGRPCServer{
		customerCreator: customerCreator,
	}
}
