package customerdata

import (
	"context"
	"database/sql"
	"tmprldemo/internal/customer/domain"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
)

const customerTable = "customer"

type CustomerCreator interface {
	Create(ctx context.Context, customer domain.Customer) (*domain.Customer, error)
}

type CustomerDBCreator struct {
	db *sql.DB
}

func NewCustomerDBCreator(db *sql.DB) *CustomerDBCreator {
	return &CustomerDBCreator{
		db: db,
	}
}

func (c *CustomerDBCreator) Create(ctx context.Context, customer domain.Customer) (*domain.Customer, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(c.db)

	// extract request id from ctx

	// TODO: Make idempotent - check for request id conflict error.
	// if request id already exists then do a get request (GetByRequestId)
	// and return the result to the caller.

	query := psql.Insert(customerTable).
		SetMap(map[string]interface{}{
			"id":             customer.ID,
			"first_name":     customer.FirstName,
			"last_name":      customer.LastName,
			"email":          customer.Email,
			"phone_number":   customer.PhoneNumber,
			"phone_verified": customer.PhoneVerified,
		}).
		Suffix(`RETURNING "id", "first_name", "last_name", "email", "phone_number", "phone_verified"`)

	logQuery(query)

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}

	var createdCustomer domain.Customer
	dbscan.ScanOne(&createdCustomer, rows)

	return &createdCustomer, nil
}
