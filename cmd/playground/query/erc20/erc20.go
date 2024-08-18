package erc20

import (
	"os"

	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/spf13/cobra"
)

// ERC20Cmd represents the query command
var ERC20Cmd = &cobra.Command{
	Use:   "erc20",
	Short: "ERC20 erc20 related data",
	Run: func(cmd *cobra.Command, _ []string) {
		filesmanager.SetHomeFolderFromCobraFlags(cmd)
		_ = cmd.Help()
		os.Exit(0)
	},
}

func init() {
	ERC20Cmd.PersistentFlags().String("url", "", "Set the url path if using external provider")
	ERC20Cmd.PersistentFlags().Bool("mainnet", false, "Set as true if the query for Evmos mainnet. This flag takes overwrite all the other provider related flags.")
}
