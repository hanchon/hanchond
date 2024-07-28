package filesmanager

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var baseDir = "/tmp"

func SetBaseDir(path string) {
	baseDir = path
}

func SetHomeFolderFromCobraFlags(cmd *cobra.Command) {
	home, err := cmd.Flags().GetString("home")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	home, _ = strings.CutSuffix(home, "/")
	SetBaseDir(home)
}

func GetBaseDir() string {
	return baseDir
}

func GetBuildsDir() string {
	return baseDir + "/evmos_build"
}

func GetTempDir() string {
	return baseDir + "/temp"
}

func GetBranchFolder(version string) string {
	return GetTempDir() + "/" + version
}

func GetEvmosdPath(version string) string {
	return GetBuildsDir() + "/evmosd" + version
}

func GetHermesBinary() string {
	return GetBuildsDir() + "/hermes"
}

func GetHermesPath() string {
	return GetTempDir() + "/hermes"
}

func CreateBuildsDir() error {
	if _, err := os.Stat(GetBuildsDir()); os.IsNotExist(err) {
		return os.Mkdir(GetBuildsDir(), os.ModePerm)
	}
	return nil
}

func CreateTempFolder(version string) error {
	return os.MkdirAll(GetBranchFolder(version), os.ModePerm)
}

func CreateHermesFolder() error {
	return os.MkdirAll(GetHermesPath(), os.ModePerm)
}

func CleanUpTempFolder() error {
	return os.RemoveAll(GetTempDir())
}
