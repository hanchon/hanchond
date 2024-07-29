package playground

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/hanchon/hanchond/playground/database"
	"github.com/spf13/cobra"
)

// stopHermesCmd represents the stop-hermes command
var stopHermesCmd = &cobra.Command{
	Use:   "stop-hermes",
	Short: "Stop the relayer",
	Long:  `It gets the PID from the database and send the kill signal to the process`,
	Run: func(cmd *cobra.Command, _ []string) {
		queries := initDBFromCmd(cmd)
		relayer, err := queries.GetRelayer(context.Background())
		if err != nil {
			fmt.Println("the relayer is not in the database", err.Error())
			os.Exit(1)
		}

		// TODO: check if the process is running checking the PID
		if relayer.IsRunning != 1 {
			fmt.Println("the relayer is not running")
			os.Exit(1)
		}

		command := exec.Command( //nolint:gosec
			"kill",
			fmt.Sprintf("%d", relayer.ProcessID),
		)

		out, err := command.CombinedOutput()
		if strings.Contains(strings.ToLower(string(out)), "no such process") {
			fmt.Println("the relayer is not running, updating the database..")
		} else if err != nil {
			fmt.Println("could not kill the process:", err.Error())
			os.Exit(1)
		}

		if err := queries.UpdateRelayer(context.Background(), database.UpdateRelayerParams{
			ProcessID: 0,
			IsRunning: 0,
		}); err != nil {
			fmt.Println("could not update the relayer database", err.Error())
			os.Exit(1)
		}

		fmt.Println("Relayer is no longer running")
	},
}

func init() {
	PlaygroundCmd.AddCommand(stopHermesCmd)
}
