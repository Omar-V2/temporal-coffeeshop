package integration_test

import (
	"context"
	customerpb "tmprldemo/internal/pb/customer/v1"

	"github.com/google/uuid"
	"go.temporal.io/api/enums/v1"
)

func (s *CustomerIntegrationTestSuite) TestCreateCustomer() {
	testCtx := context.Background()

	requestID := uuid.NewString()
	customerID := uuid.NewString()

	customerPb := customerpb.Customer{
		Id:          customerID,
		FirstName:   "Test",
		LastName:    "Customer",
		Email:       "testcustomer@gmail.com",
		PhoneNumber: "07500512380",
	}

	createCustomerRequest := &customerpb.CreateCustomerRequest{
		RequestId: requestID,
		Customer:  &customerPb,
	}

	_, err := s.customerServiceClient.CreateCustomer(testCtx, createCustomerRequest)
	s.Require().NoError(err)

	workflowExecution, err := s.temporalClient.DescribeWorkflowExecution(testCtx, customerID, "")
	s.Require().NoError(err)

	s.Equal(enums.WORKFLOW_EXECUTION_STATUS_RUNNING, workflowExecution.WorkflowExecutionInfo.Status)
}

func (s *CustomerIntegrationTestSuite) TestVerifyCustomer() {
	testCtx := context.Background()

	requestID := uuid.NewString()
	customerID := uuid.NewString()

	customerPb := customerpb.Customer{
		Id:          customerID,
		FirstName:   "Test",
		LastName:    "Customer",
		Email:       "testcustomer@gmail.com",
		PhoneNumber: "07500512380",
	}

	createCustomerRequest := &customerpb.CreateCustomerRequest{
		RequestId: requestID,
		Customer:  &customerPb,
	}

	_, err := s.customerServiceClient.CreateCustomer(testCtx, createCustomerRequest)
	s.Require().NoError(err)

	verifyRequest := &customerpb.VerifyCustomerRequest{
		CustomerId:       customerID,
		VerificationCode: "1234",
	}
	_, err = s.customerServiceClient.VerifyCustomer(testCtx, verifyRequest)
	s.Require().NoError(err)

	err = s.temporalClient.GetWorkflow(testCtx, customerID, "").Get(testCtx, nil)
	s.Require().NoError(err)
}
