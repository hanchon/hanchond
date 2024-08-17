package tx

import (
	"os"

	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/spf13/cobra"

	"github.com/hanchon/hanchond/cmd/playground/tx/solidity"
)

// TxCmd represents the tx command
var TxCmd = &cobra.Command{
	Use:     "tx",
	Aliases: []string{"t"},
	Short:   "Send transactions",
	Run: func(cmd *cobra.Command, _ []string) {
		filesmanager.SetHomeFolderFromCobraFlags(cmd)
		_ = cmd.Help()
		os.Exit(0)
	},
}

func init() {
	TxCmd.AddCommand(solidity.SolidityCmd)
	TxCmd.PersistentFlags().StringP("node", "n", "1", "Playground node that is sending the transaction")
}
