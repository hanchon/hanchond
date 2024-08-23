package solidity

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/hanchon/hanchond/lib/smartcontract"
	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/hanchon/hanchond/playground/solidity"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// deployUniswapV2FactoryCmd represents the deploy command
var deployUniswapV2FactoryCmd = &cobra.Command{
	Use:   "deploy-uniswap-v2-factory [fee_wallet]",
	Args:  cobra.ExactArgs(1),
	Short: "Deploy uniswap v2 factory",
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

		feeWallet := args[0]

		// TODO: allow mainnet as a valid endpoint
		e := evmos.NewEvmosFromDB(queries, nodeID)
		builder := e.NewTxBuilder(uint64(gasLimit))

		contractName := "/UniswapV2Factory"
		// Clone uniswap-v2-core if needed
		path, err := solidity.DownloadUniswapV2Core()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		path = path + "/contracts" + contractName + ".sol"

		// Set up temp folder
		if err := filesmanager.CleanUpTempFolder(); err != nil {
			fmt.Println("could not clean up the temp folder:", err.Error())
			os.Exit(1)
		}

		folderName := "factoryBuilder"
		if err := filesmanager.CreateTempFolder(folderName); err != nil {
			fmt.Println("could not create the temp folder:", err.Error())
			os.Exit(1)
		}

		// Compile the contract
		err = solidity.CompileWithSolc("0.5.16", path, filesmanager.GetBranchFolder(folderName))
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

		// Generate the constructor
		abiBytes, err := filesmanager.ReadFile(filesmanager.GetBranchFolder(folderName) + contractName + ".abi")
		if err != nil {
			fmt.Printf("error reading the abi file:%s\n", err.Error())
			os.Exit(1)
		}

		// Get Params
		callArgs, err := smartcontract.StringsToABIArguments([]string{fmt.Sprintf("a:%s", feeWallet)})
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

		txHash, err := builder.DeployContract(0, bytecode, uint64(gasLimit))
		if err != nil {
			fmt.Printf("error sending the transaction: %s\n", err.Error())
			os.Exit(1)
		}

		contractAddress, err := e.NewRequester().GetContractAddress(txHash)
		if err != nil {
			fmt.Printf(err.Error())
			os.Exit(1)
		}

		codeHash, err := e.NewRequester().EthCodeHash(contractAddress, "latest")
		if err != nil {
			fmt.Println("failed to get the eth code:", err.Error())
			os.Exit(1)
		}

		fmt.Printf("{\"contract_address\":\"%s\", \"code_hash\":\"%s\", \"tx_hash\":\"%s\"}\n", contractAddress, "0x"+codeHash, txHash)

		// Clean up files
		if err := filesmanager.CleanUpTempFolder(); err != nil {
			fmt.Println("could not clean up the temp folder:", err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	},
}

func init() {
	SolidityCmd.AddCommand(deployUniswapV2FactoryCmd)
	deployUniswapV2FactoryCmd.Flags().Int("gas-limit", 20_000_000, "GasLimit to be used to deploy the transaction")
}
