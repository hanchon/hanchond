package gaia

import (
	"fmt"
	"os/exec"

	"github.com/hanchon/hanchond/playground/filesmanager"
)

func GetGaiadBinary(isDarwin bool, version string) error {
	darwinURL := "https://github.com/cosmos/gaia/releases/download/" + version + "/gaiad-" + version + "-darwin-arm64"
	url := " https://github.com/cosmos/gaia/releases/download/" + version + "/gaiad-" + version + "-linux-amd64"
	if isDarwin {
		url = darwinURL
	}
	path := filesmanager.GetGaiadPath()
	cmdString := fmt.Sprintf("wget %s -O %s && chmod +x %s", url, path, path)
	command := exec.Command("bash", "-c", cmdString)
	_, err := command.CombinedOutput()
	return err
}
