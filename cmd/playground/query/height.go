package query

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/lib/requester"
	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// heightCmd represents the query height command
var heightCmd = &cobra.Command{
	Use:   "height",
	Short: "Get the current networkt height",
	Run: func(cmd *cobra.Command, _ []string) {
		queries := sql.InitDBFromCmd(cmd)
		nodeID, err := cmd.Flags().GetString("node")
		if err != nil {
			fmt.Println("node not set")
			os.Exit(1)
		}

		e := evmos.NewEvmosFromDB(queries, nodeID)
		client := requester.NewClient().WithUnsecureTendermintEndpoint(fmt.Sprintf("http://localhost:%d", e.Ports.P26657))
		height, err := client.GetCurrentHeight()
		if err != nil {
			fmt.Println("could not query the current height:", err.Error())
			os.Exit(1)
		}
		fmt.Println(height)
	},
}

func init() {
	QueryCmd.AddCommand(heightCmd)
}
