package solidity

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// compileContractCmd represents the compile command
var compileContractCmd = &cobra.Command{
	Use:     "compile-contract [path_to_solidity_file]",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"c"},
	Short:   "Compile a solidity contract",
	// Long:    "The bytecode file must have the following format: {\"bytecode\":\"60806...\",...}",
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)
		nodeID, err := cmd.Flags().GetString("node")
		if err != nil {
			fmt.Println("node not set")
			os.Exit(1)
		}
		_ = queries
		_ = nodeID

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
		paths := strings.Split(pathToSolidityCode, "/")
		fileName := strings.Split(paths[len(paths)-1], ".sol")[0]

		solcPath := filesmanager.GetSolcPath(solcVersion)

		compileCmd := exec.Command(solcPath, "--optimize", "--combined-json", "abi,bin", pathToSolidityCode, "-o", outputFolder)
		out, err := compileCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("error compiling the contract:%s. %s\n", err.Error(), string(out))
			os.Exit(1)
		}

		if err := os.Rename(outputFolder+"combined.json", outputFolder+fileName+".json"); err != nil {
			fmt.Printf("error copying the built file: %s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Contract compiled at %s%s.json\n", outputFolder, fileName)
	},
}

func init() {
	SolidityCmd.AddCommand(compileContractCmd)
	compileContractCmd.Flags().StringP("output-folder", "o", "./", "Output folder where the compile code will be saved")
	compileContractCmd.Flags().StringP("solc-version", "v", "0.8.0", "Solc version used to compile the code")
}
