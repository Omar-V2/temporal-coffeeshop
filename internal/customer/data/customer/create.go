package customerdata

import (
	"context"
	"database/sql"
	"log"
	"tmprldemo/internal/customer/domain"

	sq "github.com/Masterminds/squirrel"
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

	// TODO: Make idempotent
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

	queryString, _, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	log.Printf("Create Customer SQL Query: %s", queryString)

	var createdCustomer domain.Customer
	err = query.
		ScanContext(
			ctx,
			&createdCustomer.ID,
			&createdCustomer.FirstName,
			&createdCustomer.LastName,
			&createdCustomer.Email,
			&createdCustomer.PhoneNumber,
			&createdCustomer.PhoneVerified,
		)
	if err != nil {
		return nil, err
	}

	return &createdCustomer, nil
}
