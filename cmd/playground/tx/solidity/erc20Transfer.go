package solidity

import (
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hanchon/hanchond/lib/converter"
	"github.com/hanchon/hanchond/lib/txbuilder"
	"github.com/hanchon/hanchond/playground/cosmosdaemon"
	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/solidity"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// erc20TransferCmd represents the erc20Transfer command
var erc20TransferCmd = &cobra.Command{
	Use:   "erc20-transfer [contract] [wallet] [amount]",
	Args:  cobra.ExactArgs(3),
	Short: "Transfer erc20 coins from the validator wallet",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		queries := sql.InitDBFromCmd(cmd)

		contract := strings.TrimSpace(args[0])
		wallet := strings.TrimSpace(args[1])
		amount := strings.TrimSpace(args[2])

		wallet, err = converter.NormalizeAddressToHex(wallet)
		if err != nil {
			fmt.Println("invalid wallet")
			os.Exit(1)
		}

		nodeID, err := cmd.Flags().GetString("node")
		if err != nil {
			fmt.Println("node not set")
			os.Exit(1)
		}

		endpoint, err := cosmosdaemon.GetWeb3Endpoint(queries, cmd)
		if err != nil {
			fmt.Printf("error generting web3 endpoint: %s\n", err.Error())
			os.Exit(1)
		}

		e := evmos.NewEvmosFromDB(queries, nodeID)
		valWallet := txbuilder.NewSimpleWeb3WalletFromMnemonic(e.ValMnemonic, endpoint)

		callData, err := solidity.ERC20TransferCallData(wallet, amount)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		to := common.HexToAddress(contract)
		txhash, err := valWallet.TxBuilder.SendTx(valWallet.Address, &to, big.NewInt(0), 200_000, callData, valWallet.PrivKey)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("{\"txhash\":\"" + txhash + "\"}")
		os.Exit(0)
	},
}

func init() {
	SolidityCmd.AddCommand(erc20TransferCmd)
}
