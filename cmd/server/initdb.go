package main

import (
	"fmt"
	"os"

	"github.com/rei0721/go-scaffold/internal/app"
	"github.com/rei0721/go-scaffold/pkg/cli"
	"github.com/rei0721/go-scaffold/types/constants"
)

type InitdbCommand struct{}

func (c *InitdbCommand) Name() string {
	return constants.AppInitDBCommandName
}

func (c *InitdbCommand) Description() string {
	return "Initialize database schema and data"
}

func (c *InitdbCommand) Usage() string {
	return fmt.Sprintf("%s [--config=<path>]", constants.AppInitDBCommandName)
}

func (c *InitdbCommand) Flags() []cli.Flag {
	return []cli.Flag{
		{
			Name:        "config",
			ShortName:   "c",
			Type:        cli.FlagTypeString,
			Required:    false,
			Default:     constants.AppDefaultConfigPath,
			Description: "Config file path",
			EnvVar:      "REI_CONFIG_PATH",
		},
	}
}

func (c *InitdbCommand) Execute(ctx *cli.Context) error {
	application, err := app.New(app.Options{
		ConfigPath: ctx.GetString("config"),
		Mode:       app.ModeInitDB,
	})
	if err != nil {
		os.Stderr.WriteString("failed to initialize application: " + err.Error() + "\n")
		return err
	}

	defer func() {
		if application.Infra.Database != nil {
			_ = application.Infra.Database.Close()
		}
		if application.Core.Logger != nil {
			_ = application.Core.Logger.Sync()
		}
	}()

	return nil
}
