package customerdata

import (
	"context"
	"database/sql"
	"log"
	"tmprldemo/internal/customer/domain"

	sq "github.com/Masterminds/squirrel"
)

const customersTable = "customers"

type CustomerCreator interface {
	Create(ctx context.Context, customer domain.Customer) (*domain.Customer, error)
}

type CustomerDBCreator struct {
	db                   *sql.DB
	statementBuilderType sq.StatementBuilderType
}

func NewCustomerDBCreator(db *sql.DB, statementBuilderType sq.StatementBuilderType) *CustomerDBCreator {
	return &CustomerDBCreator{
		db:                   db,
		statementBuilderType: statementBuilderType,
	}
}

func (c *CustomerDBCreator) Create(ctx context.Context, customer domain.Customer) (*domain.Customer, error) {
	psql := c.statementBuilderType

	query := psql.Insert(customersTable).
		SetMap(map[string]interface{}{
			"id":             customer.ID,
			"first_name":     customer.FirstName,
			"last_name":      customer.LastName,
			"email":          customer.Email,
			"phone_number":   customer.PhoneNumber,
			"phone_verified": customer.PhoneVerified,
		}).
		Suffix(`RETURNING "id", "first_name", "last_name", "phone_number", "phone_verified"`)

	queryString, _, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	log.Printf("Create Customer SQL Query: %s", queryString)

	var createdCustomer domain.Customer
	err = query.QueryRowContext(ctx).Scan(&createdCustomer)
	if err != nil {
		return nil, err
	}

	return &createdCustomer, nil
}
