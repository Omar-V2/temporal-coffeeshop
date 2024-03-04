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
	env         *testsuite.TestWorkflowEnvironment
	activities  *activities
	customerID  string
	codeLength  int
	phoneNumber string
	smsMessage  string
}

func TestVerifyPhoneWorkflow(t *testing.T) {
	suite.Run(t, new(VerifyPhoneWorkflowTestSuite))
}

func (s *VerifyPhoneWorkflowTestSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
	s.activities = &activities{}
	s.customerID = uuid.NewString()
	s.codeLength = 4
	s.phoneNumber = "0123456789"
	s.smsMessage = "Thanks for signing up to GoCoffee. Please enter the following code in our app to verify your phone number: 1234"

}

func (s *VerifyPhoneWorkflowTestSuite) TearDownTest(suiteName, testName string) {
	s.env.AssertExpectations(s.T())
}

func (s *VerifyPhoneWorkflowTestSuite) TestVerifyPhoneWorkflow() {
	codeValidityDuration := time.Minute * 2

	s.env.OnActivity(s.activities.NewOneTimeCode, mock.Anything, mock.Anything).Return(
		func(length int, validityDuration time.Duration) (*OneTimeCode, error) {
			s.Equal(s.codeLength, length)
			s.Equal(codeValidityDuration, validityDuration)

			return &OneTimeCode{Code: "1234", ValidUntil: s.env.Now().Add(validityDuration)}, nil
		})

	s.env.OnActivity(s.activities.SendSMS, mock.Anything, mock.Anything).Return(
		func(phoneNumber, message string) error {
			s.Equal(s.phoneNumber, phoneNumber)
			s.Equal(s.smsMessage, message)
			return nil
		})

	s.env.OnActivity(s.activities.VerifyCustomer, mock.Anything, mock.Anything).Return(
		func(ctx context.Context, customerID string) error {
			s.Equal(s.customerID, customerID)
			return nil
		})

	// send the correct code on the first try - expecting workflow to complete thereafter.
	s.env.RegisterDelayedCallback(func() {
		s.env.SignalWorkflow(UserCodeSignal, "1234")
	}, time.Minute*1)

	params := WorkflowParams{
		PhoneNumber:          s.phoneNumber,
		MaximumAttempts:      2,
		CodeLength:           s.codeLength,
		CodeValidityDuration: codeValidityDuration,
	}

	s.env.SetStartWorkflowOptions(client.StartWorkflowOptions{ID: s.customerID})
	s.env.ExecuteWorkflow(NewVerificationWorkflow, params)
	s.True(s.env.IsWorkflowCompleted())

	res, err := s.env.QueryWorkflow(VerificationResultQueryType)
	s.NoError(err)

	var result VerificationResult
	err = res.Get(&result)
	s.NoError(err)
	s.Equal(CorrectCode, result)
}

