package evmos

import (
	"fmt"
	"strings"

	"github.com/hanchon/hanchond/playground/cosmosdaemon"
	"github.com/hanchon/hanchond/playground/filesmanager"
)

type Evmos struct {
	*cosmosdaemon.Daemon
}

func NewEvmos(moniker string, version string, homeDir string, chainID string, keyName string, denom string) *Evmos {
	daemonName := version
	if !strings.Contains(version, "evmosd") {
		daemonName = fmt.Sprintf("evmosd%s", version)
	}
	e := &Evmos{
		Daemon: cosmosdaemon.NewDameon(
			moniker,
			daemonName,
			homeDir,
			chainID,
			keyName,
			cosmosdaemon.EthAlgo,
			denom,
			"evmos",
			cosmosdaemon.EvmosSDK,
		),
	}
	e.SetBinaryPath(filesmanager.GetDaemondPath(daemonName))
	e.SetCustomConfig(e.UpdateAppFile)
	return e
}
