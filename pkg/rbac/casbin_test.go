package rbac

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestCasbinAuthorizerAuthorizesRolePolicyWithGlob(t *testing.T) {
	authorizer, err := NewCasbinAuthorizer("")
	if err != nil {
		t.Fatalf("NewCasbinAuthorizer() error = %v", err)
	}

	policies := []Policy{
		{Role: "admin", Permission: "*:*"},
		{Role: "reader", Permission: "users:read"},
		{Role: "editor", Permission: "users:*"},
	}

	allowed, err := authorizer.Authorize(context.Background(), Principal{ID: "1", Roles: []string{"reader"}}, "users:read", policies)
	if err != nil {
		t.Fatalf("Authorize() error = %v", err)
	}
	if !allowed {
		t.Fatal("reader should be allowed to read users")
	}

	allowed, err = authorizer.Authorize(context.Background(), Principal{ID: "1", Roles: []string{"reader"}}, "users:delete", policies)
	if err != nil {
		t.Fatalf("Authorize() error = %v", err)
	}
	if allowed {
		t.Fatal("reader should not be allowed to delete users")
	}

	allowed, err = authorizer.Authorize(context.Background(), Principal{ID: "2", Roles: []string{"editor"}}, "users:delete", policies)
	if err != nil {
		t.Fatalf("Authorize() error = %v", err)
	}
	if !allowed {
		t.Fatal("editor should be allowed by users:*")
	}

	allowed, err = authorizer.Authorize(context.Background(), Principal{ID: "3", Roles: []string{"admin"}}, "roles:delete", policies)
	if err != nil {
		t.Fatalf("Authorize() error = %v", err)
	}
	if !allowed {
		t.Fatal("admin should be allowed by *:*")
	}
}

func TestCasbinAuthorizerLoadsModelFile(t *testing.T) {
	modelPath := filepath.Join(t.TempDir(), "rbac_model.conf")
	if err := os.WriteFile(modelPath, []byte(DefaultCasbinModel), 0o600); err != nil {
		t.Fatalf("write model: %v", err)
	}

	authorizer, err := NewCasbinAuthorizer(modelPath)
	if err != nil {
		t.Fatalf("NewCasbinAuthorizer() error = %v", err)
	}

	allowed, err := authorizer.Authorize(context.Background(), Principal{ID: "1", Roles: []string{"admin"}}, "users:read", []Policy{{Role: "admin", Permission: "*:*"}})
	if err != nil {
		t.Fatalf("Authorize() error = %v", err)
	}
	if !allowed {
		t.Fatal("admin should be allowed with model loaded from file")
	}
}
