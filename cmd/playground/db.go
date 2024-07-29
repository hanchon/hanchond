package playground

import (
	"context"

	"github.com/hanchon/hanchond/playground/database"
	"github.com/hanchon/hanchond/playground/sql"
)

func initDB(dbPath string) (*database.Queries, error) {
	db, err := sql.InitDatabase(context.Background(), dbPath+"/playground.db")
	if err != nil {
		return nil, err
	}
	return database.New(db), nil
}
