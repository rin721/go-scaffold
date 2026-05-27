package config

import (
	"fmt"
	"strings"
	"time"
)

type PluginConfig struct {
	Enabled          bool                      `mapstructure:"enabled" envname:"PLUGIN_ENABLED" json:"enabled" yaml:"enabled" toml:"enabled"`
	DefaultTimeout   int                       `mapstructure:"default_timeout" envname:"PLUGIN_DEFAULT_TIMEOUT" json:"default_timeout" yaml:"default_timeout" toml:"default_timeout"`
	MaxResponseBytes int64                     `mapstructure:"max_response_bytes" envname:"PLUGIN_MAX_RESPONSE_BYTES" json:"max_response_bytes" yaml:"max_response_bytes" toml:"max_response_bytes"`
	Interface        PluginInterfaceConfig     `mapstructure:"interface" json:"interface" yaml:"interface" toml:"interface"`
	Registration     PluginRegistrationConfig  `mapstructure:"registration" json:"registration" yaml:"registration" toml:"registration"`
	Plugins          []PluginDefinitionConfig  `mapstructure:"plugins" json:"plugins" yaml:"plugins" toml:"plugins"`
	Hooks            []PluginHookBindingConfig `mapstructure:"hooks" json:"hooks" yaml:"hooks" toml:"hooks"`
}

type PluginInterfaceConfig struct {
	HTTP PluginHTTPInterfaceConfig `mapstructure:"http" json:"http" yaml:"http" toml:"http"`
	WS   PluginWSInterfaceConfig   `mapstructure:"ws" json:"ws" yaml:"ws" toml:"ws"`
}

type PluginHTTPInterfaceConfig struct {
	Enabled   bool   `mapstructure:"enabled" envname:"PLUGIN_INTERFACE_HTTP_ENABLED" json:"enabled" yaml:"enabled" toml:"enabled"`
	Host      string `mapstructure:"host" envname:"PLUGIN_INTERFACE_HTTP_HOST" json:"host" yaml:"host" toml:"host"`
	Port      int    `mapstructure:"port" envname:"PLUGIN_INTERFACE_HTTP_PORT" json:"port" yaml:"port" toml:"port"`
	PublicURL string `mapstructure:"public_url" envname:"PLUGIN_INTERFACE_HTTP_PUBLIC_URL" json:"public_url" yaml:"public_url" toml:"public_url"`
}

type PluginWSInterfaceConfig struct {
	PublicURL string `mapstructure:"public_url" envname:"PLUGIN_INTERFACE_WS_PUBLIC_URL" json:"public_url" yaml:"public_url" toml:"public_url"`
}

type PluginRegistrationConfig struct {
	Enabled          bool   `mapstructure:"enabled" envname:"PLUGIN_REGISTRATION_ENABLED" json:"enabled" yaml:"enabled" toml:"enabled"`
	ExposeOnMainHTTP bool   `mapstructure:"expose_on_main_http" envname:"PLUGIN_REGISTRATION_EXPOSE_ON_MAIN_HTTP" json:"expose_on_main_http" yaml:"expose_on_main_http" toml:"expose_on_main_http"`
	Token            string `mapstructure:"token" envname:"PLUGIN_REGISTRATION_TOKEN" json:"token" yaml:"token" toml:"token"`
}

type PluginDefinitionConfig struct {
	Name         string            `mapstructure:"name" json:"name" yaml:"name" toml:"name"`
	Version      string            `mapstructure:"version" json:"version" yaml:"version" toml:"version"`
	Protocol     string            `mapstructure:"protocol" json:"protocol" yaml:"protocol" toml:"protocol"`
	Endpoint     string            `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint" toml:"endpoint"`
	Timeout      int               `mapstructure:"timeout" json:"timeout" yaml:"timeout" toml:"timeout"`
	Headers      map[string]string `mapstructure:"headers" json:"headers" yaml:"headers" toml:"headers"`
	Description  string            `mapstructure:"description" json:"description" yaml:"description" toml:"description"`
	Capabilities []string          `mapstructure:"capabilities" json:"capabilities" yaml:"capabilities" toml:"capabilities"`
	Labels       map[string]string `mapstructure:"labels" json:"labels" yaml:"labels" toml:"labels"`
}

type PluginHookBindingConfig struct {
	Point    string `mapstructure:"point" json:"point" yaml:"point" toml:"point"`
	Plugin   string `mapstructure:"plugin" json:"plugin" yaml:"plugin" toml:"plugin"`
	Name     string `mapstructure:"name" json:"name" yaml:"name" toml:"name"`
	Priority int    `mapstructure:"priority" json:"priority" yaml:"priority" toml:"priority"`
}

func (c *PluginConfig) ValidateName() string {
	return AppPluginName
}

