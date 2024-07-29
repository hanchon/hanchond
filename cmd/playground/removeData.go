package playground

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/spf13/cobra"
)

// removeDataCmd represents the removeData command
var removeDataCmd = &cobra.Command{
	Use:   "remove-data",
	Short: "Removes the data folder, deleting the configuration and data for all the networks and relayers.",
	Long:  `It is a command useful when restarting the process from scratch, it will delete all the data keeping just the built binaries.`,
	Run: func(cmd *cobra.Command, _ []string) {
		_ = filesmanager.SetHomeFolderFromCobraFlags(cmd)
		if err := filesmanager.CleanUpData(); err != nil {
			fmt.Println("failed to remove the data:", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	PlaygroundCmd.AddCommand(removeDataCmd)
}
