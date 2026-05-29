package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/rei0721/go-scaffold/internal/app/dbapp"
	"github.com/rei0721/go-scaffold/internal/app/initapp"
	appconfig "github.com/rei0721/go-scaffold/internal/config"
	"github.com/rei0721/go-scaffold/pkg/cli"
	"github.com/rei0721/go-scaffold/types/constants"
)

const (
	dbCommandName       = "db"
	dbOperationDatabase = "database"
	dbOperationSchema   = "schema"
	dbOperationCreate   = "todo-create"
	dbOperationList     = "todo-list"
	dbOperationGet      = "todo-get"
	dbOperationUpdate   = "todo-update"
	dbOperationDelete   = "todo-delete"
	defaultDBOperation  = dbOperationSchema
	dbCompletedFlagName = "completed"
)

// DBCommand exposes database DDL, demo schema, and demo Todo CRUD operations
// through sqlgen-backed application services.
type DBCommand struct {
	runner func(*cli.Context, dbOptions) error
}

// dbOptions is the parsed flag state for one db command invocation.
type dbOptions struct {
	ConfigPath  string
	Operation   string
	Apply       bool
	ID          int
	Title       string
	Description string
	Completed   string
	Limit       int
	Offset      int
	PrintSQL    bool
}

// NewDBCommand builds the cmd/main db command. Keep command wiring here and
// SQL generation/execution in internal/app/dbapp so the CLI remains thin.
func NewDBCommand() *DBCommand {
	return &DBCommand{}
}

func (c *DBCommand) Name() string {
	return dbCommandName
}

func (c *DBCommand) Description() string {
	return "Run sqlgen-powered database schema and demo CRUD operations"
}

func (c *DBCommand) Usage() string {
	return fmt.Sprintf("%s [--operation=<database|schema|todo-create|todo-list|todo-get|todo-update|todo-delete>] [flags]", dbCommandName)
}

func (c *DBCommand) Flags() []cli.Flag {
	return []cli.Flag{
		{
			Name:        "config",
			Type:        cli.FlagTypeString,
			Default:     constants.AppDefaultConfigPath,
			Description: "Config file path",
			EnvVar:      appconfig.EnvConfigPathName(),
		},
		{
			Name:        "operation",
			Type:        cli.FlagTypeString,
			Default:     defaultDBOperation,
			Description: "Database operation",
		},
		{
			Name:        "apply",
			Type:        cli.FlagTypeBool,
			Default:     false,
			Description: "Apply generated schema SQL instead of printing it",
		},
		{
			Name:        "id",
			Type:        cli.FlagTypeInt,
			Default:     0,
			Description: "Demo todo ID",
		},
		{
			Name:        "title",
			Type:        cli.FlagTypeString,
			Default:     "",
			Description: "Demo todo title",
		},
		{
			Name:        "description",
			Type:        cli.FlagTypeString,
			Default:     "",
			Description: "Demo todo description",
		},
		{
			Name:        dbCompletedFlagName,
			Type:        cli.FlagTypeString,
			Default:     "",
			Description: "Demo todo completed value: true or false",
		},
		{
			Name:        "limit",
			Type:        cli.FlagTypeInt,
			Default:     dbapp.DefaultTodoLimit,
			Description: "Demo todo list limit",
		},
		{
			Name:        "offset",
			Type:        cli.FlagTypeInt,
			Default:     0,
			Description: "Demo todo list offset",
		},
		{
			Name:        "print-sql",
			Type:        cli.FlagTypeBool,
			Default:     false,
			Description: "Print generated SQL after executing the operation",
		},
	}
}

func (c *DBCommand) Execute(ctx *cli.Context) error {
	opts := dbOptions{
		ConfigPath:  ctx.GetString("config"),
		Operation:   ctx.GetString("operation"),
		Apply:       ctx.GetBool("apply"),
		ID:          ctx.GetInt("id"),
		Title:       ctx.GetString("title"),
		Description: ctx.GetString("description"),
		Completed:   ctx.GetString(dbCompletedFlagName),
		Limit:       ctx.GetInt("limit"),
		Offset:      ctx.GetInt("offset"),
		PrintSQL:    ctx.GetBool("print-sql"),
	}
	if opts.Operation == "" {
		opts.Operation = defaultDBOperation
	}
	if c.runner != nil {
		return c.runner(ctx, opts)
	}
	return runDBCommand(ctx, opts)
}

