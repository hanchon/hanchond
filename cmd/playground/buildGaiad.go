package playground

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/hanchon/hanchond/playground/gaia"
	"github.com/spf13/cobra"
)

// buildGaiadCmd represents the buildGaiad command
var buildGaiadCmd = &cobra.Command{
	Use:   "build-gaiad",
	Short: "Get the Gaiad binary from the github releases",
	Long:  `It downloads the already built gaiad binary from github, it accepts a version flag to specify any tag. It defaults to: v1.9.0.`,
	Run: func(cmd *cobra.Command, _ []string) {
		_ = filesmanager.SetHomeFolderFromCobraFlags(cmd)
		version, err := cmd.Flags().GetString("version")
		if err != nil {
			fmt.Println("could not read the version:", err.Error())
			os.Exit(1)
		}

		isDarwin, err := cmd.Flags().GetBool("is-darwin")
		if err != nil {
			fmt.Println("could not read the isDarwin:", err.Error())
			os.Exit(1)
		}

		// Create build folder if needed
		if err := filesmanager.CreateBuildsDir(); err != nil {
			fmt.Println("could not create build folder:" + err.Error())
			os.Exit(1)
		}

		fmt.Println("Downloading gaiad from github:", version)
		if err = gaia.GetGaiadBinary(isDarwin, version); err != nil {
			fmt.Println("could not get gaiad from github:" + err.Error())
			os.Exit(1)
		}
		fmt.Println("Gaiad is now available")
	},
}

func init() {
	PlaygroundCmd.AddCommand(buildGaiadCmd)
	buildGaiadCmd.PersistentFlags().StringP("version", "v", "v18.1.0", "Gaiad version to download")
	buildGaiadCmd.PersistentFlags().Bool("is-darwin", true, "Is the system MacOS arm?")
}
