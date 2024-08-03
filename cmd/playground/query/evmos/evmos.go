package evmos

import (
	"os"

	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/spf13/cobra"
)

// EvmosCmd represents the evmos command
var EvmosCmd = &cobra.Command{
	Use:   "evmos",
	Short: "evmos unique queries ",
	Run: func(cmd *cobra.Command, _ []string) {
		filesmanager.SetHomeFolderFromCobraFlags(cmd)
		_ = cmd.Help()
		os.Exit(0)
	},
}