func runDBCommand(ctx *cli.Context, opts dbOptions) error {
	core, err := initapp.NewCore(opts.ConfigPath)
	if err != nil {
		return fmt.Errorf("initialize core: %w", err)
	}
	defer func() {
		if core.Logger != nil {
			_ = core.Logger.Sync()
		}
	}()

	// Read-only DDL preview must not open a database connection.
	if (opts.Operation == dbOperationDatabase || opts.Operation == dbOperationSchema) && !opts.Apply {
		sql, err := dbSQLForPrint(opts, core.Config.Database)
		if err != nil {
			return err
		}
		fmt.Fprintln(ctx.Stdout, sql)
		return nil
	}

	db, err := initapp.NewDatabase(core.Config)
	if err != nil {
		return fmt.Errorf("initialize database: %w", err)
	}
	defer func() {
		_ = db.Close()
	}()

	switch opts.Operation {
	case dbOperationDatabase:
		sql, err := dbapp.ApplyDatabase(ctxWithFallback(ctx), db, core.Config.Database.Driver, core.Config.Database.DBName)
		if err != nil {
			return err
		}
		writeOperationResult(ctx.Stdout, "database create applied", sql, opts.PrintSQL)
	case dbOperationSchema:
		sql, err := dbapp.ApplyDemoSchema(ctxWithFallback(ctx), db, core.Config.Database.Driver)
		if err != nil {
			return err
		}
		writeOperationResult(ctx.Stdout, "schema applied", sql, opts.PrintSQL)
	case dbOperationCreate:
		completed, err := parseCompleted(opts.Completed, false)
		if err != nil {
			return err
		}
		sql, err := dbapp.CreateTodo(ctxWithFallback(ctx), db, core.Config.Database.Driver, dbapp.TodoCreateInput{
			Title:       opts.Title,
			Description: opts.Description,
			Completed:   completed,
		})
		if err != nil {
			return err
		}
		writeOperationResult(ctx.Stdout, "todo created", sql, opts.PrintSQL)
	case dbOperationList:
		todos, sql, err := dbapp.ListTodos(ctxWithFallback(ctx), db, core.Config.Database.Driver, opts.Limit, opts.Offset)
		if err != nil {
			return err
		}
		if opts.PrintSQL {
			fmt.Fprintln(ctx.Stdout, sql)
		}
		return json.NewEncoder(ctx.Stdout).Encode(todos)
	case dbOperationGet:
		todo, sql, err := dbapp.GetTodo(ctxWithFallback(ctx), db, core.Config.Database.Driver, uint(opts.ID))
		if err != nil {
			return err
		}
		if opts.PrintSQL {
			fmt.Fprintln(ctx.Stdout, sql)
		}
		return json.NewEncoder(ctx.Stdout).Encode(todo)
	case dbOperationUpdate:
		update, err := buildTodoUpdate(opts)
		if err != nil {
			return err
		}
		sql, err := dbapp.UpdateTodo(ctxWithFallback(ctx), db, core.Config.Database.Driver, update)
		if err != nil {
			return err
		}
		writeOperationResult(ctx.Stdout, "todo updated", sql, opts.PrintSQL)
	case dbOperationDelete:
		sql, err := dbapp.DeleteTodo(ctxWithFallback(ctx), db, core.Config.Database.Driver, uint(opts.ID))
		if err != nil {
			return err
		}
		writeOperationResult(ctx.Stdout, "todo deleted", sql, opts.PrintSQL)
	default:
		return fmt.Errorf("unsupported db operation: %s", opts.Operation)
	}

	return nil
}

// dbSQLForPrint centralizes operations that can render generated SQL without
// side effects. New printable operations should be added here and backed by
// sqlgen APIs in internal/app/dbapp.
func dbSQLForPrint(opts dbOptions, cfg appconfig.DatabaseConfig) (string, error) {
	switch opts.Operation {
	case dbOperationDatabase:
		return dbapp.DatabaseSQL(cfg.Driver, cfg.DBName)
	case dbOperationSchema:
		return dbapp.DemoSchemaSQL(cfg.Driver)
	default:
		return "", fmt.Errorf("unsupported db operation: %s", opts.Operation)
	}
}

func ctxWithFallback(ctx *cli.Context) context.Context {
	return context.Background()
}

func writeOperationResult(w io.Writer, message, sql string, printSQL bool) {
	fmt.Fprintln(w, message)
	if printSQL {
		fmt.Fprintln(w, sql)
	}
}

func parseCompleted(value string, defaultValue bool) (bool, error) {
	if value == "" {
		return defaultValue, nil
	}
	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return false, fmt.Errorf("invalid --%s value: %w", dbCompletedFlagName, err)
	}
	return parsed, nil
}

func buildTodoUpdate(opts dbOptions) (dbapp.TodoUpdateInput, error) {
	update := dbapp.TodoUpdateInput{ID: uint(opts.ID)}
	if opts.Title != "" {
		update.Title = &opts.Title
	}
	if opts.Description != "" {
		update.Description = &opts.Description
	}
	if opts.Completed != "" {
		completed, err := parseCompleted(opts.Completed, false)
		if err != nil {
			return update, err
		}
		update.Completed = &completed
	}
	return update, nil
}
