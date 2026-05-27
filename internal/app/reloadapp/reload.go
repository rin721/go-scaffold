package reloadapp

import (
	"context"
	"time"

	"github.com/rei0721/go-scaffold/internal/app/initapp"
	"github.com/rei0721/go-scaffold/internal/config"
	"github.com/rei0721/go-scaffold/pkg/executor"
	"github.com/rei0721/go-scaffold/pkg/storage"
)

func Reload(core *initapp.Core, infra *initapp.Infrastructure, transport *initapp.Transport, old, new *config.Config) {
	if initapp.IsRedisConfigChanged(old, new) {
		reloadCache(core, infra, new)
	}
	if initapp.IsDatabaseConfigChanged(old, new) {
		reloadDatabase(core, infra, new)
	}
	if initapp.IsLoggerConfigChanged(old, new) {
		reloadLogger(core, new)
	}
	if initapp.IsExecutorConfigChanged(old, new) {
		reloadExecutor(core, infra, new)
	}
	if initapp.IsServerConfigChanged(old, new) {
		reloadHTTPServer(core, transport, new)
	}
	if initapp.IsStorageConfigChanged(old, new) {
		reloadStorage(core, infra, new)
	}
	iamChanged := initapp.IsIAMConfigChanged(old, new)
	pluginChanged := initapp.IsPluginConfigChanged(old, new)
	if iamChanged || pluginChanged {
		reloadIAMAndPlugins(core, infra, new, iamChanged)
	}
}

func reloadCache(core *initapp.Core, infra *initapp.Infrastructure, cfg *config.Config) {
	if !cfg.Redis.Enabled {
		if infra.Cache != nil {
			_ = infra.Cache.Close()
			infra.Cache = nil
		}
		core.Logger.Info("redis disabled")
		return
	}

	if infra.Cache == nil {
		cacheClient, err := initapp.NewCache(cfg, core.Logger)
		if err != nil {
			core.Logger.Error("failed to initialize redis cache", "error", err)
			return
		}
		infra.Cache = cacheClient
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := infra.Cache.Reload(ctx, initapp.RedisCacheConfig(cfg)); err != nil {
		core.Logger.Error("failed to reload redis cache", "error", err)
		return
	}
	core.Logger.Info("redis cache reloaded")
}

func reloadDatabase(core *initapp.Core, infra *initapp.Infrastructure, cfg *config.Config) {
	if infra.Database == nil {
		db, err := initapp.NewDatabase(cfg)
		if err != nil {
			core.Logger.Error("failed to initialize database", "error", err)
			return
		}
		infra.Database = db
		core.Logger.Info("database initialized")
		return
	}

	if err := infra.Database.Reload(initapp.DatabaseConfig(cfg)); err != nil {
		core.Logger.Error("failed to reload database", "error", err)
		return
	}
	if _, err := initapp.MigrateDemoSchemaForTrigger(infra.Database, core.Logger, initapp.DemoMigrationTriggerReload); err != nil {
		core.Logger.Error("failed to migrate demo schema after database reload", "error", err)
		return
	}
	core.Logger.Info("database reloaded")
}

func reloadLogger(core *initapp.Core, cfg *config.Config) {
	if core.Logger == nil {
		return
	}
	if err := core.Logger.Reload(initapp.LoggerConfig(cfg)); err != nil {
		core.Logger.Error("failed to reload logger", "error", err)
		return
	}
	core.Logger.Info("logger reloaded")
}

func reloadExecutor(core *initapp.Core, infra *initapp.Infrastructure, cfg *config.Config) {
	if !cfg.Executor.Enabled {
		if infra.Executor != nil {
			infra.Executor.Shutdown()
			infra.Executor = nil
		}
		core.Logger.Info("executor disabled")
		return
	}

	executorConfigs := initapp.ExecutorConfigs(cfg)
	if infra.Executor == nil {
		mgr, err := executor.NewManager(executorConfigs)
		if err != nil {
			core.Logger.Error("failed to initialize executor", "error", err)
			return
		}
		infra.Executor = mgr
		core.Logger.Info("executor initialized", "pools", len(executorConfigs))
		return
	}

	if err := infra.Executor.Reload(executorConfigs); err != nil {
		core.Logger.Error("failed to reload executor", "error", err)
		return
	}
	core.Logger.Info("executor reloaded", "pools", len(executorConfigs))
}

func reloadHTTPServer(core *initapp.Core, transport *initapp.Transport, cfg *config.Config) {
	if transport.HTTPServer == nil {
		core.Logger.Warn("http server is nil, cannot reload configuration")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := transport.HTTPServer.Reload(ctx, initapp.HTTPServerConfig(cfg)); err != nil {
		core.Logger.Error("failed to reload HTTP server", "error", err)
		return
	}
	core.Logger.Info("HTTP server reloaded")
}

func reloadStorage(core *initapp.Core, infra *initapp.Infrastructure, cfg *config.Config) {
	storageCfg := initapp.StorageConfig(cfg)
	if !cfg.Storage.Enabled {
		if infra.Storage != nil {
			_ = infra.Storage.Close()
			infra.Storage = nil
		}
		core.Logger.Info("storage disabled")
		return
	}

	if infra.Storage == nil {
		storageService, err := storage.New(storageCfg)
		if err != nil {
			core.Logger.Error("failed to initialize storage", "error", err)
			return
		}
		infra.Storage = storageService
		core.Logger.Info("storage initialized")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := infra.Storage.Reload(ctx, storageCfg); err != nil {
		core.Logger.Error("failed to reload storage", "error", err)
		return
	}
	core.Logger.Info("storage reloaded")
}

func reloadIAMAndPlugins(core *initapp.Core, infra *initapp.Infrastructure, cfg *config.Config, iamChanged bool) {
	newIAM := infra.IAM
	if iamChanged {
		service, err := initapp.NewIAM(cfg, core.Logger)
		if err != nil {
			core.Logger.Error("failed to reload iam", "error", err)
			return
		}
		newIAM = service
	}

	newManager, err := initapp.NewPluginManager(cfg, core.Logger, newIAM)
	if err != nil {
		core.Logger.Error("failed to reload plugin manager", "error", err)
		return
	}
	oldManager := infra.Plugins
	infra.IAM = newIAM
	infra.Plugins = newManager
	if oldManager == nil {
		if iamChanged {
			core.Logger.Info("iam reloaded", "enabled", cfg.IAM.Enabled)
		}
		core.Logger.Info("plugin manager reloaded", "enabled", cfg.Plugin.Enabled)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := oldManager.Close(ctx); err != nil {
		core.Logger.Error("failed to close previous plugin manager", "error", err)
		return
	}
	if iamChanged {
		core.Logger.Info("iam reloaded", "enabled", cfg.IAM.Enabled)
	}
	core.Logger.Info("plugin manager reloaded", "enabled", cfg.Plugin.Enabled)
}
