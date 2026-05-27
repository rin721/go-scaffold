package initapp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rei0721/go-scaffold/internal/config"
	"github.com/rei0721/go-scaffold/internal/middleware"
	"github.com/rei0721/go-scaffold/pkg/plugin"
)

func TestNewHTTPServerDoesNotExposePluginRegistrationByDefault(t *testing.T) {
	cfg := testPluginInterfaceConfig()
	manager := plugin.NewManager()

	router, _, err := NewHTTPServer(cfg, nil, nil, nil, middleware.CORSConfig{}, manager, nil)
	if err != nil {
		t.Fatalf("NewHTTPServer() error = %v", err)
	}

	request := httptest.NewRequest(http.MethodPost, plugin.HTTPRegisterPath, bytes.NewReader([]byte(`{}`)))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("registration status = %d, want 404 when expose_on_main_http is false", recorder.Code)
	}
}

func TestNewHTTPServerCanExposePluginRegistrationOnMainHTTP(t *testing.T) {
	cfg := testPluginInterfaceConfig()
	cfg.Plugin.Registration.ExposeOnMainHTTP = true
	manager := plugin.NewManager()

	router, _, err := NewHTTPServer(cfg, nil, nil, nil, middleware.CORSConfig{}, manager, nil)
	if err != nil {
		t.Fatalf("NewHTTPServer() error = %v", err)
	}
	body, err := json.Marshal(plugin.RegistrationRequest{
		Plugin: plugin.Definition{
			Name:     "blog",
			Protocol: plugin.ProtocolHTTP,
			Endpoint: "http://127.0.0.1:18081" + plugin.HTTPInvokePath,
		},
	})
	if err != nil {
		t.Fatalf("Marshal registration: %v", err)
	}
	request := httptest.NewRequest(http.MethodPost, plugin.HTTPRegisterPath, bytes.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer registration-secret")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusCreated {
		t.Fatalf("registration status = %d, want 201 with body %s", recorder.Code, recorder.Body.String())
	}
	if _, ok := manager.Get("blog"); !ok {
		t.Fatal("registered plugin not found in manager")
	}
}

func TestNewPluginHTTPServerUsesPluginInterfaceConfig(t *testing.T) {
	cfg := testPluginInterfaceConfig()

	server, err := NewPluginHTTPServer(cfg, nil, plugin.NewManager())
	if err != nil {
		t.Fatalf("NewPluginHTTPServer() error = %v", err)
	}
	if server == nil {
		t.Fatal("NewPluginHTTPServer() = nil, want server")
	}

	got := PluginHTTPServerConfig(cfg)
	if got.Host != "127.0.0.1" || got.Port != 18080 {
		t.Fatalf("PluginHTTPServerConfig() = %#v", got)
	}
}

func testPluginInterfaceConfig() *config.Config {
	return &config.Config{
		Server: config.ServerConfig{
			Host:         "127.0.0.1",
			Port:         9999,
			Mode:         "test",
			ReadTimeout:  1,
			WriteTimeout: 1,
			IdleTimeout:  1,
		},
		Plugin: config.PluginConfig{
			Enabled: true,
			Interface: config.PluginInterfaceConfig{
				HTTP: config.PluginHTTPInterfaceConfig{
					Enabled:   true,
					Host:      "127.0.0.1",
					Port:      18080,
					PublicURL: "http://127.0.0.1:18080",
				},
				WS: config.PluginWSInterfaceConfig{
					PublicURL: "ws://127.0.0.1:18080/plugin/v1/ws",
				},
			},
			Registration: config.PluginRegistrationConfig{
				Enabled: true,
				Token:   "registration-secret",
			},
		},
	}
}
