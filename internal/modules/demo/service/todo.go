package service

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
	ErrTodoTitleRequired = errors.New("todo title is required")
	ErrTodoNotFound      = errors.New("todo not found")
)

type CreateTodoInput struct {
	Title       string
	Description string
	Completed   bool
}

type UpdateTodoInput struct {
	Title       *string
	Description *string
	Completed   *bool
}

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

func NewTodoService(db database.Database, repo repository.TodoRepository) TodoService {
	return &todoService{db: db, repo: repo}
}

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

func (s *todoService) List(ctx context.Context) ([]model.Todo, error) {
	return s.repo.List(ctx, s.db.DB())
}

func (s *todoService) Get(ctx context.Context, id uint) (*model.Todo, error) {
	todo, err := s.repo.FindByID(ctx, s.db.DB(), id)
	return todo, normalizeNotFound(err)
}

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

func (s *todoService) Delete(ctx context.Context, id uint) error {
	return s.db.WithTx(ctx, func(ctx context.Context, tx *gorm.DB) error {
		if _, err := s.repo.FindByID(ctx, tx, id); err != nil {
			return normalizeNotFound(err)
		}
		return s.repo.Delete(ctx, tx, id)
	})
}

func normalizeNotFound(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrTodoNotFound
	}
	return err
}
