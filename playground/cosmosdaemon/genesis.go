package cosmosdaemon

import "fmt"

func (d *Daemon) UpdateGenesisFile() error {
	genesis, err := d.OpenGenesisFile()
	if err != nil {
		return err
	}
	// Update the genesis
	d.setStaking(genesis)
	d.setEvm(genesis)
	d.setInflation(genesis)
	d.setCrisis(genesis)
	d.setMint(genesis)
	d.setProvider(genesis)
	d.setConsensusParams(genesis)
	d.setFeeMarket(genesis)
	d.setGovernance(genesis, true)

	return d.SaveGenesisFile(genesis)
}

func (d *Daemon) setStaking(genesis map[string]interface{}) {
	appState := genesis["app_state"].(map[string]interface{})
	if v, ok := appState["staking"]; ok {
		if v, ok := v.(map[string]interface{}); ok {
			if v, ok := v["params"]; ok {
				if v, ok := v.(map[string]interface{}); ok {
					// Base Denom
					if _, ok := v["base_denom"]; ok {
						appState["staking"].(map[string]interface{})["params"].(map[string]interface{})["bond_denom"] = d.BaseDenom
					}

					// Bond denom
					if _, ok := v["bond_denom"]; ok {
						appState["staking"].(map[string]interface{})["params"].(map[string]interface{})["bond_denom"] = d.BaseDenom
					}
				}
			}
		}
	}
}

func (d *Daemon) setEvm(genesis map[string]interface{}) {
	appState := genesis["app_state"].(map[string]interface{})
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
}

func (d *Daemon) setInflation(genesis map[string]interface{}) {
	appState := genesis["app_state"].(map[string]interface{})
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

func (d *Daemon) setCrisis(genesis map[string]interface{}) {
	appState := genesis["app_state"].(map[string]interface{})
	if v, ok := appState["crisis"]; ok {
		if v, ok := v.(map[string]interface{}); ok {
			if v, ok := v["constant_fee"]; ok {
				if v, ok := v.(map[string]interface{}); ok {
					if _, ok := v["denom"]; ok {
						appState["crisis"].(map[string]interface{})["constant_fee"].(map[string]interface{})["denom"] = d.BaseDenom
					}
				}
			}
		}
	}
}

func (d *Daemon) setMint(genesis map[string]interface{}) {
	appState := genesis["app_state"].(map[string]interface{})
	if v, ok := appState["mint"]; ok {
		if v, ok := v.(map[string]interface{}); ok {
			if v, ok := v["params"]; ok {
				if v, ok := v.(map[string]interface{}); ok {
					if _, ok := v["mint_denom"]; ok {
						appState["mint"].(map[string]interface{})["params"].(map[string]interface{})["mint_denom"] = d.BaseDenom
					}
				}
			}
		}
	}
}

func (d *Daemon) setProvider(genesis map[string]interface{}) {
	appState := genesis["app_state"].(map[string]interface{})
	if v, ok := appState["provider"]; ok {
		if v, ok := v.(map[string]interface{}); ok {
			if v, ok := v["params"]; ok {
				if v, ok := v.(map[string]interface{}); ok {
					if v, ok := v["consumer_reward_denom_registration_fee"]; ok {
						if v, ok := v.(map[string]interface{}); ok {
							if _, ok := v["denom"]; ok {
								appState["provider"].(map[string]interface{})["params"].(map[string]interface{})["consumer_reward_denom_registration_fee"].(map[string]interface{})["denom"] = d.BaseDenom
							}
						}
					}
				}
			}
		}
	}
}

func (d *Daemon) setConsensusParams(genesis map[string]interface{}) {
	var consensusParams map[string]interface{}
	if _, ok := genesis["consensus_params"]; ok {
		consensusParams = genesis["consensus_params"].(map[string]interface{})
	}

	// SDKv0.50 support
	if _, ok := genesis["consensus"]; ok {
		consensusParams = genesis["consensus"].(map[string]interface{})["params"].(map[string]interface{})
	}

	if v, ok := consensusParams["block"]; ok {
		if v, ok := v.(map[string]interface{}); ok {
			if _, ok := v["max_gas"]; ok {
				consensusParams["block"].(map[string]interface{})["max_gas"] = d.GasLimit
				fmt.Println("editing the max gas")
			}
		}
	}
}

func (d *Daemon) setFeeMarket(genesis map[string]interface{}) {
	appState := genesis["app_state"].(map[string]interface{})
	if v, ok := appState["feemarket"]; ok {
		if v, ok := v.(map[string]interface{}); ok {
			if v, ok := v["params"]; ok {
				if v, ok := v.(map[string]interface{}); ok {
					// Evmos FeeMarket
					if _, ok := v["base_fee"]; ok {
						appState["feemarket"].(map[string]interface{})["params"].(map[string]interface{})["base_fee"] = d.BaseFee
						// FeeMarket using static base fee
						appState["feemarket"].(map[string]interface{})["params"].(map[string]interface{})["base_fee_change_denominator"] = 1
						appState["feemarket"].(map[string]interface{})["params"].(map[string]interface{})["elasticity_multiplier"] = 1
						appState["feemarket"].(map[string]interface{})["params"].(map[string]interface{})["min_gas_multiplier"] = "0.0"

					}
					// SDK FeeMarket
					if _, ok := v["fee_denom"]; ok {
						appState["feemarket"].(map[string]interface{})["params"].(map[string]interface{})["fee_denom"] = d.BaseDenom
					}
				}
			}
		}
	}
}

func (d *Daemon) setGovernance(genesis map[string]interface{}, fastProposals bool) {
	appState := genesis["app_state"].(map[string]interface{})
	if v, ok := appState["gov"]; ok {
		if v, ok := v.(map[string]interface{}); ok {
			if v, ok := v["params"]; ok {
				if v, ok := v.(map[string]interface{}); ok {
					// Proposals
					if fastProposals {
						if _, ok := v["max_deposit_period"]; ok {
							appState["gov"].(map[string]interface{})["params"].(map[string]interface{})["max_deposit_period"] = "10s"
						}

						if _, ok := v["voting_period"]; ok {
							appState["gov"].(map[string]interface{})["params"].(map[string]interface{})["voting_period"] = "15s"
						}

						if _, ok := v["expedited_voting_period"]; ok {
							appState["gov"].(map[string]interface{})["params"].(map[string]interface{})["expedited_voting_period"] = "14s"
						}
					}

					//  Expedited_min_deposit
					if v, ok := v["expedited_min_deposit"]; ok {
						if v, ok := v.([]interface{}); ok {
							if len(v) > 0 {
								if v, ok := v[0].(map[string]interface{}); ok {
									if _, ok := v["denom"]; ok {
										appState["gov"].(map[string]interface{})["params"].(map[string]interface{})["expedited_min_deposit"].([]interface{})[0].(map[string]interface{})["denom"] = d.BaseDenom
									}
								}
							}
						}
					}

					// Min Deposit
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
}

func (d *Daemon) setBank() error {
	genesis, err := d.OpenGenesisFile()
	if err != nil {
		return err
	}
	appState := genesis["app_state"].(map[string]interface{})
	appState["bank"].(map[string]interface{})["supply"].([]interface{})[0].(map[string]interface{})["amount"] = d.ValidatorInitialSupply

	if err := d.SaveGenesisFile(genesis); err != nil {
		return err
	}

	return nil
}
