package initapp

import (
	"context"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/rei0721/go-scaffold/internal/config"
	"github.com/rei0721/go-scaffold/pkg/iam"
	"github.com/rei0721/go-scaffold/pkg/plugin"
)

func TestNewPluginManagerRegistersConfiguredHTTPPluginAndRemoteHook(t *testing.T) {
	var hookExecuted bool
	remote, err := plugin.NewLocal(plugin.Metadata{Name: "remote", Protocol: plugin.ProtocolLocal}, func(ctx context.Context, req plugin.Request) (*plugin.Response, error) {
		if req.Operation == plugin.OperationHooksExecute {
			hookExecuted = true
			return plugin.NewResponse(map[string]string{"ok": "true"})
		}
		return plugin.NewResponse(map[string]string{"operation": req.Operation})
	})
	if err != nil {
		t.Fatalf("NewLocal(remote): %v", err)
	}
	server := httptest.NewServer(plugin.NewHTTPServer(remote))
	defer server.Close()

	cfg := &config.Config{
		Plugin: config.PluginConfig{
			Enabled: true,
			Plugins: []config.PluginDefinitionConfig{
				{Name: "remote", Protocol: "http", Endpoint: server.URL + plugin.HTTPInvokePath},
			},
			Hooks: []config.PluginHookBindingConfig{
				{Point: string(plugin.HookAfterInvoke), Plugin: "remote", Name: "remote-audit", Priority: 1},
			},
		},
	}

	manager, err := NewPluginManager(cfg, nil, nil)
	if err != nil {
		t.Fatalf("NewPluginManager() error = %v", err)
	}
	resp, err := manager.Invoke(context.Background(), "remote", plugin.MustNewRequest("run", nil))
	if err != nil {
		t.Fatalf("Invoke() error = %v", err)
	}
	var got map[string]string
	if err := resp.DecodePayload(&got); err != nil {
		t.Fatalf("DecodePayload() error = %v", err)
	}
	if got["operation"] != "run" {
		t.Fatalf("response = %#v, want run operation", got)
	}
	if !hookExecuted {
		t.Fatal("remote hook was not executed")
	}
}

func TestPluginManagerIAMHook(t *testing.T) {
	remote, err := plugin.NewLocal(plugin.Metadata{Name: "remote", Protocol: plugin.ProtocolLocal}, func(ctx context.Context, req plugin.Request) (*plugin.Response, error) {
		return plugin.NewResponse(map[string]string{"operation": req.Operation})
	})
	if err != nil {
		t.Fatalf("NewLocal(remote): %v", err)
	}
	server := httptest.NewServer(plugin.NewHTTPServer(remote))
	defer server.Close()

	defaultDeny := true
	cfg := &config.Config{
		IAM: config.IAMConfig{
			Enabled:     true,
			Mode:        "memory",
			DefaultDeny: &defaultDeny,
			Tokens: []config.IAMTokenConfig{
				{Token: "admin-token", Principal: config.IAMPrincipalConfig{ID: "admin"}},
			},
			Policies: []config.IAMPolicyConfig{
				{Subject: "admin", Action: "run", Resource: "remote", Effect: "allow"},
			},
		},
		Plugin: config.PluginConfig{
			Enabled: true,
			Plugins: []config.PluginDefinitionConfig{
				{Name: "remote", Protocol: "http", Endpoint: server.URL + plugin.HTTPInvokePath},
			},
		},
	}
	iamService, err := NewIAM(cfg, nil)
	if err != nil {
		t.Fatalf("NewIAM() error = %v", err)
	}
	manager, err := NewPluginManager(cfg, nil, iamService)
	if err != nil {
		t.Fatalf("NewPluginManager() error = %v", err)
	}

	if _, err := manager.Invoke(context.Background(), "remote", plugin.MustNewRequest("run", nil)); !errors.Is(err, iam.ErrUnauthenticated) {
		t.Fatalf("Invoke without principal error = %v, want ErrUnauthenticated", err)
	}

	ctx := iam.ContextWithPrincipal(context.Background(), iam.Principal{ID: "admin"})
	if _, err := manager.Invoke(ctx, "remote", plugin.MustNewRequest("run", nil)); err != nil {
		t.Fatalf("Invoke with principal error = %v", err)
	}
}
