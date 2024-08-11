package evmos

import (
	"fmt"

	"github.com/hanchon/hanchond/playground/cosmosdaemon"
	"github.com/hanchon/hanchond/playground/filesmanager"
)

type Evmos struct {
	*cosmosdaemon.Daemon
}

func NewEvmos(moniker string, version string, homeDir string, chainID string, keyName string, denom string) *Evmos {
	e := &Evmos{
		Daemon: cosmosdaemon.NewDameon(
			moniker,
			fmt.Sprintf("evmosd%v", version),
			homeDir,
			chainID,
			keyName,
			cosmosdaemon.EthAlgo,
			denom,
			"evmos",
			cosmosdaemon.EvmosSDK,
		),
	}
	e.SetBinaryPath(filesmanager.GetEvmosdPath(version))
	e.SetCustomConfig(e.UpdateAppFile)
	return e
}
