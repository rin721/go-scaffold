package config

import (
	"strings"
	"testing"
)

func TestRBACConfigValidateAcceptsSeedConfig(t *testing.T) {
	cfg := validRBACConfig()
	if err := cfg.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
}

func TestRBACConfigValidateSkipsDisabledConfig(t *testing.T) {
	cfg := RBACConfig{
		Enabled: false,
		Roles:   []RBACRoleConfig{{Name: ""}},
	}
	if err := cfg.Validate(); err != nil {
		t.Fatalf("disabled Validate() error = %v", err)
	}
}

func TestRBACConfigValidateRejectsDuplicateRole(t *testing.T) {
	cfg := validRBACConfig()
	cfg.Roles = append(cfg.Roles, RBACRoleConfig{Name: "Admin"})

	err := cfg.Validate()
	if err == nil || !strings.Contains(err.Error(), "duplicate role") {
		t.Fatalf("Validate() error = %v, want duplicate role", err)
	}
}

func TestRBACConfigValidateRejectsUnknownPermissionGrant(t *testing.T) {
	cfg := validRBACConfig()
	cfg.RolePermissions[0].Permissions = append(cfg.RolePermissions[0].Permissions, "users:delete")

	err := cfg.Validate()
	if err == nil || !strings.Contains(err.Error(), "is not declared") {
		t.Fatalf("Validate() error = %v, want undeclared permission", err)
	}
}

func TestRBACConfigValidateRejectsUnknownRoleGrant(t *testing.T) {
	cfg := validRBACConfig()
	cfg.RolePermissions = append(cfg.RolePermissions, RBACRolePermissionConfig{
		Role:        "manager",
		Permissions: []string{"users:read"},
	})

	err := cfg.Validate()
	if err == nil || !strings.Contains(err.Error(), "role \"manager\" is not declared") {
		t.Fatalf("Validate() error = %v, want undeclared role", err)
	}
}

func TestOverrideRBACConfigSupportsModelPath(t *testing.T) {
	t.Setenv("RIN_APP_RBAC_MODEL_PATH", "./configs/custom_rbac_model.conf")
	cfg := RBACConfig{}

	overrideRBACConfig(&cfg)

	if cfg.ModelPath != "./configs/custom_rbac_model.conf" {
		t.Fatalf("ModelPath = %q, want env override", cfg.ModelPath)
	}
}

func validRBACConfig() RBACConfig {
	return RBACConfig{
		Enabled:      true,
		ApplyOnStart: true,
		ModelPath:    "./configs/rbac_model.conf",
		Roles: []RBACRoleConfig{
			{Name: "admin", Description: "Admin role"},
			{Name: "user", Description: "User role"},
		},
		Permissions: []RBACPermissionConfig{
			{Code: "*:*", Description: "All permissions"},
			{Code: "users:read", Description: "Read users"},
		},
		RolePermissions: []RBACRolePermissionConfig{
			{Role: "admin", Permissions: []string{"*:*", "users:read"}},
		},
	}
}
