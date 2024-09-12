package sql

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"os"

	// blank import to support sqlite3
	_ "modernc.org/sqlite"

	"github.com/hanchon/hanchond/playground/database"
	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/spf13/cobra"
)

//go:embed schema.sql
var ddl string

func InitDatabase(ctx context.Context, path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	return db, nil
}

func initDB(dbPath string) (*database.Queries, error) {
	// TODO: move the database path to the filesmanager
	db, err := InitDatabase(context.Background(), dbPath+"/playground.db")
	if err != nil {
		return nil, err
	}
	return database.New(db), nil
}

func InitDBFromCmd(cmd *cobra.Command) *database.Queries {
	home := filesmanager.SetHomeFolderFromCobraFlags(cmd)
	queries, err := initDB(home)
	if err != nil {
		fmt.Println("could not init database", err.Error())
		os.Exit(1)
	}
	return queries
}
