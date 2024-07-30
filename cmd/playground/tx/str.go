package tx

import (
	"fmt"
	"os"
	"strings"

	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// strV1ProposalCmd represents the str-v1-proposal command
var strV1ProposalCmd = &cobra.Command{
	Use:     "str-v1-proposal [denom]",
	Aliases: []string{"strv1"},
	Args:    cobra.ExactArgs(1),
	Short:   "Creates a STRv1 proposal",
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)
		nodeID, err := cmd.Flags().GetString("node")
		if err != nil {
			fmt.Println("node not set")
			os.Exit(1)
		}

		denom := args[0]

		ibcTransferCmd.Flags().IntP("exponent", "e", 18, "Exponents of the token")
		ibcTransferCmd.Flags().StringP("alias", "a", "tokenalias", "Token alias")
		ibcTransferCmd.Flags().StringP("name", "n", "tokenname", "Token name")
		ibcTransferCmd.Flags().StringP("symbol", "s", "tokensymbol", "Token symbol")

		exponent, err := cmd.Flags().GetInt("exponent")
		if err != nil {
			fmt.Println("exponent not set")
			os.Exit(1)
		}

		alias, err := cmd.Flags().GetString("alias")
		if err != nil {
			fmt.Println("alias not set")
			os.Exit(1)
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println("name not set")
			os.Exit(1)
		}

		symbol, err := cmd.Flags().GetString("symbol")
		if err != nil {
			fmt.Println("symbol not set")
			os.Exit(1)
		}

		e := evmos.NewEvmosFromDB(queries, nodeID)
		out, err := e.CreateSTRv1Proposal(evmos.STRv1{
			Denom:    denom,
			Exponent: exponent,
			Alias:    alias,
			Name:     name,
			Symbol:   symbol,
		})
		if err != nil {
			fmt.Println("error sending the transaction:", err.Error())
			os.Exit(1)
		}

		if !strings.Contains(out, "code: 0") {
			fmt.Println("transaction failed!")
			fmt.Println(out)
			os.Exit(1)
		}
		hash := strings.Split(out, "txhash: ")
		if len(hash) > 1 {
			hash[1] = strings.TrimSpace(hash[1])
		}
		fmt.Println(hash[1])
	},
}

func init() {
	TxCmd.AddCommand(strV1ProposalCmd)
	strV1ProposalCmd.Flags().IntP("exponent", "e", 18, "Exponents of the token")
	strV1ProposalCmd.Flags().StringP("alias", "a", "tokenalias", "Token alias")
	strV1ProposalCmd.Flags().String("name", "tokenname", "Token name")
	strV1ProposalCmd.Flags().StringP("symbol", "s", "tokensymbol", "Token symbol")
}
