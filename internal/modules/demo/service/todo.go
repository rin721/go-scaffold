package service

// 本文件定义 Demo Todo 的领域服务，集中处理输入归一化、事务范围和业务错误语义。

import (
	"context"
	"errors"
	"strings"

	"github.com/rei0721/go-scaffold/internal/modules/demo/model"
	"github.com/rei0721/go-scaffold/internal/modules/demo/repository"
	"github.com/rei0721/go-scaffold/pkg/database"
	"gorm.io/gorm"
)

var (
	// ErrTodoTitleRequired 表示创建或更新后标题为空，handler 会将其转换为客户端输入错误。
	ErrTodoTitleRequired = errors.New("todo title is required")
	// ErrTodoNotFound 表示按 ID 查询不到 Todo，handler 会将其转换为未找到响应。
	ErrTodoNotFound = errors.New("todo not found")
)

// CreateTodoInput 是服务层创建 Todo 的输入 DTO，字段已脱离 HTTP binding 细节。
type CreateTodoInput struct {
	Title       string
	Description string
	Completed   bool
}

// UpdateTodoInput 是服务层更新 Todo 的输入 DTO，指针字段用于区分未传入与传入零值。
type UpdateTodoInput struct {
	Title       *string
	Description *string
	Completed   *bool
}

// TodoService 定义 Demo Todo 的业务端口，集中约束事务、校验和错误语义。
type TodoService interface {
	Create(ctx context.Context, input CreateTodoInput) (*model.Todo, error)
	List(ctx context.Context) ([]model.Todo, error)
	Get(ctx context.Context, id uint) (*model.Todo, error)
	Update(ctx context.Context, id uint, input UpdateTodoInput) (*model.Todo, error)
	Delete(ctx context.Context, id uint) error
}

type todoService struct {
	db   database.Database
	repo repository.TodoRepository
}

// NewTodoService 创建 TodoService，并注入仓储、数据库和日志依赖以便替换测试。
func NewTodoService(db database.Database, repo repository.TodoRepository) TodoService {
	return &todoService{db: db, repo: repo}
}

// Create 校验并归一化标题后在事务内创建 Todo，返回业务错误或底层数据库错误。
func (s *todoService) Create(ctx context.Context, input CreateTodoInput) (*model.Todo, error) {
	title := strings.TrimSpace(input.Title)
	if title == "" {
		return nil, ErrTodoTitleRequired
	}

	todo := &model.Todo{
		Title:       title,
		Description: strings.TrimSpace(input.Description),
		Completed:   input.Completed,
	}

	if err := s.db.WithTx(ctx, func(ctx context.Context, tx *gorm.DB) error {
		return s.repo.Create(ctx, tx, todo)
	}); err != nil {
		return nil, err
	}
	return todo, nil
}

// List 读取 Todo 列表，当前示例把分页约束留给调用入口控制。
func (s *todoService) List(ctx context.Context) ([]model.Todo, error) {
	return s.repo.List(ctx, s.db.DB())
}

// Get 按 ID 查询 Todo，并把 GORM 未命中错误收敛为领域级 ErrTodoNotFound。
func (s *todoService) Get(ctx context.Context, id uint) (*model.Todo, error) {
	todo, err := s.repo.FindByID(ctx, s.db.DB(), id)
	return todo, normalizeNotFound(err)
}

// Update 在事务中应用部分字段更新，确保空标题不会被持久化。
func (s *todoService) Update(ctx context.Context, id uint, input UpdateTodoInput) (*model.Todo, error) {
	var todo *model.Todo
	err := s.db.WithTx(ctx, func(ctx context.Context, tx *gorm.DB) error {
		current, err := s.repo.FindByID(ctx, tx, id)
		if err != nil {
			return normalizeNotFound(err)
		}

		if input.Title != nil {
			title := strings.TrimSpace(*input.Title)
			if title == "" {
				return ErrTodoTitleRequired
			}
			current.Title = title
		}
		if input.Description != nil {
			current.Description = strings.TrimSpace(*input.Description)
		}
		if input.Completed != nil {
			current.Completed = *input.Completed
		}

		if err := s.repo.Update(ctx, tx, current); err != nil {
			return err
		}
		todo = current
		return nil
	})
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// Delete 在事务中删除 Todo，并把未命中删除统一收敛为领域错误。
func (s *todoService) Delete(ctx context.Context, id uint) error {
	return s.db.WithTx(ctx, func(ctx context.Context, tx *gorm.DB) error {
		if _, err := s.repo.FindByID(ctx, tx, id); err != nil {
			return normalizeNotFound(err)
		}
		return s.repo.Delete(ctx, tx, id)
	})
}

// normalizeNotFound 将 GORM 未命中错误收敛为领域级 ErrTodoNotFound，其余错误保持原样上抛。
func normalizeNotFound(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrTodoNotFound
	}
	return err
}
