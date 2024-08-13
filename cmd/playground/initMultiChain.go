package playground

import (
	"context"
	dbsql "database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hanchon/hanchond/playground/cosmosdaemon"
	"github.com/hanchon/hanchond/playground/evmos"
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

		latestChain, err := queries.GetLatestChain(context.Background())
		chainNumber := 1
		if err == nil {
			chainNumber = int(latestChain.ID) + 1
		} else if err != dbsql.ErrNoRows {
			fmt.Println("could not get the chains info from db")
			os.Exit(1)
		}

		nodes := make([]*cosmosdaemon.Daemon, amountOfValidators)
		switch strings.ToLower(strings.TrimSpace(client)) {
		case "evmos":
			chainID := fmt.Sprintf("evmos_9000-%d", chainNumber)
			for k := range nodes {
				if filesmanager.IsNodeHomeFolderInitialized(int64(chainNumber), int64(k)) {
					fmt.Printf("the home folder already exists: %d-%d\n", chainNumber, k)
					os.Exit(1)
				}
				path := filesmanager.GetNodeHomeFolder(int64(chainNumber), int64(k))
				nodes[k] = evmos.NewEvmos(
					fmt.Sprintf("moniker-%d-%d", chainNumber, k),
					version,
					path,
					chainID,
					fmt.Sprintf("validator-key-%d-%d", chainNumber, k),
					"aevmos",
				).Daemon
			}
		case "gaia":
			chainID := fmt.Sprintf("cosmoshub-%d", chainNumber)
			for k := range nodes {
				if filesmanager.IsNodeHomeFolderInitialized(int64(chainNumber), int64(k)) {
					fmt.Printf("the home folder already exists: %d-%d\n", chainNumber, k)
					os.Exit(1)
				}
				path := filesmanager.GetNodeHomeFolder(int64(chainNumber), int64(k))
				nodes[k] = gaia.NewGaia(
					fmt.Sprintf("moniker-%d-%d", chainNumber, k),
					path,
					chainID,
					fmt.Sprintf("validator-key-%d-%d", chainNumber, k),
					"icsstake",
				).Daemon
			}
		default:
			fmt.Println("invalid client")
			os.Exit(1)
		}

		chainID, err := cosmosdaemon.InitMultiNodeChain(nodes, queries)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			os.Exit(1)
		}

		fmt.Println("New chain created with id:", chainID)
	},
}

func init() {
	PlaygroundCmd.AddCommand(initMultiChainCmd)
	initMultiChainCmd.Flags().String("client", "evmos", "Client that you want to use. Options: evmos, gaia")
	initMultiChainCmd.Flags().StringP("version", "v", "local", "Version of the Evmos node that you want to use, defaults to local. Tag names are supported. If selected node is gaia, the flag is ignored.")
}