func (c *PluginConfig) ValidateRequired() bool {
	return false
}

func (c *PluginConfig) Validate() error {
	if c.DefaultTimeout < 0 {
		return fmt.Errorf("plugin: default_timeout must be non-negative")
	}
	if c.MaxResponseBytes < 0 {
		return fmt.Errorf("plugin: max_response_bytes must be non-negative")
	}
	if c.Registration.Enabled && !c.Enabled {
		return fmt.Errorf("plugin: registration requires plugin manager to be enabled")
	}
	if c.Interface.HTTP.Enabled && !c.Enabled {
		return fmt.Errorf("plugin: interface http requires plugin manager to be enabled")
	}
	if !c.Enabled {
		return nil
	}
	if c.Registration.Enabled && c.Registration.Token == "" {
		return fmt.Errorf("plugin: registration token is required when registration is enabled")
	}
	if c.Registration.ExposeOnMainHTTP && !c.Registration.Enabled {
		return fmt.Errorf("plugin: registration expose_on_main_http requires registration to be enabled")
	}
	if c.Interface.HTTP.Enabled && !c.Registration.Enabled {
		return fmt.Errorf("plugin: interface http requires registration to be enabled")
	}
	if c.Registration.Enabled && !c.Registration.ExposeOnMainHTTP && !c.Interface.HTTP.Enabled {
		return fmt.Errorf("plugin: registration requires either interface http enabled or expose_on_main_http")
	}
	if c.Interface.HTTP.Port < 0 || c.Interface.HTTP.Port > 65535 {
		return fmt.Errorf("plugin: interface http port must be between 0 and 65535")
	}
	if c.Interface.HTTP.Enabled {
		if c.Interface.HTTP.Port == 0 {
			return fmt.Errorf("plugin: interface http port is required when interface http is enabled")
		}
		if c.Interface.HTTP.PublicURL == "" {
			return fmt.Errorf("plugin: interface http public_url is required when interface http is enabled")
		}
	}
	if c.Interface.HTTP.PublicURL != "" && !hasURLScheme(c.Interface.HTTP.PublicURL, "http://", "https://") {
		return fmt.Errorf("plugin: interface http public_url must start with http:// or https://")
	}
	if c.Interface.WS.PublicURL != "" && !hasURLScheme(c.Interface.WS.PublicURL, "ws://", "wss://") {
		return fmt.Errorf("plugin: interface ws public_url must start with ws:// or wss://")
	}
	pluginNames := make(map[string]bool, len(c.Plugins))
	for i, def := range c.Plugins {
		if def.Name == "" {
			return fmt.Errorf("plugin %d: name is required", i)
		}
		if pluginNames[def.Name] {
			return fmt.Errorf("plugin %q: duplicate name", def.Name)
		}
		pluginNames[def.Name] = true
		if def.Protocol != "http" {
			return fmt.Errorf("plugin %q: only http protocol can be created from config", def.Name)
		}
		if def.Endpoint == "" {
			return fmt.Errorf("plugin %q: endpoint is required", def.Name)
		}
		if def.Timeout < 0 {
			return fmt.Errorf("plugin %q: timeout must be non-negative", def.Name)
		}
	}
	for i, binding := range c.Hooks {
		if binding.Point == "" {
			return fmt.Errorf("plugin hook %d: point is required", i)
		}
		if binding.Plugin == "" {
			return fmt.Errorf("plugin hook %d: plugin is required", i)
		}
		if !pluginNames[binding.Plugin] {
			return fmt.Errorf("plugin hook %d: plugin %q is not configured", i, binding.Plugin)
		}
	}
	return nil
}

func (c *PluginConfig) DefaultTimeoutDuration() time.Duration {
	if c.DefaultTimeout <= 0 {
		return 0
	}
	return time.Duration(c.DefaultTimeout) * time.Second
}

func hasURLScheme(value string, schemes ...string) bool {
	value = strings.TrimSpace(strings.ToLower(value))
	for _, scheme := range schemes {
		if strings.HasPrefix(value, scheme) {
			return true
		}
	}
	return false
}

func overridePluginConfig(cfg *PluginConfig) {
	overrideConfigFromEnv(cfg)
}

func copyPluginDefinitions(src []PluginDefinitionConfig) []PluginDefinitionConfig {
	if len(src) == 0 {
		return nil
	}
	dst := make([]PluginDefinitionConfig, len(src))
	for i, def := range src {
		dst[i] = def
		dst[i].Headers = copyStringMap(def.Headers)
		dst[i].Capabilities = append([]string(nil), def.Capabilities...)
		dst[i].Labels = copyStringMap(def.Labels)
	}
	return dst
}

func copyStringMap(src map[string]string) map[string]string {
	if len(src) == 0 {
		return nil
	}
	dst := make(map[string]string, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
