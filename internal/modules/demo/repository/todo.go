package repository

// 本文件定义 Demo Todo 的仓储边界，把 GORM 数据访问限制在领域服务可替换的接口之后。

import (
	"context"

	"github.com/rei0721/go-scaffold/internal/modules/demo/model"
	"gorm.io/gorm"
)

// TodoRepository 定义 Todo 持久化端口，使服务层只依赖可替换的数据访问契约。
type TodoRepository interface {
	Create(ctx context.Context, db *gorm.DB, todo *model.Todo) error
	List(ctx context.Context, db *gorm.DB) ([]model.Todo, error)
	FindByID(ctx context.Context, db *gorm.DB, id uint) (*model.Todo, error)
	Update(ctx context.Context, db *gorm.DB, todo *model.Todo) error
	Delete(ctx context.Context, db *gorm.DB, id uint) error
}

type gormTodoRepository struct{}

// NewTodoRepository 创建 GORM 版 TodoRepository，实例本身不持有连接状态。
func NewTodoRepository() TodoRepository {
	return &gormTodoRepository{}
}

// Create 在传入事务或数据库句柄上插入 Todo，底层错误由服务层决定如何映射。
func (r *gormTodoRepository) Create(ctx context.Context, db *gorm.DB, todo *model.Todo) error {
	return db.WithContext(ctx).Create(todo).Error
}

// List 按 limit 和 offset 读取 Todo 列表，调用方负责约束分页上限。
func (r *gormTodoRepository) List(ctx context.Context, db *gorm.DB) ([]model.Todo, error) {
	var todos []model.Todo
	err := db.WithContext(ctx).Order("id DESC").Find(&todos).Error
	return todos, err
}

// FindByID 按主键读取 Todo，未命中时保留 GORM 的 not found 错误语义。
func (r *gormTodoRepository) FindByID(ctx context.Context, db *gorm.DB, id uint) (*model.Todo, error) {
	var todo model.Todo
	if err := db.WithContext(ctx).First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

// Update 在当前数据库上下文中持久化 Todo 的字段变化。
func (r *gormTodoRepository) Update(ctx context.Context, db *gorm.DB, todo *model.Todo) error {
	return db.WithContext(ctx).Save(todo).Error
}

// Delete 执行 Todo 删除操作，是否软删除取决于模型与 GORM 配置。
func (r *gormTodoRepository) Delete(ctx context.Context, db *gorm.DB, id uint) error {
	return db.WithContext(ctx).Delete(&model.Todo{}, id).Error
}
