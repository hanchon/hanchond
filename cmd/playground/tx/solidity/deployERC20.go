package solidity

import (
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
	Use:   "deploy-erc20 [name] [symbol]",
	Args:  cobra.ExactArgs(2),
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

		initialAmount, err := cmd.Flags().GetString("initial-amount")
		if err != nil {
			fmt.Println("incorrect initial-amount")
			os.Exit(1)
		}

		isWrapped, err := cmd.Flags().GetBool("is-wrapped-coin")
		if err != nil {
			fmt.Println("incorrect wrapped flag")
			os.Exit(1)
		}

		name := args[0]
		symbol := args[1]

		// TODO: allow mainnet as a valid endpoint
		e := evmos.NewEvmosFromDB(queries, nodeID)
		builder := e.NewTxBuilder(uint64(gasLimit))

		txHash, err := solidity.BuildAndDeployERC20Contract(name, symbol, initialAmount, isWrapped, builder, gasLimit)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		contractAddress, err := e.NewRequester().GetContractAddress(txHash)
		if err != nil {
			fmt.Printf(err.Error())
			os.Exit(1)
		}

		fmt.Printf("{\"contract_address\":\"%s\", \"tx_hash\":\"%s\"}\n", contractAddress, txHash)

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
	deployERC20Cmd.Flags().String("initial-amount", "1000000", "Initial amout of coins sent to the deployer address")
	deployERC20Cmd.Flags().Bool("is-wrapped-coin", false, "Flag used to indenfity if the contract is representing the base denom. It uses WETH9 instead of OpenZeppelin contracts")
}
