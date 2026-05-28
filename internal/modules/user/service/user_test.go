package service

import (
	"context"
	"errors"
	"path/filepath"
	"testing"

	"github.com/rei0721/go-scaffold/internal/app/dbapp"
	"github.com/rei0721/go-scaffold/internal/modules/user/model"
	"github.com/rei0721/go-scaffold/internal/modules/user/repository"
	authapi "github.com/rei0721/go-scaffold/pkg/auth"
	passwordcrypto "github.com/rei0721/go-scaffold/pkg/crypto"
	"github.com/rei0721/go-scaffold/pkg/database"
)

func TestUserServiceRegisterLoginAndRBAC(t *testing.T) {
	ctx := context.Background()
	svc, repo, db := newUserServiceForTest(t)

	adminToken, err := svc.Register(ctx, CreateUserInput{
		Username: "admin",
		Email:    "admin@example.com",
		Password: "password123",
	})
	if err != nil {
		t.Fatalf("register admin: %v", err)
	}
	if adminToken.Token == "" || adminToken.TokenType != "Bearer" {
		t.Fatalf("expected bearer token, got %#v", adminToken)
	}
	if len(adminToken.User.Roles) != 1 || adminToken.User.Roles[0].Name != model.RoleAdmin {
		t.Fatalf("expected first user to be admin, got %#v", adminToken.User.Roles)
	}

	storedAdmin, err := repo.FindUserByUsername(ctx, db.DB(), "admin")
	if err != nil {
		t.Fatalf("find stored admin: %v", err)
	}
	if storedAdmin.PasswordHash == "password123" || storedAdmin.PasswordHash == "" {
		t.Fatalf("expected password hash to be stored securely, got %q", storedAdmin.PasswordHash)
	}

	loginToken, err := svc.Login(ctx, LoginInput{Identifier: "admin@example.com", Password: "password123"})
	if err != nil {
		t.Fatalf("login admin: %v", err)
	}
	adminPrincipal, err := svc.AuthenticateToken(ctx, loginToken.Token)
	if err != nil {
		t.Fatalf("authenticate admin token: %v", err)
	}
	if err := svc.Authorize(ctx, *adminPrincipal, PermissionUsersDelete); err != nil {
		t.Fatalf("expected admin wildcard permission, got %v", err)
	}

	memberToken, err := svc.Register(ctx, CreateUserInput{
		Username: "member",
		Email:    "member@example.com",
		Password: "password123",
	})
	if err != nil {
		t.Fatalf("register member: %v", err)
	}
	memberPrincipal, err := svc.AuthenticateToken(ctx, memberToken.Token)
	if err != nil {
		t.Fatalf("authenticate member token: %v", err)
	}
	if err := svc.Authorize(ctx, *memberPrincipal, PermissionUsersRead); !errors.Is(err, ErrPermissionDenied) {
		t.Fatalf("member authorize users:read error = %v, want ErrPermissionDenied", err)
	}
}

func TestUserServiceRolePermissionAssignment(t *testing.T) {
	ctx := context.Background()
	svc, _, _ := newUserServiceForTest(t)

	if _, err := svc.Register(ctx, CreateUserInput{
		Username: "admin",
		Email:    "admin@example.com",
		Password: "password123",
	}); err != nil {
		t.Fatalf("register admin: %v", err)
	}
	if _, err := svc.CreatePermission(ctx, CreatePermissionInput{
		Code:        PermissionUsersRead,
		Description: "Read users",
	}); err != nil {
		t.Fatalf("create permission: %v", err)
	}
	role, err := svc.CreateRole(ctx, CreateRoleInput{
		Name:        "reader",
		Description: "Can read users",
		Permissions: []string{PermissionUsersRead},
	})
	if err != nil {
		t.Fatalf("create role: %v", err)
	}
	if len(role.Permissions) != 1 || role.Permissions[0].Code != PermissionUsersRead {
		t.Fatalf("expected role permission users:read, got %#v", role.Permissions)
	}
	user, err := svc.CreateUser(ctx, CreateUserInput{
		Username: "reader",
		Email:    "reader@example.com",
		Password: "password123",
		Roles:    []string{"reader"},
	})
	if err != nil {
		t.Fatalf("create reader user: %v", err)
	}
	if len(user.Roles) != 1 || user.Roles[0].Name != "reader" {
		t.Fatalf("expected reader role, got %#v", user.Roles)
	}
	token, err := svc.Login(ctx, LoginInput{Identifier: "reader", Password: "password123"})
	if err != nil {
		t.Fatalf("login reader: %v", err)
	}
	principal, err := svc.AuthenticateToken(ctx, token.Token)
	if err != nil {
		t.Fatalf("authenticate reader: %v", err)
	}
	if err := svc.Authorize(ctx, *principal, PermissionUsersRead); err != nil {
		t.Fatalf("expected reader to authorize users:read, got %v", err)
	}
	if err := svc.Authorize(ctx, *principal, PermissionUsersDelete); !errors.Is(err, ErrPermissionDenied) {
		t.Fatalf("reader authorize users:delete error = %v, want ErrPermissionDenied", err)
	}
}

