package tx

import (
	"fmt"
	"os"
	"strings"

	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// rateLimitProposalCmd represents the rateLimit-proposal command
var rateLimitProposalCmd = &cobra.Command{
	// TODO: add flags
	Use:   "rate-limit-proposal [denom]",
	Args:  cobra.ExactArgs(1),
	Short: "Create an rate-limit propsal",
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)
		nodeID, err := cmd.Flags().GetString("node")
		if err != nil {
			fmt.Println("node not set")
			os.Exit(1)
		}

		channel, err := cmd.Flags().GetString("channel")
		if err != nil {
			fmt.Println("channel not set")
			os.Exit(1)
		}

		duration, err := cmd.Flags().GetString("duration")
		if err != nil {
			fmt.Println("duration not set")
			os.Exit(1)
		}

		maxSend, err := cmd.Flags().GetString("max-send")
		if err != nil {
			fmt.Println("max-send not set")
			os.Exit(1)
		}

		maxRecv, err := cmd.Flags().GetString("max-recv")
		if err != nil {
			fmt.Println("max-recv not set")
			os.Exit(1)
		}

		e := evmos.NewEvmosFromDB(queries, nodeID)
		txhash, err := e.CreateRateLimitProposal(evmos.RateLimitParams{
			Channel:  channel,
			Denom:    strings.TrimSpace(args[0]),
			MaxSend:  maxSend,
			MaxRecv:  maxRecv,
			Duration: duration,
		})
		if err != nil {
			fmt.Println("error sending the transaction:", err.Error())
			os.Exit(1)
		}

		fmt.Println(txhash)
	},
}

func init() {
	TxCmd.AddCommand(rateLimitProposalCmd)
	rateLimitProposalCmd.Flags().StringP("channel", "c", "channel-0", "IBC channel")
	rateLimitProposalCmd.Flags().String("max-send", "10", "Max send rate")
	rateLimitProposalCmd.Flags().String("max-recv", "10", "Max recv rate")
	rateLimitProposalCmd.Flags().String("duration", "24", "Duration in hours")
}
