package gaia

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"

	"github.com/hanchon/hanchond/playground/filesmanager"
)

func GetGaiadBinary(isDarwin bool, version string) error {
	arch := runtime.GOARCH
	if arch != "arm64" {
		arch = "amd64"
	}
	systemOS := "darwin"
	if !isDarwin {
		systemOS = "linux"
	}

	url := fmt.Sprintf("https://github.com/cosmos/gaia/releases/download/%s/gaiad-%s-%s-%s", version, version, systemOS, arch)

	path := filesmanager.GetGaiadPath()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
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

	err = os.Chmod(path, 0o755)
	if err != nil {
		return fmt.Errorf("failed to set file permissions: %s", err)
	}

	return nil
}
