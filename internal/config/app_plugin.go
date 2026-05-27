package config

import (
	"fmt"
	"time"
)

type PluginConfig struct {
	Enabled          bool                      `mapstructure:"enabled" envname:"PLUGIN_ENABLED" json:"enabled" yaml:"enabled" toml:"enabled"`
	DefaultTimeout   int                       `mapstructure:"default_timeout" envname:"PLUGIN_DEFAULT_TIMEOUT" json:"default_timeout" yaml:"default_timeout" toml:"default_timeout"`
	MaxResponseBytes int64                     `mapstructure:"max_response_bytes" envname:"PLUGIN_MAX_RESPONSE_BYTES" json:"max_response_bytes" yaml:"max_response_bytes" toml:"max_response_bytes"`
	Registration     PluginRegistrationConfig  `mapstructure:"registration" json:"registration" yaml:"registration" toml:"registration"`
	Plugins          []PluginDefinitionConfig  `mapstructure:"plugins" json:"plugins" yaml:"plugins" toml:"plugins"`
	Hooks            []PluginHookBindingConfig `mapstructure:"hooks" json:"hooks" yaml:"hooks" toml:"hooks"`
}

type PluginRegistrationConfig struct {
	Enabled bool   `mapstructure:"enabled" envname:"PLUGIN_REGISTRATION_ENABLED" json:"enabled" yaml:"enabled" toml:"enabled"`
	Token   string `mapstructure:"token" envname:"PLUGIN_REGISTRATION_TOKEN" json:"token" yaml:"token" toml:"token"`
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
	if !c.Enabled {
		return nil
	}
	if c.Registration.Enabled && c.Registration.Token == "" {
		return fmt.Errorf("plugin: registration token is required when registration is enabled")
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
