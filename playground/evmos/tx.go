package evmos

import (
	"fmt"
	"os/exec"

	"github.com/hanchon/hanchond/playground/filesmanager"
)

func (e *Evmos) GetTransaction(txhash string) (string, error) {
	command := exec.Command( //nolint:gosec
		filesmanager.GetEvmosdPath(e.Version),
		"q",
		"tx",
		"--type=hash",
		txhash,
		"--home",
		e.HomeDir,
		"--node",
		fmt.Sprintf("http://localhost:%d", e.Ports.P26657),
	)
	out, err := command.CombinedOutput()
	return string(out), err
}
