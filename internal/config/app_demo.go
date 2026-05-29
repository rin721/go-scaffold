package config

type DemoConfig struct {
	Enabled            *bool `mapstructure:"enabled" envname:"DEMO_ENABLED" json:"enabled" yaml:"enabled" toml:"enabled"`
	ApplySchemaOnStart *bool `mapstructure:"apply_schema_on_start" envname:"DEMO_APPLY_SCHEMA_ON_START" json:"apply_schema_on_start" yaml:"apply_schema_on_start" toml:"apply_schema_on_start"`
}

func (c *DemoConfig) ValidateName() string {
	return AppDemoName
}

func (c *DemoConfig) ValidateRequired() bool {
	return false
}

func (c *DemoConfig) Validate() error {
	return nil
}

func (c DemoConfig) EnabledValue() bool {
	if c.Enabled == nil {
		return true
	}
	return *c.Enabled
}

func (c DemoConfig) ApplySchemaOnStartValue() bool {
	if c.ApplySchemaOnStart == nil {
		return c.EnabledValue()
	}
	return *c.ApplySchemaOnStart
}

func copyDemoConfig(src DemoConfig) DemoConfig {
	dst := src
	if src.Enabled != nil {
		value := *src.Enabled
		dst.Enabled = &value
	}
	if src.ApplySchemaOnStart != nil {
		value := *src.ApplySchemaOnStart
		dst.ApplySchemaOnStart = &value
	}
	return dst
}
