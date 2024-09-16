package playground

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/hanchon/hanchond/playground/database"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// stopChainCmd represents the stopChain command
var stopChainCmd = &cobra.Command{
	Use:   "stop-chain [chain_id]",
	Args:  cobra.ExactArgs(1),
	Short: "Stops all the running validators for the given Chain ID",
	Long:  `Stops all the nodes using the PID stored in the database`,
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)

		chainNumber, err := strconv.Atoi(strings.TrimSpace(args[0]))
		if err != nil {
			fmt.Println("invalid chain id:", err.Error())
			os.Exit(1)
		}
		nodes, err := queries.GetAllNodesForChainID(context.Background(), int64(chainNumber))
		if err != nil {
			fmt.Println("could not find the chain:", err.Error())
			os.Exit(1)
		}

		for _, v := range nodes {
			if v.IsRunning != 1 {
				fmt.Printf("The node %d is not running\n", v.ID)
				continue
			}

			command := exec.Command( //nolint:gosec
				"kill",
				fmt.Sprintf("%d", v.ProcessID),
			)
			out, err := command.CombinedOutput()

			if strings.Contains(strings.ToLower(string(out)), "no such process") {
				fmt.Printf("Process is not running for node %d, updating the database..\n", v.ID)
			} else if err != nil {
				fmt.Println("could not kill the process:", err.Error())
				os.Exit(1)
			}

			if err = queries.SetProcessID(context.Background(), database.SetProcessIDParams{
				ProcessID: 0,
				IsRunning: 0,
				ID:        v.ID,
			}); err != nil {
				fmt.Println("could not update the database:", err.Error())
				os.Exit(1)
			}
		}

		fmt.Println("Chain is stopped")
	},
}

func init() {
	PlaygroundCmd.AddCommand(stopChainCmd)
}
