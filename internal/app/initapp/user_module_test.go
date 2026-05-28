package initapp

import (
	"context"
	"testing"
	"time"

	"github.com/rei0721/go-scaffold/internal/config"
	userservice "github.com/rei0721/go-scaffold/internal/modules/user/service"
	authapi "github.com/rei0721/go-scaffold/pkg/auth"
	"github.com/rei0721/go-scaffold/pkg/database"
)

func TestNewUserModuleUsesConfiguredAuthTokenSecret(t *testing.T) {
	secret := "0123456789abcdef0123456789abcdef"
	module, err := NewUserModule(nil, nil, config.AuthConfig{
		TokenSecret: secret,
		TokenTTL:    60,
	}, config.RBACConfig{})
	if err != nil {
		t.Fatalf("NewUserModule() error = %v", err)
	}

	issued, err := module.Tokens.Issue(context.Background(), authapi.Claims{Subject: "42", Username: "ada"})
	if err != nil {
		t.Fatalf("Issue() error = %v", err)
	}
	token := issued.Value
	expiresAt := issued.ExpiresAt
	if time.Until(expiresAt) > 2*time.Minute {
		t.Fatalf("expiresAt = %v, want configured short ttl", expiresAt)
	}

	verifier, err := authapi.NewJWTService(authapi.JWTConfig{
		Secret: []byte(secret),
		TTL:    time.Minute,
		Issuer: "go-scaffold",
	})
	if err != nil {
		t.Fatalf("NewJWTService(verifier) error = %v", err)
	}
	claims, err := verifier.Verify(context.Background(), token)
	if err != nil {
		t.Fatalf("Verify() with configured secret error = %v", err)
	}
	if claims.Subject != "42" || claims.Username != "ada" {
		t.Fatalf("claims = %#v, want user 42/ada", claims)
	}
}

func TestNewModulesAppliesConfiguredRBACSeed(t *testing.T) {
	ctx := context.Background()
	db := newTestDatabase(t)
	core := Core{
		Config: &config.Config{
			Database: config.DatabaseConfig{Driver: string(database.DriverSQLite)},
			Auth: config.AuthConfig{
				TokenSecret: "0123456789abcdef0123456789abcdef",
				TokenTTL:    60,
			},
			RBAC: config.RBACConfig{
				Enabled:      true,
				ApplyOnStart: true,
				Roles: []config.RBACRoleConfig{
					{Name: "admin", Description: "Admin role"},
					{Name: "user", Description: "User role"},
					{Name: "reader", Description: "Reader role"},
				},
				Permissions: []config.RBACPermissionConfig{
					{Code: userservice.PermissionAll, Description: "All permissions"},
					{Code: userservice.PermissionUsersRead, Description: "Read users"},
				},
				RolePermissions: []config.RBACRolePermissionConfig{
					{Role: "admin", Permissions: []string{userservice.PermissionAll}},
					{Role: "reader", Permissions: []string{userservice.PermissionUsersRead}},
				},
			},
		},
		Logger: testLogger{},
	}

	modules, err := NewModules(core, Infrastructure{Database: db})
	if err != nil {
		t.Fatalf("NewModules() error = %v", err)
	}
	roles, err := modules.User.Service.ListRoles(ctx)
	if err != nil {
		t.Fatalf("ListRoles() error = %v", err)
	}
	reader := findInitRoleByName(roles, "reader")
	if reader == nil || len(reader.Permissions) != 1 || reader.Permissions[0].Code != userservice.PermissionUsersRead {
		t.Fatalf("reader role = %#v, want users:read permission", reader)
	}
}

func findInitRoleByName(roles []userservice.RoleDTO, name string) *userservice.RoleDTO {
	for i := range roles {
		if roles[i].Name == name {
			return &roles[i]
		}
	}
	return nil
}
