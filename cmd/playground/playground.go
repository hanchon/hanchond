package playground

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// PlaygroundCmd represents the playground command
var PlaygroundCmd = &cobra.Command{
	Use:   "playground",
	Short: "Cosmos chain runner",
	Long:  `Tooling to set up your local cosmos network.`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := cmd.Flags().GetString("home")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("playground called", home)
	},
}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("could not find home folder:" + err.Error())
	}
	PlaygroundCmd.PersistentFlags().String("home", fmt.Sprintf("%s/.hanchond", home), "Home folder for the playground")
}
