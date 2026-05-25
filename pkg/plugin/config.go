package plugin

import (
	"fmt"
	"net/url"
	"time"
)

// Config contains plugin manager configuration.
type Config struct {
	Plugins          []Definition  `json:"plugins" yaml:"plugins" mapstructure:"plugins"`
	DefaultTimeout   time.Duration `json:"defaultTimeout" yaml:"defaultTimeout" mapstructure:"defaultTimeout"`
	MaxResponseBytes int64         `json:"maxResponseBytes" yaml:"maxResponseBytes" mapstructure:"maxResponseBytes"`
}

// Definition describes how to create one plugin instance.
type Definition struct {
	Name         string            `json:"name" yaml:"name" mapstructure:"name"`
	Version      string            `json:"version,omitempty" yaml:"version,omitempty" mapstructure:"version"`
	Protocol     Protocol          `json:"protocol" yaml:"protocol" mapstructure:"protocol"`
	Endpoint     string            `json:"endpoint,omitempty" yaml:"endpoint,omitempty" mapstructure:"endpoint"`
	Timeout      time.Duration     `json:"timeout,omitempty" yaml:"timeout,omitempty" mapstructure:"timeout"`
	Headers      map[string]string `json:"headers,omitempty" yaml:"headers,omitempty" mapstructure:"headers"`
	Description  string            `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description"`
	Capabilities []string          `json:"capabilities,omitempty" yaml:"capabilities,omitempty" mapstructure:"capabilities"`
	Labels       map[string]string `json:"labels,omitempty" yaml:"labels,omitempty" mapstructure:"labels"`
}

// DefaultConfig returns a plugin manager configuration with safe defaults.
func DefaultConfig() *Config {
	return &Config{
		DefaultTimeout:   DefaultTimeout,
		MaxResponseBytes: DefaultMaxResponseBytes,
	}
}

// Validate validates the full manager configuration.
func (c *Config) Validate() error {
	if c == nil {
		return ErrInvalidConfig
	}
	if c.DefaultTimeout < 0 {
		return fmt.Errorf("%w: default timeout must be non-negative", ErrInvalidConfig)
	}
	if c.MaxResponseBytes < 0 {
		return fmt.Errorf("%w: max response bytes must be non-negative", ErrInvalidConfig)
	}
	for i := range c.Plugins {
		if err := c.Plugins[i].Validate(); err != nil {
			return fmt.Errorf("plugin %d: %w", i, err)
		}
	}
	return nil
}

// Validate verifies a single plugin definition.
func (d Definition) Validate() error {
	if d.Name == "" {
		return fmt.Errorf("%w: name is required", ErrInvalidDefinition)
	}
	switch d.Protocol {
	case ProtocolLocal:
		return nil
	case ProtocolHTTP:
		if d.Endpoint == "" {
			return fmt.Errorf("%w: http endpoint is required", ErrInvalidDefinition)
		}
		u, err := url.Parse(d.Endpoint)
		if err != nil || u.Scheme == "" || u.Host == "" {
			return fmt.Errorf("%w: invalid http endpoint", ErrInvalidDefinition)
		}
		if u.Scheme != "http" && u.Scheme != "https" {
			return fmt.Errorf("%w: http endpoint must use http or https", ErrInvalidDefinition)
		}
		return nil
	case ProtocolRPC, ProtocolWS:
		return nil
	default:
		return fmt.Errorf("%w: %s", ErrUnsupportedProtocol, d.Protocol)
	}
}

// Metadata returns public metadata derived from the definition.
func (d Definition) Metadata() Metadata {
	return d.metadata()
}

func (d Definition) metadata() Metadata {
	return Metadata{
		Name:         d.Name,
		Version:      d.Version,
		Protocol:     d.Protocol,
		Description:  d.Description,
		Capabilities: copyStrings(d.Capabilities),
		Labels:       copyStringMap(d.Labels),
	}
}

func (d Definition) timeout(defaultTimeout time.Duration) time.Duration {
	if d.Timeout > 0 {
		return d.Timeout
	}
	if defaultTimeout > 0 {
		return defaultTimeout
	}
	return DefaultTimeout
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

func copyStrings(src []string) []string {
	if len(src) == 0 {
		return nil
	}
	dst := make([]string, len(src))
	copy(dst, src)
	return dst
}
