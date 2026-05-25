// Package app is the composition root of the service.
//
// It follows an onion-style dependency direction:
//
//	Core -> Infrastructure -> Modules -> Transport
//
// Inner rings know nothing about outer rings. This package is the only place
// that wires all rings together.
package app

import (
	"context"
	"fmt"

	"github.com/rei0721/go-scaffold/internal/app/initapp"
	"github.com/rei0721/go-scaffold/internal/app/lifecycleapp"
	"github.com/rei0721/go-scaffold/internal/app/modeapp"
	"github.com/rei0721/go-scaffold/internal/app/reloadapp"
	"github.com/rei0721/go-scaffold/internal/config"
	"github.com/rei0721/go-scaffold/pkg/logger"
)

type App struct {
	Options Options

	Core      CoreLayer
	Infra     InfrastructureLayer
	Modules   ModulesLayer
	Transport TransportLayer
}

type Options struct {
	ConfigPath string
	Mode       AppMode
}

func New(opts Options) (*App, error) {
	if opts.Mode == "" {
		opts.Mode = ModeServer
	}

	app := &App{Options: opts}
	core, err := initapp.NewCore(opts.ConfigPath)
	if err != nil {
		return nil, fmt.Errorf("initialize core: %w", err)
	}
	app.Core = core

	app.Core.ConfigManager.RegisterLogger(func() logger.Logger {
		return app.Core.Logger
	})

	initapp.DebugConfig(app.Core, opts.ConfigPath)

	result, err := modeapp.Build(opts.Mode, app.Core, app.handleConfigChange)
	if err != nil {
		return nil, err
	}
	app.Infra = result.Infra
	app.Modules = result.Modules
	app.Transport = result.Transport

	return app, nil
}

func (a *App) Run() error {
	return a.Start(context.Background())
}

func (a *App) Start(ctx context.Context) error {
	return lifecycleapp.Start(ctx, a.Transport)
}

func (a *App) Shutdown(ctx context.Context) error {
	return lifecycleapp.Shutdown(ctx, a.Core, a.Infra, a.Transport)
}

func (a *App) handleConfigChange(old, new *config.Config) {
	a.Core.Logger.Info("configuration file changed, processing updates")
	reloadapp.Reload(&a.Core, &a.Infra, &a.Transport, old, new)
	a.Core.Config = new
	a.Core.Logger.Info("configuration update completed")
}
