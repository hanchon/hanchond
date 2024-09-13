package gaia

import (
	"fmt"
	"github.com/hanchon/hanchond/playground/filesmanager"
	"io"
	"net/http"
	"os"
)

func GetGaiadBinary(isDarwin bool, version string) error {
	darwinURL := "https://github.com/cosmos/gaia/releases/download/" + version + "/gaiad-" + version + "-darwin-arm64"
	url := " https://github.com/cosmos/gaia/releases/download/" + version + "/gaiad-" + version + "-linux-amd64"
	if isDarwin {
		url = darwinURL
	}
	path := filesmanager.GetGaiadPath()
	//cmdString := fmt.Sprintf("wget %s -O %s && chmod +x %s", url, path, path)
	//command := exec.Command("bash", "-c", cmdString)
	//_, err := command.CombinedOutput()

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download Gaia: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download gaiad binary: status code %d", resp.StatusCode)
	}

	outFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save gaiad binary: %w", err)
	}

	err = os.Chmod(path, 0755)
	if err != nil {
		return fmt.Errorf("failed to set file permissions: %v", err)
	}

	return nil
}
