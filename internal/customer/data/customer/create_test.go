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

type CustomerDBCreatorTestSuite struct {
	suite.Suite
	testCtx           context.Context
	postgresContainer testcontainers.Container
	db                *sql.DB
	creator           *CustomerDBCreator
	getter            *CustomerDBGetter
}

func (s *CustomerDBCreatorTestSuite) SetupSuite() {
	s.testCtx = context.Background()

	s.postgresContainer, s.db = testutils.MustNewPostgresInstance(
		s.testCtx,
		"customer",
		migration.Customer,
	)

	s.creator = NewCustomerDBCreator(s.db)
	s.getter = NewCustomerDBGetter(s.db)
}

func (s *CustomerDBCreatorTestSuite) TearDownTest() {
	s.db.Exec("TRUNCATE TABLE customer")
}

func (s *CustomerDBCreatorTestSuite) TearDownSuite() {
	s.postgresContainer.Terminate(s.testCtx)
}

func (s *CustomerDBCreatorTestSuite) TestCreate() {
	customerID := uuid.New()
	customerToCreate := domain.Customer{
		ID:            customerID,
		FirstName:     "Sasuke",
		LastName:      "Uchiha",
		Email:         "sasuke@leafvillage.com",
		PhoneNumber:   "07799234235",
		PhoneVerified: false,
	}

	createdCustomer, err := s.creator.Create(context.Background(), customerToCreate)
	s.Require().NoError(err)

	s.Equal(createdCustomer.ID.String(), customerID.String())

	fetchedCustomer, err := s.getter.Get(context.Background(), customerID.String())
	s.Require().NoError(err)

	s.Equal(customerToCreate, *fetchedCustomer)
}

func TestCustomerDBCreatorTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerDBCreatorTestSuite))
}
