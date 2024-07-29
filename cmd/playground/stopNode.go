package playground

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/hanchon/hanchond/playground/database"
	"github.com/spf13/cobra"
)

// stopNodeCmd represents the stopNode command
var stopNodeCmd = &cobra.Command{
	Use:   "stop-node id",
	Args:  cobra.ExactArgs(1),
	Short: "Stops a running node with the given ID",
	Long:  `Stops the node using the PID stored in the database`,
	Run: func(cmd *cobra.Command, args []string) {
		queries := initDBFromCmd(cmd)

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

		if node.IsRunning != 1 {
			fmt.Println("the node is not running")
			os.Exit(1)
		}
		command := exec.Command( //nolint:gosec
			"kill",
			fmt.Sprintf("%d", node.ProcessID),
		)

		out, err := command.CombinedOutput()
		if strings.Contains(strings.ToLower(string(out)), "no such process") {
			fmt.Println("process is not running, updating the database..")
		} else if err != nil {
			fmt.Println("could not kill the process:", err.Error())
			os.Exit(1)
		}

		if err = queries.SetProcessID(context.Background(), database.SetProcessIDParams{
			ProcessID: 0,
			IsRunning: 0,
			ID:        idNumber,
		}); err != nil {
			fmt.Println("could not update the database:", err.Error())
			os.Exit(1)
		}

		fmt.Println("Node is no longer running")
	},
}

func init() {
	PlaygroundCmd.AddCommand(stopNodeCmd)
}
