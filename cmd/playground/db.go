package playground

import (
	"context"
	"fmt"
	"os"

	"github.com/hanchon/hanchond/playground/database"
	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

func initDB(dbPath string) (*database.Queries, error) {
	db, err := sql.InitDatabase(context.Background(), dbPath+"/playground.db")
	if err != nil {
		return nil, err
	}
	return database.New(db), nil
}

func initDBFromCmd(cmd *cobra.Command) *database.Queries {
	home := filesmanager.SetHomeFolderFromCobraFlags(cmd)
	queries, err := initDB(home)
	if err != nil {
		fmt.Println("could not init database", err.Error())
		os.Exit(1)
	}
	return queries
}
