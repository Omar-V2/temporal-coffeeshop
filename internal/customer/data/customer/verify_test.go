package customerdata

import (
	"context"
	"database/sql"
	"testing"

	"tmprldemo/internal/customer/domain"
	migration "tmprldemo/internal/customer/migrations"
	"tmprldemo/pkg/testutils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
)

type CustomerDBVerifierTestSuite struct {
	suite.Suite
	testCtx           context.Context
	postgresContainer testcontainers.Container
	db                *sql.DB
	getter            *CustomerDBGetter
	creator           *CustomerDBCreator
	verifier          *CustomerDBVerifier
}

func (s *CustomerDBVerifierTestSuite) SetupTest() {
	s.testCtx = context.Background()
	s.postgresContainer, s.db = testutils.MustNewPostgresInstance(
		s.testCtx,
		"customer",
		migration.Customer,
	)

	s.creator = NewCustomerDBCreator(s.db)
	s.getter = NewCustomerDBGetter(s.db)
	s.verifier = NewCustomerDBVerifier(s.db)
}

func (s *CustomerDBVerifierTestSuite) TearDownSuite() {
	s.postgresContainer.Terminate(s.testCtx)
}

func (s *CustomerDBVerifierTestSuite) TearDownTest() {
	s.db.Exec("TRUNCATE TABLE customer")
}

func (s *CustomerDBVerifierTestSuite) TestVerify() {
	customerID := uuid.New()
	customer := domain.Customer{
		ID:            customerID,
		FirstName:     "Madara",
		LastName:      "Uchiha",
		Email:         "madara@akatsuki.com",
		PhoneNumber:   "072312490",
		PhoneVerified: false,
	}

	_, err := s.creator.Create(context.Background(), customer)
	s.Require().NoError(err)

	err = s.verifier.Verify(context.Background(), customerID.String())
	s.Require().NoError(err)

	verifiedCustomer, err := s.getter.Get(context.Background(), customerID.String())
	s.Require().NoError(err)

	s.True(verifiedCustomer.PhoneVerified)
}

func TestCustomerDBVerifierTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerDBVerifierTestSuite))

}
