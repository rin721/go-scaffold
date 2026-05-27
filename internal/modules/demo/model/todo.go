package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint           `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Title       string         `gorm:"column:title;size:200;not null" json:"title"`
	Description string         `gorm:"column:description;size:1000" json:"description,omitempty"`
	Completed   bool           `gorm:"column:completed;not null;default:false" json:"completed"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
}

func (Todo) TableName() string {
	return "demo_todos"
}
