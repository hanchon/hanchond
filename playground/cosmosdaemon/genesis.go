package cosmosdaemon

import "strings"

func (d *Daemon) initNode() error {
	if err := d.ConfigKeyring(); err != nil {
		return err
	}

	if err := d.ConfigChainID(); err != nil {
		return err
	}

	if err := d.NodeInit(); err != nil {
		return err
	}
	return nil
}

func (d *Daemon) InitGenesisPart1() error {
	if err := d.initNode(); err != nil {
		return err
	}

	if err := d.AddValidatorKey(); err != nil {
		return err
	}
	return nil
}

func (d *Daemon) InitGenesis() error {
	if err := d.InitGenesisPart1(); err != nil {
		return err
	}
	return d.InitGenesisPart2()
}

func (d *Daemon) InitGenesisPart2() error {
	genesis, err := d.openGenesisFile()
	if err != nil {
		return err
	}
	// Update the genesis
	d.UpdateGenesisDenom(genesis)
	d.SetGenesisGasLimit(genesis)
	d.SetGenesisFeeMarketBaseFee(genesis)
	d.SetGenesisFastProposals(genesis)

	if err := d.saveGenesisFile(genesis); err != nil {
		return err
	}

	// ConfigFile
	configFile, err := d.openConfigFile()
	if err != nil {
		return err
	}
	// TODO: renable this if we want slow blocks
	// configFile = e.UpdateConfigFile(configFile)
	if err := d.saveConfigFile(configFile); err != nil {
		return err
	}

	//  Pruning
	appFile, err := d.openAppFile()
	if err != nil {
		return err
	}
	appFile = d.SetPruningInAppFile(true, appFile)
	if err := d.saveAppFile(appFile); err != nil {
		return err
	}

	if err := d.AddGenesisAccount(); err != nil {
		return err
	}

	if err := d.UpdateTotalSupply(); err != nil {
		return err
	}

	if err := d.ValidatorGenTx(); err != nil {
		return err
	}

	if err := d.CollectGenTxs(); err != nil {
		return err
	}

	if err := d.ValidateGenesis(); err != nil {
		return err
	}

	// Make backups of the AppFile and ConfigFile
	if err := d.backupConfigFiles(); err != nil {
		return err
	}

	return nil
}

func (d *Daemon) UpdateGenesisDenom(genesis map[string]interface{}) {
	// We are doing the validation for each elemen to run the same function for any cosmos chain
	// If the module is not enable it will not update the value
	appState := genesis["app_state"].(map[string]interface{})
	// Staking
	if v, ok := appState["staking"]; ok {
		if v, ok := v.(map[string]interface{}); ok {
			if v, ok := v["params"]; ok {
				if v, ok := v.(map[string]interface{}); ok {
					if _, ok := v["base_denom"]; ok {
						appState["staking"].(map[string]interface{})["params"].(map[string]interface{})["bond_denom"] = d.BaseDenom
					}
				}
			}
		}
	}

	// Gov
	if v, ok := appState["dov"]; ok {
		if v, ok := v.(map[string]interface{}); ok {
			if v, ok := v["params"]; ok {
				if v, ok := v.(map[string]interface{}); ok {
					if v, ok := v["min_deposit"]; ok {
						if v, ok := v.([]interface{}); ok {
							if len(v) > 0 {
								if v, ok := v[0].(map[string]interface{}); ok {
									if _, ok := v["denom"]; ok {
										appState["gov"].(map[string]interface{})["params"].(map[string]interface{})["min_deposit"].([]interface{})[0].(map[string]interface{})["denom"] = d.BaseDenom
									}
								}
							}
						}
					}
				}
			}
		}
	}

	// EVM
	if v, ok := appState["evm"]; ok {
		if v, ok := v.(map[string]interface{}); ok {
			if v, ok := v["params"]; ok {
				if v, ok := v.(map[string]interface{}); ok {
					if _, ok := v["evm_denom"]; ok {
						appState["evm"].(map[string]interface{})["params"].(map[string]interface{})["evm_denom"] = d.BaseDenom
					}
				}
			}
		}
	}

	// Inflation
	if v, ok := appState["inflation"]; ok {
		if v, ok := v.(map[string]interface{}); ok {
			if v, ok := v["params"]; ok {
				if v, ok := v.(map[string]interface{}); ok {
					if _, ok := v["mint_denom"]; ok {
						appState["inflation"].(map[string]interface{})["params"].(map[string]interface{})["mint_denom"] = d.BaseDenom
					}
				}
			}
		}
	}
}

func (d *Daemon) SetGenesisGasLimit(genesis map[string]interface{}) {
	consensusParams := genesis["consensus_params"].(map[string]interface{})
	if v, ok := consensusParams["block"]; ok {
		if v, ok := v.(map[string]interface{}); ok {
			if _, ok := v["max_gas"]; ok {
				consensusParams["block"].(map[string]interface{})["max_gas"] = d.GasLimit
			}
		}
	}
}

func (d *Daemon) SetGenesisFeeMarketBaseFee(genesis map[string]interface{}) {
	appState := genesis["app_state"].(map[string]interface{})
	if v, ok := appState["feemarket"]; ok {
		if v, ok := v.(map[string]interface{}); ok {
			if v, ok := v["params"]; ok {
				if v, ok := v.(map[string]interface{}); ok {
					if _, ok := v["base_fee"]; ok {
						appState["feemarket"].(map[string]interface{})["params"].(map[string]interface{})["base_fee"] = d.BaseFee
					}
				}
			}
		}
	}
}

func (d *Daemon) SetGenesisFastProposals(genesis map[string]interface{}) {
	appState := genesis["app_state"].(map[string]interface{})
	if v, ok := appState["gov"]; ok {
		if v, ok := v.(map[string]interface{}); ok {
			if v, ok := v["params"]; ok {
				if v, ok := v.(map[string]interface{}); ok {
					if _, ok := v["max_deposit_period"]; ok {
						appState["gov"].(map[string]interface{})["params"].(map[string]interface{})["max_deposit_period"] = "10s"
					}

					if _, ok := v["voting_period"]; ok {
						appState["gov"].(map[string]interface{})["params"].(map[string]interface{})["voting_period"] = "15s"
					}
				}
			}
		}
	}
}

func (d *Daemon) UpdateConfigFile(config []byte) []byte {
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

func (d *Daemon) EnableWeb3API(config []byte) []byte {
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

func (d *Daemon) SetPruningInAppFile(pruningEnabled bool, config []byte) []byte {
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

func (d *Daemon) UpdateTotalSupply() error {
	genesis, err := d.openGenesisFile()
	if err != nil {
		return err
	}
	appState := genesis["app_state"].(map[string]interface{})
	appState["bank"].(map[string]interface{})["supply"].([]interface{})[0].(map[string]interface{})["amount"] = d.ValidatorInitialSupply

	if err := d.saveGenesisFile(genesis); err != nil {
		return err
	}

	return nil
}
