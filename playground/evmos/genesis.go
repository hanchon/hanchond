package evmos

import (
	"fmt"
	"strings"
)

func (e *Evmos) initNode() error {
	if err := e.ConfigKeyring(); err != nil {
		return err
	}

	if err := e.ConfigChainID(); err != nil {
		return err
	}

	if err := e.EvmosdInit(); err != nil {
		return err
	}
	return nil
}

func (e *Evmos) InitArchiveNode(origin *Evmos) error {
	if err := e.initNode(); err != nil {
		return err
	}
	// Copy genesis file
	if err := e.copyGenesisFile(origin.getGenesisPath()); err != nil {
		return err
	}

	//  Pruning
	appFile, err := e.openAppFile()
	if err != nil {
		return err
	}
	appFile = e.SetPruningInAppFile(false, appFile)
	// Enable API
	appFile = e.EnableAPI(appFile)
	if err := e.saveAppFile(appFile); err != nil {
		return err
	}

	// Update ports
	if err := e.SetPorts(); err != nil {
		return err
	}

	// Add persistent peer
	nodeID, err := origin.EvmosdShowNodeID()
	if err != nil {
		return err
	}

	configFile, err := e.openConfigFile()
	if err != nil {
		return err
	}

	configFile = []byte(strings.Replace(
		string(configFile),
		"persistent_peers = \"\"",
		fmt.Sprintf("persistent_peers = \"%s@localhost:%d\"", nodeID, origin.Ports.P26656),
		1,
	))

	if err := e.saveConfigFile(configFile); err != nil {
		return err
	}

	return nil
}

func (e *Evmos) InitGenesis() error {
	if err := e.initNode(); err != nil {
		return err
	}

	if err := e.AddValidatorKey(); err != nil {
		return err
	}

	genesis, err := e.openGenesisFile()
	if err != nil {
		return err
	}
	// Update the genesis
	e.UpdateGenesisDenom(genesis)
	e.SetGenesisGasLimit(genesis)
	e.SetGenesisBaseFee(genesis)
	e.SetGenesisFastProposals(genesis)

	if err := e.saveGenesisFile(genesis); err != nil {
		return err
	}

	// ConfigFile
	configFile, err := e.openConfigFile()
	if err != nil {
		return err
	}
	// TODO: renable this if we want slow blocks
	// configFile = e.UpdateConfigFile(configFile)
	if err := e.saveConfigFile(configFile); err != nil {
		return err
	}

	//  Pruning
	appFile, err := e.openAppFile()
	if err != nil {
		return err
	}
	appFile = e.SetPruningInAppFile(true, appFile)
	if err := e.saveAppFile(appFile); err != nil {
		return err
	}

	if err := e.AddGenesisAccount(); err != nil {
		return err
	}

	if err := e.UpdateTotalSupply(); err != nil {
		return err
	}

	if err := e.ValidatorGenTx(); err != nil {
		return err
	}

	if err := e.CollectGenTxs(); err != nil {
		return err
	}

	if err := e.ValidateGenesis(); err != nil {
		return err
	}

	// Make backups of the AppFile and ConfigFile
	if err := e.backupConfigFiles(); err != nil {
		return err
	}

	return nil
}

func (e *Evmos) UpdateGenesisDenom(genesis map[string]interface{}) {
	appState := genesis["app_state"].(map[string]interface{})
	// Staking
	appState["staking"].(map[string]interface{})["params"].(map[string]interface{})["bond_denom"] = e.BaseDenom
	// Gov
	appState["gov"].(map[string]interface{})["params"].(map[string]interface{})["min_deposit"].([]interface{})[0].(map[string]interface{})["denom"] = e.BaseDenom
	// Evm
	appState["evm"].(map[string]interface{})["params"].(map[string]interface{})["evm_denom"] = e.BaseDenom
	// Inflation
	appState["inflation"].(map[string]interface{})["params"].(map[string]interface{})["mint_denom"] = e.BaseDenom
}

func (e *Evmos) SetGenesisGasLimit(genesis map[string]interface{}) {
	consensusParams := genesis["consensus_params"].(map[string]interface{})
	consensusParams["block"].(map[string]interface{})["max_gas"] = e.GasLimit
}

