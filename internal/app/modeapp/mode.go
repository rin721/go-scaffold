package modeapp

import (
	"github.com/rei0721/go-scaffold/internal/app/initapp"
	"github.com/rei0721/go-scaffold/internal/config"
)

type Mode string

const (
	ModeServer Mode = "server"
	ModeInitDB Mode = "initdb"
)

type ConfigChangeHandler = config.HookHandler

type BuildResult struct {
	Infra     initapp.Infrastructure
	Modules   initapp.Modules
	Transport initapp.Transport
}

func Build(mode Mode, core initapp.Core, onConfigChange ConfigChangeHandler) (BuildResult, error) {
	switch mode {
	case ModeInitDB:
		return BuildInitDB(core)
	case ModeServer:
		return BuildServer(core, onConfigChange)
	default:
		return BuildServer(core, onConfigChange)
	}
}

func BuildServer(core initapp.Core, onConfigChange ConfigChangeHandler) (BuildResult, error) {
	infra, err := initapp.NewInfrastructure(core)
	if err != nil {
		return BuildResult{}, err
	}

	modules, err := initapp.NewModules(core, infra)
	if err != nil {
		return BuildResult{}, err
	}

	transport, err := initapp.NewTransport(core, infra, modules)
	if err != nil {
		return BuildResult{}, err
	}

	WatchConfig(core, onConfigChange)
	core.Logger.Info("application initialized successfully")

	return BuildResult{
		Infra:     infra,
		Modules:   modules,
		Transport: transport,
	}, nil
}

func BuildInitDB(core initapp.Core) (BuildResult, error) {
	db, err := initapp.NewDatabase(core.Config)
	if err != nil {
		return BuildResult{}, err
	}

	infra := initapp.Infrastructure{Database: db}
	if _, err := initapp.MigrateDemoSchemaForTrigger(infra.Database, core.Logger, initapp.DemoMigrationTriggerInitDB); err != nil {
		return BuildResult{}, err
	}

	core.Logger.Info("database schema initialized")
	return BuildResult{Infra: infra}, nil
}

func WatchConfig(core initapp.Core, onConfigChange ConfigChangeHandler) {
	if err := core.ConfigManager.Watch(); err != nil {
		core.Logger.Warn("failed to start config watcher", "error", err)
		return
	}
	core.Logger.Debug("config watcher started")

	if onConfigChange != nil {
		core.ConfigManager.RegisterHook(onConfigChange)
	}
}
