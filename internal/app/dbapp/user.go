package dbapp

import (
	"context"
	"fmt"
	"strings"

	"github.com/rei0721/go-scaffold/internal/modules/user/model"
	"github.com/rei0721/go-scaffold/pkg/database"
)

func UserSchemaSQL(driver string) (string, error) {
	statements, err := userSchemaStatements(driver)
	if err != nil {
		return "", err
	}
	return strings.Join(statements, "\n"), nil
}

func ApplyUserSchema(ctx context.Context, db database.Database, driver string) (string, error) {
	if db == nil {
		return "", ErrMissingDatabase
	}
	statements, err := userSchemaStatements(driver)
	if err != nil {
		return "", err
	}
	for _, statement := range statements {
		if err := db.DB().WithContext(ctx).Exec(statement).Error; err != nil {
			return strings.Join(statements, "\n"), fmt.Errorf("apply user schema: %w", err)
		}
	}
	return strings.Join(statements, "\n"), nil
}

func userSchemaStatements(driver string) ([]string, error) {
	gen, err := NewGenerator(driver)
	if err != nil {
		return nil, err
	}
	models := []interface{}{
		&model.User{},
		&model.Role{},
		&model.Permission{},
		&model.UserRole{},
		&model.RolePermission{},
	}
	statements := make([]string, 0, len(models)+5)
	for _, item := range models {
		sql, err := gen.TableIfNotExists(item)
		if err != nil {
			return nil, err
		}
		statements = append(statements, sql)
	}
	statements = append(statements, userUniqueIndexStatements(driver)...)
	return statements, nil
}

func userUniqueIndexStatements(driver string) []string {
	switch database.Driver(driver) {
	case database.DriverSQLite, database.DriverPostgres:
		return []string{
			`CREATE UNIQUE INDEX IF NOT EXISTS "uk_users_username" ON "users" ("username");`,
			`CREATE UNIQUE INDEX IF NOT EXISTS "uk_users_email" ON "users" ("email");`,
			`CREATE UNIQUE INDEX IF NOT EXISTS "uk_roles_name" ON "roles" ("name");`,
			`CREATE UNIQUE INDEX IF NOT EXISTS "uk_permissions_code" ON "permissions" ("code");`,
			`CREATE UNIQUE INDEX IF NOT EXISTS "uk_user_roles_user_role" ON "user_roles" ("user_id", "role_id");`,
			`CREATE UNIQUE INDEX IF NOT EXISTS "uk_role_permissions_role_permission" ON "role_permissions" ("role_id", "permission_id");`,
		}
	default:
		return nil
	}
}
