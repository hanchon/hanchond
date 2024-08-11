package evmos

import (
	"fmt"
)

func (e *Evmos) Start(name string) (int, error) {
	logFile := e.HomeDir + "/run.log"
	cmd := fmt.Sprintf("%s start --chain-id %s --home %s --json-rpc.api eth,txpool,personal,net,debug,web3 --api.enable --grpc.enable >> %s 2>&1",
		e.BinaryPath,
		e.ChainID,
		e.HomeDir,
		logFile,
	)
	return e.Daemon.Start(cmd)
}
