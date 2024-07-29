package playground

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/hanchon/hanchond/playground/database"
	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/spf13/cobra"
)

// startNodeCmd represents the startNode command
var startNodeCmd = &cobra.Command{
	Use:   "start-node id",
	Args:  cobra.ExactArgs(1),
	Short: "Starts a node with the given ID",
	Long:  `It will run the node in a subprocess, saving the pid in the database in case it needs to be stoped in the future`,
	Run: func(cmd *cobra.Command, args []string) {
		home := filesmanager.SetHomeFolderFromCobraFlags(cmd)
		queries, err := initDB(home)
		if err != nil {
			fmt.Println("could not init database", err.Error())
			os.Exit(1)
		}

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

		e := evmos.NewEvmos(node.BinaryVersion, node.ConfigFolder, chain.ChainID, node.ValidatorKeyName)
		pID, err := e.Start(node.Moniker)
		if err != nil {
			fmt.Println("could not start the node:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Evmos is running with id:", pID)

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
