package model

import "time"

const (
	UserStatusActive   = "active"
	UserStatusDisabled = "disabled"

	RoleAdmin = "admin"
	RoleUser  = "user"
)

type User struct {
	ID           uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Username     string    `gorm:"column:username;size:64;not null;uniqueIndex:uk_users_username" json:"username"`
	Email        string    `gorm:"column:email;size:254;not null;uniqueIndex:uk_users_email" json:"email"`
	PasswordHash string    `gorm:"column:password_hash;size:255;not null" json:"-"`
	DisplayName  string    `gorm:"column:display_name;size:100" json:"displayName,omitempty"`
	Status       string    `gorm:"column:status;size:32;not null;default:'active'" json:"status"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (User) TableName() string {
	return "users"
}

type Role struct {
	ID          uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"column:name;size:64;not null;uniqueIndex:uk_roles_name" json:"name"`
	Description string    `gorm:"column:description;size:255" json:"description,omitempty"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (Role) TableName() string {
	return "roles"
}

type Permission struct {
	ID          uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Code        string    `gorm:"column:code;size:128;not null;uniqueIndex:uk_permissions_code" json:"code"`
	Description string    `gorm:"column:description;size:255" json:"description,omitempty"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (Permission) TableName() string {
	return "permissions"
}

type UserRole struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"column:user_id;not null;uniqueIndex:uk_user_roles_user_role" json:"userId"`
	RoleID    uint      `gorm:"column:role_id;not null;uniqueIndex:uk_user_roles_user_role" json:"roleId"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (UserRole) TableName() string {
	return "user_roles"
}

type RolePermission struct {
	ID           uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RoleID       uint      `gorm:"column:role_id;not null;uniqueIndex:uk_role_permissions_role_permission" json:"roleId"`
	PermissionID uint      `gorm:"column:permission_id;not null;uniqueIndex:uk_role_permissions_role_permission" json:"permissionId"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
