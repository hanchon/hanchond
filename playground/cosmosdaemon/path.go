package cosmosdaemon

import (
	"os/exec"

	"github.com/hanchon/hanchond/playground/filesmanager"
)

func (d *Daemon) getGenesisPath() string {
	return d.HomeDir + "/config/genesis.json"
}

func (d *Daemon) getConfigPath() string {
	return d.HomeDir + "/config/config.toml"
}

func (d *Daemon) getAppPath() string {
	return d.HomeDir + "/config/app.toml"
}

func (d *Daemon) openGenesisFile() (map[string]interface{}, error) {
	return readJSONFile(d.getGenesisPath())
}

func (d *Daemon) saveGenesisFile(genesis map[string]interface{}) error {
	return saveJSONFile(genesis, d.getGenesisPath())
}

func (d *Daemon) openConfigFile() ([]byte, error) {
	return filesmanager.ReadFile(d.getConfigPath())
}

func (d *Daemon) saveConfigFile(configFile []byte) error {
	return filesmanager.SaveFile(configFile, d.getConfigPath())
}

func (d *Daemon) openAppFile() ([]byte, error) {
	return filesmanager.ReadFile(d.getAppPath())
}

func (d *Daemon) saveAppFile(appFile []byte) error {
	return filesmanager.SaveFile(appFile, d.getAppPath())
}

func (d *Daemon) backupConfigFiles() error {
	cmd := exec.Command("cp", d.getAppPath(), d.getAppPath()+".bkup") //nolint:gosec
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	cmd = exec.Command("cp", d.getConfigPath(), d.getConfigPath()+".bkup") //nolint:gosec
	_, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

// func (d *Daemon) copyGenesisFile(genesisPath string) error {
// 	cmd := exec.Command("cp", genesisPath, d.getGenesisPath()) //nolint:gosec
// 	_, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
