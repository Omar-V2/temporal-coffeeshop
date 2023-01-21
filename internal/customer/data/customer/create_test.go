package customerdata

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"tmprldemo/internal/customer/domain"
	"tmprldemo/pkg/testutils"

	"github.com/google/uuid"
	"github.com/orlangure/gnomock"
	"github.com/stretchr/testify/suite"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type CustomerDBCreatorTestSuite struct {
	suite.Suite
	postgresContaniner *gnomock.Container
	db                 *sql.DB
	creator            *CustomerDBCreator
	getter             *CustomerDBGetter
}

func (s *CustomerDBCreatorTestSuite) SetupSuite() {
	container, db, err := testutils.NewPostgresInstance(
		"customer",
		"/Users/omardiab/code3/temporal-coffeeshop/internal/customer/migrations/init.sql",
	)
	s.Require().NoError(err)

	s.postgresContaniner = container
	s.db = db
	s.creator = NewCustomerDBCreator(db)
	s.getter = NewCustomerDBGetter(db)
}

func (s *CustomerDBCreatorTestSuite) TearDownTest() {
	s.db.Exec("TRUNCATE TABLE customer")
}

func (s *CustomerDBCreatorTestSuite) TearDownSuite() {
	gnomock.Stop(s.postgresContaniner)
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

	fmt.Println("created customer ID is ", createdCustomer.ID.String())
	s.Equal(createdCustomer.ID.String(), customerID.String())

	fetchedCustomer, err := s.getter.Get(context.Background(), customerID.String())
	s.Require().NoError(err)

	s.Equal(customerToCreate, *fetchedCustomer)
}

func TestCustomerDBCreatorTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerDBCreatorTestSuite))
}
