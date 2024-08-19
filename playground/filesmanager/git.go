package filesmanager

import (
	"os/exec"
)

func GitCloneEvmosBranch(version string) error {
	return GitCloneBranch(version, GetBranchFolder(version), "https://github.com/evmos/evmos")
}

func GitCloneHermesBranch(version string) error {
	return GitCloneBranch(version, GetBranchFolder(version), "https://github.com/informalsystems/hermes")
}

func GitCloneBranch(version string, dstFolder string, repoURL string) error {
	cmd := exec.Command("git", "clone", "--depth", "1", "--branch", version, repoURL, dstFolder)
	_, err := cmd.Output()
	return err
}
