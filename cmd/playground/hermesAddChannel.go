package playground

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/hermes"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// hermesAddChannelCmd represents the hermesAddChannel command
var hermesAddChannelCmd = &cobra.Command{
	Use:   "hermes-add-channel id1 id2",
	Args:  cobra.ExactArgs(2),
	Short: "It uses the hermes client to open an IBC channel between two chains",
	Long:  `This command requires that Hermes was already built and at least one node for each chain running.`,
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)

		node1 := args[0]
		node2 := args[1]
		fmt.Println("Getting first node data...")
		firstNode := evmos.GetNodeFromDB(queries, node1)
		fmt.Println("Getting second node data...")
		secondNode := evmos.GetNodeFromDB(queries, node2)
		// TODO: make sure that the nodes are running checking for the PID
		if firstNode.Node.IsRunning != 1 {
			fmt.Println("first node is not running")
			os.Exit(1)

		}
		if secondNode.Node.IsRunning != 1 {
			fmt.Println("second node is not running")
			os.Exit(1)

		}
		fmt.Println("Both chains are running")

		h := hermes.NewHermes()
		fmt.Println("Relayer initialized")

		if err := h.AddEvmosChain(
			firstNode.Chain.ChainID,
			firstNode.Ports.P26657,
			firstNode.Ports.P9090,
			firstNode.Node.ValidatorKeyName,
			firstNode.Node.ValidatorKey,
		); err != nil {
			fmt.Println("error adding first chain to the relayer:", err.Error())
			os.Exit(1)
		}
		fmt.Println("First chain added")

		if err := h.AddEvmosChain(
			secondNode.Chain.ChainID,
			secondNode.Ports.P26657,
			secondNode.Ports.P9090,
			secondNode.Node.ValidatorKeyName,
			secondNode.Node.ValidatorKey,
		); err != nil {
			fmt.Println("error adding second chain to the relayer:", err.Error())
			os.Exit(1)
		}

		fmt.Println("Second chain added")

		fmt.Println("Calling create channel")
		err := h.CreateChannel(firstNode.Chain.ChainID, secondNode.Chain.ChainID)
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
