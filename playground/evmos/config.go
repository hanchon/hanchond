package evmos

import (
	"os/exec"
	"strings"

	"github.com/hanchon/hanchond/playground/filesmanager"
)

func (e *Evmos) ConfigKeyring() error {
	command := exec.Command( //nolint:gosec
		filesmanager.GetEvmosdPath(e.Version),
		"config",
		"keyring-backend",
		e.KeyringBackend,
		"--home",
		e.HomeDir,
	)
	_, err := command.CombinedOutput()
	return err
}

func (e *Evmos) ConfigChainID() error {
	command := exec.Command( //nolint:gosec
		filesmanager.GetEvmosdPath(e.Version),
		"config",
		"chain-id",
		e.ChainID,
		"--home",
		e.HomeDir,
	)
	_, err := command.CombinedOutput()
	return err
}

func (e *Evmos) EvmosdInit() error {
	command := exec.Command( //nolint:gosec
		filesmanager.GetEvmosdPath(e.Version),
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

func (e *Evmos) EvmosdShowNodeID() (string, error) {
	command := exec.Command( //nolint:gosec
		filesmanager.GetEvmosdPath(e.Version),
		"tendermint",
		"show-node-id",
		"--home",
		e.HomeDir,
	)
	o, err := command.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(o)), nil
}
