package gaia

func (g *Gaia) UpdateGenesisFile() error {
	// Gaia extra config
	genesis, err := g.Daemon.OpenGenesisFile()
	if err != nil {
		return err
	}

	g.setUnbondingTime(genesis)
	// Maybe we need to update the `genesis_time` but I am not sure why

	return g.Daemon.SaveGenesisFile(genesis)
}

func (g *Gaia) setUnbondingTime(genesis map[string]interface{}) {
	appState := genesis["app_state"].(map[string]interface{})
	if v, ok := appState["staking"]; ok {
		if v, ok := v.(map[string]interface{}); ok {
			if v, ok := v["params"]; ok {
				if v, ok := v.(map[string]interface{}); ok {
					// Base Denom
					if _, ok := v["unbonding_time"]; ok {
						appState["staking"].(map[string]interface{})["params"].(map[string]interface{})["unbonding_time"] = "1814400s"
					}
				}
			}
		}
	}
}
