package memory

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/rei0721/go-scaffold/pkg/iam"
)

func TestAuthenticate(t *testing.T) {
	service := newTestService(t, true, []iam.Policy{})

	principal, err := service.Authenticate(context.Background(), iam.Credential{Scheme: iam.CredentialSchemeBearer, Token: "admin-token"})
	if err != nil {
		t.Fatalf("Authenticate() error = %v", err)
	}
	if principal.ID != "admin" {
		t.Fatalf("Principal.ID = %q, want admin", principal.ID)
	}

	if _, err := service.Authenticate(context.Background(), iam.Credential{Token: "missing"}); !errors.Is(err, iam.ErrInvalidCredential) {
		t.Fatalf("missing token error = %v, want ErrInvalidCredential", err)
	}
	if _, err := service.Authenticate(context.Background(), iam.Credential{}); !errors.Is(err, iam.ErrUnauthenticated) {
		t.Fatalf("empty token error = %v, want ErrUnauthenticated", err)
	}
}

func TestAuthorizeDefaultDenyAllowDenyAndWildcard(t *testing.T) {
	service := newTestService(t, true, []iam.Policy{
		{Subject: "admin", Action: "read", Resource: "plugin:echo", Effect: iam.EffectAllow},
		{Subject: "admin", Action: "delete", Resource: "*", Effect: iam.EffectDeny},
		{Subject: "*", Action: "health", Resource: "*", Effect: iam.EffectAllow},
	})
	admin := iam.Principal{ID: "admin"}

	if decision, err := service.Authorize(context.Background(), admin, "read", "plugin:echo"); err != nil || !decision.Allowed {
		t.Fatalf("Authorize(read) = %#v, %v; want allowed", decision, err)
	}
	if decision, err := service.Authorize(context.Background(), admin, "delete", "plugin:echo"); !errors.Is(err, iam.ErrPermissionDenied) || decision.Allowed {
		t.Fatalf("Authorize(delete) = %#v, %v; want denied", decision, err)
	}
	if decision, err := service.Authorize(context.Background(), iam.Principal{ID: "guest"}, "health", "plugin:any"); err != nil || !decision.Allowed {
		t.Fatalf("Authorize(wildcard) = %#v, %v; want allowed", decision, err)
	}
	if decision, err := service.Authorize(context.Background(), admin, "write", "plugin:echo"); !errors.Is(err, iam.ErrPermissionDenied) || decision.Allowed {
		t.Fatalf("Authorize(write) = %#v, %v; want default deny", decision, err)
	}
}

func TestAuthorizeDefaultAllowAndExpiredPolicy(t *testing.T) {
	service := newTestService(t, false, []iam.Policy{
		{Subject: "admin", Action: "write", Resource: "plugin:echo", Effect: iam.EffectDeny, ExpiresAt: time.Now().Add(-time.Minute)},
	})
	decision, err := service.Authorize(context.Background(), iam.Principal{ID: "admin"}, "write", "plugin:echo")
	if err != nil || !decision.Allowed {
		t.Fatalf("Authorize(expired deny) = %#v, %v; want default allow", decision, err)
	}
}

func TestContextHelpers(t *testing.T) {
	principal := iam.Principal{ID: "admin"}
	ctx := iam.ContextWithPrincipal(context.Background(), principal)
	got, ok := iam.PrincipalFromContext(ctx)
	if !ok || got.ID != "admin" {
		t.Fatalf("PrincipalFromContext() = %#v, %v", got, ok)
	}
	required, err := iam.RequirePrincipal(ctx)
	if err != nil || required.ID != "admin" {
		t.Fatalf("RequirePrincipal() = %#v, %v", required, err)
	}
	if _, err := iam.RequirePrincipal(context.Background()); !errors.Is(err, iam.ErrUnauthenticated) {
		t.Fatalf("RequirePrincipal(empty) error = %v, want ErrUnauthenticated", err)
	}
}

func newTestService(t *testing.T, defaultDeny bool, policies []iam.Policy) *Service {
	t.Helper()
	service, err := NewService(Config{
		DefaultDeny: defaultDeny,
		Principals: map[string]iam.Principal{
			"admin-token": {ID: "admin", Name: "Admin"},
		},
		Policies: policies,
	})
	if err != nil {
		t.Fatalf("NewService() error = %v", err)
	}
	return service
}
