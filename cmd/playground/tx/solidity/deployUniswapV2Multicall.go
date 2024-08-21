package solidity

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/hanchon/hanchond/playground/solidity"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// deployUniswapV2MulticallyCmd represents the deploy command
var deployUniswapV2MulticallyCmd = &cobra.Command{
	Use:   "deploy-uniswap-v2-multicall",
	Args:  cobra.ExactArgs(0),
	Short: "Deploy uniswap v2 multicall",
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

		// TODO: allow mainnet as a valid endpoint
		e := evmos.NewEvmosFromDB(queries, nodeID)
		builder := e.NewTxBuilder(uint64(gasLimit))

		contractName := "/Multicall"
		// Clone v2-minified if needed
		path, err := solidity.DownloadUniswapV2Minified()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// Keep working with the main contract
		path = path + "/contracts" + contractName + ".sol"

		// Set up temp folder
		if err := filesmanager.CleanUpTempFolder(); err != nil {
			fmt.Println("could not clean up the temp folder:", err.Error())
			os.Exit(1)
		}

		folderName := "multicallBuilder"
		if err := filesmanager.CreateTempFolder(folderName); err != nil {
			fmt.Println("could not create the temp folder:", err.Error())
			os.Exit(1)
		}

		// Compile the contract
		err = solidity.CompileWithSolc("0.5.0", path, filesmanager.GetBranchFolder(folderName))
		if err != nil {
			fmt.Println("could not compile the erc20 contract:", err.Error())
			os.Exit(1)
		}

		bytecode, err := filesmanager.ReadFile(filesmanager.GetBranchFolder(folderName) + contractName + ".bin")
		if err != nil {
			fmt.Printf("error reading the bytecode file:%s\n", err.Error())
			os.Exit(1)
		}

		bytecode, err = hex.DecodeString(string(bytecode))
		if err != nil {
			fmt.Println("error converting bytecode to []byte:", err.Error())
			os.Exit(1)
		}

		txHash, err := builder.DeployContract(0, bytecode, uint64(gasLimit))
		if err != nil {
			fmt.Printf("error sending the transaction: %s\n", err.Error())
			os.Exit(1)
		}

		receipt, err := e.NewRequester().GetTransactionReceiptWithRetry(txHash, 15)
		if err != nil {
			fmt.Printf("error getting the tx receipt:%s\n", err.Error())
		}

		trace, err := e.NewRequester().GetTransactionTrace(txHash)
		if err != nil {
			fmt.Printf("error getting the tx trace:%s\n", err.Error())
		}
		if trace.Result.Error != "" {
			fmt.Println("failed to execute the transaction:", trace.Result.Error)
			os.Exit(1)
		}

		codeHash, err := e.NewRequester().EthCodeHash(receipt.Result.ContractAddress, "latest")
		if err != nil {
			fmt.Println("failed to get the eth code:", err.Error())
			os.Exit(1)
		}

		fmt.Printf("{\"contract_address\":\"%s\", \"code_hash\":\"%s\", \"tx_hash\":\"%s\"}\n", receipt.Result.ContractAddress, "0x"+codeHash, txHash)

		// Clean up files
		if err := filesmanager.CleanUpTempFolder(); err != nil {
			fmt.Println("could not clean up the temp folder:", err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	},
}

func init() {
	SolidityCmd.AddCommand(deployUniswapV2MulticallyCmd)
	deployUniswapV2MulticallyCmd.Flags().Int("gas-limit", 20_000_000, "GasLimit to be used to deploy the transaction")
}
