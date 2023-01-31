package testutils

import (
	"database/sql"
	"embed"
	"fmt"

	"tmprldemo/pkg/database"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
)

// MustNewPostgresInstance starts a container running postgres and returns the container and the db instance.
// This function accepts a migrations argument which it executes against the db instance running inside the container.
func MustNewPostgresInstance(dbName string, migrations embed.FS) (*gnomock.Container, *sql.DB) {
	p := postgres.Preset(
		postgres.WithDatabase(dbName),
	)

	container, err := gnomock.Start(p)
	if err != nil {
		panic(err)
	}

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		"postgres", "password", container.Host, container.DefaultPort(), dbName,
	)

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		panic(err)
	}

	migrator, err := database.NewPostgresMigrator(migrations, db)
	if err != nil {
		panic(err)
	}

	err = migrator.Up()
	if err != nil {
		panic(err)
	}

	return container, db
}
