package customerdata

import (
	"context"
	"database/sql"
	"log"
	"tmprldemo/internal/customer/domain"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
)

type CustomerDBGetter struct {
	db *sql.DB
}

func NewCustomerDBGetter(db *sql.DB) *CustomerDBGetter {
	return &CustomerDBGetter{db: db}
}

// TODO: handling SQL errors, not found, conflicts etc.
func (g *CustomerDBGetter) Get(ctx context.Context, customerID string) (*domain.Customer, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(g.db)

	query := psql.
		Select("*").
		From(customerTable).
		Where(sq.Eq{"id": customerID})

	queryString, _ := query.MustSql()
	log.Printf("Get Customer SQL Query: %s", queryString)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var c domain.Customer
	if err = dbscan.ScanOne(&c, rows); err != nil {
		return nil, err
	}

	return &c, nil
}

func (g *CustomerDBGetter) BatchGet(ctx context.Context, customerIDs []string) (domain.Customers, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(g.db)

	query := psql.Select("*").
		From(customerTable).
		Where(sq.Eq{"id": customerIDs})

	queryString, _ := query.MustSql()
	log.Printf("Batch Get Customer SQL Query: %s", queryString)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var customers domain.Customers
	if err = dbscan.ScanAll(&customers, rows); err != nil {
		return nil, err
	}

	return customers, nil
}
