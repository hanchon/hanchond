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
