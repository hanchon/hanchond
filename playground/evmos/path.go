package evmos

import (
	"os/exec"

	"github.com/hanchon/hanchond/playground/filesmanager"
)

func (e *Evmos) getGenesisPath() string {
	return e.HomeDir + "/config/genesis.json"
}

func (e *Evmos) getConfigPath() string {
	return e.HomeDir + "/config/config.toml"
}

func (e *Evmos) getAppPath() string {
	return e.HomeDir + "/config/app.toml"
}

func (e *Evmos) openGenesisFile() (map[string]interface{}, error) {
	return readJSONFile(e.getGenesisPath())
}

func (e *Evmos) saveGenesisFile(genesis map[string]interface{}) error {
	return saveJSONFile(genesis, e.getGenesisPath())
}

func (e *Evmos) openConfigFile() ([]byte, error) {
	return filesmanager.ReadFile(e.getConfigPath())
}

func (e *Evmos) saveConfigFile(configFile []byte) error {
	return filesmanager.SaveFile(configFile, e.getConfigPath())
}

func (e *Evmos) openAppFile() ([]byte, error) {
	return filesmanager.ReadFile(e.getAppPath())
}

func (e *Evmos) saveAppFile(appFile []byte) error {
	return filesmanager.SaveFile(appFile, e.getAppPath())
}

func (e *Evmos) backupConfigFiles() error {
	cmd := exec.Command("cp", e.getAppPath(), e.getAppPath()+".bkup") //nolint:gosec
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	cmd = exec.Command("cp", e.getConfigPath(), e.getConfigPath()+".bkup") //nolint:gosec
	_, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

func (e *Evmos) copyGenesisFile(genesisPath string) error {
	cmd := exec.Command("cp", genesisPath, e.getGenesisPath()) //nolint:gosec
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}
