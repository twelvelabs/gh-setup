package main

import (
	"fmt"
	"os"

	"github.com/twelvelabs/gh-setup/internal/cmd"
	"github.com/twelvelabs/gh-setup/internal/core"
)

func main() {
	app, err := core.NewApp()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	command := cmd.NewRootCmd(app)
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
