package gaia

import (
	"fmt"
)

func (g *Gaia) Start(name string) (int, error) {
	logFile := g.HomeDir + "/run.log"
	cmd := fmt.Sprintf("%s start --home %s --api.enable --grpc.enable >> %s 2>&1",
		g.BinaryPath,
		g.HomeDir,
		logFile,
	)
	return g.Daemon.Start(cmd)
}
