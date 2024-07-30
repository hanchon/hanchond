package query

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// represents the query command
var balanceCmd = &cobra.Command{
	Use:   "balance [wallet]",
	Args:  cobra.ExactArgs(1),
	Short: "Get the wallet balance",
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)
		nodeID, err := cmd.Flags().GetString("node")
		if err != nil {
			fmt.Println("node not set")
			os.Exit(1)
		}

		wallet := args[0]

		e := evmos.NewEvmosFromDB(queries, nodeID)
		balance, err := e.CheckBalance(wallet)
		if err != nil {
			fmt.Println("could not get the balance:", err.Error())
			os.Exit(1)
		}
		fmt.Println(balance)
	},
}

func init() {
	QueryCmd.AddCommand(balanceCmd)
}
