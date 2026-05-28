package config

import (
	"strings"
	"testing"
	"time"
)

func TestAuthConfigValidateAllowsLocalFallback(t *testing.T) {
	cfg := AuthConfig{}
	if err := cfg.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
	if got := cfg.TokenSecretBytes(); got != nil {
		t.Fatalf("TokenSecretBytes() = %q, want nil fallback", string(got))
	}
	if got := cfg.TokenTTLDuration(); got != 0 {
		t.Fatalf("TokenTTLDuration() = %v, want zero default signal", got)
	}
}

func TestAuthConfigValidateRejectsShortSecret(t *testing.T) {
	cfg := AuthConfig{TokenSecret: "short"}
	if err := cfg.Validate(); err == nil || !strings.Contains(err.Error(), "token_secret") {
		t.Fatalf("Validate() error = %v, want token_secret error", err)
	}
}

func TestAuthConfigValidateRejectsNegativeTTL(t *testing.T) {
	cfg := AuthConfig{TokenTTL: -1}
	if err := cfg.Validate(); err == nil || !strings.Contains(err.Error(), "token_ttl") {
		t.Fatalf("Validate() error = %v, want token_ttl error", err)
	}
}

func TestAuthConfigHelpersReturnConfiguredValues(t *testing.T) {
	secret := "0123456789abcdef0123456789abcdef"
	cfg := AuthConfig{TokenSecret: secret, TokenTTL: 3600}
	if err := cfg.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
	if got := string(cfg.TokenSecretBytes()); got != secret {
		t.Fatalf("TokenSecretBytes() = %q, want %q", got, secret)
	}
	if got := cfg.TokenTTLDuration(); got != time.Hour {
		t.Fatalf("TokenTTLDuration() = %v, want %v", got, time.Hour)
	}
}
