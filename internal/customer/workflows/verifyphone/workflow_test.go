package verifyphone

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/testsuite"
)

type VerifyPhoneWorkflowTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
	env        *testsuite.TestWorkflowEnvironment
	activities *activities
}

func TestVerifyPhoneWorkflow(t *testing.T) {
	suite.Run(t, new(VerifyPhoneWorkflowTestSuite))
}

func (s *VerifyPhoneWorkflowTestSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
	s.activities = &activities{}
}

func (s *VerifyPhoneWorkflowTestSuite) TearDownTest(suiteName, testName string) {
	s.env.AssertExpectations(s.T())
}

func (s *VerifyPhoneWorkflowTestSuite) TestVerifyPhoneWorkflow() {
	testCustomerID := uuid.NewString()
	testPhoneNumber := "012345678"
	testMessage := "Thanks for signing up to GoCoffee. Please enter the following code in our app to verify your phone number: 1234"

	s.env.OnActivity(s.activities.SendSMS, mock.Anything, mock.Anything).Return(
		func(phoneNumber, message string) error {
			s.Equal(testPhoneNumber, phoneNumber)
			s.Equal(testMessage, message)
			return nil
		})

	s.env.OnActivity(s.activities.VerifyCustomer, mock.Anything, mock.Anything).Return(
		func(ctx context.Context, customerID string) error {
			s.Equal(testCustomerID, customerID)
			return nil
		})

	// send the correct code on the first try - expecting workflow to complete thereafter.
	s.env.RegisterDelayedCallback(func() {
		s.env.SignalWorkflow(UserCodeSignal, "1234")
	}, time.Minute*1)

	params := WorkflowParams{
		PhoneNumber:          testPhoneNumber,
		MaximumAttempts:      2,
		CodeValidityDuration: time.Minute * 2,
	}

	s.env.SetStartWorkflowOptions(client.StartWorkflowOptions{ID: testCustomerID})
	s.env.ExecuteWorkflow(NewWorkflow, params)
	s.True(s.env.IsWorkflowCompleted())

	res, err := s.env.QueryWorkflow(VerificationResultQueryType)
	s.NoError(err)

	var result VerificationResult
	err = res.Get(&result)
	s.NoError(err)
	s.Equal(CorrectCode, result)
}

func (s *VerifyPhoneWorkflowTestSuite) TestVerifyPhoneWorkflowAllowsMultipleTries() {
	testCustomerID := uuid.NewString()
	testPhoneNumber := "012345678"
	testMessage := "Thanks for signing up to GoCoffee. Please enter the following code in our app to verify your phone number: 1234"

	s.env.OnActivity(s.activities.SendSMS, mock.Anything, mock.Anything).Return(
		func(phoneNumber, message string) error {
			s.Equal(testPhoneNumber, phoneNumber)
			s.Equal(testMessage, message)
			return nil
		}).Twice()

	s.env.OnActivity(s.activities.VerifyCustomer, mock.Anything, mock.Anything).Return(
		func(ctx context.Context, customerID string) error {
			s.Equal(testCustomerID, customerID)
			return nil
		})

	// send the incorrect code on the first try
	s.env.RegisterDelayedCallback(func() {
		s.env.SignalWorkflow(UserCodeSignal, "2345")
	}, time.Minute*1)

	// query the workflow for the result of the most recent attempt
	s.env.RegisterDelayedCallback(func() {
		res, err := s.env.QueryWorkflow(VerificationResultQueryType)
		s.NoError(err)

		var result VerificationResult
		err = res.Get(&result)
		s.Require().NoError(err)
		s.Equal(IncorrectCode, result)
	}, time.Minute*1+time.Second*10)

	// send the correct code on the second try
	s.env.RegisterDelayedCallback(func() {
		s.env.SignalWorkflow(UserCodeSignal, "1234")
	}, time.Minute*2)

	params := WorkflowParams{
		PhoneNumber:          testPhoneNumber,
		MaximumAttempts:      2,
		CodeValidityDuration: time.Minute * 3,
	}

	s.env.SetStartWorkflowOptions(client.StartWorkflowOptions{ID: testCustomerID})
	s.env.ExecuteWorkflow(NewWorkflow, params)
	s.True(s.env.IsWorkflowCompleted())

	res, err := s.env.QueryWorkflow(VerificationResultQueryType)
	s.NoError(err)

	var result VerificationResult
	err = res.Get(&result)
	s.NoError(err)
	s.Equal(CorrectCode, result)
}

