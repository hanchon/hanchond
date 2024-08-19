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

// deployERC20Cmd represents the deploy command
var deployERC20Cmd = &cobra.Command{
	Use:   "deploy-erc20 [name] [symbol] [initial_amount]",
	Args:  cobra.ExactArgs(3),
	Short: "Deploy an erc20 contract",
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

		name := args[0]
		symbol := args[1]
		initialAmount := args[2]

		// TODO: allow mainnet as a valid endpoint
		e := evmos.NewEvmosFromDB(queries, nodeID)
		builder := e.NewTxBuilder(uint64(gasLimit))

		// Clone openzeppelin if needed
		path, err := solidity.DownloadDep("https://github.com/OpenZeppelin/openzeppelin-contracts", "v5.0.2", "openzeppelin")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// Set up temp folder
		if err := filesmanager.CleanUpTempFolder(); err != nil {
			fmt.Println("could not clean up the temp folder:", err.Error())
			os.Exit(1)
		}

		folderName := "erc20builder"
		if err := filesmanager.CreateTempFolder(folderName); err != nil {
			fmt.Println("could not create the temp folder:", err.Error())
			os.Exit(1)
		}

		// Create the contract
		contract := solidity.GenerateERC20Contract(path, name, symbol, initialAmount)
		contractPath := filesmanager.GetBranchFolder(folderName) + "/mycontract.sol"
		if err := filesmanager.SaveFile([]byte(contract), contractPath); err != nil {
			fmt.Println("could not save the contract file:", err.Error())
			os.Exit(1)
		}

		// Compile the contract
		err = solidity.CompileWithSolc("0.8.25", contractPath, filesmanager.GetBranchFolder(folderName))
		if err != nil {
			fmt.Println("could not compile the erc20 contract:", err.Error())
			os.Exit(1)
		}

		bytecode, err := filesmanager.ReadFile(filesmanager.GetBranchFolder(folderName) + "/" + solidity.StringToTitle(name) + ".bin")
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

		fmt.Printf("{\"contract_address\":\"%s\", \"tx_hash\":\"%s\"}\n", receipt.Result.ContractAddress, txHash)

		// Clean up files
		if err := filesmanager.CleanUpTempFolder(); err != nil {
			fmt.Println("could not clean up the temp folder:", err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	},
}

func init() {
	SolidityCmd.AddCommand(deployERC20Cmd)
	deployERC20Cmd.Flags().Int("gas-limit", 2_000_000, "GasLimit to be used to deploy the transaction")
}
