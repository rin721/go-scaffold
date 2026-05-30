package reloadapp

import (
	"context"
	"testing"

	"github.com/rei0721/go-scaffold/internal/app/initapp"
	"github.com/rei0721/go-scaffold/internal/config"
	"github.com/rei0721/go-scaffold/pkg/cache"
	"github.com/rei0721/go-scaffold/pkg/database"
	"github.com/rei0721/go-scaffold/pkg/executor"
	"github.com/rei0721/go-scaffold/pkg/httpserver"
	"github.com/rei0721/go-scaffold/pkg/logger"
	storagepkg "github.com/rei0721/go-scaffold/pkg/storage"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestReloadSkipsUnchangedComponents(t *testing.T) {
	clearReloadEnv(t)

	oldCfg := baseReloadConfig()
	newCfg := cloneReloadConfig(oldCfg)
	core, infra, transport, fakes := newReloadFixture(t)

	Reload(&core, &infra, &transport, oldCfg, newCfg)

	assertReloadCounts(t, fakes, reloadCounts{})
	if fakes.cache.closes != 0 {
		t.Fatalf("expected cache not to close, got %d", fakes.cache.closes)
	}
	if fakes.executor.shutdowns != 0 {
		t.Fatalf("expected executor not to shutdown, got %d", fakes.executor.shutdowns)
	}
	if fakes.storage.closes != 0 {
		t.Fatalf("expected storage not to close, got %d", fakes.storage.closes)
	}
	if fakes.database.dbCalls != 0 {
		t.Fatalf("expected unchanged database config not to inspect gorm DB, got %d calls", fakes.database.dbCalls)
	}
}

func TestReloadReloadsOnlyChangedComponent(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*config.Config)
		want   reloadCounts
	}{
		{
			name: "redis",
			mutate: func(cfg *config.Config) {
				cfg.Redis.Host = "redis.changed.local"
			},
			want: reloadCounts{cache: 1},
		},
		{
			name: "database",
			mutate: func(cfg *config.Config) {
				cfg.Database.DBName = "file:changed-database?mode=memory&cache=shared"
			},
			want: reloadCounts{database: 1},
		},
		{
			name: "logger",
			mutate: func(cfg *config.Config) {
				cfg.Logger.Level = "debug"
			},
			want: reloadCounts{logger: 1},
		},
		{
			name: "executor",
			mutate: func(cfg *config.Config) {
				cfg.Executor.Pools[0].Size = 3
			},
			want: reloadCounts{executor: 1},
		},
		{
			name: "http-server",
			mutate: func(cfg *config.Config) {
				cfg.Server.Port = 18082
			},
			want: reloadCounts{httpServer: 1},
		},
		{
			name: "storage",
			mutate: func(cfg *config.Config) {
				cfg.Storage.BasePath = "./changed-storage"
			},
			want: reloadCounts{storage: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearReloadEnv(t)

			oldCfg := baseReloadConfig()
			newCfg := cloneReloadConfig(oldCfg)
			tt.mutate(newCfg)
			core, infra, transport, fakes := newReloadFixture(t)

			Reload(&core, &infra, &transport, oldCfg, newCfg)

			assertReloadCounts(t, fakes, tt.want)
			if tt.name == "database" && fakes.database.dbCalls != 0 {
				t.Fatalf("expected database reload not to trigger demo schema migration DB() call, got %d calls", fakes.database.dbCalls)
			}
		})
	}
}

func TestReloadDisablesOptionalComponents(t *testing.T) {
	clearReloadEnv(t)

	oldCfg := baseReloadConfig()
	newCfg := cloneReloadConfig(oldCfg)
	newCfg.Redis.Enabled = false
	newCfg.Executor.Enabled = false
	newCfg.Storage.Enabled = false
	core, infra, transport, fakes := newReloadFixture(t)

	Reload(&core, &infra, &transport, oldCfg, newCfg)

	assertReloadCounts(t, fakes, reloadCounts{})
	if fakes.cache.closes != 1 {
		t.Fatalf("expected cache close once, got %d", fakes.cache.closes)
	}
	if infra.Cache != nil {
		t.Fatal("expected disabled redis cache to be nil")
	}
	if fakes.executor.shutdowns != 1 {
		t.Fatalf("expected executor shutdown once, got %d", fakes.executor.shutdowns)
	}
	if infra.Executor != nil {
		t.Fatal("expected disabled executor to be nil")
	}
	if fakes.storage.closes != 1 {
		t.Fatalf("expected storage close once, got %d", fakes.storage.closes)
	}
	if infra.Storage != nil {
		t.Fatal("expected disabled storage to be nil")
	}
}

