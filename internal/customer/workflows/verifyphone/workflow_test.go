package verifyphone

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
)

type VerifyPhoneWorkflowTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
	env           *testsuite.TestWorkflowEnvironment
	mockSmsSender *MockSMSSender
	smsSender     SMSSender
}

func (s *VerifyPhoneWorkflowTestSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
	s.mockSmsSender = &MockSMSSender{}
	s.smsSender = SMSSender{Sender: s.mockSmsSender}
}

func (s *VerifyPhoneWorkflowTestSuite) TearDownTest(suiteName, testName string) {
	s.env.AssertExpectations(s.T())
}

func (s *VerifyPhoneWorkflowTestSuite) TestVerifyPhoneWorkflow() {
	testPhoneNumber := "012345678"
	testMessage := "Thanks for signing up to GoCoffee. Please enter the following code in our app to verify your phone number: 1234"

	s.env.OnActivity(s.smsSender.SendSMS, mock.Anything, mock.Anything).
		Return(
			func(phoneNumber, message string) error {
				s.Equal(testPhoneNumber, phoneNumber)
				s.Equal(testMessage, message)
				return nil
			},
		)

	params := WorkflowParams{
		PhoneNumber:          testPhoneNumber,
		MaximumAttempts:      2,
		CodeValidityDuration: time.Minute * 2,
	}

	// send the correct code on the first try - expecting workflow to complete thereafter.
	s.env.RegisterDelayedCallback(func() {
		s.env.SignalWorkflow(UserCodeSignal, "1234")
	}, time.Minute*1)

	s.env.ExecuteWorkflow(NewWorkflow, params)
	s.True(s.env.IsWorkflowCompleted())

	res, err := s.env.QueryWorkflow(VerificationStateQueryType)
	s.NoError(err)

	var state VerificationState
	err = res.Get(&state)
	s.NoError(err)
	s.Equal(state, CorrectCode)
}

func (s *VerifyPhoneWorkflowTestSuite) TestVerifyPhoneWorkflowAllowsMultipleTries() {
	testPhoneNumber := "012345678"
	testMessage := "Thanks for signing up to GoCoffee. Please enter the following code in our app to verify your phone number: 1234"

	s.env.OnActivity(s.smsSender.SendSMS, mock.Anything, mock.Anything).
		Return(
			func(phoneNumber, message string) error {
				s.Equal(testPhoneNumber, phoneNumber)
				s.Equal(testMessage, message)
				return nil
			},
		).Twice()

	params := WorkflowParams{
		PhoneNumber:          testPhoneNumber,
		MaximumAttempts:      2,
		CodeValidityDuration: time.Minute * 3,
	}

	// send the incorrect code on the first try
	s.env.RegisterDelayedCallback(func() {
		s.env.SignalWorkflow(UserCodeSignal, "2345")
	}, time.Minute*1)

	// send the correct code on the second try
	s.env.RegisterDelayedCallback(func() {
		s.env.SignalWorkflow(UserCodeSignal, "1234")
	}, time.Minute*2)

	s.env.ExecuteWorkflow(NewWorkflow, params)
	s.True(s.env.IsWorkflowCompleted())

	res, err := s.env.QueryWorkflow(VerificationStateQueryType)
	s.NoError(err)

	var state VerificationState
	err = res.Get(&state)
	s.NoError(err)
	s.Equal(state, CorrectCode)
}

func (s *VerifyPhoneWorkflowTestSuite) TestVerifyPhoneWorkflowErrorsOnMaximumAttempts() {
	testPhoneNumber := "012345678"
	testMessage := "Thanks for signing up to GoCoffee. Please enter the following code in our app to verify your phone number: 1234"

	s.env.OnActivity(s.smsSender.SendSMS, mock.Anything, mock.Anything).
		Return(
			func(phoneNumber, message string) error {
				s.Equal(testPhoneNumber, phoneNumber)
				s.Equal(testMessage, message)
				return nil
			},
		)

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
	}, time.Minute*1)

	s.env.ExecuteWorkflow(NewWorkflow, params)
	s.True(s.env.IsWorkflowCompleted())

	res, err := s.env.QueryWorkflow(VerificationStateQueryType)
	s.NoError(err)

	var state VerificationState
	err = res.Get(&state)
	s.NoError(err)
	s.Equal(state, MaxAttemptsReached)

	err = s.env.GetWorkflowError()
	s.ErrorContains(err, "too many attempts")
}

func (s *VerifyPhoneWorkflowTestSuite) TestVerifyPhoneWorkflowCodeExpiration() {
	testPhoneNumber := "012345678"
	testMessage := "Thanks for signing up to GoCoffee. Please enter the following code in our app to verify your phone number: 1234"

	s.env.OnActivity(s.smsSender.SendSMS, mock.Anything, mock.Anything).
		Return(
			func(phoneNumber, message string) error {
				s.Equal(testPhoneNumber, phoneNumber)
				s.Equal(testMessage, message)
				return nil
			},
		)

	params := WorkflowParams{
		PhoneNumber:          testPhoneNumber,
		MaximumAttempts:      1,
		CodeValidityDuration: time.Minute * 1,
	}

	// send the correct code after one minute, which is after it has expired.
	s.env.RegisterDelayedCallback(func() {
		s.env.SignalWorkflow(UserCodeSignal, "1234")
	}, time.Minute*2)

	s.env.ExecuteWorkflow(NewWorkflow, params)
	s.True(s.env.IsWorkflowCompleted())

	res, err := s.env.QueryWorkflow(VerificationStateQueryType)
	s.NoError(err)

	var state VerificationState
	err = res.Get(&state)
	s.NoError(err)
	s.Equal(state, MaxAttemptsReached)

	err = s.env.GetWorkflowError()
	s.ErrorContains(err, "too many attempts")
}

func TestVerifyPhoneWorkflow(t *testing.T) {
	suite.Run(t, new(VerifyPhoneWorkflowTestSuite))
}
