package query

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// txCmd represents the query tx command
var txCmd = &cobra.Command{
	Use:   "tx [txhash]",
	Args:  cobra.ExactArgs(1),
	Short: "Get the transaction info",
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)
		nodeID, err := cmd.Flags().GetString("node")
		if err != nil {
			fmt.Println("node not set")
			os.Exit(1)
		}

		txhash := args[0]

		e := evmos.NewEvmosFromDB(queries, nodeID)
		resp, err := e.GetTransaction(txhash)
		if err != nil {
			fmt.Println("could not get the balance:", err.Error())
			os.Exit(1)
		}
		fmt.Println(resp)
	},
}

func init() {
	QueryCmd.AddCommand(txCmd)
}
