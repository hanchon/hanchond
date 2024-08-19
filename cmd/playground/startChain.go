package playground

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hanchon/hanchond/playground/database"
	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/gaia"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// startChainCmd represents the startChainCmd
var startChainCmd = &cobra.Command{
	Use:   "start-chain [chain_id]",
	Args:  cobra.ExactArgs(1),
	Short: "Start all the validators of the chain",
	Long:  `Start all the required processes to run the chain`,
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)

		chainNumber, err := strconv.Atoi(strings.TrimSpace(args[0]))
		if err != nil {
			fmt.Println("invalid chain id:", err.Error())
			os.Exit(1)
		}
		nodes, err := queries.GetAllChainNodes(context.Background(), int64(chainNumber))
		if err != nil {
			fmt.Println("could not find the chain:", err.Error())
			os.Exit(1)
		}

		for _, v := range nodes {
			version := strings.ToLower(strings.TrimSpace(v.BinaryVersion))
			var pID int
			var err error
			switch {
			case strings.Contains(version, "gaia"):
				d := gaia.NewGaia(v.Moniker, v.ConfigFolder, v.ChainID_2, v.ValidatorKeyName, v.Denom)
				pID, err = d.Start()
			case strings.Contains(version, "evmos"):
				d := evmos.NewEvmos(v.Moniker, v.BinaryVersion, v.ConfigFolder, v.ChainID_2, v.ValidatorKeyName, v.Denom)
				pID, err = d.Start()
			default:
				fmt.Println("incorrect binary name")
				os.Exit(1)
			}

			if err != nil {
				fmt.Println("could not start the node:", err.Error())
				os.Exit(1)
			}

			fmt.Println("Node is running with pID:", pID)
			err = queries.SetProcessID(context.Background(), database.SetProcessIDParams{
				ProcessID: int64(pID),
				IsRunning: 1,
				ID:        v.ID,
			})
			if err != nil {
				fmt.Println("could not save the process ID to the db:", err.Error())
				os.Exit(1)
			}
		}
	},
}

func init() {
	PlaygroundCmd.AddCommand(startChainCmd)
}
