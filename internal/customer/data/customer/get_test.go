package customerdata

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"tmprldemo/internal/customer/domain"
	"tmprldemo/pkg/testutils"

	"github.com/georgysavva/scany/v2/dbscan"
	"github.com/google/uuid"
	"github.com/orlangure/gnomock"
	"github.com/stretchr/testify/suite"
)

type CustomerDBGetterTestSuite struct {
	suite.Suite
	db                *sql.DB
	postgresContainer *gnomock.Container
	getter            *CustomerDBGetter
	creator           *CustomerDBCreator
}

func (s *CustomerDBGetterTestSuite) SetupTest() {
	container, db, err := testutils.NewPostgresInstance(
		"customer",
		"/Users/omardiab/code3/temporal-coffeeshop/internal/customer/migrations/init.sql",
	)
	s.Require().NoError(err)

	s.postgresContainer = container
	s.db = db
	s.creator = NewCustomerDBCreator(db)
	s.getter = NewCustomerDBGetter(db)
}

func (s *CustomerDBGetterTestSuite) TestGetReturnsExistingCustomer() {
	customerID := uuid.New()
	customerToCreate := domain.Customer{
		ID:            customerID,
		FirstName:     "Itachi",
		LastName:      "Uchiha",
		Email:         "itachi@leafvillage.com",
		PhoneNumber:   "07799234235",
		PhoneVerified: false,
	}

	_, err := s.creator.Create(context.Background(), customerToCreate)
	s.Require().NoError(err)

	fetchedCustomer, err := s.getter.Get(context.Background(), customerID.String())

	s.Equal(customerToCreate, *fetchedCustomer)
}

func (s *CustomerDBGetterTestSuite) TestGetReturnsErrorWhenIDNotFound() {
	fetchedCustomer, err := s.getter.Get(context.Background(), uuid.NewString())
	s.Error(err)
	s.ErrorIs(err, dbscan.ErrNotFound)
	s.Nil(fetchedCustomer)
}

func (s *CustomerDBGetterTestSuite) TestBatchGetReturnsExistingCustomers() {
	customerOneID := uuid.New()
	customerTwoID := uuid.New()
	customerThreeID := uuid.New()

	customerIDs := []uuid.UUID{customerOneID, customerTwoID, customerThreeID}

	// create three customers.
	for i, customerID := range customerIDs {
		customerToCreate := domain.Customer{
			ID:            customerID,
			FirstName:     fmt.Sprintf("Itachi%d", i),
			LastName:      fmt.Sprintf("Uchiha%d", i),
			Email:         fmt.Sprintf("itachi%d@leafvillage.com", i),
			PhoneNumber:   fmt.Sprintf("07799234235%d", i),
			PhoneVerified: false,
		}
		_, err := s.creator.Create(context.Background(), customerToCreate)
		s.Require().NoError(err)
	}

	// call batch get only issuing two of the three ids of the created customers.
	customers, err := s.getter.BatchGet(context.Background(), []string{customerOneID.String(), customerThreeID.String()})
	s.Require().NoError(err)

	var fetchedCustomerIDs []uuid.UUID
	for _, customer := range customers {
		fetchedCustomerIDs = append(fetchedCustomerIDs, customer.ID)
	}

	// assert that only the customers with the issued ids are retrieved.
	s.Contains(fetchedCustomerIDs, customerOneID)
	s.Contains(fetchedCustomerIDs, customerThreeID)
	s.NotContains(fetchedCustomerIDs, customerTwoID)

}

func TestCustomerDBGetterTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerDBGetterTestSuite))
}
