package cosmosdaemon

import "os/exec"

func (e *Daemon) ConfigKeyring() error {
	command := exec.Command( //nolint:gosec
		e.BinaryPath,
		"config",
		"keyring-backend",
		e.KeyringBackend,
		"--home",
		e.HomeDir,
	)
	_, err := command.CombinedOutput()
	return err
}

func (e *Daemon) ConfigChainID() error {
	command := exec.Command( //nolint:gosec
		e.BinaryPath,
		"config",
		"chain-id",
		e.ChainID,
		"--home",
		e.HomeDir,
	)
	_, err := command.CombinedOutput()
	return err
}

func (e *Daemon) NodeInit() error {
	command := exec.Command( //nolint:gosec
		e.BinaryPath,
		"init",
		e.Moniker,
		"--chain-id",
		e.ChainID,
		"--home",
		e.HomeDir,
	)
	_, err := command.CombinedOutput()
	return err
}
