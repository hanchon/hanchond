package evmos

import (
	"os"

	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/spf13/cobra"
)

// EvmosCmd represents the evmos command
var EvmosCmd = &cobra.Command{
	Use:   "evmos",
	Short: "evmos unique queries",
	Run: func(cmd *cobra.Command, _ []string) {
		filesmanager.SetHomeFolderFromCobraFlags(cmd)
		_ = cmd.Help()
		os.Exit(0)
	},
}

func init() {
	EvmosCmd.PersistentFlags().String("url", "", "Set the url path if using external provider")
	EvmosCmd.PersistentFlags().Bool("mainnet", false, "Set as true if the query for Evmos mainnet. This flag takes overwrite all the other provider related flags.")
}
