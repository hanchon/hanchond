package sql

import (
	"context"
	"database/sql"

	_ "embed"
	// blank import to support sqlite3
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

func InitDatabase(ctx context.Context, path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	return db, nil
}
