package playground

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/hanchon/hanchond/playground/solidity"
	"github.com/spf13/cobra"
)

// buildSolcCmd represents the buildSolc command
var buildSolcCmd = &cobra.Command{
	Use:   "build-solc",
	Short: "Build an specific version of Solc",
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

		fmt.Println("Downloading solidity from github:", version)
		if err := solidity.DownloadSolcBinary(isDarwin, version); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("Solc %s is now available\n", version)
	},
}

func init() {
	PlaygroundCmd.AddCommand(buildSolcCmd)
	buildSolcCmd.PersistentFlags().StringP("version", "v", "0.8.0", "Solc version to download")
	buildSolcCmd.PersistentFlags().Bool("is-darwin", true, "Is the system MacOS arm?")
}
