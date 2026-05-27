package iam

import (
	"context"
	"time"
)

type contextKey string

const principalContextKey contextKey = "iam.principal"

const (
	CredentialSchemeToken  = "token"
	CredentialSchemeBearer = "bearer"
)

// Principal identifies an authenticated caller.
type Principal struct {
	ID         string            `json:"id" yaml:"id" mapstructure:"id"`
	Name       string            `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name"`
	Roles      []string          `json:"roles,omitempty" yaml:"roles,omitempty" mapstructure:"roles"`
	Attributes map[string]string `json:"attributes,omitempty" yaml:"attributes,omitempty" mapstructure:"attributes"`
}

// Credential carries authentication input.
type Credential struct {
	Scheme string `json:"scheme" yaml:"scheme" mapstructure:"scheme"`
	Token  string `json:"token" yaml:"token" mapstructure:"token"`
}

// Action identifies an operation to authorize.
type Action string

// Resource identifies a protected resource.
type Resource string

// Decision is the result of authorization.
type Decision struct {
	Allowed   bool      `json:"allowed"`
	Reason    string    `json:"reason,omitempty"`
	Principal Principal `json:"principal,omitempty"`
}

// Authenticator verifies credentials.
type Authenticator interface {
	Authenticate(ctx context.Context, credential Credential) (Principal, error)
}

// Authorizer checks whether a principal can perform an action on a resource.
type Authorizer interface {
	Authorize(ctx context.Context, principal Principal, action Action, resource Resource) (Decision, error)
}

// Service combines authentication and authorization.
type Service interface {
	Authenticator
	Authorizer
}

// Effect is the policy effect.
type Effect string

const (
	EffectAllow Effect = "allow"
	EffectDeny  Effect = "deny"
)

// Policy is a simple authorization rule.
type Policy struct {
	Subject   string    `json:"subject" yaml:"subject" mapstructure:"subject"`
	Action    Action    `json:"action" yaml:"action" mapstructure:"action"`
	Resource  Resource  `json:"resource" yaml:"resource" mapstructure:"resource"`
	Effect    Effect    `json:"effect" yaml:"effect" mapstructure:"effect"`
	ExpiresAt time.Time `json:"expires_at,omitempty" yaml:"expires_at,omitempty" mapstructure:"expires_at"`
}

// ContextWithPrincipal stores a principal in context.
func ContextWithPrincipal(ctx context.Context, principal Principal) context.Context {
	return context.WithValue(ctx, principalContextKey, principal)
}

// PrincipalFromContext returns the principal stored in context.
func PrincipalFromContext(ctx context.Context) (Principal, bool) {
	principal, ok := ctx.Value(principalContextKey).(Principal)
	return principal, ok
}

// RequirePrincipal returns a principal or ErrUnauthenticated.
func RequirePrincipal(ctx context.Context) (Principal, error) {
	principal, ok := PrincipalFromContext(ctx)
	if !ok || principal.ID == "" {
		return Principal{}, ErrUnauthenticated
	}
	return principal, nil
}