func baseReloadConfig() *config.Config {
	return &config.Config{
		Server: config.ServerConfig{
			Host:         "127.0.0.1",
			Port:         18081,
			Mode:         "test",
			ReadTimeout:  1,
			WriteTimeout: 1,
			IdleTimeout:  1,
		},
		Database: config.DatabaseConfig{
			Driver:       "sqlite",
			DBName:       "file:reload-old?mode=memory&cache=shared",
			MaxOpenConns: 1,
			MaxIdleConns: 1,
		},
		Redis: config.RedisConfig{
			Enabled:      true,
			Host:         "127.0.0.1",
			Port:         6379,
			DB:           0,
			PoolSize:     1,
			MinIdleConns: 0,
			MaxRetries:   0,
			DialTimeout:  1,
			ReadTimeout:  1,
			WriteTimeout: 1,
		},
		Logger: config.LoggerConfig{
			Level:  "info",
			Format: "console",
			Output: "stdout",
		},
		I18n: config.I18nConfig{
			Default:   "zh-CN",
			Supported: []string{"zh-CN", "en-US"},
		},
		Executor: config.ExecutorConfig{
			Enabled: true,
			Pools: []config.ExecutorPoolConfig{
				{Name: "background", Size: 1, Expiry: 1, NonBlocking: true},
			},
		},
		Storage: config.StorageConfig{
			Enabled:         true,
			FSType:          "memory",
			BasePath:        "./storage",
			EnableWatch:     false,
			WatchBufferSize: 1,
		},
		CORS: config.CORSConfig{
			Enabled:      true,
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET", "POST", "OPTIONS"},
			AllowHeaders: []string{"Origin", "Content-Type", "X-Request-ID"},
			ExposeHeaders: []string{
				"X-Request-ID",
			},
			MaxAge: 60,
		},
	}
}

func cloneReloadConfig(src *config.Config) *config.Config {
	dst := *src
	dst.I18n.Supported = append([]string(nil), src.I18n.Supported...)
	dst.Executor.Pools = append([]config.ExecutorPoolConfig(nil), src.Executor.Pools...)
	dst.CORS.AllowOrigins = append([]string(nil), src.CORS.AllowOrigins...)
	dst.CORS.AllowMethods = append([]string(nil), src.CORS.AllowMethods...)
	dst.CORS.AllowHeaders = append([]string(nil), src.CORS.AllowHeaders...)
	dst.CORS.ExposeHeaders = append([]string(nil), src.CORS.ExposeHeaders...)
	return &dst
}

type reloadFixture struct {
	cache      *fakeCache
	database   *fakeDatabase
	logger     *fakeLogger
	executor   *fakeExecutor
	httpServer *fakeHTTPServer
	storage    *fakeStorage
}

func newReloadFixture(t *testing.T) (initapp.Core, initapp.Infrastructure, initapp.Transport, reloadFixture) {
	t.Helper()

	fakes := reloadFixture{
		cache:      &fakeCache{},
		database:   newFakeDatabase(t),
		logger:     &fakeLogger{},
		executor:   &fakeExecutor{},
		httpServer: &fakeHTTPServer{},
		storage:    &fakeStorage{},
	}

	return initapp.Core{Logger: fakes.logger},
		initapp.Infrastructure{
			Cache:    fakes.cache,
			Database: fakes.database,
			Executor: fakes.executor,
			Storage:  fakes.storage,
		},
		initapp.Transport{HTTPServer: fakes.httpServer},
		fakes
}

type reloadCounts struct {
	cache      int
	database   int
	logger     int
	executor   int
	httpServer int
	storage    int
}

