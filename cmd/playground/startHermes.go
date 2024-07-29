package playground

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/hanchon/hanchond/playground/database"
	"github.com/hanchon/hanchond/playground/hermes"
	"github.com/spf13/cobra"
)

// startHermesCmd represents the start-hermes command
var startHermesCmd = &cobra.Command{
	Use:   "start-hermes",
	Short: "Starts the relayer",
	Long:  `The command assumes that the relayer was already built and that there is a channel enabled between 2 chains`,
	Run: func(cmd *cobra.Command, _ []string) {
		queries := initDBFromCmd(cmd)
		relayer, err := queries.GetRelayer(context.Background())
		if err == sql.ErrNoRows {
			if err := queries.InitRelayer(context.Background()); err != nil {
				fmt.Println("could not init the relayer's database:", err.Error())
				os.Exit(1)
			}
		}

		// TODO: check if the process is running checking the PID
		if relayer.IsRunning == 1 {
			fmt.Println("the relayer is already running")
			os.Exit(1)
		}

		pid, err := hermes.NewHermes().Start()
		if err != nil {
			fmt.Println("could not start the relayer:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Hermes running with PID:", pid)

		if err := queries.UpdateRelayer(context.Background(), database.UpdateRelayerParams{
			ProcessID: int64(pid),
			IsRunning: 1,
		}); err != nil {
			fmt.Println("could not update the relayer database", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	PlaygroundCmd.AddCommand(startHermesCmd)
}
