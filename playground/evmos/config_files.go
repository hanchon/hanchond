package evmos

import (
	"strings"
)

func (e *Evmos) UpdateAppFile() error {
	//  Pruning
	appFile, err := e.Daemon.OpenAppFile()
	if err != nil {
		return err
	}
	appFile = e.enableWeb3API(appFile)
	return e.Daemon.SaveAppFile(appFile)
}

func (e *Evmos) enableWeb3API(config []byte) []byte {
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
