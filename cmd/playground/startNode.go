package playground

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hanchon/hanchond/playground/database"
	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/gaia"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// startNodeCmd represents the startNode command
var startNodeCmd = &cobra.Command{
	Use:   "start-node [node_id]",
	Args:  cobra.ExactArgs(1),
	Short: "Starts a node with the given ID",
	Long:  `It will run the node in a subprocess, saving the pid in the database in case it needs to be stoped in the future`,
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

		chain, err := queries.GetChain(context.Background(), node.ChainID)
		if err != nil {
			fmt.Println("could not get the chain:", err.Error())
			os.Exit(1)
		}

		var pID int
		switch {
		case strings.Contains(node.BinaryVersion, "evmos"):
			d := evmos.NewEvmos(
				node.Moniker,
				node.BinaryVersion,
				node.ConfigFolder,
				chain.ChainID,
				node.ValidatorKeyName,
				chain.Denom,
			)
			pID, err = d.Start()
		case strings.Contains(node.BinaryVersion, "gaia"):
			d := gaia.NewGaia(
				node.Moniker,
				node.ConfigFolder,
				chain.ChainID,
				node.ValidatorKeyName,
				node.ValidatorKeyName,
			)
			pID, err = d.Start()
		}
		if err != nil {
			fmt.Println("could not start the node:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Node is running with pID:", pID)

		err = queries.SetProcessID(context.Background(), database.SetProcessIDParams{
			ProcessID: int64(pID),
			IsRunning: 1,
			ID:        node.ID,
		})
		if err != nil {
			fmt.Println("could not save the process ID to the db:", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	PlaygroundCmd.AddCommand(startNodeCmd)
}
