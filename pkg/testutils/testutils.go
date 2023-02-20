package testutils

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"time"

	"tmprldemo/pkg/database"

	"github.com/docker/go-connections/nat"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/phayes/freeport"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// MustNewPostgresInstance starts a container running postgres and returns the container and the db instance.
// This function accepts a migrations argument which it executes against the db instance running inside the container.
func MustNewPostgresInstance(ctx context.Context, dbName string, migrations embed.FS) (testcontainers.Container, *sql.DB) {
	req := testcontainers.ContainerRequest{
		Image: "postgres:14-alpine",
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "password",
			"POSTGRES_DB":       dbName,
		},
		ExposedPorts: []string{"5432"},
		Cmd:          []string{"postgres", "-c", "fsync=off"},
		WaitingFor: wait.ForAll(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5 * time.Second),
		).WithDeadline(1 * time.Minute),
	}

	postgresContainer, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		},
	)
	if err != nil {
		panic(err)
	}

	host, err := postgresContainer.Host(ctx)
	if err != nil {
		panic(err)
	}

	natPort, err := nat.NewPort("tcp", "5432")
	if err != nil {
		panic(err)
	}
	port, err := postgresContainer.MappedPort(ctx, natPort)
	if err != nil {
		panic(err)
	}

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		"postgres", "password", host, port.Int(), dbName,
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

	return postgresContainer, db
}

func GetFreeAddress(host string) (string, error) {
	port, err := freeport.GetFreePort()
	if err != nil {
		return "", fmt.Errorf("failed to allocate free port: %w", err)
	}

	return fmt.Sprintf("%s:%d", host, port), nil
}
