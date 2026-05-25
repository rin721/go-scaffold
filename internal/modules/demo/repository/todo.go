package repository

import (
	"context"

	"github.com/rei0721/go-scaffold/internal/modules/demo/model"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(ctx context.Context, db *gorm.DB, todo *model.Todo) error
	List(ctx context.Context, db *gorm.DB) ([]model.Todo, error)
	FindByID(ctx context.Context, db *gorm.DB, id uint) (*model.Todo, error)
	Update(ctx context.Context, db *gorm.DB, todo *model.Todo) error
	Delete(ctx context.Context, db *gorm.DB, id uint) error
}

type gormTodoRepository struct{}

func NewTodoRepository() TodoRepository {
	return &gormTodoRepository{}
}

func (r *gormTodoRepository) Create(ctx context.Context, db *gorm.DB, todo *model.Todo) error {
	return db.WithContext(ctx).Create(todo).Error
}

func (r *gormTodoRepository) List(ctx context.Context, db *gorm.DB) ([]model.Todo, error) {
	var todos []model.Todo
	err := db.WithContext(ctx).Order("id DESC").Find(&todos).Error
	return todos, err
}

func (r *gormTodoRepository) FindByID(ctx context.Context, db *gorm.DB, id uint) (*model.Todo, error) {
	var todo model.Todo
	if err := db.WithContext(ctx).First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *gormTodoRepository) Update(ctx context.Context, db *gorm.DB, todo *model.Todo) error {
	return db.WithContext(ctx).Save(todo).Error
}

func (r *gormTodoRepository) Delete(ctx context.Context, db *gorm.DB, id uint) error {
	return db.WithContext(ctx).Delete(&model.Todo{}, id).Error
}
