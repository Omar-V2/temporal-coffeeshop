package api

import (
	customerdata "tmprldemo/internal/customer/data/customer"
	customerpb "tmprldemo/internal/pb/customer/v1"

	"go.temporal.io/sdk/client"
)

type customerServiceGRPCServer struct {
	customerpb.UnimplementedCustomerServiceServer
	customerCreator customerdata.CustomerCreator
	customerGetter  customerdata.CustomerGetter
	temporalClient  client.Client
}

// NewCustomerServiceGRPCServer creates and returns the gRPC server
// responsible for serving requests for the Customer service.
func NewCustomerServiceGRPCServer(
	customerCreator customerdata.CustomerCreator,
	customerGetter customerdata.CustomerGetter,
	temporalClient client.Client,
) *customerServiceGRPCServer {
	return &customerServiceGRPCServer{
		customerCreator: customerCreator,
		customerGetter:  customerGetter,
		temporalClient:  temporalClient,
	}
}
