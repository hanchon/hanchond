package solidity

import (
	"fmt"
	"os"
	"strings"

	"github.com/hanchon/hanchond/lib/requester"
	"github.com/hanchon/hanchond/lib/smartcontract"
	"github.com/hanchon/hanchond/playground/cosmosdaemon"
	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

var params []string

// callContractViewCmd represents the callContractView command
var callContractViewCmd = &cobra.Command{
	Use:   "call-contract-view [contract] [abi_path] [method]",
	Args:  cobra.ExactArgs(3),
	Short: "Call a contract view with eth_call",
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)

		height, err := cmd.Flags().GetString("height")
		if err != nil {
			fmt.Println("could not read height value:", err.Error())
			os.Exit(1)
		}

		contract := strings.TrimSpace(args[0])
		abiPath := strings.TrimSpace(args[1])
		method := strings.TrimSpace(args[2])

		abiBytes, err := filesmanager.ReadFile(abiPath)
		if err != nil {
			fmt.Printf("error reading the abi file:%s\n", err.Error())
			os.Exit(1)
		}

		endpoint, err := cosmosdaemon.GetWeb3Endpoint(queries, cmd)
		if err != nil {
			fmt.Printf("error generting web3 endpoint: %s\n", err.Error())
			os.Exit(1)
		}

		callArgs, err := smartcontract.StringsToABIArguments(params)
		if err != nil {
			fmt.Printf("error converting arguments: %s\n", err.Error())
			os.Exit(1)
		}

		client := requester.NewClient().WithUnsecureWeb3Endpoint(endpoint)

		callData, err := smartcontract.ABIPack(abiBytes, method, callArgs...)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		resp, err := client.EthCall(contract, callData, height)
		if err != nil {
			fmt.Println("error on eth call", err.Error())
			os.Exit(1)
		}
		fmt.Println(string(resp))
		os.Exit(0)
	},
}

func init() {
	SolidityCmd.AddCommand(callContractViewCmd)
	callContractViewCmd.Flags().String("height", "latest", "Query at the given height.")
	callContractViewCmd.Flags().StringSliceVarP(&params, "params", "p", []string{}, "A list of params. If the param is an address, prefix with `a:0x123...`")
}
