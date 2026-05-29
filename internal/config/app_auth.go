package config

import (
	"fmt"
	"time"
)

const minimumAuthTokenSecretBytes = 32

type AuthConfig struct {
	TokenSecret        string `mapstructure:"token_secret" envname:"AUTH_TOKEN_SECRET" json:"token_secret" yaml:"token_secret" toml:"token_secret"`
	TokenTTL           int    `mapstructure:"token_ttl" envname:"AUTH_TOKEN_TTL" json:"token_ttl" yaml:"token_ttl" toml:"token_ttl"`
	PublicRegistration *bool  `mapstructure:"public_registration" envname:"AUTH_PUBLIC_REGISTRATION" json:"public_registration" yaml:"public_registration" toml:"public_registration"`
}

func (c *AuthConfig) ValidateName() string {
	return AppAuthName
}

func (c *AuthConfig) ValidateRequired() bool {
	return false
}

func (c *AuthConfig) Validate() error {
	if c.TokenTTL < 0 {
		return fmt.Errorf("auth: token_ttl must be non-negative")
	}
	if c.TokenSecret != "" && len([]byte(c.TokenSecret)) < minimumAuthTokenSecretBytes {
		return fmt.Errorf("auth: token_secret must be at least %d bytes when set", minimumAuthTokenSecretBytes)
	}
	return nil
}

func (c AuthConfig) TokenSecretBytes() []byte {
	if c.TokenSecret == "" {
		return nil
	}
	return []byte(c.TokenSecret)
}

func (c AuthConfig) PublicRegistrationEnabled() bool {
	if c.PublicRegistration == nil {
		return true
	}
	return *c.PublicRegistration
}

func (c AuthConfig) TokenTTLDuration() time.Duration {
	if c.TokenTTL <= 0 {
		return 0
	}
	return time.Duration(c.TokenTTL) * time.Second
}

func overrideAuthConfig(cfg *AuthConfig) {
	overrideConfigFromEnv(cfg)
}
