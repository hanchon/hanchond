package playground

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/hanchon/hanchond/playground/database"
	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/hanchon/hanchond/playground/gaia"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// initGaiaCmd represents the initGaia command
var initGaiaCmd = &cobra.Command{
	Use:   "init-gaia id",
	Args:  cobra.ExactArgs(1),
	Short: "Init the genesis file for a new chain",
	Long:  `Set up the data and config folder for the new chain`,
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)
		_ = queries

		chainid, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("invalid chain id, it must be integer", err.Error())
			os.Exit(1)
		}

		if filesmanager.IsNodeHomeFolderInitialized(chainid) {
			fmt.Println("the home folder for this node was already created")
			os.Exit(1)
		}

		chainID := "cosmoshub-99"

		nodes := make([]*gaia.Gaia, 2)

		var chainDB database.Chain

		for k := range nodes {
			path := filesmanager.GetNodeHomeFolder(chainid + int64(k))
			g := gaia.NewGaia(fmt.Sprintf("moniker-%d-%d", chainid, k), path, chainID, "validator-key", "icsstake")
			// Init the config files
			if err := g.InitNode(); err != nil {
				panic(err)
			}
			// Update general parameters in the genesis file
			if err := g.UpdateGenesisFile(); err != nil {
				panic(err)
			}
			if err := g.UpdateConfigFile(false); err != nil {
				panic(err)
			}
			if err := g.UpdateAppFile(); err != nil {
				panic(err)
			}
			if err := g.CreateGenTx(); err != nil {
				panic(err)
			}
			// Assign random and unique ports
			if err := g.AssignPorts(queries); err != nil {
				panic(err)
			}
			// Update the Config Files
			if err := g.UpdateConfigPorts(); err != nil {
				panic(err)
			}

			nodes[k] = g
			if k == 0 {
				chainDB, err = g.SaveChainToDB(queries)
				if err != nil {
					panic(err)
				}
			}
			_, err := g.SaveNodeToDB(chainDB, queries)
			if err != nil {
				panic(err)
			}
		}

		// Join genesis transactions
		for k, v := range nodes {
			if k == 0 {
				continue
			}
			files, err := filepath.Glob(v.HomeDir + "/config/gentx/*.json")
			if err != nil {
				panic("no files: " + err.Error())
			}
			if len(files) == 0 {
				panic("no files 2: " + err.Error())
			}

			if err := filesmanager.CopyFile(
				files[0],
				nodes[0].HomeDir+"/config/gentx",
			); err != nil {
				panic(err)
			}
			addr, err := v.GetValidatorAddress()
			if err != nil {
				panic(err)
			}
			if err := nodes[0].AddGenesisAccount(addr); err != nil {
				panic(err)
			}
		}

		if err := nodes[0].CollectGenTxs(); err != nil {
			panic(err)
		}
		if err := nodes[0].ValidateGenesis(); err != nil {
			panic(err)
		}

		peers := []string{}
		for k := range nodes {
			peerInfo, err := nodes[k].GetPeerInfo()
			if err != nil {
				panic(err)
			}
			peers = append(peers, peerInfo)

			if k == 0 {
				continue
			}
			if err := filesmanager.CopyFile(
				nodes[0].HomeDir+"/config/genesis.json",
				nodes[k].HomeDir+"/config/genesis.json",
			); err != nil {
				panic(err)
			}
		}

		for k := range nodes {
			if err := nodes[k].AddPersistenPeers(peers); err != nil {
				panic(err)
			}
		}

		// {
		// 	// TODO: create all the other gaia instances, set the ports and genesis transactions
		// 	// Copy genTxns to node 0 folder
		// 	// for each validator:
		// 	// if err := d.AddGenesisAccount(validatorAddr); err != nil {
		// 	// 	return err
		// 	// }
		//
		// 	if err := g.CollectGenTxs(); err != nil {
		// 		panic(err)
		// 	}
		// 	if err := g.ValidateGenesis(); err != nil {
		// 		panic(err)
		// 	}
		//
		// 	// Copy the genesis file to all the other validators
		// 	// Update the configuration files with the peers
		// }

		// 	e := evmos.NewEvmos(version, path, chainID, fmt.Sprintf("mykey%d", chainid))
		// 	if err := e.InitGaia(); err != nil {
		// 		fmt.Println("could not init the genesis file", err.Error())
		// 		os.Exit(1)
		// 	}
		// 	if err := e.SetPorts(); err != nil {
		// 		fmt.Println("could not set the ports", err.Error())
		// 		os.Exit(1)
		// 	}
		//
		// 	row, err := queries.InsertChain(context.Background(), database.InsertChainParams{
		// 		Name:          fmt.Sprintf("chain%d", chainid),
		// 		ChainID:       e.ChainID,
		// 		BinaryVersion: e.Version,
		// 	})
		// 	if err != nil {
		// 		fmt.Println("could not insert chain. ", err.Error())
		// 		os.Exit(1)
		// 	}
		//
		// 	nodeID, err := queries.InsertNode(context.Background(), database.InsertNodeParams{
		// 		ChainID:          row.ID,
		// 		ConfigFolder:     path,
		// 		Moniker:          e.Moniker,
		// 		ValidatorKey:     e.ValMnemonic,
		// 		ValidatorKeyName: e.ValKeyName,
		// 		BinaryVersion:    e.Version,
		// 		ProcessID:        0,
		// 		IsValidator:      1,
		// 		IsArchive:        0,
		// 		IsRunning:        0,
		// 	})
		// 	if err != nil {
		// 		fmt.Println("could not insert node", err.Error())
		// 		os.Exit(1)
		// 	}
		//
		// 	err = queries.InsertPorts(context.Background(), database.InsertPortsParams{
		// 		NodeID: nodeID,
		// 		P1317:  int64(e.Ports.P1317),
		// 		P8080:  int64(e.Ports.P8080),
		// 		P9090:  int64(e.Ports.P9090),
		// 		P9091:  int64(e.Ports.P9091),
		// 		P8545:  int64(e.Ports.P8545),
		// 		P8546:  int64(e.Ports.P8546),
		// 		P6065:  int64(e.Ports.P6065),
		// 		P26658: int64(e.Ports.P26658),
		// 		P26657: int64(e.Ports.P26657),
		// 		P6060:  int64(e.Ports.P6060),
		// 		P26656: int64(e.Ports.P26656),
		// 		P26660: int64(e.Ports.P26660),
		// 	})
		// 	if err != nil {
		// 		fmt.Println("could not insert ports", err.Error())
		// 		os.Exit(1)
		// 	}
		//
		// 	fmt.Println("Node added with id:", nodeID)
	},
}

func init() {
	PlaygroundCmd.AddCommand(initGaiaCmd)
	initGaiaCmd.Flags().StringP("version", "v", "local", "Version of the Evmos node that you want to use, defaults to local. Tag names are supported.")
}
