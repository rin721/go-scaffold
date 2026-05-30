package config

import "fmt"

type Configurable interface {
	Validate() error
}

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Logger   LoggerConfig   `mapstructure:"logger"`
	I18n     I18nConfig     `mapstructure:"i18n"`
	Executor ExecutorConfig `mapstructure:"executor"`
	Storage  StorageConfig  `mapstructure:"storage"`
	Demo     DemoConfig     `mapstructure:"demo"`
	IAM      IAMConfig      `mapstructure:"iam"`
	Auth     AuthConfig     `mapstructure:"auth"`
	RBAC     RBACConfig     `mapstructure:"rbac"`
	CORS     CORSConfig     `mapstructure:"cors"`
}

type Validator interface {
	Validate() error
	ValidateName() string
	ValidateRequired() bool
}

func (c *Config) Validate() error {
	validators := []Validator{
		&c.Server,
		&c.Database,
		&c.Redis,
		&c.Logger,
		&c.I18n,
		&c.Executor,
		&c.Storage,
		&c.Demo,
		&c.IAM,
		&c.Auth,
		&c.RBAC,
		&c.CORS,
	}
	for _, validator := range validators {
		if validator == nil {
			continue
		}
		if err := validator.Validate(); err != nil {
			return fmt.Errorf("%s config: %w", validator.ValidateName(), err)
		}
	}
	return nil
}

func (c *Config) ValidateOld() error {
	if err := c.Server.Validate(); err != nil {
		return fmt.Errorf("server config: %w", err)
	}
	if err := c.Database.Validate(); err != nil {
		return fmt.Errorf("database config: %w", err)
	}
	if err := c.Redis.Validate(); err != nil {
		return fmt.Errorf("redis config: %w", err)
	}
	if err := c.Logger.Validate(); err != nil {
		return fmt.Errorf("logger config: %w", err)
	}
	if err := c.I18n.Validate(); err != nil {
		return fmt.Errorf("i18n config: %w", err)
	}
	if err := c.Executor.Validate(); err != nil {
		return fmt.Errorf("executor config: %w", err)
	}
	if err := c.Auth.Validate(); err != nil {
		return fmt.Errorf("auth config: %w", err)
	}
	if err := c.RBAC.Validate(); err != nil {
		return fmt.Errorf("rbac config: %w", err)
	}
	return nil
}
