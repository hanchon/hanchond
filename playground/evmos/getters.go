package evmos

import (
	"os/exec"
	"strings"

	"github.com/hanchon/hanchond/playground/filesmanager"
)

// Returns bech32 encoded validator addresss
func (e *Evmos) GetValidatorAddress() (string, error) {
	command := exec.Command( //nolint:gosec
		filesmanager.GetEvmosdPath(e.Version),
		"keys",
		"show",
		"-a",
		e.ValKeyName,
		"--keyring-backend",
		e.KeyringBackend,
		"--home",
		e.HomeDir,
	)
	o, err := command.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(o)), nil
}
