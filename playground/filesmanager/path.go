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

func SetHomeFolderFromCobraFlags(cmd *cobra.Command) string {
	home, err := cmd.Flags().GetString("home")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	home, _ = strings.CutSuffix(home, "/")
	SetBaseDir(home)
	// Ensure that the folder exists
	if _, err := os.Stat(home); os.IsNotExist(err) {
		if err := os.Mkdir(home, os.ModePerm); err != nil {
			// We panic here because if we can not create the folder we should inmediately stop
			panic(err)
		}
	}
	return home
}

func GetDatabaseFile() string {
	return fmt.Sprintf("%s/playground.db", GetBaseDir())
}

func GetDataFolder() string {
	return fmt.Sprintf("%s/data", GetBaseDir())
}

func getNodeHomePath(chainID int64, nodeID int64) string {
	return fmt.Sprintf("%s/%d-%d", GetDataFolder(), chainID, nodeID)
}

func GetNodeHomeFolder(chainID, nodeID int64) string {
	if _, err := os.Stat(GetDataFolder()); os.IsNotExist(err) {
		if err := os.Mkdir(GetDataFolder(), os.ModePerm); err != nil {
			// We panic here because if we can not create the folder we should inmediately stop
			panic(err)
		}
	}
	return getNodeHomePath(chainID, nodeID)
}

func IsNodeHomeFolderInitialized(chainID int64, nodeID int64) bool {
	return DoesFileExist(getNodeHomePath(chainID, nodeID))
}

func GetBaseDir() string {
	return baseDir
}

func GetBuildsDir() string {
	return baseDir + "/builds"
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

func GetDaemondPath(binaryName string) string {
	return GetBuildsDir() + "/" + binaryName
}

func GetGaiadPath() string {
	return GetBuildsDir() + "/gaiad"
}

func DoesEvmosdPathExist(version string) bool {
	return DoesFileExist(GetBuildsDir() + "/evmosd" + version)
}

func GetHermesBinary() string {
	return GetBuildsDir() + "/hermes"
}

func GetHermesPath() string {
	return GetDataFolder() + "/hermes"
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

func CleanUpData() error {
	_ = os.RemoveAll(GetDatabaseFile())
	return os.RemoveAll(GetDataFolder())
}

func GetSolcPath(version string) string {
	return GetBuildsDir() + "/solc" + version
}
