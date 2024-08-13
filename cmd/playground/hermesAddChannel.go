package playground

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hanchon/hanchond/playground/database"
	"github.com/hanchon/hanchond/playground/hermes"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// hermesAddChannelCmd represents the hermesAddChannel command
var hermesAddChannelCmd = &cobra.Command{
	Use:   "hermes-add-channel [chain_id] [chain_id]",
	Args:  cobra.ExactArgs(2),
	Short: "It uses the hermes client to open an IBC channel between two chains",
	Long:  `This command requires that Hermes was already built and at least one node for each chain running.`,
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)

		chainOne := args[0]
		chainOneID, err := strconv.Atoi(chainOne)
		if err != nil {
			fmt.Println("invalid chain id")
			os.Exit(1)
		}
		chainTwo := args[1]
		chainTwoID, err := strconv.Atoi(chainTwo)
		if err != nil {
			fmt.Println("invalid chain id")
			os.Exit(1)
		}
		chains := make([]database.GetAllChainNodesRow, 2)
		nodesChainOne, err := queries.GetAllChainNodes(context.Background(), int64(chainOneID))
		if err != nil {
			fmt.Println("could not find nodes for chain:", chainOne)
			os.Exit(1)
		}
		chains[0] = nodesChainOne[0]

		nodesChainTwo, err := queries.GetAllChainNodes(context.Background(), int64(chainTwoID))
		if err != nil {
			fmt.Println("could not find nodes for chain:", chainTwo)
			os.Exit(1)
		}
		chains[1] = nodesChainTwo[0]

		h := hermes.NewHermes()
		fmt.Println("Relayer initialized")

		for _, v := range chains {
			if v.IsRunning != 1 {
				fmt.Println("the node is not running, chain id:", v.ChainID)
			}

			switch {
			case strings.Contains(v.BinaryVersion, "gaia"):
				// d := gaia.NewGaia(v.Moniker, v.ConfigFolder, v.ChainID_2, v.ValidatorKeyName, v.Denom)
				// pID, err = d.Start(v.Moniker)
			case strings.Contains(v.BinaryVersion, "evmos"):
				fmt.Println("adding evmos key")
				if err := h.AddEvmosChain(
					v.ChainID_2,
					v.P26657,
					v.P9090,
					v.ValidatorKeyName,
					v.ValidatorKey,
				); err != nil {
					fmt.Println("error adding first chain to the relayer:", err.Error())
					os.Exit(1)
				}
				// d := evmos.NewEvmos(v.Moniker, v.BinaryVersion, v.ConfigFolder, v.ChainID_2, v.ValidatorKeyName, v.Denom)
				// pID, err = d.Start(v.Moniker)
			default:
				fmt.Println("incorrect binary name")
				os.Exit(1)
			}

		}

		// if err := h.AddEvmosChain(
		// 	secondNode.Chain.ChainID,
		// 	secondNode.Ports.P26657,
		// 	secondNode.Ports.P9090,
		// 	secondNode.Node.ValidatorKeyName,
		// 	secondNode.Node.ValidatorKey,
		// ); err != nil {
		// 	fmt.Println("error adding second chain to the relayer:", err.Error())
		// 	os.Exit(1)
		// }

		// fmt.Println("Second chain added")

		fmt.Println("Calling create channel")
		err = h.CreateChannel(chains[0].ChainID_2, chains[1].ChainID_2)
		if err != nil {
			fmt.Println("error creating channel", err.Error())
			os.Exit(1)
		}
		fmt.Println("Channel created")
	},
}

func init() {
	PlaygroundCmd.AddCommand(hermesAddChannelCmd)
}
