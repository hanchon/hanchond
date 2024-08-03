package evmos

import (
	"fmt"
	"os/exec"

	"github.com/hanchon/hanchond/playground/filesmanager"
)

func (e *Evmos) SendIBC(port, channel, receiver, amount string) (string, error) {
	command := exec.Command( //nolint:gosec
		filesmanager.GetEvmosdPath(e.Version),
		"tx",
		"ibc-transfer",
		"transfer",
		port,
		channel,
		receiver,
		amount,
		"--keyring-backend",
		e.KeyringBackend,
		"--home",
		e.HomeDir,
		"--node",
		fmt.Sprintf("http://localhost:%d", e.Ports.P26657),
		"--from",
		e.ValKeyName,
		"--fees",
		fmt.Sprintf("1000000000000%s", e.BaseDenom),
		"-y",
	)

	out, err := command.CombinedOutput()
	return string(out), err
}
