package main

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/cmd"
	"github.com/hanchon/hanchond/cmd/ui"
)

func main() {
	// UI Code
	if len(os.Args) == 2 && os.Args[1] == "ui" {
		p := ui.CreateExplorerTUI()
		if _, err := p.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	// Headless code with cobra-cli
	cmd.Execute()
}
