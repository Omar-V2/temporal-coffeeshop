package api

import customerpb "tmprldemo/internal/customer/pb/customer/v1"

type customerServiceGRPCServer struct {
	customerpb.UnimplementedCustomerServiceServer
	// creator, verifier, getter and batch getter interfaces go here
}

func NewCustomerServiceGRPCServer() *customerServiceGRPCServer {
	return &customerServiceGRPCServer{}
}
