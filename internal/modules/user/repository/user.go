package repository

import (
	"context"

	"github.com/rei0721/go-scaffold/internal/modules/user/model"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(ctx context.Context, db *gorm.DB, user *model.User) error
	ListUsers(ctx context.Context, db *gorm.DB) ([]model.User, error)
	FindUserByID(ctx context.Context, db *gorm.DB, id uint) (*model.User, error)
	FindUserByUsername(ctx context.Context, db *gorm.DB, username string) (*model.User, error)
	FindUserByEmail(ctx context.Context, db *gorm.DB, email string) (*model.User, error)
	UpdateUser(ctx context.Context, db *gorm.DB, user *model.User) error
	DeleteUser(ctx context.Context, db *gorm.DB, id uint) error
	CountUsers(ctx context.Context, db *gorm.DB) (int64, error)

	CreateRole(ctx context.Context, db *gorm.DB, role *model.Role) error
	ListRoles(ctx context.Context, db *gorm.DB) ([]model.Role, error)
	FindRoleByID(ctx context.Context, db *gorm.DB, id uint) (*model.Role, error)
	FindRoleByName(ctx context.Context, db *gorm.DB, name string) (*model.Role, error)
	UpdateRole(ctx context.Context, db *gorm.DB, role *model.Role) error
	DeleteRole(ctx context.Context, db *gorm.DB, id uint) error

	CreatePermission(ctx context.Context, db *gorm.DB, permission *model.Permission) error
	ListPermissions(ctx context.Context, db *gorm.DB) ([]model.Permission, error)
	FindPermissionByID(ctx context.Context, db *gorm.DB, id uint) (*model.Permission, error)
	FindPermissionByCode(ctx context.Context, db *gorm.DB, code string) (*model.Permission, error)
	UpdatePermission(ctx context.Context, db *gorm.DB, permission *model.Permission) error
	DeletePermission(ctx context.Context, db *gorm.DB, id uint) error

	AssignRoleToUser(ctx context.Context, db *gorm.DB, userID uint, roleID uint) error
	RemoveRoleFromUser(ctx context.Context, db *gorm.DB, userID uint, roleID uint) error
	ReplaceUserRoles(ctx context.Context, db *gorm.DB, userID uint, roleIDs []uint) error
	ListRolesForUser(ctx context.Context, db *gorm.DB, userID uint) ([]model.Role, error)

	AssignPermissionToRole(ctx context.Context, db *gorm.DB, roleID uint, permissionID uint) error
	RemovePermissionFromRole(ctx context.Context, db *gorm.DB, roleID uint, permissionID uint) error
	ReplaceRolePermissions(ctx context.Context, db *gorm.DB, roleID uint, permissionIDs []uint) error
	ListPermissionsForRole(ctx context.Context, db *gorm.DB, roleID uint) ([]model.Permission, error)
	ListPermissionsForUser(ctx context.Context, db *gorm.DB, userID uint) ([]model.Permission, error)
	ListRolePermissionPolicies(ctx context.Context, db *gorm.DB) ([]RolePermissionPolicy, error)
}

type RolePermissionPolicy struct {
	RoleName       string `gorm:"column:role_name"`
	PermissionCode string `gorm:"column:permission_code"`
}

type gormRepository struct{}

func NewRepository() Repository {
	return &gormRepository{}
}

func (r *gormRepository) CreateUser(ctx context.Context, db *gorm.DB, user *model.User) error {
	return db.WithContext(ctx).Create(user).Error
}

func (r *gormRepository) ListUsers(ctx context.Context, db *gorm.DB) ([]model.User, error) {
	var users []model.User
	err := db.WithContext(ctx).Order("id DESC").Find(&users).Error
	return users, err
}

