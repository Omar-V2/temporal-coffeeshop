package customerdata

import (
	"context"
	"database/sql"
	"tmprldemo/internal/customer/domain"

	sq "github.com/Masterminds/squirrel"
)

const customersTable = "customers"

type CustomerCreator interface {
	Create(ctx context.Context, customer domain.Customer) (domain.Customer, error)
}

type CustomerDBCreator struct {
	db *sql.DB
}

func NewCustomerDBCreator(db *sql.DB) *CustomerDBCreator {
	return &CustomerDBCreator{db: db}
}

func (c *CustomerDBCreator) Create(ctx context.Context, customer domain.Customer) (domain.Customer, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(c.db)

	// var c domain.Customer
	result, err := psql.Insert(customersTable).
		SetMap(map[string]interface{}{
			"id":             customer.ID,
			"first_name":     customer.FirstName,
			"last_name":      customer.LastName,
			"phone_number":   customer.PhoneNumber,
			"phone_verified": customer.PhoneVerified,
		}).
		ExecContext(ctx)

	return domain.Customer{}, nil
}
