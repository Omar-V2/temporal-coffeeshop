package api

import (
	"context"
	"fmt"

	"tmprldemo/internal/customer/workflows/verifyphone"
	customerpb "tmprldemo/internal/pb/customer/v1"
)

func (s customerServiceGRPCServer) VerifyCustomer(ctx context.Context, request *customerpb.VerifyCustomerRequest) (*customerpb.VerifyCustomerResponse, error) {
	workflowID := request.CustomerId
	err := s.temporalClient.SignalWorkflow(ctx, workflowID, "", verifyphone.UserCodeSignal, request.VerificationCode)
	if err != nil {
		return nil, fmt.Errorf("failed to signal verify phone workflow with ID: %s, err: %w", workflowID, err)
	}

	response, err := s.temporalClient.QueryWorkflow(ctx, workflowID, "", verifyphone.VerificationResultQueryType)
	if err != nil {
		return nil, fmt.Errorf("failed to  query verify phone workflow with ID %s, err: %w", workflowID, err)
	}

	var result verifyphone.VerificationResult
	err = response.Get(&result)
	if err != nil {
		return nil, err
	}

	return &customerpb.VerifyCustomerResponse{
		CustomerId: request.CustomerId,
		Result:     convertWorkflowResultToPb(result),
	}, nil
}
