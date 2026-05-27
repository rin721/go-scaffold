package initapp

import (
	"context"
	"fmt"

	"github.com/rei0721/go-scaffold/internal/app/dbapp"
	demohandler "github.com/rei0721/go-scaffold/internal/modules/demo/handler"
	demorepository "github.com/rei0721/go-scaffold/internal/modules/demo/repository"
	demoservice "github.com/rei0721/go-scaffold/internal/modules/demo/service"
	"github.com/rei0721/go-scaffold/pkg/database"
	"github.com/rei0721/go-scaffold/pkg/logger"
)

func NewModules(core Core, infra Infrastructure) (Modules, error) {
	if _, err := ApplyDemoSchemaForTrigger(infra.Database, core.Config.Database.Driver, core.Logger, DemoSchemaTriggerServerStart); err != nil {
		return Modules{}, err
	}
	return Modules{
		Demo: NewDemoModule(infra.Database, core.Logger),
	}, nil
}

type DemoSchemaTrigger string

const (
	DemoSchemaTriggerServerStart DemoSchemaTrigger = "server-start"
	DemoSchemaTriggerReload      DemoSchemaTrigger = "reload"
)

type DemoSchemaPolicy struct {
	Trigger DemoSchemaTrigger
	Apply   bool
	Reason  string
}

func DemoSchemaPolicyFor(trigger DemoSchemaTrigger) DemoSchemaPolicy {
	switch trigger {
	case DemoSchemaTriggerServerStart:
		return DemoSchemaPolicy{
			Trigger: trigger,
			Apply:   true,
			Reason:  "demo server startup keeps the local development schema ready through sqlgen",
		}
	case DemoSchemaTriggerReload:
		return DemoSchemaPolicy{
			Trigger: trigger,
			Apply:   false,
			Reason:  "database reload must not perform implicit schema changes",
		}
	default:
		return DemoSchemaPolicy{
			Trigger: trigger,
			Apply:   false,
			Reason:  "unknown schema trigger requires an explicit policy",
		}
	}
}

func ApplyDemoSchema(db database.Database, driver string, log logger.Logger) error {
	_, err := ApplyDemoSchemaForTrigger(db, driver, log, DemoSchemaTriggerServerStart)
	return err
}

func ApplyDemoSchemaForTrigger(db database.Database, driver string, log logger.Logger, trigger DemoSchemaTrigger) (DemoSchemaPolicy, error) {
	policy := DemoSchemaPolicyFor(trigger)
	if !policy.Apply {
		logDemoSchemaSkipped(log, policy)
		return policy, nil
	}
	if db == nil {
		return policy, nil
	}
	if _, err := dbapp.ApplyDemoSchema(context.Background(), db, driver); err != nil {
		return policy, fmt.Errorf("apply demo schema: %w", err)
	}
	logDemoSchemaApplied(log, policy)
	return policy, nil
}

func logDemoSchemaApplied(log logger.Logger, policy DemoSchemaPolicy) {
	if log == nil {
		return
	}
	log.Info("demo schema applied", "trigger", policy.Trigger, "reason", policy.Reason)
}

func logDemoSchemaSkipped(log logger.Logger, policy DemoSchemaPolicy) {
	if log == nil {
		return
	}
	log.Debug("demo schema apply skipped", "trigger", policy.Trigger, "reason", policy.Reason)
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
