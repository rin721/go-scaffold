package config

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	rbacRoleNamePattern   = regexp.MustCompile(`^[A-Za-z0-9_.-]{2,64}$`)
	rbacPermissionPattern = regexp.MustCompile(`^[A-Za-z0-9_.*-]+:[A-Za-z0-9_.*-]+$`)
)

type RBACConfig struct {
	Enabled         bool                       `mapstructure:"enabled" envname:"RBAC_ENABLED" json:"enabled" yaml:"enabled" toml:"enabled"`
	ApplyOnStart    bool                       `mapstructure:"apply_on_start" envname:"RBAC_APPLY_ON_START" json:"apply_on_start" yaml:"apply_on_start" toml:"apply_on_start"`
	ModelPath       string                     `mapstructure:"model_path" envname:"RBAC_MODEL_PATH" json:"model_path" yaml:"model_path" toml:"model_path"`
	Roles           []RBACRoleConfig           `mapstructure:"roles" json:"roles" yaml:"roles" toml:"roles"`
	Permissions     []RBACPermissionConfig     `mapstructure:"permissions" json:"permissions" yaml:"permissions" toml:"permissions"`
	RolePermissions []RBACRolePermissionConfig `mapstructure:"role_permissions" json:"role_permissions" yaml:"role_permissions" toml:"role_permissions"`
}

type RBACRoleConfig struct {
	Name        string `mapstructure:"name" json:"name" yaml:"name" toml:"name"`
	Description string `mapstructure:"description" json:"description" yaml:"description" toml:"description"`
}

type RBACPermissionConfig struct {
	Code        string `mapstructure:"code" json:"code" yaml:"code" toml:"code"`
	Description string `mapstructure:"description" json:"description" yaml:"description" toml:"description"`
}

type RBACRolePermissionConfig struct {
	Role        string   `mapstructure:"role" json:"role" yaml:"role" toml:"role"`
	Permissions []string `mapstructure:"permissions" json:"permissions" yaml:"permissions" toml:"permissions"`
}

func (c *RBACConfig) ValidateName() string {
	return AppRBACName
}

func (c *RBACConfig) ValidateRequired() bool {
	return false
}

func (c *RBACConfig) Validate() error {
	if !c.Enabled {
		return nil
	}

	roles := make(map[string]struct{}, len(c.Roles))
	for i, role := range c.Roles {
		name, err := validateRBACRoleName(role.Name)
		if err != nil {
			return fmt.Errorf("rbac role %d: %w", i, err)
		}
		if _, ok := roles[name]; ok {
			return fmt.Errorf("rbac role %d: duplicate role %q", i, name)
		}
		roles[name] = struct{}{}
	}

	permissions := make(map[string]struct{}, len(c.Permissions))
	for i, permission := range c.Permissions {
		code, err := validateRBACPermissionCode(permission.Code)
		if err != nil {
			return fmt.Errorf("rbac permission %d: %w", i, err)
		}
		if _, ok := permissions[code]; ok {
			return fmt.Errorf("rbac permission %d: duplicate permission %q", i, code)
		}
		permissions[code] = struct{}{}
	}

	for i, grant := range c.RolePermissions {
		role, err := validateRBACRoleName(grant.Role)
		if err != nil {
			return fmt.Errorf("rbac role_permissions %d: %w", i, err)
		}
		if _, ok := roles[role]; !ok {
			return fmt.Errorf("rbac role_permissions %d: role %q is not declared", i, role)
		}
		if len(grant.Permissions) == 0 {
			return fmt.Errorf("rbac role_permissions %d: permissions are required", i)
		}
		seen := map[string]struct{}{}
		for j, permission := range grant.Permissions {
			code, err := validateRBACPermissionCode(permission)
			if err != nil {
				return fmt.Errorf("rbac role_permissions %d permission %d: %w", i, j, err)
			}
			if _, ok := permissions[code]; !ok {
				return fmt.Errorf("rbac role_permissions %d permission %d: permission %q is not declared", i, j, code)
			}
			if _, ok := seen[code]; ok {
				return fmt.Errorf("rbac role_permissions %d permission %d: duplicate permission %q", i, j, code)
			}
			seen[code] = struct{}{}
		}
	}
	return nil
}

func overrideRBACConfig(cfg *RBACConfig) {
	overrideConfigFromEnv(cfg)
}

func copyRBACConfig(src RBACConfig) RBACConfig {
	dst := src
	dst.Roles = append([]RBACRoleConfig(nil), src.Roles...)
	dst.Permissions = append([]RBACPermissionConfig(nil), src.Permissions...)
	dst.RolePermissions = make([]RBACRolePermissionConfig, len(src.RolePermissions))
	for i, grant := range src.RolePermissions {
		dst.RolePermissions[i] = grant
		dst.RolePermissions[i].Permissions = append([]string(nil), grant.Permissions...)
	}
	return dst
}

func validateRBACRoleName(value string) (string, error) {
	name := strings.ToLower(strings.TrimSpace(value))
	if name == "" {
		return "", fmt.Errorf("role name is required")
	}
	if !rbacRoleNamePattern.MatchString(name) {
		return "", fmt.Errorf("role name must be 2-64 characters and contain only letters, numbers, dot, underscore or dash")
	}
	return name, nil
}

func validateRBACPermissionCode(value string) (string, error) {
	code := strings.ToLower(strings.TrimSpace(value))
	if code == "" {
		return "", fmt.Errorf("permission code is required")
	}
	if !rbacPermissionPattern.MatchString(code) {
		return "", fmt.Errorf("permission code must use resource:action format")
	}
	return code, nil
}
