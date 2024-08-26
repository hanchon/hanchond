package relayer

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/playground/hermes"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// represents the addChainConfigCmd command
var addChainConfigCmd = &cobra.Command{
	Use:   "add-chain-config",
	Args:  cobra.ExactArgs(0),
	Short: "Add chain config to hermes, it is ignored if the chain id already exists",
	Run: func(cmd *cobra.Command, _ []string) {
		_ = sql.InitDBFromCmd(cmd)

		h := hermes.NewHermes()
		fmt.Println("Relayer initialized")

		chainID, err := cmd.Flags().GetString("chainid")
		if err != nil || chainID == "" {
			fmt.Println("missing chainid value")
			os.Exit(1)
		}

		p26657, err := cmd.Flags().GetString("p26657")
		if err != nil || chainID == "" {
			fmt.Println("missing p26657 value")
			os.Exit(1)
		}

		p9090, err := cmd.Flags().GetString("p9090")
		if err != nil || chainID == "" {
			fmt.Println("missing p9090 value")
			os.Exit(1)
		}

		keyname, err := cmd.Flags().GetString("keyname")
		if err != nil || chainID == "" {
			fmt.Println("missing keyname value")
			os.Exit(1)
		}

		keymnemonic, err := cmd.Flags().GetString("keymnemonic")
		if err != nil || chainID == "" {
			fmt.Println("missing keymnemonic value")
			os.Exit(1)
		}

		prefix, err := cmd.Flags().GetString("prefix")
		if err != nil || chainID == "" {
			fmt.Println("missing prefix value")
			os.Exit(1)
		}

		denom, err := cmd.Flags().GetString("denom")
		if err != nil || chainID == "" {
			fmt.Println("missing denom value")
			os.Exit(1)
		}

		isEvm, err := cmd.Flags().GetBool("is-evm")
		if err != nil || chainID == "" {
			fmt.Println("missing is-evm value")
			os.Exit(1)
		}

		switch isEvm {
		case false:
			fmt.Println("Adding a NOT EVM chain")
			if err := h.AddCosmosChain(
				chainID,
				p26657,
				p9090,
				keyname,
				keymnemonic,
				prefix,
				denom,
			); err != nil {
				fmt.Println("error adding first chain to the relayer:", err.Error())
				os.Exit(1)
			}
		case true:
			fmt.Println("Adding a EVM chain")
			if err := h.AddEvmosChain(
				chainID,
				p26657,
				p9090,
				keyname,
				keymnemonic,
				prefix,
				denom,
			); err != nil {
				fmt.Println("error adding first chain to the relayer:", err.Error())
				os.Exit(1)
			}
		}
	},
}

func init() {
	RelayerCmd.AddCommand(addChainConfigCmd)
	addChainConfigCmd.Flags().String("chainid", "", "Chain-id, i.e., evmos_9001-2")
	addChainConfigCmd.Flags().String("p26657", "", "Endpoint where the port 26657 is exposed, i.e., http://localhost:26657")
	addChainConfigCmd.Flags().String("p9090", "", "Endpoint where the port 9090 is exposed, i.e., http://localhost:9090")
	addChainConfigCmd.Flags().String("keyname", "", "Key name, it's used to identify the files inside hermes, i.e., relayerkey")
	addChainConfigCmd.Flags().String("keymnemonic", "", "Key mnemonic, mnemonic for the wallet")
	addChainConfigCmd.Flags().String("prefix", "", "Prefix for the bech32 address, i.e, osmo")
	addChainConfigCmd.Flags().String("denom", "", "Denom of the base token, i.e, aevmos")
	addChainConfigCmd.Flags().Bool("is-evm", false, "If the chain is evm compatible, this is used to determinate the type of wallet.")
}
