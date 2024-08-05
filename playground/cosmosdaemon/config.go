package cosmosdaemon

import "os/exec"

func (d *Daemon) ConfigKeyring() error {
	command := exec.Command( //nolint:gosec
		d.BinaryPath,
		"config",
		"keyring-backend",
		d.KeyringBackend,
		"--home",
		d.HomeDir,
	)
	_, err := command.CombinedOutput()
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
	_, err := command.CombinedOutput()
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
	_, err := command.CombinedOutput()
	return err
}
