package erc20

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hanchon/hanchond/lib/requester"
	"github.com/hanchon/hanchond/lib/smartcontract/erc20"
	"github.com/hanchon/hanchond/playground/cosmosdaemon"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// supplyCmd represents the supply command
var supplyCmd = &cobra.Command{
	Use:   "supply [contract]",
	Args:  cobra.ExactArgs(1),
	Short: "Get the wallet supply",
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)
		contract := strings.TrimSpace(args[0])

		endpoint, err := cosmosdaemon.GetWeb3Endpoint(queries, cmd)
		if err != nil {
			fmt.Printf("error generting web3 endpoint: %s\n", err.Error())
			os.Exit(1)
		}
		client := requester.NewClient().WithUnsecureWeb3Endpoint(endpoint)

		height, _ := cmd.Flags().GetString("height")
		heightInt := erc20.Latest
		if height != "latest" {
			temp, err := strconv.ParseInt(height, 10, 64)
			if err != nil {
				fmt.Printf("invalid height: %s\n", err.Error())
				os.Exit(1)
			}
			heightInt = int(temp)
		}

		supply, err := client.GetTotalSupply(contract, heightInt)
		if err != nil {
			fmt.Println("could not get the supply:", err.Error())
			os.Exit(1)
		}
		fmt.Println(supply)
	},
}

func init() {
	ERC20Cmd.AddCommand(supplyCmd)
	supplyCmd.Flags().String("height", "latest", "Query at the given height.")
}
