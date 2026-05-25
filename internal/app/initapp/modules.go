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
	if _, err := MigrateDemoSchemaForTrigger(infra.Database, core.Logger, DemoMigrationTriggerServerStart); err != nil {
		return Modules{}, err
	}
	return Modules{
		Demo: NewDemoModule(infra.Database, core.Logger),
	}, nil
}

type DemoMigrationTrigger string

const (
	DemoMigrationTriggerServerStart DemoMigrationTrigger = "server-start"
	DemoMigrationTriggerInitDB      DemoMigrationTrigger = "initdb"
	DemoMigrationTriggerReload      DemoMigrationTrigger = "reload"
)

type DemoMigrationPolicy struct {
	Trigger     DemoMigrationTrigger
	AutoMigrate bool
	Reason      string
}

func DemoMigrationPolicyFor(trigger DemoMigrationTrigger) DemoMigrationPolicy {
	switch trigger {
	case DemoMigrationTriggerServerStart:
		return DemoMigrationPolicy{
			Trigger:     trigger,
			AutoMigrate: true,
			Reason:      "demo server startup keeps the local development schema ready",
		}
	case DemoMigrationTriggerInitDB:
		return DemoMigrationPolicy{
			Trigger:     trigger,
			AutoMigrate: true,
			Reason:      "initdb is the explicit demo bootstrap command",
		}
	case DemoMigrationTriggerReload:
		return DemoMigrationPolicy{
			Trigger:     trigger,
			AutoMigrate: false,
			Reason:      "database reload must not perform implicit schema changes",
		}
	default:
		return DemoMigrationPolicy{
			Trigger:     trigger,
			AutoMigrate: false,
			Reason:      "unknown migration trigger requires an explicit policy",
		}
	}
}

func MigrateDemoSchema(db database.Database, log logger.Logger) error {
	_, err := MigrateDemoSchemaForTrigger(db, log, DemoMigrationTriggerServerStart)
	return err
}

func MigrateDemoSchemaForTrigger(db database.Database, log logger.Logger, trigger DemoMigrationTrigger) (DemoMigrationPolicy, error) {
	policy := DemoMigrationPolicyFor(trigger)
	if !policy.AutoMigrate {
		logDemoMigrationSkipped(log, policy)
		return policy, nil
	}
	if db == nil {
		return policy, nil
	}
	if err := db.DB().AutoMigrate(&model.Todo{}); err != nil {
		return policy, fmt.Errorf("migrate demo schema: %w", err)
	}
	logDemoMigrationApplied(log, policy)
	return policy, nil
}

func logDemoMigrationApplied(log logger.Logger, policy DemoMigrationPolicy) {
	if log == nil {
		return
	}
	log.Info("demo schema migrated", "trigger", policy.Trigger, "reason", policy.Reason)
}

func logDemoMigrationSkipped(log logger.Logger, policy DemoMigrationPolicy) {
	if log == nil {
		return
	}
	log.Debug("demo schema migration skipped", "trigger", policy.Trigger, "reason", policy.Reason)
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
