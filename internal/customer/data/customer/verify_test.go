package customerdata

import (
	"context"
	"database/sql"
	"testing"

	"tmprldemo/internal/customer/domain"
	migration "tmprldemo/internal/customer/migrations"
	"tmprldemo/pkg/testutils"

	"github.com/google/uuid"
	"github.com/orlangure/gnomock"
	"github.com/stretchr/testify/suite"
)

type CustomerDBVerifierTestSuite struct {
	suite.Suite
	db                *sql.DB
	postgresContainer *gnomock.Container
	getter            *CustomerDBGetter
	creator           *CustomerDBCreator
	verifier          *CustomerDBVerifier
}

func (s *CustomerDBVerifierTestSuite) SetupTest() {
	container, db := testutils.MustNewPostgresInstance(
		"customer",
		migration.Customer,
	)

	s.postgresContainer = container
	s.db = db
	s.creator = NewCustomerDBCreator(db)
	s.getter = NewCustomerDBGetter(db)
	s.verifier = NewCustomerDBVerifier(db)
}

func (s *CustomerDBVerifierTestSuite) TearDownTest() {
	s.db.Exec("TRUNCATE TABLE customer")
}

func (s *CustomerDBVerifierTestSuite) TearDownSuite() {
	gnomock.Stop(s.postgresContainer)
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
