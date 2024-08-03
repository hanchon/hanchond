package erc20

import (
	"fmt"
	"os"
	"strings"

	"github.com/hanchon/hanchond/lib/converter"
	"github.com/hanchon/hanchond/lib/requester"
	"github.com/hanchon/hanchond/lib/smartcontract/erc20"
	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance [contract] [wallet]",
	Args:  cobra.ExactArgs(2),
	Short: "Get the wallet erc20 balance",
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)
		nodeID, err := cmd.Flags().GetString("node")
		if err != nil {
			fmt.Println("node not set")
			os.Exit(1)
		}
		contract := strings.TrimSpace(args[0])
		wallet := strings.TrimSpace(args[1])
		wallet, err = converter.NormalizeAddressToHex(wallet)
		if err != nil {
			fmt.Println("could not convert address to hex encoding")
			os.Exit(1)
		}

		e := evmos.NewEvmosFromDB(queries, nodeID)
		client := requester.NewClient().WithUnsecureWeb3Endpoint(fmt.Sprintf("http://localhost:%d", e.Ports.P8545))
		balance, err := client.GetBalanceERC20(contract, wallet, erc20.Latest)
		if err != nil {
			fmt.Println("could not get the balance:", err.Error())
			os.Exit(1)
		}
		fmt.Println(balance)
	},
}

func init() {
	ERC20Cmd.AddCommand(balanceCmd)
}
