package main

import (
	"fmt"

	appconfig "github.com/rei0721/go-scaffold/internal/config"
	"github.com/rei0721/go-scaffold/pkg/cli"
	"github.com/rei0721/go-scaffold/types/constants"
)

type AppCommand struct{}

func NewAppCommand() *AppCommand {
	return &AppCommand{}
}

func (c *AppCommand) Name() string {
	return constants.AppServerCommandName
}

func (c *AppCommand) Description() string {
	return "Run server"
}

func (c *AppCommand) Usage() string {
	return fmt.Sprintf("%s [--config=<name>]", constants.AppServerCommandName)
}

func (c *AppCommand) Flags() []cli.Flag {
	return []cli.Flag{
		{
			Name:        "config",
			ShortName:   "c",
			Type:        cli.FlagTypeString,
			Required:    false,
			Default:     constants.AppDefaultConfigPath,
			Description: "Config file path",
			EnvVar:      appconfig.EnvConfigPathName(),
		},
	}
}

func (c *AppCommand) Execute(ctx *cli.Context) error {
	runApp(ctx.GetString("config"))
	return nil
}
