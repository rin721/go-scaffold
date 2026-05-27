package config

import "fmt"

type IAMConfig struct {
	Enabled     bool              `mapstructure:"enabled" envname:"IAM_ENABLED" json:"enabled" yaml:"enabled" toml:"enabled"`
	Mode        string            `mapstructure:"mode" envname:"IAM_MODE" json:"mode" yaml:"mode" toml:"mode"`
	DefaultDeny *bool             `mapstructure:"default_deny" envname:"IAM_DEFAULT_DENY" json:"default_deny" yaml:"default_deny" toml:"default_deny"`
	Tokens      []IAMTokenConfig  `mapstructure:"tokens" json:"tokens" yaml:"tokens" toml:"tokens"`
	Policies    []IAMPolicyConfig `mapstructure:"policies" json:"policies" yaml:"policies" toml:"policies"`
}

type IAMTokenConfig struct {
	Token     string             `mapstructure:"token" json:"token" yaml:"token" toml:"token"`
	Principal IAMPrincipalConfig `mapstructure:"principal" json:"principal" yaml:"principal" toml:"principal"`
}

type IAMPrincipalConfig struct {
	ID         string            `mapstructure:"id" json:"id" yaml:"id" toml:"id"`
	Name       string            `mapstructure:"name" json:"name" yaml:"name" toml:"name"`
	Roles      []string          `mapstructure:"roles" json:"roles" yaml:"roles" toml:"roles"`
	Attributes map[string]string `mapstructure:"attributes" json:"attributes" yaml:"attributes" toml:"attributes"`
}

type IAMPolicyConfig struct {
	Subject   string `mapstructure:"subject" json:"subject" yaml:"subject" toml:"subject"`
	Action    string `mapstructure:"action" json:"action" yaml:"action" toml:"action"`
	Resource  string `mapstructure:"resource" json:"resource" yaml:"resource" toml:"resource"`
	Effect    string `mapstructure:"effect" json:"effect" yaml:"effect" toml:"effect"`
	ExpiresAt string `mapstructure:"expires_at" json:"expires_at" yaml:"expires_at" toml:"expires_at"`
}

func (c *IAMConfig) ValidateName() string {
	return AppIAMName
}

func (c *IAMConfig) ValidateRequired() bool {
	return false
}

func (c *IAMConfig) Validate() error {
	if !c.Enabled {
		return nil
	}
	if c.Mode == "" {
		c.Mode = "memory"
	}
	if c.Mode != "memory" {
		return fmt.Errorf("iam: unsupported mode %q", c.Mode)
	}
	for i, token := range c.Tokens {
		if token.Token == "" {
			return fmt.Errorf("iam token %d: token is required", i)
		}
		if token.Principal.ID == "" {
			return fmt.Errorf("iam token %d: principal.id is required", i)
		}
	}
	for i, policy := range c.Policies {
		if policy.Subject == "" || policy.Action == "" || policy.Resource == "" {
			return fmt.Errorf("iam policy %d: subject, action and resource are required", i)
		}
		if policy.Effect != "allow" && policy.Effect != "deny" {
			return fmt.Errorf("iam policy %d: effect must be allow or deny", i)
		}
	}
	return nil
}

func (c *IAMConfig) DefaultDenyEnabled() bool {
	if c.DefaultDeny == nil {
		return true
	}
	return *c.DefaultDeny
}

func overrideIAMConfig(cfg *IAMConfig) {
	overrideConfigFromEnv(cfg)
}

func copyIAMTokens(src []IAMTokenConfig) []IAMTokenConfig {
	if len(src) == 0 {
		return nil
	}
	dst := make([]IAMTokenConfig, len(src))
	for i, token := range src {
		dst[i] = token
		dst[i].Principal.Roles = append([]string(nil), token.Principal.Roles...)
		dst[i].Principal.Attributes = copyStringMap(token.Principal.Attributes)
	}
	return dst
}
