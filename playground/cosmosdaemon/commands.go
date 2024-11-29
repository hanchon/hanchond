package cosmosdaemon

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/hanchon/hanchond/playground/filesmanager"
)

func (d *Daemon) AddGenesisAccount(validatorAddr string) error {
	args := []string{
		"add-genesis-account",
		validatorAddr,
		d.ValidatorInitialSupply + "poa," + d.ValidatorInitialSupply + d.BaseDenom,
		"--keyring-backend",
		d.KeyringBackend,
		"--home",
		d.HomeDir,
	}
	if d.SDKVersion == GaiaSDK {
		args = append([]string{"genesis"}, args...)
	}

	command := exec.Command( //nolint:gosec
		d.BinaryPath,
		args...,
	)

	out, err := command.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("error %s: %s", err.Error(), string(out))
	}
	return err
}

func (d *Daemon) AddAuthorityAccount(validatorAddr string) error {
	args := []string{
		"add-genesis-account",
		validatorAddr,
		"1000000000000000000000000" + "poa," + d.ValidatorInitialSupply + d.BaseDenom,
		"--keyring-backend",
		d.KeyringBackend,
		"--home",
		d.HomeDir,
	}
	if d.SDKVersion == GaiaSDK {
		args = append([]string{"genesis"}, args...)
	}

	command := exec.Command( //nolint:gosec
		d.BinaryPath,
		args...,
	)

	out, err := command.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("error %s: %s", err.Error(), string(out))
	}
	return err
}

func (d *Daemon) ValidatorGenTx() error {
	args := []string{
		"gentx",
		d.ValKeyName,
		d.ValidatorInitialSupply + "poa",
		"--commission-max-rate=1.00",
		"--commission-rate=1.00",
		"--gas-prices",
		d.BaseFee + d.BaseDenom,
		"--chain-id",
		d.ChainID,
		"--keyring-backend",
		d.KeyringBackend,
		"--home",
		d.HomeDir,
	}

	if d.SDKVersion == GaiaSDK {
		args = append([]string{"genesis"}, args...)
	}

	command := exec.Command( //nolint:gosec
		d.BinaryPath,
		args...,
	)
	out, err := command.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("error %s: %s", err.Error(), string(out))
	}
	return err
}

func (d *Daemon) CollectGenTxs() error {
	args := []string{
		"collect-gentxs",
		"--home",
		d.HomeDir,
	}

	if d.SDKVersion == GaiaSDK {
		args = append([]string{"genesis"}, args...)
	}
	command := exec.Command( //nolint:gosec
		d.BinaryPath,
		args...,
	)
	out, err := command.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("error %s: %s", err.Error(), string(out))
	}
	return err
}

func (d *Daemon) ValidateGenesis() error {
	args := []string{
		"validate-genesis",
		"--home",
		d.HomeDir,
	}
	if d.SDKVersion == GaiaSDK {
		args = append([]string{"genesis"}, args...)
	}
	command := exec.Command( //nolint:gosec
		d.BinaryPath,
		args...,
	)
	out, err := command.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("error %s: %s", err.Error(), string(out))
	}
	return err
}

// Returns bech32 encoded validator addresss
func (d *Daemon) GetValidatorAddress() (string, error) {
	command := exec.Command( //nolint:gosec
		d.BinaryPath,
		"keys",
		"show",
		"-a",
		d.ValKeyName,
		"--keyring-backend",
		d.KeyringBackend,
		"--home",
		d.HomeDir,
	)
	o, err := command.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("error %s: %s", err.Error(), string(o))
		return "", err
	}
	return strings.TrimSpace(string(o)), nil
}

func (d *Daemon) Start(startCmd string) (int, error) {
	command := exec.Command("bash", "-c", startCmd)
	// Deattach the program
	command.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
	err := command.Start()
	if err != nil {
		return 0, err
	}
	time.Sleep(2 * time.Second)
	id, err := filesmanager.GetChildPID(command.Process.Pid)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d *Daemon) GetNodeID() (string, error) {
	command := exec.Command( //nolint:gosec
		d.BinaryPath,
		"tendermint",
		"show-node-id",
		"--home",
		d.HomeDir,
	)
	o, err := command.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("error %s: %s", err.Error(), string(o))
		return "", err
	}
	return strings.TrimSpace(string(o)), nil
}
