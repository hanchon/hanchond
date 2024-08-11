package playground

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hanchon/hanchond/playground/cosmosdaemon"
	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/hanchon/hanchond/playground/gaia"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// initMultiChainCmd represents the initMultiChainCmd
var initMultiChainCmd = &cobra.Command{
	Use:   "init-multi-chain [amount_of_validators]",
	Args:  cobra.ExactArgs(1),
	Short: "Init the genesis and configurations files for a new chain",
	Long:  `Set up the validators nodes for the new chain.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := cmd.Flags().GetString("client")
		if err != nil {
			fmt.Println("client flag was not set")
			os.Exit(1)
		}
		version, err := cmd.Flags().GetString("version")
		if err != nil {
			fmt.Println("version flag was not set")
			os.Exit(1)
		}

		amountOfValidators, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("invalid amount of validators")
			os.Exit(1)
		}

		queries := sql.InitDBFromCmd(cmd)

		// GetNextChainID
		chainNumber := 1

		nodes := make([]*cosmosdaemon.Daemon, amountOfValidators)
		switch strings.ToLower(strings.TrimSpace(client)) {
		case "evmos":
			fmt.Println("using", version)
		case "gaia":
			chainID := fmt.Sprintf("cosmoshub-%d", chainNumber)
			for k := range nodes {
				path := filesmanager.GetNodeHomeFolder(int64(chainNumber) + int64(k))
				nodes[k] = gaia.NewGaia(
					fmt.Sprintf("moniker-%d-%d", chainNumber, k),
					path,
					chainID,
					"validator-key",
					"icsstake",
				).Daemon

			}

			fmt.Println(nodes)
			if err := cosmosdaemon.InitMultiNodeChain(nodes, queries); err != nil {
				fmt.Printf("error: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Println(nodes)
		default:
			fmt.Println("invalid client")
			os.Exit(1)
		}
	},
}

func init() {
	PlaygroundCmd.AddCommand(initMultiChainCmd)
	initMultiChainCmd.Flags().String("client", "gaia", "Client that you want to use. Options: evmos, gaia")
	initMultiChainCmd.Flags().StringP("version", "v", "local", "Version of the Evmos node that you want to use, defaults to local. Tag names are supported. If selected node is gaia, the flag is ignored.")
}
