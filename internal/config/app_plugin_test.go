package config

import (
	"strings"
	"testing"
)

func TestPluginConfigValidateRequiresExplicitRegistrationExposure(t *testing.T) {
	cfg := PluginConfig{
		Enabled: true,
		Registration: PluginRegistrationConfig{
			Enabled: true,
			Token:   "registration-secret",
		},
	}

	err := cfg.Validate()
	if err == nil {
		t.Fatal("Validate() error = nil, want explicit exposure error")
	}
	if !strings.Contains(err.Error(), "registration requires either interface http enabled or expose_on_main_http") {
		t.Fatalf("Validate() error = %v", err)
	}
}

func TestPluginConfigValidateAcceptsDedicatedPluginInterface(t *testing.T) {
	cfg := PluginConfig{
		Enabled: true,
		Interface: PluginInterfaceConfig{
			HTTP: PluginHTTPInterfaceConfig{
				Enabled:   true,
				Host:      "127.0.0.1",
				Port:      18080,
				PublicURL: "http://127.0.0.1:18080",
			},
			WS: PluginWSInterfaceConfig{
				PublicURL: "ws://127.0.0.1:18080/plugin/v1/ws",
			},
		},
		Registration: PluginRegistrationConfig{
			Enabled: true,
			Token:   "registration-secret",
		},
	}

	if err := cfg.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
}

func TestPluginConfigValidateRejectsHTTPInterfaceWithoutRegistration(t *testing.T) {
	cfg := PluginConfig{
		Enabled: true,
		Interface: PluginInterfaceConfig{
			HTTP: PluginHTTPInterfaceConfig{
				Enabled:   true,
				Host:      "127.0.0.1",
				Port:      18080,
				PublicURL: "http://127.0.0.1:18080",
			},
		},
	}

	err := cfg.Validate()
	if err == nil {
		t.Fatal("Validate() error = nil, want registration requirement")
	}
	if !strings.Contains(err.Error(), "interface http requires registration to be enabled") {
		t.Fatalf("Validate() error = %v", err)
	}
}
