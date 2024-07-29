package playground

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/spf13/cobra"
)

const LocalVesrsion = "local"

// buildEvmosCmd represents the buildEvmos command
var buildEvmosCmd = &cobra.Command{
	Use:   "build-evmos",
	Short: "Build an specific version of Evmos (hanchond playground build-evmos v18.0.0), it also supports local repositories (hanchond playground build-evmos --path /home/hanchon/evmos)",
	Long:  `It downloads, builds and clean up temp files for any Evmos tag. Using the --path flag will build you local repo`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = filesmanager.SetHomeFolderFromCobraFlags(cmd)

		path, err := cmd.Flags().GetString("path")
		// Local build
		if err == nil && path != "" {
			version := LocalVesrsion
			if path[len(path)-1] == '/' {
				path = path[0 : len(path)-2]
			}
			fmt.Println("Building evmos...")
			if err := filesmanager.BuildEvmos(path); err != nil {
				fmt.Println("error building evmos:", err.Error())
				os.Exit(1)
			}
			fmt.Println("Moving built binary...")
			if err := filesmanager.CopyFile(path+"/build/evmosd", filesmanager.GetEvmosdPath(version)); err != nil {
				fmt.Println("could not move the built binary:", err.Error())
				os.Exit(1)
			}
			os.Exit(0)
		}

		// Clone from github
		if len(args) == 0 {
			fmt.Println("version is missing. Usage: hanchond playground build-evmosd v18.1.0")
			os.Exit(1)
		}
		version := args[0]
		if err := filesmanager.CreateTempFolder(version); err != nil {
			fmt.Println("could not create temp folder:" + err.Error())
			os.Exit(1)
		}
		fmt.Println("Cloning evmos version:", version)
		if err := filesmanager.GitCloneEvmosBranch(version); err != nil {
			fmt.Println("could not clone the evmos version: ", err)
			os.Exit(1)
		}
		fmt.Println("Building evmos...")
		if err := filesmanager.BuildEvmosVersion(version); err != nil {
			fmt.Println("error building evmos:", err)
			os.Exit(1)
		}
		fmt.Println("Moving built binary...")
		if err := filesmanager.SaveEvmosBuiltVersion(version); err != nil {
			fmt.Println("could not move the built binary:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Cleaning up...")
		if err := filesmanager.CleanUpTempFolder(); err != nil {
			fmt.Println("could not remove temp folder", err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	},
}

func init() {
	PlaygroundCmd.AddCommand(buildEvmosCmd)
	buildEvmosCmd.Flags().StringP("path", "p", "", "Path to you local clone of Evmos")
}