func assertReloadCounts(t *testing.T, fakes reloadFixture, want reloadCounts) {
	t.Helper()

	if fakes.cache.reloads != want.cache {
		t.Fatalf("expected cache reloads %d, got %d", want.cache, fakes.cache.reloads)
	}
	if fakes.database.reloads != want.database {
		t.Fatalf("expected database reloads %d, got %d", want.database, fakes.database.reloads)
	}
	if fakes.logger.reloads != want.logger {
		t.Fatalf("expected logger reloads %d, got %d", want.logger, fakes.logger.reloads)
	}
	if fakes.executor.reloads != want.executor {
		t.Fatalf("expected executor reloads %d, got %d", want.executor, fakes.executor.reloads)
	}
	if fakes.httpServer.reloads != want.httpServer {
		t.Fatalf("expected http server reloads %d, got %d", want.httpServer, fakes.httpServer.reloads)
	}
	if fakes.storage.reloads != want.storage {
		t.Fatalf("expected storage reloads %d, got %d", want.storage, fakes.storage.reloads)
	}
}

type fakeCache struct {
	cache.Cache
	reloads    int
	closes     int
	lastConfig *cache.Config
}

func (c *fakeCache) Reload(_ context.Context, cfg *cache.Config) error {
	c.reloads++
	c.lastConfig = cfg
	return nil
}

func (c *fakeCache) Close() error {
	c.closes++
	return nil
}

type fakeDatabase struct {
	database.Database
	reloads    int
	dbCalls    int
	lastConfig *database.Config
	gormDB     *gorm.DB
}

func newFakeDatabase(t *testing.T) *fakeDatabase {
	t.Helper()

	gormDB, err := gorm.Open(sqlite.Open("file:reload-fake?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("create fake gorm database: %v", err)
	}
	return &fakeDatabase{gormDB: gormDB}
}

func (d *fakeDatabase) DB() *gorm.DB {
	d.dbCalls++
	return d.gormDB
}

func (d *fakeDatabase) Reload(cfg *database.Config) error {
	d.reloads++
	d.lastConfig = cfg
	return nil
}

func (d *fakeDatabase) WithTx(ctx context.Context, fn database.TxFunc) error {
	return fn(ctx, d.gormDB)
}

func (d *fakeDatabase) WithTxOptions(ctx context.Context, _ *database.TxOptions, fn database.TxFunc) error {
	return fn(ctx, d.gormDB)
}

func (d *fakeDatabase) Close() error {
	sqlDB, err := d.gormDB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (d *fakeDatabase) Ping(context.Context) error {
	return nil
}

type fakeLogger struct {
	reloads    int
	lastConfig *logger.Config
}

func (l *fakeLogger) Debug(string, ...interface{}) {}

func (l *fakeLogger) Info(string, ...interface{}) {}

func (l *fakeLogger) Warn(string, ...interface{}) {}

func (l *fakeLogger) Error(string, ...interface{}) {}

func (l *fakeLogger) Fatal(string, ...interface{}) {}

func (l *fakeLogger) With(...interface{}) logger.Logger {
	return l
}

func (l *fakeLogger) Sync() error {
	return nil
}

func (l *fakeLogger) Reload(cfg *logger.Config) error {
	l.reloads++
	l.lastConfig = cfg
	return nil
}

type fakeExecutor struct {
	executor.Manager
	reloads     int
	shutdowns   int
	lastConfigs []executor.Config
}

func (e *fakeExecutor) Reload(configs []executor.Config) error {
	e.reloads++
	e.lastConfigs = configs
	return nil
}

func (e *fakeExecutor) Shutdown() {
	e.shutdowns++
}

type fakeHTTPServer struct {
	httpserver.HTTPServer
	reloads    int
	lastConfig *httpserver.Config
}

func (s *fakeHTTPServer) Reload(_ context.Context, cfg *httpserver.Config) error {
	s.reloads++
	s.lastConfig = cfg
	return nil
}

type fakeStorage struct {
	storagepkg.Storage
	reloads    int
	closes     int
	lastConfig *storagepkg.Config
}

func (s *fakeStorage) Reload(_ context.Context, cfg *storagepkg.Config) error {
	s.reloads++
	s.lastConfig = cfg
	return nil
}

func (s *fakeStorage) Close() error {
	s.closes++
	return nil
}

func clearReloadEnv(t *testing.T) {
	t.Helper()

	keys := []string{
		"STORAGE_ENABLED",
		"STORAGE_FS_TYPE",
		"STORAGE_BASE_PATH",
		"STORAGE_ENABLE_WATCH",
		"STORAGE_WATCH_BUFFER_SIZE",
	}
	for _, key := range keys {
		t.Setenv(key, "")
		t.Setenv(config.EnvPrefixJoin(key), "")
	}
}
