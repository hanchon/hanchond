package solidity

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// deployContractCmd represents the deploy command
var deployContractCmd = &cobra.Command{
	Use:     "deploy-contract [path_to_bytecode]",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"d"},
	Short:   "Deploy a solidity contract",
	Long:    "The bytecode file must have the following format: {\"bytecode\":\"60806...\",...}",
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

		type Bytecode struct {
			Bytecode string `json:"bytecode"`
		}
		var b Bytecode
		if err := json.Unmarshal(bytecode, &b); err != nil {
			fmt.Printf("error unmarshing the json file: %s\n", err.Error())
			os.Exit(1)
		}

		bytes, err := hex.DecodeString(b.Bytecode)
		if err != nil {
			fmt.Printf("error decoding the bytecode: %s\n", err.Error())
			os.Exit(1)
		}

		txHash, err := builder.DeployContract(0, bytes, uint64(gasLimit))
		if err != nil {
			fmt.Printf("error sending the transaction: %s\n", err.Error())
			os.Exit(1)
		}

		receipt, err := e.NewRequester().GetTransactionReceiptWithRetry(txHash, 15)
		if err != nil {
			fmt.Printf("error getting the tx receipt:%s\n", err.Error())
		}

		fmt.Printf("{\"contract_address\":\"%s\", \"tx_hash\":\"%s\"}\n", receipt.Result.ContractAddress, txHash)
		os.Exit(0)
	},
}

func init() {
	SolidityCmd.AddCommand(deployContractCmd)
	deployContractCmd.Flags().Int("gas-limit", 2_000_000, "GasLimit to be used to deploy the transaction")
}
