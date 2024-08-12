package cosmosdaemon

import (
	"fmt"
	"path/filepath"

	"github.com/hanchon/hanchond/playground/database"
	"github.com/hanchon/hanchond/playground/filesmanager"
)

func InitMultiNodeChain(nodes []*Daemon, queries *database.Queries) error {
	if err := InitFilesAndDB(nodes, queries); err != nil {
		return err
	}
	if err := JoinGenesisTransactions(nodes, queries); err != nil {
		return err
	}
	if err := CollectGenTxns(nodes, queries); err != nil {
		return err
	}
	if err := UpdatePeers(nodes, queries); err != nil {
		return err
	}
	return nil
}

func InitFilesAndDB(nodes []*Daemon, queries *database.Queries) error {
	var chainDB database.Chain
	var err error

	for k := range nodes {
		// Init the config files
		if err := nodes[k].InitNode(); err != nil {
			return err
		}
		// Update general parameters in the genesis file
		if err := nodes[k].UpdateGenesisFile(); err != nil {
			return err
		}
		if err := nodes[k].UpdateConfigFile(false); err != nil {
			return err
		}
		if err := nodes[k].UpdateAppFile(); err != nil {
			return err
		}
		if err := nodes[k].CreateGenTx(); err != nil {
			return err
		}
		// Assign random and unique ports
		if err := nodes[k].AssignPorts(queries); err != nil {
			return err
		}
		// Update the Config Files
		if err := nodes[k].UpdateConfigPorts(); err != nil {
			return err
		}

		// Apply client specific configurations
		if err := nodes[k].ExecuteCustomConfig(); err != nil {
			return err
		}

		if k == 0 {
			chainDB, err = nodes[k].SaveChainToDB(queries)
			if err != nil {
				return err
			}
		}
		nodeID, err := nodes[k].SaveNodeToDB(chainDB, queries)
		if err != nil {
			return err
		}
		fmt.Printf("Node added with ID: %d\n", nodeID)
	}
	return nil
}

func JoinGenesisTransactions(nodes []*Daemon, queries *database.Queries) error {
	_ = queries
	for k, v := range nodes {
		// Node 0 will be the only the one that creates the genesis
		if k == 0 {
			continue
		}
		files, err := filepath.Glob(v.HomeDir + "/config/gentx/*.json")
		if err != nil {
			return err
		}
		if len(files) == 0 {
			return err
		}

		if err := filesmanager.CopyFile(
			files[0],
			nodes[0].HomeDir+"/config/gentx",
		); err != nil {
			return err
		}
		addr, err := v.GetValidatorAddress()
		if err != nil {
			return err
		}
		if err := nodes[0].AddGenesisAccount(addr); err != nil {
			return err
		}
	}
	return nil
}

func CollectGenTxns(nodes []*Daemon, queries *database.Queries) error {
	_ = queries
	if err := nodes[0].CollectGenTxs(); err != nil {
		return err
	}
	if err := nodes[0].ValidateGenesis(); err != nil {
		return err
	}
	return nil
}

func UpdatePeers(nodes []*Daemon, queries *database.Queries) error {
	_ = queries
	peers := []string{}

	for k := range nodes {
		peerInfo, err := nodes[k].GetPeerInfo()
		if err != nil {
			return err
		}
		peers = append(peers, peerInfo)
		if k != 0 {
			if err := filesmanager.CopyFile(
				nodes[0].HomeDir+"/config/genesis.json",
				nodes[k].HomeDir+"/config/genesis.json",
			); err != nil {
				return err
			}
		}
	}

	for k := range nodes {
		if k == 0 {
			continue
		}
		if err := nodes[k].AddPersistenPeers([]string{peers[k-1]}); err != nil {
			return err
		}
	}
	return nil
}
