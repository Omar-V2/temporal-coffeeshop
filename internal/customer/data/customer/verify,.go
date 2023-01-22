package customerdata

import (
	"context"
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
)

type CustomerVerifier interface {
	Verify(ctc context.Context, customerID string) error
}

type CustomerDBVerifier struct {
	db *sql.DB
}

func NewCustomerDBVerifier(db *sql.DB) *CustomerDBVerifier {
	return &CustomerDBVerifier{db: db}
}

func (v CustomerDBVerifier) Verify(ctx context.Context, customerID string) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(v.db)

	query := psql.Update(customerTable).
		Where(sq.Eq{"id": customerID}).
		Set("phone_verified", true)

	queryString, _, err := query.ToSql()
	if err != nil {
		return err
	}
	log.Printf("Verify Customer SQL Query: %s", queryString)

	_, err = query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
