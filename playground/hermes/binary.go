package hermes

import (
	"fmt"
	"os/exec"
	"syscall"
	"time"

	"github.com/hanchon/hanchond/playground/filesmanager"
)

func (h *Hermes) GetHermesBinary() string {
	return filesmanager.GetHermesBinary()
}

func (h *Hermes) AddRelayerKey(chainID string, mnemonic string, isEthWallet bool) error {
	hdPath := " "
	if isEthWallet {
		hdPath = " --hd-path \"m/44'/60'/0'/0/0\" "
	}
	cmd := fmt.Sprintf(
		"echo \"%s\" | %s --config %s keys add %s --mnemonic-file /dev/stdin --chain %s >> %s 2>&1",
		mnemonic,
		h.GetHermesBinary(),
		h.GetConfigFile(),
		hdPath,
		chainID,
		filesmanager.GetHermesPath()+"/logs_keys"+chainID,
	)
	command := exec.Command("bash", "-c", cmd)
	_, err := command.CombinedOutput()
	return err
}

func (h *Hermes) CreateChannel(firstChainID, secondChainID string) error {
	cmd := fmt.Sprintf(
		"%s --config %s create channel --a-chain %s --b-chain %s --a-port transfer --b-port transfer --new-client-connection --yes >> %s 2>&1",
		h.GetHermesBinary(),
		h.GetConfigFile(),
		firstChainID,
		secondChainID,
		filesmanager.GetHermesPath()+"/logs_channel"+firstChainID+secondChainID,
	)
	command := exec.Command("bash", "-c", cmd)
	out, err := command.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("error %s: %s", err.Error(), string(out))
	}
	return err
}

func (h *Hermes) Start() (int, error) {
	cmd := fmt.Sprintf(
		"%s --config %s start >> %s 2>&1",
		h.GetHermesBinary(),
		h.GetConfigFile(),
		filesmanager.GetHermesPath()+"/run.log",
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

	// Let hermes start
	time.Sleep(2 * time.Second)

	id, err := filesmanager.GetChildPID(command.Process.Pid)
	if err != nil {
		return 0, err
	}

	return id, nil
}
