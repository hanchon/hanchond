package evmos

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/lib/requester"
	"github.com/hanchon/hanchond/playground/cosmosdaemon"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// ethCodeCmd represents the ethCode command
var ethCodeCmd = &cobra.Command{
	Use:   "eth-code [address]",
	Args:  cobra.ExactArgs(1),
	Short: "Get the smartcontract registered eth code",
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)

		height, err := cmd.Flags().GetString("height")
		if err != nil {
			fmt.Println("error getting the request height:", err.Error())
			os.Exit(1)
		}

		endpoint, err := cosmosdaemon.GetWeb3Endpoint(queries, cmd)
		if err != nil {
			fmt.Printf("error generting web3 endpoint: %s\n", err.Error())
			os.Exit(1)
		}
		client := requester.NewClient().WithUnsecureWeb3Endpoint(endpoint)

		code, err := client.EthCode(args[0], height)
		if err != nil {
			fmt.Println("could not get the ethCode:", err.Error())
			os.Exit(1)
		}

		fmt.Println(string(code))
	},
}

func init() {
	EvmosCmd.AddCommand(ethCodeCmd)
	ethCodeCmd.Flags().String("height", "latest", "Query at the given height.")
}
