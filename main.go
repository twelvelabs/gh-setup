package main

import (
	"os"

	"github.com/twelvelabs/gh-setup/internal/cmd"
	"github.com/twelvelabs/gh-setup/internal/core"
)

func main() {
	app := core.NewApp()
	command := cmd.NewRootCmd(app)
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
