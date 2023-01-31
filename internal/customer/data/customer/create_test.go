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

type CustomerDBCreatorTestSuite struct {
	suite.Suite
	postgresContainer *gnomock.Container
	db                *sql.DB
	creator           *CustomerDBCreator
	getter            *CustomerDBGetter
}

func (s *CustomerDBCreatorTestSuite) SetupSuite() {
	container, db := testutils.MustNewPostgresInstance(
		"customer",
		migration.Customer,
	)

	s.postgresContainer = container
	s.db = db
	s.creator = NewCustomerDBCreator(db)
	s.getter = NewCustomerDBGetter(db)
}

func (s *CustomerDBCreatorTestSuite) TearDownTest() {
	s.db.Exec("TRUNCATE TABLE customer")
}

func (s *CustomerDBCreatorTestSuite) TearDownSuite() {
	gnomock.Stop(s.postgresContainer)
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
