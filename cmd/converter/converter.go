package converter

import (
	"os"

	"github.com/spf13/cobra"
)

// ConverterCmd represents the converter command
var ConverterCmd = &cobra.Command{
	Use:   "converter",
	Short: "converter utils",
	Long:  `Convert wallets, coins and numbers`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			os.Exit(0)
		}
	},
}