func (r *gormRepository) FindUserByID(ctx context.Context, db *gorm.DB, id uint) (*model.User, error) {
	var user model.User
	if err := db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *gormRepository) FindUserByUsername(ctx context.Context, db *gorm.DB, username string) (*model.User, error) {
	var user model.User
	if err := db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *gormRepository) FindUserByEmail(ctx context.Context, db *gorm.DB, email string) (*model.User, error) {
	var user model.User
	if err := db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *gormRepository) UpdateUser(ctx context.Context, db *gorm.DB, user *model.User) error {
	return db.WithContext(ctx).Save(user).Error
}

func (r *gormRepository) DeleteUser(ctx context.Context, db *gorm.DB, id uint) error {
	if err := db.WithContext(ctx).Where("user_id = ?", id).Delete(&model.UserRole{}).Error; err != nil {
		return err
	}
	return db.WithContext(ctx).Delete(&model.User{}, id).Error
}

func (r *gormRepository) CountUsers(ctx context.Context, db *gorm.DB) (int64, error) {
	var count int64
	err := db.WithContext(ctx).Model(&model.User{}).Count(&count).Error
	return count, err
}

func (r *gormRepository) CreateRole(ctx context.Context, db *gorm.DB, role *model.Role) error {
	return db.WithContext(ctx).Create(role).Error
}

func (r *gormRepository) ListRoles(ctx context.Context, db *gorm.DB) ([]model.Role, error) {
	var roles []model.Role
	err := db.WithContext(ctx).Order("id ASC").Find(&roles).Error
	return roles, err
}

func (r *gormRepository) FindRoleByID(ctx context.Context, db *gorm.DB, id uint) (*model.Role, error) {
	var role model.Role
	if err := db.WithContext(ctx).First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *gormRepository) FindRoleByName(ctx context.Context, db *gorm.DB, name string) (*model.Role, error) {
	var role model.Role
	if err := db.WithContext(ctx).Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *gormRepository) UpdateRole(ctx context.Context, db *gorm.DB, role *model.Role) error {
	return db.WithContext(ctx).Save(role).Error
}

func (r *gormRepository) DeleteRole(ctx context.Context, db *gorm.DB, id uint) error {
	if err := db.WithContext(ctx).Where("role_id = ?", id).Delete(&model.UserRole{}).Error; err != nil {
		return err
	}
	if err := db.WithContext(ctx).Where("role_id = ?", id).Delete(&model.RolePermission{}).Error; err != nil {
		return err
	}
	return db.WithContext(ctx).Delete(&model.Role{}, id).Error
}

func (r *gormRepository) CreatePermission(ctx context.Context, db *gorm.DB, permission *model.Permission) error {
	return db.WithContext(ctx).Create(permission).Error
}

func (r *gormRepository) ListPermissions(ctx context.Context, db *gorm.DB) ([]model.Permission, error) {
	var permissions []model.Permission
	err := db.WithContext(ctx).Order("code ASC").Find(&permissions).Error
	return permissions, err
}

func (r *gormRepository) FindPermissionByID(ctx context.Context, db *gorm.DB, id uint) (*model.Permission, error) {
	var permission model.Permission
	if err := db.WithContext(ctx).First(&permission, id).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *gormRepository) FindPermissionByCode(ctx context.Context, db *gorm.DB, code string) (*model.Permission, error) {
	var permission model.Permission
	if err := db.WithContext(ctx).Where("code = ?", code).First(&permission).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *gormRepository) UpdatePermission(ctx context.Context, db *gorm.DB, permission *model.Permission) error {
	return db.WithContext(ctx).Save(permission).Error
}

func (r *gormRepository) DeletePermission(ctx context.Context, db *gorm.DB, id uint) error {
	if err := db.WithContext(ctx).Where("permission_id = ?", id).Delete(&model.RolePermission{}).Error; err != nil {
		return err
	}
	return db.WithContext(ctx).Delete(&model.Permission{}, id).Error
}

func (r *gormRepository) AssignRoleToUser(ctx context.Context, db *gorm.DB, userID uint, roleID uint) error {
	return db.WithContext(ctx).Where(model.UserRole{UserID: userID, RoleID: roleID}).FirstOrCreate(&model.UserRole{
		UserID: userID,
		RoleID: roleID,
	}).Error
}

