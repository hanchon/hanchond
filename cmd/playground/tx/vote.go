package tx

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// vote represents the vote command
var voteCmd = &cobra.Command{
	Use:   "vote",
	Short: "Vote on all the active proposals",
	Run: func(cmd *cobra.Command, _ []string) {
		queries := sql.InitDBFromCmd(cmd)
		nodeID, err := cmd.Flags().GetString("node")
		if err != nil {
			fmt.Println("node not set")
			os.Exit(1)
		}

		option, err := cmd.Flags().GetString("option")
		if err != nil {
			fmt.Println("option not set")
			os.Exit(1)
		}

		e := evmos.NewEvmosFromDB(queries, nodeID)
		txhashes, err := e.VoteOnAllTheProposals(option)
		if err != nil {
			fmt.Println("error sending the transaction:", err.Error())
			os.Exit(1)
		}
		for _, v := range txhashes {
			fmt.Println(v)
		}
	},
}

func init() {
	TxCmd.AddCommand(voteCmd)
	voteCmd.Flags().StringP("option", "o", "yes", "Vote option")
}
