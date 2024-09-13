package gaia

import (
	"fmt"
	"github.com/hanchon/hanchond/playground/filesmanager"
	"io"
	"net/http"
	"os"
	"runtime"
)

func main() {
	err := GetGaiadBinary(false, "v19.2.0")
	if err != nil {
		return
	}
}

func GetGaiadBinary(isDarwin bool, version string) error {
	url := "https://github.com/cosmos/gaia/releases/download/" + version + "/gaiad-" + version + "-linux-amd64"
	arch := runtime.GOARCH
	if isDarwin {
		if arch == "arm64" {
			url = "https://github.com/cosmos/gaia/releases/download/" + version + "/gaiad-" + version + "-darwin-arm64"
		} else {
			url = "https://github.com/cosmos/gaia/releases/download/" + version + "/gaiad-" + version + "-darwin-amd64"
		}
	} else {
		if arch == "arm64" {
			url = "https://github.com/cosmos/gaia/releases/download/" + version + "/gaiad-" + version + "-linux-arm64"
		} else {
			url = "https://github.com/cosmos/gaia/releases/download/" + version + "/gaiad-" + version + "-linux-amd64"
		}
	}
	path := filesmanager.GetGaiadPath()
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download Gaia: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download gaiad binary: status code %d", resp.StatusCode)
	}

	outFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %s", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save gaiad binary: %s", err)
	}

	err = os.Chmod(path, 0755)
	if err != nil {
		return fmt.Errorf("failed to set file permissions: %s", err)
	}

	return nil
}
