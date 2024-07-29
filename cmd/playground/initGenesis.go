package playground

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/hanchon/hanchond/playground/database"
	"github.com/hanchon/hanchond/playground/evmos"
	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/spf13/cobra"
)

// initGenesisCmd represents the initGenesis command
var initGenesisCmd = &cobra.Command{
	Use:   "init-genesis id",
	Args:  cobra.ExactArgs(1),
	Short: "Init the genesis file for a new chain",
	Long:  `Set up the data and config folder for the new chain`,
	Run: func(cmd *cobra.Command, args []string) {
		home := filesmanager.SetHomeFolderFromCobraFlags(cmd)
		queries, err := initDB(home)
		if err != nil {
			fmt.Println("could not init database", err.Error())
			os.Exit(1)
		}

		version, err := cmd.Flags().GetString("version")
		if err != nil {
			fmt.Println("version flag was not set")
			os.Exit(1)
		}

		chainid, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("invalid chain id, it must be integer", err.Error())
			os.Exit(1)
		}

		if !filesmanager.DoesEvmosdPathExist(version) {
			fmt.Println("the evmos version was not found in the built folder", version)
			os.Exit(1)

		}

		if filesmanager.IsNodeHomeFolderInitialized(chainid) {
			fmt.Println("the home folder for this node was already created")
			os.Exit(1)
		}

		path := filesmanager.GetNodeHomeFolder(chainid)
		chainID := fmt.Sprintf("evmos_9001-%d", chainid)

		e := evmos.NewEvmos(version, path, chainID, fmt.Sprintf("mykey%d", chainid))
		if err := e.InitGenesis(); err != nil {
			fmt.Println("could not init the genesis file", err.Error())
			os.Exit(1)
		}
		if err := e.SetPorts(); err != nil {
			fmt.Println("could not set the ports", err.Error())
			os.Exit(1)
		}

		row, err := queries.InsertChain(context.Background(), database.InsertChainParams{
			Name:          fmt.Sprintf("chain%d", chainid),
			ChainID:       e.ChainID,
			BinaryVersion: e.Version,
		})
		if err != nil {
			fmt.Println("could not insert chain. ", err.Error())
			os.Exit(1)
		}

		nodeID, err := queries.InsertNode(context.Background(), database.InsertNodeParams{
			ChainID:          row.ID,
			ConfigFolder:     path,
			Moniker:          e.Moniker,
			ValidatorKey:     e.ValMnemonic,
			ValidatorKeyName: e.ValKeyName,
			BinaryVersion:    e.Version,
			ProcessID:        0,
			IsValidator:      1,
			IsArchive:        0,
			IsRunning:        0,
		})
		if err != nil {
			fmt.Println("could not insert node", err.Error())
			os.Exit(1)
		}

		err = queries.InsertPorts(context.Background(), database.InsertPortsParams{
			NodeID: nodeID,
			P1317:  int64(e.Ports.P1317),
			P8080:  int64(e.Ports.P8080),
			P9090:  int64(e.Ports.P9090),
			P9091:  int64(e.Ports.P9091),
			P8545:  int64(e.Ports.P8545),
			P8546:  int64(e.Ports.P8546),
			P6065:  int64(e.Ports.P6065),
			P26658: int64(e.Ports.P26658),
			P26657: int64(e.Ports.P26657),
			P6060:  int64(e.Ports.P6060),
			P26656: int64(e.Ports.P26656),
			P26660: int64(e.Ports.P26660),
		})
		if err != nil {
			fmt.Println("could not insert ports", err.Error())
			os.Exit(1)
		}

		fmt.Println("Node added with id:", nodeID)
	},
}

func init() {
	PlaygroundCmd.AddCommand(initGenesisCmd)
	initGenesisCmd.Flags().StringP("version", "v", "local", "Version of the Evmos node that you want to use, defaults to local. Tag names are supported.")
}
