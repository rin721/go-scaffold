package dbapp

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/rei0721/go-scaffold/internal/modules/demo/model"
	"github.com/rei0721/go-scaffold/pkg/database"
	"github.com/rei0721/go-scaffold/pkg/sqlgen"
	"gorm.io/gorm"
)

const (
	DefaultTodoLimit = 100
	MaxTodoLimit     = 1000
)

var (
	ErrUnsupportedDriver = errors.New("unsupported database driver")
	ErrMissingDatabase   = errors.New("database is nil")
	ErrMissingTodoTitle  = errors.New("todo title is required")
	ErrMissingTodoID     = errors.New("todo id is required")
	ErrNoTodoUpdate      = errors.New("no todo fields to update")
)

type TodoCreateInput struct {
	Title       string
	Description string
	Completed   bool
}

type TodoUpdateInput struct {
	ID          uint
	Title       *string
	Description *string
	Completed   *bool
}

func DialectForDriver(driver string) (sqlgen.Dialect, error) {
	switch database.Driver(driver) {
	case database.DriverSQLite:
		return sqlgen.SQLite, nil
	case database.DriverMySQL:
		return sqlgen.MySQL, nil
	case database.DriverPostgres:
		return sqlgen.PostgreSQL, nil
	default:
		return "", fmt.Errorf("%w: %s", ErrUnsupportedDriver, driver)
	}
}

func NewGenerator(driver string) (*sqlgen.Generator, error) {
	dialect, err := DialectForDriver(driver)
	if err != nil {
		return nil, err
	}
	return sqlgen.New(&sqlgen.Config{
		Dialect:       dialect,
		SoftDelete:    true,
		SkipZeroValue: true,
	}), nil
}

func DemoSchemaSQL(driver string) (string, error) {
	gen, err := NewGenerator(driver)
	if err != nil {
		return "", err
	}
	return gen.TableIfNotExists(&model.Todo{})
}

func DatabaseSQL(driver, name string) (string, error) {
	gen, err := NewGenerator(driver)
	if err != nil {
		return "", err
	}
	return gen.DatabaseIfNotExists(name)
}

func ApplyDatabase(ctx context.Context, db database.Database, driver, name string) (string, error) {
	if db == nil {
		return "", ErrMissingDatabase
	}
	sql, err := DatabaseSQL(driver, name)
	if err != nil {
		return "", err
	}
	if err := db.DB().WithContext(ctx).Exec(sql).Error; err != nil {
		return sql, fmt.Errorf("apply database create: %w", err)
	}
	return sql, nil
}

func ApplyDemoSchema(ctx context.Context, db database.Database, driver string) (string, error) {
	if db == nil {
		return "", ErrMissingDatabase
	}
	sql, err := DemoSchemaSQL(driver)
	if err != nil {
		return "", err
	}
	if err := db.DB().WithContext(ctx).Exec(sql).Error; err != nil {
		return sql, fmt.Errorf("apply demo schema: %w", err)
	}
	return sql, nil
}

func CreateTodo(ctx context.Context, db database.Database, driver string, input TodoCreateInput) (string, error) {
	if db == nil {
		return "", ErrMissingDatabase
	}
	if input.Title == "" {
		return "", ErrMissingTodoTitle
	}

	now := time.Now().UTC()
	todo := model.Todo{
		Title:       input.Title,
		Description: input.Description,
		Completed:   input.Completed,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	gen, err := NewGenerator(driver)
	if err != nil {
		return "", err
	}
	sql, err := gen.Omit("ID", "DeletedAt").Create(&todo)
	if err != nil {
		return "", err
	}
	if err := db.DB().WithContext(ctx).Exec(sql).Error; err != nil {
		return sql, fmt.Errorf("create demo todo: %w", err)
	}
	return sql, nil
}

func ListTodos(ctx context.Context, db database.Database, driver string, limit, offset int) ([]model.Todo, string, error) {
	if db == nil {
		return nil, "", ErrMissingDatabase
	}
	limit = normalizeLimit(limit)
	if offset < 0 {
		offset = 0
	}

	gen, err := NewGenerator(driver)
	if err != nil {
		return nil, "", err
	}
	sql, err := gen.Order("id DESC").Limit(limit).Offset(offset).Find(&model.Todo{})
	if err != nil {
		return nil, "", err
	}

	var todos []model.Todo
	if err := db.DB().WithContext(ctx).Raw(sql).Scan(&todos).Error; err != nil {
		return nil, sql, fmt.Errorf("list demo todos: %w", err)
	}
	return todos, sql, nil
}

func GetTodo(ctx context.Context, db database.Database, driver string, id uint) (*model.Todo, string, error) {
	if db == nil {
		return nil, "", ErrMissingDatabase
	}
	if id == 0 {
		return nil, "", ErrMissingTodoID
	}

	gen, err := NewGenerator(driver)
	if err != nil {
		return nil, "", err
	}
	sql, err := gen.First(&model.Todo{}, id)
	if err != nil {
		return nil, "", err
	}

	var todo model.Todo
	result := db.DB().WithContext(ctx).Raw(sql).Scan(&todo)
	if result.Error != nil {
		return nil, sql, fmt.Errorf("get demo todo: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, sql, gorm.ErrRecordNotFound
	}
	return &todo, sql, nil
}

func UpdateTodo(ctx context.Context, db database.Database, driver string, input TodoUpdateInput) (string, error) {
	if db == nil {
		return "", ErrMissingDatabase
	}
	if input.ID == 0 {
		return "", ErrMissingTodoID
	}

	values := map[string]interface{}{}
	if input.Title != nil {
		values["title"] = *input.Title
	}
	if input.Description != nil {
		values["description"] = *input.Description
	}
	if input.Completed != nil {
		values["completed"] = *input.Completed
	}
	if len(values) == 0 {
		return "", ErrNoTodoUpdate
	}
	values["updated_at"] = time.Now().UTC()

	gen, err := NewGenerator(driver)
	if err != nil {
		return "", err
	}
	sql, err := gen.Model(&model.Todo{}).Where("id = ?", input.ID).Updates(values)
	if err != nil {
		return "", err
	}
	if err := db.DB().WithContext(ctx).Exec(sql).Error; err != nil {
		return sql, fmt.Errorf("update demo todo: %w", err)
	}
	return sql, nil
}

func DeleteTodo(ctx context.Context, db database.Database, driver string, id uint) (string, error) {
	if db == nil {
		return "", ErrMissingDatabase
	}
	if id == 0 {
		return "", ErrMissingTodoID
	}

	gen, err := NewGenerator(driver)
	if err != nil {
		return "", err
	}
	sql, err := gen.Delete(&model.Todo{}, id)
	if err != nil {
		return "", err
	}
	if err := db.DB().WithContext(ctx).Exec(sql).Error; err != nil {
		return sql, fmt.Errorf("delete demo todo: %w", err)
	}
	return sql, nil
}

func normalizeLimit(limit int) int {
	if limit <= 0 {
		return DefaultTodoLimit
	}
	if limit > MaxTodoLimit {
		return MaxTodoLimit
	}
	return limit
}
