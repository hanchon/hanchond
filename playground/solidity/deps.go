package solidity

import (
	"fmt"

	"github.com/hanchon/hanchond/playground/filesmanager"
)

func DownloadDep(repoURL, version, name string) (string, error) {
	path := filesmanager.GetDepsDir(name)
	// Already downloaded
	if filesmanager.DoesFileExist(path) {
		return path, nil
	}
	if err := filesmanager.CreateDepsFolder(); err != nil {
		return "", fmt.Errorf("could not create deps folder:%s", err.Error())
	}
	if err := filesmanager.GitCloneBranch(version, path, repoURL); err != nil {
		return "", fmt.Errorf("could not clone the repo:%s", err.Error())
	}

	return path, nil
}

func DownloadUniswapV2Core() (string, error) {
	return DownloadDep("https://github.com/Uniswap/uniswap-v2-core", "master", "uniswapv2")
}

func DownloadUniswapV2Periphery() (string, error) {
	return DownloadDep("https://github.com/Uniswap/v2-periphery", "master", "v2-periphery")
}

func DownloadUniswapV2Minified() (string, error) {
	return DownloadDep("https://github.com/casweeney/minified-uniswapv2-contracts", "main", "v2-minified")
}
