package gaia

import (
	"github.com/hanchon/hanchond/playground/cosmosdaemon"
	"github.com/hanchon/hanchond/playground/filesmanager"
)

type Gaia struct {
	*cosmosdaemon.Daemon
}

func NewGaia(moniker string, homeDir string, chainID string, keyName string, denom string) *Gaia {
	g := &Gaia{
		Daemon: cosmosdaemon.NewDameon(moniker, "gaia", homeDir, chainID, keyName, cosmosdaemon.CosmosAlgo, denom, "cosmos", cosmosdaemon.GaiaSDK),
	}
	g.SetBinaryPath(filesmanager.GetGaiadPath())
	g.SetCustomConfig(g.UpdateGenesisFile)
	return g
}
