package initapp

import (
	"fmt"

	"github.com/rei0721/go-scaffold/internal/config"
	"github.com/rei0721/go-scaffold/pkg/cache"
	"github.com/rei0721/go-scaffold/pkg/database"
	"github.com/rei0721/go-scaffold/pkg/executor"
	"github.com/rei0721/go-scaffold/pkg/logger"
	"github.com/rei0721/go-scaffold/pkg/storage"
)

func NewInfrastructure(core Core) (Infrastructure, error) {
	db, err := NewDatabase(core.Config)
	if err != nil {
		return Infrastructure{}, err
	}

	cacheClient, err := NewCache(core.Config, core.Logger)
	if err != nil {
		return Infrastructure{}, err
	}

	executorManager, err := NewExecutor(core.Config, core.Logger)
	if err != nil {
		return Infrastructure{}, err
	}

	storageService, err := NewStorage(core.Config, core.Logger)
	if err != nil {
		return Infrastructure{}, err
	}

	return Infrastructure{
		Database: db,
		Cache:    cacheClient,
		Executor: executorManager,
		Storage:  storageService,
	}, nil
}

func NewDatabase(cfg *config.Config) (database.Database, error) {
	db, err := database.New(DatabaseConfig(cfg))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

func NewCache(cfg *config.Config, log logger.Logger) (cache.Cache, error) {
	if !cfg.Redis.Enabled {
		log.Info("redis cache disabled")
		return nil, nil
	}

	cacheClient, err := cache.NewRedis(RedisCacheConfig(cfg), log)
	if err != nil {
		log.Warn("failed to connect to redis, running without cache", "error", err)
		return nil, nil
	}

	log.Info("redis cache connected successfully")
	return cacheClient, nil
}

func NewExecutor(cfg *config.Config, log logger.Logger) (executor.Manager, error) {
	if !cfg.Executor.Enabled {
		log.Info("executor disabled")
		return nil, nil
	}

	executorConfigs := ExecutorConfigs(cfg)
	mgr, err := executor.NewManager(executorConfigs)
	if err != nil {
		return nil, fmt.Errorf("failed to create executor manager: %w", err)
	}

	log.Info("executor initialized", "pools", len(executorConfigs))
	return mgr, nil
}

func NewStorage(cfg *config.Config, log logger.Logger) (storage.Storage, error) {
	storageCfg := NormalizedStorageConfig(cfg)
	if err := storageCfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid storage config: %w", err)
	}

	if !storageCfg.Enabled {
		log.Info("storage disabled")
		return nil, nil
	}

	storageService, err := storage.New(storageCfg.ToPkgConfig())
	if err != nil {
		return nil, fmt.Errorf("failed to create storage: %w", err)
	}

	log.Info(
		"storage initialized",
		"fs_type", storageCfg.FSType,
		"base_path", storageCfg.BasePath,
		"enable_watch", storageCfg.EnableWatch,
	)
	return storageService, nil
}
