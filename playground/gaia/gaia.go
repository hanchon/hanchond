package gaia

import (
	"github.com/hanchon/hanchond/playground/cosmosdaemon"
)

type Gaia struct {
	cosmosdaemon.Daemon
}

// func NewGaia(version string, homeDir string, chainID string, keyName string, algo string, denom string) *Gaia {
// 	g := &Gaia{
// 		Daemon: *cosmosdaemon.NewDameon(version, homeDir, chainID, keyName, algo, denom),
// 	}
// 	g.SetBinaryPath(filesmanager.GetGaiadPath())
// 	return g
// }
