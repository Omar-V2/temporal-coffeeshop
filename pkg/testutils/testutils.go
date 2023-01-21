package testutils

import (
	"database/sql"
	"fmt"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
)

func NewPostgresInstance(dbName, queriesFile string) (*gnomock.Container, *sql.DB, error) {
	p := postgres.Preset(
		postgres.WithDatabase(dbName),
		// TODO: replace with migrate command in code to create schemas
		postgres.WithQueriesFile(queriesFile),
	)

	container, err := gnomock.Start(p)
	if err != nil {
		return nil, nil, err
	}

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		"postgres", "password", container.Host, container.DefaultPort(), dbName,
	)

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, nil, err
	}

	return container, db, nil
}
