package query

import (
	"os"

	"github.com/hanchon/hanchond/cmd/playground/query/erc20"
	"github.com/hanchon/hanchond/cmd/playground/query/evmos"
	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/spf13/cobra"
)

// QueryCmd represents the query command
var QueryCmd = &cobra.Command{
	Use:     "query",
	Aliases: []string{"q"},
	Short:   "Query the blockchain data",
	Run: func(cmd *cobra.Command, _ []string) {
		filesmanager.SetHomeFolderFromCobraFlags(cmd)
		_ = cmd.Help()
		os.Exit(0)
	},
}

func init() {
	QueryCmd.AddCommand(erc20.ERC20Cmd)
	QueryCmd.AddCommand(evmos.EvmosCmd)
	QueryCmd.PersistentFlags().StringP("node", "n", "1", "Playground node used to get the information")
}