func (r *gormRepository) RemoveRoleFromUser(ctx context.Context, db *gorm.DB, userID uint, roleID uint) error {
	return db.WithContext(ctx).Where("user_id = ? AND role_id = ?", userID, roleID).Delete(&model.UserRole{}).Error
}

func (r *gormRepository) ReplaceUserRoles(ctx context.Context, db *gorm.DB, userID uint, roleIDs []uint) error {
	if err := db.WithContext(ctx).Where("user_id = ?", userID).Delete(&model.UserRole{}).Error; err != nil {
		return err
	}
	for _, roleID := range roleIDs {
		if err := r.AssignRoleToUser(ctx, db, userID, roleID); err != nil {
			return err
		}
	}
	return nil
}

func (r *gormRepository) ListRolesForUser(ctx context.Context, db *gorm.DB, userID uint) ([]model.Role, error) {
	var roles []model.Role
	err := db.WithContext(ctx).
		Table((&model.Role{}).TableName()).
		Select("roles.*").
		Joins("JOIN user_roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", userID).
		Order("roles.name ASC").
		Scan(&roles).Error
	return roles, err
}

func (r *gormRepository) AssignPermissionToRole(ctx context.Context, db *gorm.DB, roleID uint, permissionID uint) error {
	return db.WithContext(ctx).Where(model.RolePermission{RoleID: roleID, PermissionID: permissionID}).FirstOrCreate(&model.RolePermission{
		RoleID:       roleID,
		PermissionID: permissionID,
	}).Error
}

func (r *gormRepository) RemovePermissionFromRole(ctx context.Context, db *gorm.DB, roleID uint, permissionID uint) error {
	return db.WithContext(ctx).Where("role_id = ? AND permission_id = ?", roleID, permissionID).Delete(&model.RolePermission{}).Error
}

func (r *gormRepository) ReplaceRolePermissions(ctx context.Context, db *gorm.DB, roleID uint, permissionIDs []uint) error {
	if err := db.WithContext(ctx).Where("role_id = ?", roleID).Delete(&model.RolePermission{}).Error; err != nil {
		return err
	}
	for _, permissionID := range permissionIDs {
		if err := r.AssignPermissionToRole(ctx, db, roleID, permissionID); err != nil {
			return err
		}
	}
	return nil
}

func (r *gormRepository) ListPermissionsForRole(ctx context.Context, db *gorm.DB, roleID uint) ([]model.Permission, error) {
	var permissions []model.Permission
	err := db.WithContext(ctx).
		Table((&model.Permission{}).TableName()).
		Select("permissions.*").
		Joins("JOIN role_permissions ON role_permissions.permission_id = permissions.id").
		Where("role_permissions.role_id = ?", roleID).
		Order("permissions.code ASC").
		Scan(&permissions).Error
	return permissions, err
}

func (r *gormRepository) ListPermissionsForUser(ctx context.Context, db *gorm.DB, userID uint) ([]model.Permission, error) {
	var permissions []model.Permission
	err := db.WithContext(ctx).
		Table((&model.Permission{}).TableName()).
		Select("DISTINCT permissions.*").
		Joins("JOIN role_permissions ON role_permissions.permission_id = permissions.id").
		Joins("JOIN user_roles ON user_roles.role_id = role_permissions.role_id").
		Where("user_roles.user_id = ?", userID).
		Order("permissions.code ASC").
		Scan(&permissions).Error
	return permissions, err
}

func (r *gormRepository) ListRolePermissionPolicies(ctx context.Context, db *gorm.DB) ([]RolePermissionPolicy, error) {
	var policies []RolePermissionPolicy
	err := db.WithContext(ctx).
		Table((&model.RolePermission{}).TableName()).
		Select("roles.name AS role_name, permissions.code AS permission_code").
		Joins("JOIN roles ON roles.id = role_permissions.role_id").
		Joins("JOIN permissions ON permissions.id = role_permissions.permission_id").
		Order("roles.name ASC, permissions.code ASC").
		Scan(&policies).Error
	return policies, err
}
