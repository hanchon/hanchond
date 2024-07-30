package tx

import (
	"fmt"
	"os"
	"strings"

	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// ibcTransferCmd represents the ibc-transfer command
var ibcTransferCmd = &cobra.Command{
	Use:     "ibc-transfer wallet amount",
	Args:    cobra.ExactArgs(2),
	Aliases: []string{"it"},
	Short:   "Sends an ibc transaction",
	Long:    `It sends an IBC transfer from the validator wallet to the destination wallet`,
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)
		nodeID, err := cmd.Flags().GetString("node")
		if err != nil {
			fmt.Println("node not set")
			os.Exit(1)
		}

		channel, err := cmd.Flags().GetString("channel")
		if err != nil {
			fmt.Println("ibc channel not set")
			os.Exit(1)
		}

		dstWallet := args[0]
		amount := args[1]

		e := evmos.NewEvmosFromDB(queries, nodeID)
		denom, err := cmd.Flags().GetString("denom")
		if err != nil {
			fmt.Println("denom not set")
			os.Exit(1)
		}
		if denom == "" {
			denom = e.BaseDenom
		}
		out, err := e.SendIBC("transfer", channel, dstWallet, amount+denom)
		if err != nil {
			fmt.Println("error sending the transaction:", err.Error())
			os.Exit(1)
		}
		if !strings.Contains(out, "code: 0") {
			fmt.Println("transaction failed!")
			fmt.Println(out)
			os.Exit(1)
		}
		hash := strings.Split(out, "txhash: ")
		if len(hash) > 1 {
			hash[1] = strings.TrimSpace(hash[1])
		}
		fmt.Println(hash[1])
	},
}

func init() {
	TxCmd.AddCommand(ibcTransferCmd)
	ibcTransferCmd.Flags().StringP("channel", "c", "channel-0", "IBC channel")
	ibcTransferCmd.Flags().StringP("denom", "d", "", "Denom that you are sending, it defaults to the base denom of the chain")
}
