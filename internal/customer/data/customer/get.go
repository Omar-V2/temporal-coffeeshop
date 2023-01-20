package customerdata

import (
	"context"
	"database/sql"
	"log"
	"tmprldemo/internal/customer/domain"

	sq "github.com/Masterminds/squirrel"
)

type CustomerGetter interface {
	Get(ctx context.Context, customerID string) (*domain.Customer, error)
	BatchGet(ctx context.Context, customerIDs []string) (domain.Customers, error)
}

type CustomerDBGetter struct {
	db *sql.DB
}

func NewCustomerDBGetter(db *sql.DB) *CustomerDBGetter {
	return &CustomerDBGetter{
		db: db,
	}
}

// TODO: handling SQL errors, not found, conflicts etc.
func (g *CustomerDBGetter) Get(ctx context.Context, customerID string) (*domain.Customer, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(g.db)

	query := psql.
		Select("*").
		From(customerTable).
		Where(sq.Eq{"id": customerID})

	queryString, _, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	log.Printf("Get Customer SQL Query: %s", queryString)

	var c domain.Customer
	err = query.
		ScanContext(
			ctx,
			&c.ID,
			&c.FirstName,
			&c.LastName,
			&c.Email,
			&c.PhoneNumber,
			&c.PhoneVerified,
		)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (g *CustomerDBGetter) BatchGet(ctx context.Context, customerIDs []string) (domain.Customers, error) {
	return nil, nil
}
