package customerdata

import (
	"context"
	"database/sql"
	"log"
	"tmprldemo/internal/customer/domain"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
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

	logQuery(query)

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}

	var c domain.Customer
	dbscan.ScanOne(c, rows)

	return &c, nil
}

func (g *CustomerDBGetter) BatchGet(ctx context.Context, customerIDs []string) (domain.Customers, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(g.db)

	query := psql.Select("*").
		From(customerTable).
		Where(sq.Eq{"id": customerIDs})

	logQuery(query)

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}

	var customers domain.Customers
	dbscan.ScanAll(&customers, rows)

	return customers, nil
}

type convertableQuery interface {
	MustSql() (string, []interface{})
}

func logQuery(query convertableQuery) {
	queryString, _ := query.MustSql()
	log.Printf("Get Customer SQL Query: %s", queryString)
}
