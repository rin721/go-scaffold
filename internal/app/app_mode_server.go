package app

import "github.com/rei0721/go-scaffold/internal/config"

func (app *App) runModeServer() (*App, error) {
	if err := app.initDatabase(); err != nil {
		return nil, err
	}
	if err := app.initDemoSchema(); err != nil {
		return nil, err
	}
	if err := app.initCache(); err != nil {
		return nil, err
	}
	if err := app.initExecutor(); err != nil {
		return nil, err
	}
	if err := app.initCORS(); err != nil {
		return nil, err
	}
	if err := app.initHTTPServer(); err != nil {
		return nil, err
	}

	// Start config file watching for hot-reload
	if err := app.ConfigManager.Watch(); err != nil {
		app.Logger.Warn("failed to start config watcher", "error", err)
	}
	app.Logger.Debug("config watcher started")

	// Register config change hook
	// 当配置文件变化时自动调用
	app.ConfigManager.RegisterHook(func(old, new *config.Config) {
		app.Logger.Info("configuration file changed, processing updates...")

		// 重载 app
		app.reload(old, new)

		// 更新应用配置引用
		app.Config = new
		app.Logger.Info("configuration update completed")
	})

	app.Logger.Info("application initialized successfully")
	return app, nil
}
