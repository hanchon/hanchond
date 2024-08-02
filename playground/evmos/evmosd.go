package evmos

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/hanchon/hanchond/playground/filesmanager"
)

type Evmos struct {
	ValKeyName     string
	ValMnemonic    string
	KeyringBackend string
	HomeDir        string
	Version        string

	ChainID string
	Moniker string

	BaseDenom string
	GasLimit  string
	BaseFee   string

	ValidatorInitialSupply string

	Ports Ports
}

func NewEvmos(version string, homeDir string, chainID string, keyName string) *Evmos {
	// TODO: remove hardcoded values
	return &Evmos{
		ValKeyName:     keyName,
		ValMnemonic:    "gesture inject test cycle original hollow east ridge hen combine junk child bacon zero hope comfort vacuum milk pitch cage oppose unhappy lunar seat",
		KeyringBackend: "test",
		HomeDir:        homeDir,
		Version:        version,

		ChainID: chainID,
		Moniker: "moniker",

		BaseDenom: "aevmos",
		GasLimit:  "10000000",
		BaseFee:   "1000000000",

		ValidatorInitialSupply: "100000000000000000000000000",

		Ports: *NewPorts(),
	}
}

func (e *Evmos) KeyAdd(name string, mnemonic string) error {
	cmd := fmt.Sprintf("echo \"%s\" | %s keys add %s --recover --keyring-backend %s --algo eth_secp256k1 --home %s",
		mnemonic,
		filesmanager.GetEvmosdPath(e.Version),
		name,
		e.KeyringBackend,
		e.HomeDir,
	)
	command := exec.Command("bash", "-c", cmd)
	o, err := command.CombinedOutput()
	if strings.Contains(string(o), "duplicated") {
		return fmt.Errorf("duplicated address")
	}
	return err
}

func (e *Evmos) AddValidatorKey() error {
	return e.KeyAdd(e.ValKeyName, e.ValMnemonic)
}

func (e *Evmos) AddGenesisAccount() error {
	validatorAddr, err := e.GetValidatorAddress()
	if err != nil {
		return err
	}
	command := exec.Command( //nolint:gosec
		filesmanager.GetEvmosdPath(e.Version),
		"add-genesis-account",
		validatorAddr,
		e.ValidatorInitialSupply+e.BaseDenom,
		"--keyring-backend",
		e.KeyringBackend,
		"--home",
		e.HomeDir,
	)
	_, err = command.CombinedOutput()
	return err
}

func (e *Evmos) ValidatorGenTx() error {
	command := exec.Command( //nolint:gosec
		filesmanager.GetEvmosdPath(e.Version),
		"gentx",
		e.ValKeyName,
		e.ValidatorInitialSupply[0:len(e.ValidatorInitialSupply)-4]+e.BaseDenom,
		"--gas-prices",
		e.BaseFee+"aevmos",
		"--chain-id",
		e.ChainID,
		"--keyring-backend",
		e.KeyringBackend,
		"--home",
		e.HomeDir,
	)
	_, err := command.CombinedOutput()
	return err
}

func (e *Evmos) CollectGenTxs() error {
	command := exec.Command( //nolint:gosec
		filesmanager.GetEvmosdPath(e.Version),
		"collect-gentxs",
		"--home",
		e.HomeDir,
	)
	_, err := command.CombinedOutput()
	return err
}

func (e *Evmos) ValidateGenesis() error {
	command := exec.Command( //nolint:gosec
		filesmanager.GetEvmosdPath(e.Version),
		"validate-genesis",
		"--home",
		e.HomeDir,
	)
	_, err := command.CombinedOutput()
	return err
}

func (e *Evmos) Start(name string) (int, error) {
	// TODO: do I need the name here?
	_ = name
	logFile := e.HomeDir + "/run.log"

	cmd := fmt.Sprintf("%s start --chain-id %s --home %s --json-rpc.api eth,txpool,personal,net,debug,web3 --api.enable --grpc.enable >> %s 2>&1",
		filesmanager.GetEvmosdPath(e.Version),
		e.ChainID,
		e.HomeDir,
		logFile,
	)
	command := exec.Command("bash", "-c", cmd)

	// Deattach the program
	command.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	err := command.Start()
	if err != nil {
		return 0, err
	}

	// Let evmosd start
	time.Sleep(2 * time.Second)

	id, err := filesmanager.GetChildPID(command.Process.Pid)
	if err != nil {
		return 0, err
	}

	return id, nil
}
