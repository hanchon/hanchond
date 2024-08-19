package evmos

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// ethCodeCmd represents the ethCode command
var ethCodeCmd = &cobra.Command{
	Use:   "eth-code [address]",
	Args:  cobra.ExactArgs(1),
	Short: "Get the smartcontract registered eth code",
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)
		nodeID, err := cmd.Flags().GetString("node")
		if err != nil {
			fmt.Println("node not set")
			os.Exit(1)
		}
		height, err := cmd.Flags().GetString("height")
		if err != nil {
			fmt.Println("error getting the request height:", err.Error())
			os.Exit(1)
		}

		e := evmos.NewEvmosFromDB(queries, nodeID)
		client := e.NewRequester()

		code, err := client.EthCode(args[0], height)
		if err != nil {
			fmt.Println("could not get the ethCode:", err.Error())
			os.Exit(1)
		}

		fmt.Println(string(code))
	},
}

func init() {
	EvmosCmd.AddCommand(ethCodeCmd)
	ethCodeCmd.Flags().String("height", "latest", "Query at the given height.")
}
