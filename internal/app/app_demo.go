package app

import "github.com/rei0721/go-scaffold/internal/modules/demo/model"

func (app *App) initDemoSchema() error {
	if app.DB == nil {
		return nil
	}
	if err := app.DB.DB().AutoMigrate(&model.Todo{}); err != nil {
		return err
	}
	app.Logger.Info("demo schema migrated")
	return nil
}
