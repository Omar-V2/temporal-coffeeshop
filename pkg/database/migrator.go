package database

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

type PostgresMigrator struct {
	db *sql.DB
}

func NewPostgresMigrator(migrations embed.FS, db *sql.DB) (*PostgresMigrator, error) {
	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("pgx"); err != nil {
		return nil, fmt.Errorf("failed to set dialect: %w", err)
	}

	return &PostgresMigrator{db: db}, nil
}

func (m *PostgresMigrator) Up() error {
	return goose.Up(m.db, ".")
}

func (m *PostgresMigrator) Down() error {
	return goose.Down(m.db, ".")
}