func TestUserServiceApplyRBACSeedCreatesConfiguredEntriesIdempotently(t *testing.T) {
	ctx := context.Background()
	svc, _, _ := newUserServiceForTest(t)
	seed := RBACSeed{
		Roles: []RBACRoleSeed{
			{Name: "reader", Description: "Can read users"},
			{Name: "auditor", Description: "Can read roles"},
		},
		Permissions: []RBACPermissionSeed{
			{Code: PermissionUsersRead, Description: "Read users"},
			{Code: PermissionRolesRead, Description: "Read roles"},
		},
		RolePermissions: []RBACRolePermissionSeed{
			{Role: "reader", Permissions: []string{PermissionUsersRead}},
			{Role: "auditor", Permissions: []string{PermissionRolesRead}},
		},
	}

	if err := svc.ApplyRBACSeed(ctx, seed); err != nil {
		t.Fatalf("ApplyRBACSeed() error = %v", err)
	}
	if err := svc.ApplyRBACSeed(ctx, seed); err != nil {
		t.Fatalf("second ApplyRBACSeed() error = %v", err)
	}

	roles, err := svc.ListRoles(ctx)
	if err != nil {
		t.Fatalf("ListRoles() error = %v", err)
	}
	if got := countRolesByName(roles, "reader"); got != 1 {
		t.Fatalf("reader role count = %d, want 1", got)
	}
	reader := findRoleByName(roles, "reader")
	if reader == nil || len(reader.Permissions) != 1 || reader.Permissions[0].Code != PermissionUsersRead {
		t.Fatalf("reader permissions = %#v, want users:read", reader)
	}

	permissions, err := svc.ListPermissions(ctx)
	if err != nil {
		t.Fatalf("ListPermissions() error = %v", err)
	}
	if got := countPermissionsByCode(permissions, PermissionUsersRead); got != 1 {
		t.Fatalf("users:read permission count = %d, want 1", got)
	}

	user, err := svc.CreateUser(ctx, CreateUserInput{
		Username: "reader",
		Email:    "reader@example.com",
		Password: "password123",
		Roles:    []string{"reader"},
	})
	if err != nil {
		t.Fatalf("create reader user: %v", err)
	}
	if len(user.Roles) != 1 || user.Roles[0].Name != "reader" {
		t.Fatalf("created user roles = %#v, want reader", user.Roles)
	}
	token, err := svc.Login(ctx, LoginInput{Identifier: "reader", Password: "password123"})
	if err != nil {
		t.Fatalf("login reader: %v", err)
	}
	principal, err := svc.AuthenticateToken(ctx, token.Token)
	if err != nil {
		t.Fatalf("authenticate reader: %v", err)
	}
	if err := svc.Authorize(ctx, *principal, PermissionUsersRead); err != nil {
		t.Fatalf("reader authorize users:read error = %v", err)
	}
	if err := svc.Authorize(ctx, *principal, PermissionRolesRead); !errors.Is(err, ErrPermissionDenied) {
		t.Fatalf("reader authorize roles:read error = %v, want ErrPermissionDenied", err)
	}
}

func newUserServiceForTest(t *testing.T) (UserService, repository.Repository, database.Database) {
	t.Helper()

	db, err := database.New(&database.Config{
		Driver: database.DriverSQLite,
		DBName: filepath.Join(t.TempDir(), "users.db"),
	})
	if err != nil {
		t.Fatalf("create sqlite database: %v", err)
	}
	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Fatalf("close sqlite database: %v", err)
		}
	})
	if _, err := dbapp.ApplyUserSchema(context.Background(), db, string(database.DriverSQLite)); err != nil {
		t.Fatalf("apply user schema: %v", err)
	}
	repo := repository.NewRepository()
	hasher, err := passwordcrypto.NewBcrypt(passwordcrypto.WithBcryptCost(passwordcrypto.MinBcryptCost))
	if err != nil {
		t.Fatalf("create password hasher: %v", err)
	}
	tokens, err := authapi.NewJWTService(authapi.JWTConfig{
		Secret: []byte("0123456789abcdef0123456789abcdef"),
		Issuer: "go-scaffold",
	})
	if err != nil {
		t.Fatalf("create token service: %v", err)
	}
	return NewUserService(db, repo, hasher, tokens, nil), repo, db
}

func findRoleByName(roles []RoleDTO, name string) *RoleDTO {
	for i := range roles {
		if roles[i].Name == name {
			return &roles[i]
		}
	}
	return nil
}

func countRolesByName(roles []RoleDTO, name string) int {
	count := 0
	for _, role := range roles {
		if role.Name == name {
			count++
		}
	}
	return count
}

func countPermissionsByCode(permissions []PermissionDTO, code string) int {
	count := 0
	for _, permission := range permissions {
		if permission.Code == code {
			count++
		}
	}
	return count
}
