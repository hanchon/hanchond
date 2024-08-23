package solidity

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/hanchon/hanchond/lib/smartcontract"
	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// deployContractCmd represents the deploy command
var deployContractCmd = &cobra.Command{
	Use:     "deploy-contract [path_to_bin_file]",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"d"},
	Short:   "Deploy a solidity contract",
	Long:    "The bytecode file must contain just the hex data",
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)
		nodeID, err := cmd.Flags().GetString("node")
		if err != nil {
			fmt.Println("node not set")
			os.Exit(1)
		}

		gasLimit, err := cmd.Flags().GetInt("gas-limit")
		if err != nil {
			fmt.Println("incorrect gas limit")
			os.Exit(1)
		}

		pathToBytecode := args[0]

		e := evmos.NewEvmosFromDB(queries, nodeID)
		builder := e.NewTxBuilder(uint64(gasLimit))

		bytecode, err := filesmanager.ReadFile(pathToBytecode)
		if err != nil {
			fmt.Printf("error reading the bytecode file:%s\n", err.Error())
			os.Exit(1)
		}

		bytecode, err = hex.DecodeString(string(bytecode))
		if err != nil {
			fmt.Println("error converting bytecode to []byte:", err.Error())
			os.Exit(1)
		}

		abiPath, err := cmd.Flags().GetString("abi")
		if err != nil {
			fmt.Println("could not read abi path:", err.Error())
			os.Exit(1)
		}

		if abiPath != "" {
			// It requires a constructor
			abiBytes, err := filesmanager.ReadFile(abiPath)
			if err != nil {
				fmt.Printf("error reading the abi file:%s\n", err.Error())
				os.Exit(1)
			}
			// Get Params
			callArgs, err := smartcontract.StringsToABIArguments(params)
			if err != nil {
				fmt.Printf("error converting arguments: %s\n", err.Error())
				os.Exit(1)
			}

			callData, err := smartcontract.ABIPackRaw(abiBytes, "", callArgs...)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			bytecode = append(bytecode, callData...)
		}

		txHash, err := builder.DeployContract(0, bytecode, uint64(gasLimit))
		if err != nil {
			fmt.Printf("error sending the transaction: %s\n", err.Error())
			os.Exit(1)
		}

		contractAddress, err := e.NewRequester().GetContractAddress(txHash)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Printf("{\"contract_address\":\"%s\", \"tx_hash\":\"%s\"}\n", contractAddress, txHash)
		os.Exit(0)
	},
}

func init() {
	SolidityCmd.AddCommand(deployContractCmd)
	deployContractCmd.Flags().Int("gas-limit", 2_000_000, "GasLimit to be used to deploy the transaction")
	deployContractCmd.Flags().String("abi", "", "ABI file if the contract has a contronstructor that needs params")
	deployContractCmd.Flags().StringSliceVarP(&params, "params", "p", []string{}, "A list of params. If the param is an address, prefix with `a:0x123...`")
}
