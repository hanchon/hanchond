package playground

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// getNodeCmd represents the getNode command
var getNodeCmd = &cobra.Command{
	Use:   "get-node [id]",
	Args:  cobra.ExactArgs(1),
	Short: "Get the node configuration",
	Long:  `It will return the node configuration that is stored in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)

		id := args[0]
		idNumber, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println("could not parse the ID:", err.Error())
			os.Exit(1)
		}

		node, err := queries.GetNode(context.Background(), idNumber)
		if err != nil {
			fmt.Println("could not get the node:", err.Error())
			os.Exit(1)
		}

		ports, err := queries.GetNodePorts(context.Background(), idNumber)
		if err != nil {
			fmt.Println("could not get the ports:", err.Error())
			os.Exit(1)
		}

		chain, err := queries.GetChain(context.Background(), node.ChainID)
		if err != nil {
			fmt.Println("could not get the chain:", err.Error())
			os.Exit(1)
		}

		fmt.Printf(`Node: %d
General Configuration:
    - Binary: %s
    - ChainID: %s
Process:
    - IsRunning: %d
    - ProcessID: %d
Keys:
    - KeyName: %s
    - Mnemonic: %s
Ports:
    - 8545(web3): %d
    - 26657(cli): %d
`,
			idNumber,
			node.BinaryVersion,
			chain.ChainID,
			node.IsRunning,
			node.ProcessID,
			node.ValidatorKeyName,
			node.ValidatorKey,
			ports.P8545,
			ports.P26657,
		)
	},
}

func init() {
	PlaygroundCmd.AddCommand(getNodeCmd)
}
