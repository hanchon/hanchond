package playground

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/cmd/playground/explorer"
	"github.com/hanchon/hanchond/cmd/playground/query"
	"github.com/hanchon/hanchond/cmd/playground/tx"
	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/spf13/cobra"
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("could not find home folder:" + err.Error())
	}
	PlaygroundCmd.PersistentFlags().String("home", fmt.Sprintf("%s/.hanchond", home), "Home folder for the playground")
}

// PlaygroundCmd represents the playground command
var PlaygroundCmd = &cobra.Command{
	Use:     "playground",
	Aliases: []string{"p"},
	Short:   "Cosmos chain runner",
	Long:    `Tooling to set up your local cosmos network.`,
	Run: func(cmd *cobra.Command, _ []string) {
		filesmanager.SetHomeFolderFromCobraFlags(cmd)
		_ = cmd.Help()
		os.Exit(0)
	},
}

func init() {
	PlaygroundCmd.AddCommand(tx.TxCmd)
	PlaygroundCmd.AddCommand(query.QueryCmd)
	PlaygroundCmd.AddCommand(explorer.ExplorerCmd)
}
