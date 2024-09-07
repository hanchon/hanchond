package explorer

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/explorer"
	"github.com/hanchon/hanchond/playground/explorer/explorerui"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// ui represents the query command
var uiCmd = &cobra.Command{
	Use:   "ui",
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
		// TODO: support mainnet and testnet endpoints
		ex := explorer.NewLocalExplorerClient(e.Ports.P8545, e.Ports.P1317, e.HomeDir)

		p := explorerui.CreateExplorerTUI(startingHeight, ex)
		if _, err := p.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	},
}

func init() {
	ExplorerCmd.AddCommand(uiCmd)
	uiCmd.Flags().Int("starting-height", 1, "Starting height to index the chain.")
}
