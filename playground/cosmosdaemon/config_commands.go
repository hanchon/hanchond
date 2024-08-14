package cosmosdaemon

import (
	"fmt"
	"os/exec"
)

func (d *Daemon) ConfigKeyring() error {
	command := exec.Command( //nolint:gosec
		d.BinaryPath,
		"config",
		"keyring-backend",
		d.KeyringBackend,
		"--home",
		d.HomeDir,
	)
	out, err := command.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("error %s: %s", err.Error(), string(out))
	}
	return err
}

func (d *Daemon) ConfigChainID() error {
	command := exec.Command( //nolint:gosec
		d.BinaryPath,
		"config",
		"chain-id",
		d.ChainID,
		"--home",
		d.HomeDir,
	)
	out, err := command.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("error %s: %s", err.Error(), string(out))
	}
	return err
}

func (d *Daemon) NodeInit() error {
	command := exec.Command( //nolint:gosec
		d.BinaryPath,
		"init",
		d.Moniker,
		"--chain-id",
		d.ChainID,
		"--home",
		d.HomeDir,
	)
	out, err := command.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("error %s: %s", err.Error(), string(out))
	}
	return err
}
