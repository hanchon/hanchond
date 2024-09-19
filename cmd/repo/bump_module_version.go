package repo

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/hanchon/hanchond/playground/filesmanager"
	"github.com/spf13/cobra"
)

// BumpModuleVersionCmd represents the query command
var BumpModuleVersionCmd = &cobra.Command{
	Use:   "bump-module-version [path] [version]",
	Args:  cobra.ExactArgs(2),
	Short: "Bump the version of a go module, i.e., hanchond repo bump-version /tmp/repo v21",
	Run: func(_ *cobra.Command, args []string) {
		path := args[0]
		// Make sure that the path does not end with `/`
		path = strings.TrimSuffix(path, "/")

		version := args[1]
		// If the version arg is missing the `v` prefix, we add it
		if !strings.HasPrefix(version, "v") {
			version = fmt.Sprintf("v%s", version)
		}

		// Find the current version
		goModPath := fmt.Sprintf("%s/go.mod", path)
		fmt.Println("using go.mod path as:", goModPath)
		goModFile, err := filesmanager.ReadFile(goModPath)
		if err != nil {
			fmt.Println("error reading the go.mod file:", err.Error())
			os.Exit(1)
		}

		// Get the current version
		re := regexp.MustCompile(`(?m)^module\s+(\S+)$`)
		modules := re.FindAllStringSubmatch(string(goModFile), -1)
		if len(modules) == 0 {
			fmt.Println("the go.mod file does not define the module name")
			os.Exit(1)
		}
		currentVersion := modules[0][1]
		fmt.Println("the current version is:", currentVersion)

		// Create the new version with all the parts of the currentVersion but overwritting the last segment
		newVersion := ""
		parts := strings.Split(currentVersion, "/")
		for k, v := range parts {
			if k == len(parts)-1 {
				newVersion += version
				break
			}
			newVersion += v + "/"
		}

		fmt.Println("the new version is:", newVersion)

		// Walk through the root directory recursively
		if err = filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				fmt.Println("error reading the directory", path)
				os.Exit(1)
			}

			// Only process regular files
			if info.IsDir() {
				return nil
			}

			// Read the file
			content, err := filesmanager.ReadFile(path)
			if err != nil {
				fmt.Println("failed reading the file:", path)
				os.Exit(1)
			}

			fileContent := string(content)

			// Replace all occurrences
			re := regexp.MustCompile(regexp.QuoteMeta(currentVersion))
			updatedContent := re.ReplaceAllString(fileContent, newVersion)

			// Only write if the file was modified
			if updatedContent != fileContent {
				fmt.Printf("updating file: %s\n", path)
				err := filesmanager.SaveFileWithMode([]byte(updatedContent), path, info.Mode())
				if err != nil {
					fmt.Println("failed saving the file:", path)
					os.Exit(1)
				}
			}

			return nil
		}); err != nil {
			fmt.Println("error walking the directory:", err.Error())
			os.Exit(1)
		}

		os.Exit(0)
	},
}