func (s *VerifyPhoneWorkflowTestSuite) TestVerifyPhoneWorkflowAllowsMultipleTries() {
	codeValidityDuration := time.Minute * 3

	s.env.OnActivity(s.activities.NewOneTimeCode, mock.Anything, mock.Anything).Return(
		func(length int, validityDuration time.Duration) (*OneTimeCode, error) {
			s.Equal(s.codeLength, length)
			s.Equal(codeValidityDuration, validityDuration)

			return &OneTimeCode{Code: "1234", ValidUntil: s.env.Now().Add(validityDuration)}, nil
		}).Twice()

	s.env.OnActivity(s.activities.SendSMS, mock.Anything, mock.Anything).Return(
		func(phoneNumber, message string) error {
			s.Equal(s.phoneNumber, phoneNumber)
			s.Equal(s.smsMessage, message)
			return nil
		}).Twice()

	s.env.OnActivity(s.activities.VerifyCustomer, mock.Anything, mock.Anything).Return(
		func(ctx context.Context, customerID string) error {
			s.Equal(s.customerID, customerID)
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
		PhoneNumber:          s.phoneNumber,
		MaximumAttempts:      2,
		CodeLength:           s.codeLength,
		CodeValidityDuration: codeValidityDuration,
	}

	s.env.SetStartWorkflowOptions(client.StartWorkflowOptions{ID: s.customerID})
	s.env.ExecuteWorkflow(NewVerificationWorkflow, params)
	s.True(s.env.IsWorkflowCompleted())

	res, err := s.env.QueryWorkflow(VerificationResultQueryType)
	s.NoError(err)

	var result VerificationResult
	err = res.Get(&result)
	s.NoError(err)
	s.Equal(CorrectCode, result)
}

func (s *VerifyPhoneWorkflowTestSuite) TestVerifyPhoneWorkflowMaximumAttemptsReached() {
	codeValidityDuration := time.Minute * 3

	s.env.OnActivity(s.activities.NewOneTimeCode, mock.Anything, mock.Anything).Return(
		func(length int, validityDuration time.Duration) (*OneTimeCode, error) {
			s.Equal(s.codeLength, length)
			s.Equal(codeValidityDuration, validityDuration)

			return &OneTimeCode{Code: "1234", ValidUntil: s.env.Now().Add(validityDuration)}, nil
		}).Twice()

	s.env.OnActivity(s.activities.SendSMS, mock.Anything, mock.Anything).Return(
		func(phoneNumber, message string) error {
			s.Equal(s.phoneNumber, phoneNumber)
			s.Equal(s.smsMessage, message)
			return nil
		}).Twice()

	// send the incorrect code twice - hence exceeded max attempts and causing the wf to terminate
	s.env.RegisterDelayedCallback(func() {
		s.env.SignalWorkflow(UserCodeSignal, "2345")
	}, time.Minute*1)

	s.env.RegisterDelayedCallback(func() {
		s.env.SignalWorkflow(UserCodeSignal, "4567")
	}, time.Minute*2)

	params := WorkflowParams{
		PhoneNumber:          s.phoneNumber,
		MaximumAttempts:      2,
		CodeLength:           s.codeLength,
		CodeValidityDuration: codeValidityDuration,
	}

	s.env.ExecuteWorkflow(NewVerificationWorkflow, params)
	s.True(s.env.IsWorkflowCompleted())

	res, err := s.env.QueryWorkflow(VerificationResultQueryType)
	s.NoError(err)

	var result VerificationResult
	err = res.Get(&result)
	s.NoError(err)
	s.Equal(MaxAttemptsReached, result)
}

func (s *VerifyPhoneWorkflowTestSuite) TestVerifyPhoneWorkflowCodeExpiration() {
	codeValidityDuration := time.Minute * 1

	s.env.OnActivity(s.activities.NewOneTimeCode, mock.Anything, mock.Anything).Return(
		func(length int, validityDuration time.Duration) (*OneTimeCode, error) {
			s.Equal(s.codeLength, length)
			s.Equal(codeValidityDuration, validityDuration)

			return &OneTimeCode{Code: "1234", ValidUntil: s.env.Now().Add(validityDuration)}, nil
		}).Twice()

	s.env.OnActivity(s.activities.SendSMS, mock.Anything, mock.Anything).Return(
		func(phoneNumber, message string) error {
			s.Equal(s.phoneNumber, phoneNumber)
			s.Equal(s.smsMessage, message)
			return nil
		}).Twice()

	s.env.OnActivity(s.activities.VerifyCustomer, mock.Anything, mock.Anything).Return(
		func(ctx context.Context, customerID string) error {
			s.Equal(s.customerID, customerID)
			return nil
		})

	// send the correct code after two minutes have elapsed, which is after the expiry time.
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

	params := WorkflowParams{
		PhoneNumber:          s.phoneNumber,
		MaximumAttempts:      2,
		CodeLength:           s.codeLength,
		CodeValidityDuration: codeValidityDuration,
	}

	s.env.SetStartWorkflowOptions(client.StartWorkflowOptions{ID: s.customerID})
	s.env.ExecuteWorkflow(NewVerificationWorkflow, params)
	s.True(s.env.IsWorkflowCompleted())

	res, err := s.env.QueryWorkflow(VerificationResultQueryType)
	s.NoError(err)

	var result VerificationResult
	err = res.Get(&result)
	s.NoError(err)
	s.Equal(CorrectCode, result)
}
