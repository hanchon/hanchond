package explorer

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// represents the query command
var startCmd = &cobra.Command{
	Use:   "start",
	Args:  cobra.ExactArgs(0),
	Short: "Start the node explorer",
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)
		nodeID, err := cmd.Flags().GetString("node")
		if err != nil {
			fmt.Println("node not set")
			os.Exit(1)
		}

		// TODO: move the newFromDB to cosmos daemon
		e := evmos.NewEvmosFromDB(queries, nodeID)

		fmt.Println(e.Ports.P8545)
		fmt.Println(e.Ports.P1317)

	},
}

func init() {
	ExplorerCmd.AddCommand(startCmd)
}
