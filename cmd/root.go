package cmd

import (
	"os"

	"github.com/hanchon/hanchond/cmd/convert"
	playground "github.com/hanchon/hanchond/cmd/playground"
	"github.com/hanchon/hanchond/cmd/repo"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hanchond",
	Short: "Hanchon's toolkit",
	Long:  `Stop re-writing the same scripts one million times`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(convert.ConvertCmd)
	rootCmd.AddCommand(playground.PlaygroundCmd)
	rootCmd.AddCommand(repo.RepoCmd)
}
