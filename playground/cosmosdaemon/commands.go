package cosmosdaemon

import (
	"os/exec"
	"strings"
)

func (d *Daemon) AddGenesisAccount() error {
	validatorAddr, err := d.GetValidatorAddress()
	if err != nil {
		return err
	}
	command := exec.Command( //nolint:gosec
		d.BinaryPath,
		"add-genesis-account",
		validatorAddr,
		d.ValidatorInitialSupply+d.BaseDenom,
		"--keyring-backend",
		d.KeyringBackend,
		"--home",
		d.HomeDir,
	)
	_, err = command.CombinedOutput()
	return err
}

func (d *Daemon) ValidatorGenTx() error {
	command := exec.Command( //nolint:gosec
		d.BinaryPath,
		"gentx",
		d.ValKeyName,
		d.ValidatorInitialSupply[0:len(d.ValidatorInitialSupply)-4]+d.BaseDenom,
		"--gas-prices",
		d.BaseFee+d.BaseDenom,
		"--chain-id",
		d.ChainID,
		"--keyring-backend",
		d.KeyringBackend,
		"--home",
		d.HomeDir,
	)
	_, err := command.CombinedOutput()
	return err
}

func (d *Daemon) CollectGenTxs() error {
	command := exec.Command( //nolint:gosec
		d.BinaryPath,
		"collect-gentxs",
		"--home",
		d.HomeDir,
	)
	_, err := command.CombinedOutput()
	return err
}

func (d *Daemon) ValidateGenesis() error {
	command := exec.Command( //nolint:gosec
		d.BinaryPath,
		"validate-genesis",
		"--home",
		d.HomeDir,
	)
	_, err := command.CombinedOutput()
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
		return "", err
	}
	return strings.TrimSpace(string(o)), nil
}
