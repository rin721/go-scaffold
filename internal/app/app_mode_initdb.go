package app

// runModeInitDB initdb 模式
func (app *App) runModeInitDB() (*App, error) {
	if err := app.initDatabase(); err != nil {
		return nil, err
	}
	if err := app.initDemoSchema(); err != nil {
		return nil, err
	}
	app.Logger.Info("database schema initialized")
	return app, nil
}