func (e *Evmos) SetGenesisBaseFee(genesis map[string]interface{}) {
	appState := genesis["app_state"].(map[string]interface{})
	appState["feemarket"].(map[string]interface{})["params"].(map[string]interface{})["base_fee"] = e.BaseFee
}

func (e *Evmos) SetGenesisFastProposals(genesis map[string]interface{}) {
	appState := genesis["app_state"].(map[string]interface{})
	appState["gov"].(map[string]interface{})["params"].(map[string]interface{})["max_deposit_period"] = "30s"
	appState["gov"].(map[string]interface{})["params"].(map[string]interface{})["voting_period"] = "30s"
}

func (e *Evmos) UpdateConfigFile(config []byte) []byte {
	configValues := string(config)
	// Tendermint Values
	configValues = strings.Replace(
		configValues,
		"timeout_propose = \"3s\"",
		"timeout_propose = \"10s\"",
		1,
	)
	configValues = strings.Replace(
		configValues,
		"timeout_propose_delta = \"500ms\"",
		"timeout_propose_delta = \"1s\"",
		1,
	)
	configValues = strings.Replace(
		configValues,
		"timeout_prevote = \"1s\"",
		"timeout_prevote = \"5s\"",
		1,
	)
	configValues = strings.Replace(
		configValues,
		"timeout_prevote_delta = \"500ms\"",
		"timeout_prevote_delta = \"1s\"",
		1,
	)
	configValues = strings.Replace(
		configValues,
		"timeout_precommit = \"1s\"",
		"timeout_precommit = \"5s\"",
		1,
	)
	configValues = strings.Replace(
		configValues,
		"timeout_precommit_delta = \"500ms\"",
		"timeout_precommit_delta = \"1s\"",
		1,
	)
	configValues = strings.Replace(
		configValues,
		"timeout_commit = \"3s\"",
		"timeout_commit = \"5s\"",
		1,
	)
	configValues = strings.Replace(
		configValues,
		"timeout_broadcast_tx_commit = \"10s\"",
		"timeout_broadcast_tx_commit = \"15s\"",
		1,
	)
	return []byte(configValues)
}

func (e *Evmos) EnableAPI(config []byte) []byte {
	configValues := string(config)
	configValues = strings.Replace(
		configValues,
		`# Enable defines if the JSONRPC server should be enabled.
enable = false`,
		"enable = true",
		1,
	)
	return []byte(configValues)
}

func (e *Evmos) SetPruningInAppFile(pruningEnabled bool, config []byte) []byte {
	configValues := string(config)
	if pruningEnabled {
		configValues = strings.Replace(
			configValues,
			"pruning = \"default\"",
			"pruning = \"custom\"",
			1,
		)
		configValues = strings.Replace(
			configValues,
			"pruning-keep-recent = \"0\"",
			"pruning-keep-recent = \"2\"",
			1,
		)
		configValues = strings.Replace(
			configValues,
			"pruning-interval = \"0\"",
			"pruning-interval = \"10\"",
			1,
		)
		return []byte(configValues)
	}

	// NoPrunning
	configValues = strings.Replace(
		configValues,
		"pruning = \"default\"",
		"pruning = \"nothing\"",
		1,
	)
	configValues = strings.Replace(
		configValues,
		"pruning = \"custom\"",
		"pruning = \"nothing\"",
		1,
	)
	configValues = strings.Replace(
		configValues,
		"pruning = \"everything\"",
		"pruning = \"nothing\"",
		1,
	)
	return []byte(configValues)
}

func (e *Evmos) UpdateTotalSupply() error {
	genesis, err := e.openGenesisFile()
	if err != nil {
		return err
	}
	appState := genesis["app_state"].(map[string]interface{})
	appState["bank"].(map[string]interface{})["supply"].([]interface{})[0].(map[string]interface{})["amount"] = e.ValidatorInitialSupply

	if err := e.saveGenesisFile(genesis); err != nil {
		return err
	}

	return nil
}
