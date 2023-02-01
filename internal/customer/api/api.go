package api

import (
	"context"
	"tmprldemo/internal/customer/domain"
	customerpb "tmprldemo/internal/pb/customer/v1"

	"go.temporal.io/sdk/client"
)

// CustomerCreator provides a method for creating a customer in persistent storage.
type CustomerCreator interface {
	Create(ctx context.Context, customer domain.Customer) (*domain.Customer, error)
}

// CustomerGetter provides methods for getting customers from persistent storage.
type CustomerGetter interface {
	Get(ctx context.Context, customerID string) (*domain.Customer, error)
	BatchGet(ctx context.Context, customerIDs []string) (domain.Customers, error)
}

type customerServiceGRPCServer struct {
	customerpb.UnimplementedCustomerServiceServer
	customerCreator CustomerCreator
	customerGetter  CustomerGetter
	temporalClient  client.Client
}

// NewCustomerServiceGRPCServer creates and returns the gRPC server
// responsible for serving requests for the Customer service.
func NewCustomerServiceGRPCServer(
	customerCreator CustomerCreator,
	customerGetter CustomerGetter,
	temporalClient client.Client,
) *customerServiceGRPCServer {
	return &customerServiceGRPCServer{
		customerCreator: customerCreator,
		customerGetter:  customerGetter,
		temporalClient:  temporalClient,
	}
}
