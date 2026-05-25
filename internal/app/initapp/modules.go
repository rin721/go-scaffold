package initapp

import (
	"fmt"

	demohandler "github.com/rei0721/go-scaffold/internal/modules/demo/handler"
	"github.com/rei0721/go-scaffold/internal/modules/demo/model"
	demorepository "github.com/rei0721/go-scaffold/internal/modules/demo/repository"
	demoservice "github.com/rei0721/go-scaffold/internal/modules/demo/service"
	"github.com/rei0721/go-scaffold/pkg/database"
	"github.com/rei0721/go-scaffold/pkg/logger"
)

func NewModules(core Core, infra Infrastructure) (Modules, error) {
	if err := MigrateDemoSchema(infra.Database, core.Logger); err != nil {
		return Modules{}, err
	}
	return Modules{
		Demo: NewDemoModule(infra.Database, core.Logger),
	}, nil
}

func MigrateDemoSchema(db database.Database, log logger.Logger) error {
	if db == nil {
		return nil
	}
	if err := db.DB().AutoMigrate(&model.Todo{}); err != nil {
		return fmt.Errorf("migrate demo schema: %w", err)
	}
	log.Info("demo schema migrated")
	return nil
}

func NewDemoModule(db database.Database, log logger.Logger) DemoModule {
	todoRepo := demorepository.NewTodoRepository()
	todoService := demoservice.NewTodoService(db, todoRepo)
	todoHandler := demohandler.NewTodoHandler(todoService, log)

	return DemoModule{
		TodoRepository: todoRepo,
		TodoService:    todoService,
		TodoHandler:    todoHandler,
	}
}
