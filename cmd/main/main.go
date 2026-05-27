// Package main is the command entry point for the scaffold service.
package main

import (
	"fmt"
	"os"

	"github.com/rei0721/go-scaffold/pkg/cli"
	"github.com/rei0721/go-scaffold/types/constants"
)

func main() {
	app := cli.NewApp(constants.AppName)
	app.SetVersion(constants.AppVersion)
	app.SetDescription(constants.AppDescription)

	app.AddCommand(NewAppCommand())
	app.AddCommand(NewDBCommand())

	if err := app.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(cli.GetExitCode(err))
	}
}
