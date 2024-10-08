package erc20

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hanchon/hanchond/lib/converter"
	"github.com/hanchon/hanchond/lib/requester"
	"github.com/hanchon/hanchond/lib/smartcontract/erc20"
	"github.com/hanchon/hanchond/playground/cosmosdaemon"
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

		height, _ := cmd.Flags().GetString("height")
		contract := strings.TrimSpace(args[0])
		wallet := strings.TrimSpace(args[1])
		wallet, err := converter.NormalizeAddressToHex(wallet)
		if err != nil {
			fmt.Println("could not convert address to hex encoding")
			os.Exit(1)
		}

		endpoint, err := cosmosdaemon.GetWeb3Endpoint(queries, cmd)
		if err != nil {
			fmt.Printf("error generting web3 endpoint: %s\n", err.Error())
			os.Exit(1)
		}

		client := requester.NewClient().WithUnsecureWeb3Endpoint(endpoint)
		heightInt := erc20.Latest
		if height != "latest" {
			temp, err := strconv.ParseInt(height, 10, 64)
			if err != nil {
				fmt.Printf("invalid height: %s\n", err.Error())
				os.Exit(1)
			}
			heightInt = int(temp)
		}
		balance, err := client.GetBalanceERC20(contract, wallet, heightInt)
		if err != nil {
			fmt.Println("could not get the balance:", err.Error())
			os.Exit(1)
		}
		fmt.Println(balance)
	},
}

func init() {
	ERC20Cmd.AddCommand(balanceCmd)
	balanceCmd.Flags().String("height", "latest", "Query at the given height.")
}
