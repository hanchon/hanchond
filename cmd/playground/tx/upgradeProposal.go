package tx

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// upgradeProposalCmd represents the upgrade-proposal command
var upgradeProposalCmd = &cobra.Command{
	Use:   "upgrade-proposal [version] [height]",
	Args:  cobra.ExactArgs(2),
	Short: "Create an upgrade propsal",
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)
		nodeID, err := cmd.Flags().GetString("node")
		if err != nil {
			fmt.Println("node not set")
			os.Exit(1)
		}

		version := args[0]
		height := args[1]

		e := evmos.NewEvmosFromDB(queries, nodeID)
		txhash, err := e.CreateUpgradeProposal(version, height)
		if err != nil {
			fmt.Println("error sending the transaction:", err.Error())
			os.Exit(1)
		}

		fmt.Println(txhash)
	},
}

func init() {
	TxCmd.AddCommand(upgradeProposalCmd)
}
