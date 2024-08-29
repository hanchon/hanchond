package playground

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/spf13/cobra"
)

// buildHermesCmd represents the buildHermes command
var buildHermesCmd = &cobra.Command{
	Use:   "build-hermes",
	Short: "Build the Hermes relayer binary",
	Long:  `It builds the relayer from source, it accepts a version flag to specify any tag. It defaults to: v1.9.0.`,
	Run: func(cmd *cobra.Command, _ []string) {
		// TODO: download from release page instead of building from source
		_ = filesmanager.SetHomeFolderFromCobraFlags(cmd)
		version, err := cmd.Flags().GetString("version")
		if err != nil {
			fmt.Println("could not read the version:", err.Error())
			os.Exit(1)
		}
		// Clone and build
		if err := filesmanager.CreateTempFolder(version); err != nil {
			fmt.Println("could not create temp folder:" + err.Error())
			os.Exit(1)
		}
		fmt.Println("Cloning hermes version:", version)
		if err := filesmanager.GitCloneHermesBranch(version); err != nil {
			fmt.Println("could not clone the hermes version: ", err)
			os.Exit(1)
		}

		fmt.Println("Building hermes...")
		if err := filesmanager.BuildHermes(version); err != nil {
			fmt.Println("error building hermes:", err.Error())
			os.Exit(1)
		}

		fmt.Println("Moving built binary...")
		if err := filesmanager.SaveHermesBuiltVersion(version); err != nil {
			fmt.Println("could not move the built binary:", err.Error())
			os.Exit(1)
		}

		fmt.Println("Cleaning up...")
		if err := filesmanager.CleanUpTempFolder(); err != nil {
			fmt.Println("could not remove temp folder", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	PlaygroundCmd.AddCommand(buildHermesCmd)
	buildHermesCmd.PersistentFlags().StringP("version", "v", "v1.10.3", "Hermes version to build")
}
