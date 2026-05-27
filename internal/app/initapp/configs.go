package initapp

import (
	"reflect"
	"time"

	"github.com/rei0721/go-scaffold/internal/config"
	"github.com/rei0721/go-scaffold/pkg/cache"
	"github.com/rei0721/go-scaffold/pkg/database"
	"github.com/rei0721/go-scaffold/pkg/executor"
	"github.com/rei0721/go-scaffold/pkg/httpserver"
	"github.com/rei0721/go-scaffold/pkg/logger"
	"github.com/rei0721/go-scaffold/pkg/storage"
)

func IsRedisConfigChanged(oldCfg, newCfg *config.Config) bool {
	if oldCfg == newCfg {
		return false
	}
	return oldCfg.Redis != newCfg.Redis
}

func IsDatabaseConfigChanged(oldCfg, newCfg *config.Config) bool {
	if oldCfg == newCfg {
		return false
	}
	return oldCfg.Database != newCfg.Database
}

func IsServerConfigChanged(oldCfg, newCfg *config.Config) bool {
	if oldCfg == newCfg {
		return false
	}
	return oldCfg.Server != newCfg.Server
}

func IsLoggerConfigChanged(oldCfg, newCfg *config.Config) bool {
	if oldCfg == newCfg {
		return false
	}
	return oldCfg.Logger != newCfg.Logger
}

func IsExecutorConfigChanged(oldCfg, newCfg *config.Config) bool {
	if oldCfg == newCfg {
		return false
	}
	return oldCfg.Executor.Enabled != newCfg.Executor.Enabled ||
		!reflect.DeepEqual(oldCfg.Executor.Pools, newCfg.Executor.Pools)
}

func IsStorageConfigChanged(oldCfg, newCfg *config.Config) bool {
	if oldCfg == newCfg {
		return false
	}
	return oldCfg.Storage != newCfg.Storage
}

func IsPluginConfigChanged(oldCfg, newCfg *config.Config) bool {
	if oldCfg == newCfg {
		return false
	}
	return !reflect.DeepEqual(oldCfg.Plugin, newCfg.Plugin)
}

func IsIAMConfigChanged(oldCfg, newCfg *config.Config) bool {
	if oldCfg == newCfg {
		return false
	}
	return !reflect.DeepEqual(oldCfg.IAM, newCfg.IAM)
}

func RedisCacheConfig(cfg *config.Config) *cache.Config {
	return &cache.Config{
		Host:         cfg.Redis.Host,
		Port:         cfg.Redis.Port,
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB,
		PoolSize:     cfg.Redis.PoolSize,
		MinIdleConns: cfg.Redis.MinIdleConns,
		MaxRetries:   cfg.Redis.MaxRetries,
		DialTimeout:  time.Duration(cfg.Redis.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.Redis.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Redis.WriteTimeout) * time.Second,
	}
}

func DatabaseConfig(cfg *config.Config) *database.Config {
	return &database.Config{
		Driver:       database.Driver(cfg.Database.Driver),
		Host:         cfg.Database.Host,
		Port:         cfg.Database.Port,
		User:         cfg.Database.User,
		Password:     cfg.Database.Password,
		DBName:       cfg.Database.DBName,
		MaxOpenConns: cfg.Database.MaxOpenConns,
		MaxIdleConns: cfg.Database.MaxIdleConns,
	}
}

func LoggerConfig(cfg *config.Config) *logger.Config {
	return &logger.Config{
		Level:         cfg.Logger.Level,
		Format:        cfg.Logger.Format,
		ConsoleFormat: cfg.Logger.ConsoleFormat,
		FileFormat:    cfg.Logger.FileFormat,
		Output:        cfg.Logger.Output,
		FilePath:      cfg.Logger.FilePath,
		MaxSize:       cfg.Logger.MaxSize,
		MaxBackups:    cfg.Logger.MaxBackups,
		MaxAge:        cfg.Logger.MaxAge,
	}
}

func ExecutorConfigs(cfg *config.Config) []executor.Config {
	configs := make([]executor.Config, 0, len(cfg.Executor.Pools))
	for _, poolCfg := range cfg.Executor.Pools {
		configs = append(configs, executor.Config{
			Name:        executor.PoolName(poolCfg.Name),
			Size:        poolCfg.Size,
			Expiry:      time.Duration(poolCfg.Expiry) * time.Second,
			NonBlocking: poolCfg.NonBlocking,
		})
	}
	return configs
}

func HTTPServerConfig(cfg *config.Config) *httpserver.Config {
	return &httpserver.Config{
		Host:         cfg.Server.Host,
		Port:         cfg.Server.Port,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.Server.IdleTimeout) * time.Second,
	}
}

func PluginHTTPServerConfig(cfg *config.Config) *httpserver.Config {
	return &httpserver.Config{
		Host:         cfg.Plugin.Interface.HTTP.Host,
		Port:         cfg.Plugin.Interface.HTTP.Port,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.Server.IdleTimeout) * time.Second,
	}
}

func NormalizedStorageConfig(cfg *config.Config) config.StorageConfig {
	storageCfg := cfg.Storage
	storageCfg.DefaultConfig()
	storageCfg.OverrideConfig()
	return storageCfg
}

func StorageConfig(cfg *config.Config) *storage.Config {
	storageCfg := NormalizedStorageConfig(cfg)
	return storageCfg.ToPkgConfig()
}
