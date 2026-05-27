package memory

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/rei0721/go-scaffold/pkg/iam"
)

// Config configures the in-memory IAM service.
type Config struct {
	DefaultDeny bool
	Principals  map[string]iam.Principal
	Policies    []iam.Policy
}

// DefaultConfig returns conservative defaults.
func DefaultConfig() Config {
	return Config{DefaultDeny: true}
}

// Service implements iam.Service with in-memory data.
type Service struct {
	mu          sync.RWMutex
	defaultDeny bool
	principals  map[string]iam.Principal
	policies    []iam.Policy
}

// NewService creates an in-memory IAM service.
func NewService(cfg Config) (*Service, error) {
	principals := make(map[string]iam.Principal, len(cfg.Principals))
	for token, principal := range cfg.Principals {
		if strings.TrimSpace(token) == "" {
			return nil, iam.ErrInvalidCredential
		}
		if principal.ID == "" {
			return nil, iam.ErrInvalidPrincipal
		}
		principals[token] = copyPrincipal(principal)
	}
	policies := make([]iam.Policy, len(cfg.Policies))
	for i, policy := range cfg.Policies {
		if policy.Subject == "" || policy.Action == "" || policy.Resource == "" {
			return nil, iam.ErrInvalidPolicy
		}
		if policy.Effect != iam.EffectAllow && policy.Effect != iam.EffectDeny {
			return nil, iam.ErrInvalidPolicy
		}
		policies[i] = policy
	}
	return &Service{
		defaultDeny: cfg.DefaultDeny,
		principals:  principals,
		policies:    policies,
	}, nil
}

// Authenticate verifies a token credential.
func (s *Service) Authenticate(ctx context.Context, credential iam.Credential) (iam.Principal, error) {
	if err := ctx.Err(); err != nil {
		return iam.Principal{}, err
	}
	if credential.Token == "" {
		return iam.Principal{}, iam.ErrUnauthenticated
	}
	scheme := strings.ToLower(strings.TrimSpace(credential.Scheme))
	if scheme != "" && scheme != iam.CredentialSchemeToken && scheme != iam.CredentialSchemeBearer {
		return iam.Principal{}, iam.ErrInvalidCredential
	}

	s.mu.RLock()
	defer s.mu.RUnlock()
	principal, ok := s.principals[credential.Token]
	if !ok {
		return iam.Principal{}, iam.ErrInvalidCredential
	}
	return copyPrincipal(principal), nil
}

// Authorize checks policies for a principal/action/resource tuple.
func (s *Service) Authorize(ctx context.Context, principal iam.Principal, action iam.Action, resource iam.Resource) (iam.Decision, error) {
	if err := ctx.Err(); err != nil {
		return iam.Decision{}, err
	}
	if principal.ID == "" {
		return iam.Decision{}, iam.ErrUnauthenticated
	}
	now := time.Now().UTC()

	s.mu.RLock()
	defer s.mu.RUnlock()

	allowed := !s.defaultDeny
	reason := "default allow"
	if s.defaultDeny {
		reason = "default deny"
	}
	for _, policy := range s.policies {
		if policyExpired(policy, now) ||
			!matches(policy.Subject, principal.ID) ||
			!matches(string(policy.Action), string(action)) ||
			!matches(string(policy.Resource), string(resource)) {
			continue
		}
		if policy.Effect == iam.EffectDeny {
			return iam.Decision{Allowed: false, Reason: "policy denied", Principal: copyPrincipal(principal)}, iam.ErrPermissionDenied
		}
		allowed = true
		reason = "policy allowed"
	}
	if !allowed {
		return iam.Decision{Allowed: false, Reason: reason, Principal: copyPrincipal(principal)}, iam.ErrPermissionDenied
	}
	return iam.Decision{Allowed: true, Reason: reason, Principal: copyPrincipal(principal)}, nil
}

func policyExpired(policy iam.Policy, now time.Time) bool {
	return !policy.ExpiresAt.IsZero() && !policy.ExpiresAt.After(now)
}

func matches(pattern, value string) bool {
	return pattern == "*" || pattern == value
}

func copyPrincipal(src iam.Principal) iam.Principal {
	dst := src
	if len(src.Roles) > 0 {
		dst.Roles = append([]string(nil), src.Roles...)
	}
	if len(src.Attributes) > 0 {
		dst.Attributes = make(map[string]string, len(src.Attributes))
		for k, v := range src.Attributes {
			dst.Attributes[k] = v
		}
	}
	return dst
}
