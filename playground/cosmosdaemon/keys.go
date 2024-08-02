package cosmosdaemon

import (
	"fmt"
	"os/exec"
	"strings"
)

func (e *Daemon) AddValidatorKey() error {
	return e.KeyAdd(e.ValKeyName, e.ValMnemonic)
}

func (e *Daemon) KeyAdd(name string, mnemonic string) error {
	cmd := fmt.Sprintf("echo \"%s\" | %s keys add %s --recover --keyring-backend %s --home %s --key-type %s",
		mnemonic,
		e.BinaryPath,
		name,
		e.KeyringBackend,
		e.HomeDir,
		e.KeyType,
	)
	command := exec.Command("bash", "-c", cmd)
	o, err := command.CombinedOutput()
	if strings.Contains(string(o), "duplicated") {
		return fmt.Errorf("duplicated address")
	}
	return err
}
