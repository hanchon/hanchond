package explorer

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/explorer"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// represents the query command
var startCmd = &cobra.Command{
	Use:   "start",
	Args:  cobra.ExactArgs(0),
	Short: "Start the node explorer",
	Run: func(cmd *cobra.Command, _ []string) {
		queries := sql.InitDBFromCmd(cmd)
		nodeID, err := cmd.Flags().GetString("node")
		if err != nil {
			fmt.Println("node not set")
			os.Exit(1)
		}

		startingHeight, err := cmd.Flags().GetInt("starting-height")
		if err != nil {
			fmt.Println("starting height not set")
			os.Exit(1)
		}

		// TODO: move the newFromDB to cosmos daemon
		e := evmos.NewEvmosFromDB(queries, nodeID)

		ex := explorer.NewLocalExplorerClient(e.Ports.P8545, e.Ports.P1317, e.HomeDir)
		if err := ex.ProcessMissingBlocks(int64(startingHeight)); err != nil {
			fmt.Println("error getting missing blocks: ", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	ExplorerCmd.AddCommand(startCmd)
	startCmd.Flags().Int("starting-height", 1, "Starting height to index the chain.")
}