func (s *VerifyPhoneWorkflowTestSuite) TestVerifyPhoneWorkflowMaximumAttemptsReached() {
	testPhoneNumber := "012345678"
	testMessage := "Thanks for signing up to GoCoffee. Please enter the following code in our app to verify your phone number: 1234"

	s.env.OnActivity(s.activities.SendSMS, mock.Anything, mock.Anything).Return(
		func(phoneNumber, message string) error {
			s.Equal(testPhoneNumber, phoneNumber)
			s.Equal(testMessage, message)
			return nil
		}).Twice()

	params := WorkflowParams{
		PhoneNumber:          testPhoneNumber,
		MaximumAttempts:      2,
		CodeValidityDuration: time.Minute * 3,
	}

	// send the incorrect code twice - hence exceeded max attempts and causing the wf to terminate
	s.env.RegisterDelayedCallback(func() {
		s.env.SignalWorkflow(UserCodeSignal, "2345")
	}, time.Minute*1)

	s.env.RegisterDelayedCallback(func() {
		s.env.SignalWorkflow(UserCodeSignal, "4567")
	}, time.Minute*2)

	s.env.ExecuteWorkflow(NewWorkflow, params)
	s.True(s.env.IsWorkflowCompleted())

	res, err := s.env.QueryWorkflow(VerificationResultQueryType)
	s.NoError(err)

	var result VerificationResult
	err = res.Get(&result)
	s.NoError(err)
	s.Equal(MaxAttemptsReached, result)
}

func (s *VerifyPhoneWorkflowTestSuite) TestVerifyPhoneWorkflowCodeExpiration() {
	testCustomerID := uuid.NewString()
	testPhoneNumber := "012345678"
	testMessage := "Thanks for signing up to GoCoffee. Please enter the following code in our app to verify your phone number: 1234"

	s.env.OnActivity(s.activities.SendSMS, mock.Anything, mock.Anything).Return(
		func(phoneNumber, message string) error {
			s.Equal(testPhoneNumber, phoneNumber)
			s.Equal(testMessage, message)
			return nil
		}).Twice()

	s.env.OnActivity(s.activities.VerifyCustomer, mock.Anything, mock.Anything).Return(
		func(ctx context.Context, customerID string) error {
			s.Equal(testCustomerID, customerID)
			return nil
		})

	params := WorkflowParams{
		PhoneNumber:          testPhoneNumber,
		MaximumAttempts:      2,
		CodeValidityDuration: time.Minute * 1,
	}

	// send the correct code after one minute, which is after it has expired.
	s.env.RegisterDelayedCallback(func() {
		s.env.SignalWorkflow(UserCodeSignal, "1234")
	}, time.Minute*2)

	// query the workflow for the result of the most recent attempt
	s.env.RegisterDelayedCallback(func() {
		res, err := s.env.QueryWorkflow(VerificationResultQueryType)
		s.NoError(err)

		var result VerificationResult
		err = res.Get(&result)
		s.Require().NoError(err)
		s.Equal(CodeExpired, result)
	}, time.Minute*2+time.Second*10)

	// send the correct code before expiry to complete the workflow.
	s.env.RegisterDelayedCallback(func() {
		s.env.SignalWorkflow(UserCodeSignal, "1234")
	}, time.Minute*3)

	s.env.SetStartWorkflowOptions(client.StartWorkflowOptions{ID: testCustomerID})
	s.env.ExecuteWorkflow(NewWorkflow, params)
	s.True(s.env.IsWorkflowCompleted())

	res, err := s.env.QueryWorkflow(VerificationResultQueryType)
	s.NoError(err)

	var result VerificationResult
	err = res.Get(&result)
	s.NoError(err)
	s.Equal(CorrectCode, result)
}
