package api

import (
	"context"
	"fmt"
	"time"

	customerpb "tmprldemo/internal/pb/customer/v1"

	"tmprldemo/internal/customer/workflows/verifyphone"

	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
)

// CreateCustomer creates a new customer in the coffee shop system.
// It also executes a workflow responsible for the phone verification process for the newly created customer.
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

	workflowOptions := client.StartWorkflowOptions{
		ID:                       createdCustomer.ID.String(),
		TaskQueue:                "TEMPORAL_COFFEE_SHOP_TASK_QUEUE",
		WorkflowExecutionTimeout: time.Hour * 24,
		WorkflowIDReusePolicy:    enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY,
	}
	params := verifyphone.WorkflowParams{
		PhoneNumber:          createdCustomer.PhoneNumber,
		MaximumAttempts:      3,
		CodeLength:           4,
		CodeValidityDuration: time.Second * 30,
	}
	_, err = s.temporalClient.ExecuteWorkflow(ctx, workflowOptions, "VerifyPhoneWorkflow", params)
	if err != nil {
		return nil, fmt.Errorf("failed to execute verify phone workflow: %w", err)
	}

	return &customerpb.CreateCustomerResponse{
		Customer: ConvertFromCustomerToPb(*createdCustomer),
	}, nil
}
