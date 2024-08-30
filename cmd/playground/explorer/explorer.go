package explorer

import (
	"os"

	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/spf13/cobra"
)

// ExplorerCmd represents the explorer command
var ExplorerCmd = &cobra.Command{
	Use:     "explorer",
	Aliases: []string{"e"},
	Short:   "Explorer for the node",
	Run: func(cmd *cobra.Command, _ []string) {
		filesmanager.SetHomeFolderFromCobraFlags(cmd)
		_ = cmd.Help()
		os.Exit(0)
	},
}

func init() {
	ExplorerCmd.PersistentFlags().StringP("node", "n", "1", "Playground node used to get the information")
}
