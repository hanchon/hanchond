package solidity

import (
	"fmt"
	"os/exec"

	"github.com/hanchon/hanchond/playground/filesmanager"
)

func CompileWithSolc(solcVersion, pathToSolidityCode, outputFolder string) error {
	// TODO: build the solc version if not exists
	solcPath := filesmanager.GetSolcPath(solcVersion)
	compileCmd := exec.Command(solcPath, "--optimize", "--abi", "--bin", pathToSolidityCode, "-o", outputFolder, "--allow-paths", filesmanager.GetBuildsDir())
	out, err := compileCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error compiling the contract:%s. %s", err.Error(), string(out))
	}
	return nil
}
