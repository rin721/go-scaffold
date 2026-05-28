package rbac

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
)

// DefaultCasbinModel is the built-in role-permission model used when no model
// path is supplied.
const DefaultCasbinModel = `[request_definition]
r = sub, perm

[policy_definition]
p = sub, perm

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && globMatch(r.perm, p.perm)
`

// Principal is the subject and role set used for authorization checks.
type Principal struct {
	ID    string
	Roles []string
}

// Policy grants a permission pattern to a role.
type Policy struct {
	Role       string
	Permission string
}

// Authorizer evaluates whether a principal has a permission.
type Authorizer interface {
	Authorize(ctx context.Context, principal Principal, permission string, policies []Policy) (bool, error)
}

// CasbinAuthorizer evaluates RBAC policies through Casbin.
type CasbinAuthorizer struct {
	modelText string
}

// NewCasbinAuthorizer creates a Casbin-backed authorizer.
//
// When modelPath is empty, DefaultCasbinModel is used. When modelPath is set,
// the file is read once and validated during construction.
func NewCasbinAuthorizer(modelPath string) (*CasbinAuthorizer, error) {
	modelText := DefaultCasbinModel
	if path := strings.TrimSpace(modelPath); path != "" {
		content, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("read casbin model %q: %w", path, err)
		}
		modelText = string(content)
	}
	return NewCasbinAuthorizerFromModel(modelText)
}

// NewCasbinAuthorizerFromModel creates a Casbin-backed authorizer from model
// text and validates it before returning.
func NewCasbinAuthorizerFromModel(modelText string) (*CasbinAuthorizer, error) {
	if _, err := casbinmodel.NewModelFromString(modelText); err != nil {
		return nil, fmt.Errorf("parse casbin model: %w", err)
	}
	return &CasbinAuthorizer{modelText: modelText}, nil
}

// NewDefaultAuthorizer creates a Casbin authorizer with DefaultCasbinModel.
func NewDefaultAuthorizer() *CasbinAuthorizer {
	return &CasbinAuthorizer{modelText: DefaultCasbinModel}
}

// Authorize evaluates a principal, a permission, and a policy set.
func (a *CasbinAuthorizer) Authorize(ctx context.Context, principal Principal, permission string, policies []Policy) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	permission = normalize(permission)
	if permission == "" || strings.TrimSpace(principal.ID) == "" {
		return false, nil
	}

	enforcer, err := a.newEnforcer()
	if err != nil {
		return false, err
	}
	subject := "user:" + strings.TrimSpace(principal.ID)
	for _, role := range uniqueNormalized(principal.Roles) {
		if _, err := enforcer.AddGroupingPolicy(subject, role); err != nil {
			return false, fmt.Errorf("add casbin grouping policy: %w", err)
		}
	}
	for _, policy := range uniquePolicies(policies) {
		if _, err := enforcer.AddPolicy(policy.Role, policy.Permission); err != nil {
			return false, fmt.Errorf("add casbin permission policy: %w", err)
		}
	}

	allowed, err := enforcer.Enforce(subject, permission)
	if err != nil {
		return false, fmt.Errorf("enforce casbin policy: %w", err)
	}
	return allowed, nil
}

func (a *CasbinAuthorizer) newEnforcer() (*casbin.Enforcer, error) {
	modelText := DefaultCasbinModel
	if a != nil && strings.TrimSpace(a.modelText) != "" {
		modelText = a.modelText
	}
	model, err := casbinmodel.NewModelFromString(modelText)
	if err != nil {
		return nil, fmt.Errorf("parse casbin model: %w", err)
	}
	enforcer, err := casbin.NewEnforcer(model)
	if err != nil {
		return nil, fmt.Errorf("create casbin enforcer: %w", err)
	}
	enforcer.EnableAutoSave(false)
	return enforcer, nil
}

func uniqueNormalized(values []string) []string {
	seen := make(map[string]struct{}, len(values))
	result := make([]string, 0, len(values))
	for _, value := range values {
		value = normalize(value)
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}

func uniquePolicies(policies []Policy) []Policy {
	seen := make(map[string]struct{}, len(policies))
	result := make([]Policy, 0, len(policies))
	for _, policy := range policies {
		policy.Role = normalize(policy.Role)
		policy.Permission = normalize(policy.Permission)
		if policy.Role == "" || policy.Permission == "" {
			continue
		}
		key := policy.Role + "\x00" + policy.Permission
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		result = append(result, policy)
	}
	return result
}

func normalize(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}
