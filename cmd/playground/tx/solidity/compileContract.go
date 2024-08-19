package solidity

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/hanchon/hanchond/playground/solidity"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// compileContractCmd represents the compile command
var compileContractCmd = &cobra.Command{
	Use:     "compile-contract [path_to_solidity_file]",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"c"},
	Short:   "Compile a solidity contract",
	Run: func(cmd *cobra.Command, args []string) {
		_ = sql.InitDBFromCmd(cmd)

		outputFolder, err := cmd.Flags().GetString("output-folder")
		if err != nil {
			fmt.Println("incorrect output folder")
			os.Exit(1)
		}
		if outputFolder[len(outputFolder)-1] != '/' {
			outputFolder += "/"
		}

		// TODO: read from pragma the correct version and use it automatically
		solcVersion, err := cmd.Flags().GetString("solc-version")
		if err != nil {
			fmt.Println("incorrect solc version")
			os.Exit(1)
		}

		pathToSolidityCode := args[0]

		if err := filesmanager.CleanUpTempFolder(); err != nil {
			fmt.Printf("could not clean up temp folder:%s\n", err.Error())
			os.Exit(1)
		}

		folderName := "compiler"
		if err := filesmanager.CreateTempFolder(folderName); err != nil {
			fmt.Printf("could not create up temp folder:%s\n", err.Error())
			os.Exit(1)
		}

		err = solidity.CompileWithSolc(solcVersion, pathToSolidityCode, filesmanager.GetBranchFolder(folderName))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if err := moveFiles(filesmanager.GetBranchFolder(folderName), outputFolder, "abi"); err != nil {
			fmt.Printf("error copying the built files: %s\n", err.Error())
			os.Exit(1)
		}

		if err := moveFiles(filesmanager.GetBranchFolder(folderName), outputFolder, "bin"); err != nil {
			fmt.Printf("error copying the built files: %s\n", err.Error())
			os.Exit(1)
		}

		if err := filesmanager.CleanUpTempFolder(); err != nil {
			fmt.Printf("could not clean up temp folder:%s\n", err.Error())
			os.Exit(1)
		}

		fmt.Printf("Contract compiled at %s\n", outputFolder)
	},
}

func init() {
	SolidityCmd.AddCommand(compileContractCmd)
	compileContractCmd.Flags().StringP("output-folder", "o", "./", "Output folder where the compile code will be saved")
	compileContractCmd.Flags().StringP("solc-version", "v", "0.8.0", "Solc version used to compile the code")
}

func moveFiles(in, out, extension string) error {
	files, err := filepath.Glob(in + "/*." + extension)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return err
	}

	for _, v := range files {
		if err := filesmanager.CopyFile(
			v,
			out,
		); err != nil {
			return err
		}
	}

	return nil
}
