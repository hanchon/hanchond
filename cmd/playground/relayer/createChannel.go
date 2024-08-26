package relayer

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/playground/hermes"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// represents the createChannelCmd command
var createChannelCmd = &cobra.Command{
	Use:   "create-channel [chain_id] [chain_id]",
	Args:  cobra.ExactArgs(2),
	Short: "Create an IBC channel between two chains. The chains must be previously registered",
	Run: func(cmd *cobra.Command, args []string) {
		_ = sql.InitDBFromCmd(cmd)

		h := hermes.NewHermes()
		fmt.Println("Relayer initialized")

		chain1 := args[0]
		chain2 := args[1]

		fmt.Println("Calling create channel")
		err := h.CreateChannel(chain1, chain2)
		if err != nil {
			fmt.Println("error creating channel", err.Error())
			os.Exit(1)
		}
		fmt.Println("Channel created")
	},
}

func init() {
	RelayerCmd.AddCommand(createChannelCmd)
}
