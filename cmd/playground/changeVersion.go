package playground

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hanchon/hanchond/playground/database"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// changeVersionCmd represents the changeVersion command
var changeVersionCmd = &cobra.Command{
	Use:   "change-version [id] [version]",
	Args:  cobra.ExactArgs(2),
	Short: "Change the binary version of the given node",
	Long:  `It will update the database entry for the node, you need to manually stop and re-start the node for it to take effect on the running chain.`,
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)
		isChainID, _ := cmd.Flags().GetBool("is-chain-id")
		id := args[0]
		idNumber, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println("could not parse the ID:", err.Error())
			os.Exit(1)
		}
		binaryVersion := strings.TrimSpace(args[1])

		if isChainID {
			// Update the chain and all its nodes
			if err := queries.SetChainBinaryVersion(context.Background(), database.SetChainBinaryVersionParams{
				BinaryVersion: binaryVersion,
				ID:            idNumber,
			}); err != nil {
				fmt.Println("could not update the chain binary version:", err.Error())
				os.Exit(1)
			}

			nodes, err := queries.GetAllNodesForChainID(context.Background(), idNumber)
			if err != nil {
				fmt.Println("could not get chain nodes:", err.Error())
				os.Exit(1)
			}

			for _, v := range nodes {
				updateNodeVersion(queries, v.ID, binaryVersion)
			}
		} else {
			// Update just the node
			updateNodeVersion(queries, idNumber, binaryVersion)
		}
	},
}

func init() {
	PlaygroundCmd.AddCommand(changeVersionCmd)
	changeVersionCmd.Flags().Bool("is-chain-id", false, "If the flag is yes, it will assume that the ID is the chain ID. If it is set as false, the ID will be used just for the node.")
}

func updateNodeVersion(queries *database.Queries, nodeID int64, binaryVersion string) {
	_, err := queries.GetNode(context.Background(), nodeID)
	if err != nil {
		fmt.Println("could not get the node:", err.Error())
		os.Exit(1)
	}

	err = queries.SetBinaryVersion(context.Background(), database.SetBinaryVersionParams{
		BinaryVersion: binaryVersion,
		ID:            nodeID,
	})
	if err != nil {
		fmt.Println("could not update the binary version:", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Node %d updated to version %s\n", nodeID, binaryVersion)
}
