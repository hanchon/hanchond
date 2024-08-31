package evmos

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hanchon/hanchond/lib/requester"
	"github.com/hanchon/hanchond/playground/cosmosdaemon"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// ethTraceCmd represents the ethTrace command
var ethTraceCmd = &cobra.Command{
	Use:   "eth-trace [tx_hash]",
	Args:  cobra.ExactArgs(1),
	Short: "Get the trace for the given tx hash",
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)

		endpoint, err := cosmosdaemon.GetWeb3Endpoint(queries, cmd)
		if err != nil {
			fmt.Printf("error generting web3 endpoint: %s\n", err.Error())
			os.Exit(1)
		}
		client := requester.NewClient().WithUnsecureWeb3Endpoint(endpoint)

		receipt, err := client.GetTransactionTrace(args[0])
		if err != nil {
			fmt.Println("could not get the ethTrace:", err.Error())
			os.Exit(1)
		}

		val, err := json.Marshal(receipt.Result)
		if err != nil {
			fmt.Println("could not process the ethTrace:", err.Error())
			os.Exit(1)
		}

		fmt.Println(string(val))
	},
}

func init() {
	EvmosCmd.AddCommand(ethTraceCmd)
}
