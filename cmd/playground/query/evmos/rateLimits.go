package evmos

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hanchon/hanchond/lib/requester"
	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// rateLimitsCmd represents the ibc-rate-limits command
var rateLimitsCmd = &cobra.Command{
	Use:   "ibc-rate-limits",
	Short: "Get all active IBC rate limits",
	Run: func(cmd *cobra.Command, _ []string) {
		queries := sql.InitDBFromCmd(cmd)
		nodeID, err := cmd.Flags().GetString("node")
		if err != nil {
			fmt.Println("node not set")
			os.Exit(1)
		}

		e := evmos.NewEvmosFromDB(queries, nodeID)
		client := requester.NewClient().WithUnsecureRestEndpoint(fmt.Sprintf("http://localhost:%d", e.Ports.P1317))
		rateLimits, err := client.GetIBCRateLimits()
		if err != nil {
			fmt.Println("could not get the rateLimits:", err.Error())
			os.Exit(1)
		}
		values, err := json.Marshal(rateLimits)
		if err != nil {
			fmt.Println("could not marshal response:", err.Error())
			os.Exit(1)
		}

		fmt.Println(string(values))
	},
}

func init() {
	EvmosCmd.AddCommand(rateLimitsCmd)
}
