package config

import (
	"reflect"
	"testing"
)

func TestCopyConfigCoversAllFieldsAndDeepCopiesSlices(t *testing.T) {
	t.Parallel()

	src := testCompleteConfig()
	m := &manager{}

	got := m.copyConfig(src)
	if !reflect.DeepEqual(got, src) {
		t.Fatalf("copyConfig() did not preserve all fields\nwant: %#v\ngot:  %#v", src, got)
	}

	got.I18n.Supported[0] = "ja-JP"
	got.Executor.Pools[0].Name = "changed"
	got.CORS.AllowOrigins[0] = "https://changed.example.com"
	got.CORS.AllowMethods[0] = "PATCH"
	got.CORS.AllowHeaders[0] = "X-Changed"
	got.CORS.ExposeHeaders[0] = "X-Changed-Expose"

	if src.I18n.Supported[0] == got.I18n.Supported[0] {
		t.Fatal("copyConfig() shares I18n.Supported slice with source")
	}
	if src.Executor.Pools[0].Name == got.Executor.Pools[0].Name {
		t.Fatal("copyConfig() shares Executor.Pools slice with source")
	}
	if src.CORS.AllowOrigins[0] == got.CORS.AllowOrigins[0] {
		t.Fatal("copyConfig() shares CORS.AllowOrigins slice with source")
	}
	if src.CORS.AllowMethods[0] == got.CORS.AllowMethods[0] {
		t.Fatal("copyConfig() shares CORS.AllowMethods slice with source")
	}
	if src.CORS.AllowHeaders[0] == got.CORS.AllowHeaders[0] {
		t.Fatal("copyConfig() shares CORS.AllowHeaders slice with source")
	}
	if src.CORS.ExposeHeaders[0] == got.CORS.ExposeHeaders[0] {
		t.Fatal("copyConfig() shares CORS.ExposeHeaders slice with source")
	}
}

func TestUpdatePreservesUntouchedFields(t *testing.T) {
	t.Parallel()

	src := testCompleteConfig()
	m := &manager{}
	m.config.Store(src)

	if err := m.Update(func(cfg *Config) {
		cfg.Server.Port = 9090
	}); err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	got := m.Get()
	want := testCompleteConfig()
	want.Server.Port = 9090
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Update() did not preserve untouched fields\nwant: %#v\ngot:  %#v", want, got)
	}

	if src.Server.Port != 8080 {
		t.Fatalf("Update() mutated source config, got source port %d", src.Server.Port)
	}
}

func TestOverrideWithEnvUsesDocumentedDatabaseNames(t *testing.T) {
	cfg := testCompleteConfig()

	t.Setenv(EnvDBDriver, "postgres")
	t.Setenv(EnvDBHost, "db.example.com")
	t.Setenv(EnvDBPort, "15432")
	t.Setenv(EnvDBUser, "app")
	t.Setenv(EnvDBPassword, "secret")
	t.Setenv(EnvDBName, "appdb")
	t.Setenv(EnvDBMaxOpenConns, "42")
	t.Setenv(EnvDBMaxIdleConns, "21")
	t.Setenv(EnvPrefixJoin(EnvDBDriver), "mysql")
	t.Setenv(EnvPrefixJoin(EnvDBHost), "legacy.example.com")

	OverrideWithEnv(cfg)

	if cfg.Database.Driver != "postgres" {
		t.Fatalf("Database.Driver = %q, want documented DB_* variable to win", cfg.Database.Driver)
	}
	if cfg.Database.Host != "db.example.com" {
		t.Fatalf("Database.Host = %q, want db.example.com", cfg.Database.Host)
	}
	if cfg.Database.Port != 15432 {
		t.Fatalf("Database.Port = %d, want 15432", cfg.Database.Port)
	}
	if cfg.Database.User != "app" {
		t.Fatalf("Database.User = %q, want app", cfg.Database.User)
	}
	if cfg.Database.Password != "secret" {
		t.Fatalf("Database.Password = %q, want secret", cfg.Database.Password)
	}
	if cfg.Database.DBName != "appdb" {
		t.Fatalf("Database.DBName = %q, want appdb", cfg.Database.DBName)
	}
	if cfg.Database.MaxOpenConns != 42 {
		t.Fatalf("Database.MaxOpenConns = %d, want 42", cfg.Database.MaxOpenConns)
	}
	if cfg.Database.MaxIdleConns != 21 {
		t.Fatalf("Database.MaxIdleConns = %d, want 21", cfg.Database.MaxIdleConns)
	}
}

func TestOverrideWithEnvKeepsLegacyDatabasePrefixFallback(t *testing.T) {
	cfg := testCompleteConfig()

	t.Setenv(EnvPrefixJoin(EnvDBDriver), "mysql")
	t.Setenv(EnvPrefixJoin(EnvDBHost), "legacy.example.com")
	t.Setenv(EnvPrefixJoin(EnvDBPort), "3306")
	t.Setenv(EnvPrefixJoin(EnvDBUser), "legacy")
	t.Setenv(EnvPrefixJoin(EnvDBPassword), "legacy-secret")
	t.Setenv(EnvPrefixJoin(EnvDBName), "legacydb")
	t.Setenv(EnvPrefixJoin(EnvDBMaxOpenConns), "17")
	t.Setenv(EnvPrefixJoin(EnvDBMaxIdleConns), "9")

	OverrideWithEnv(cfg)

	if cfg.Database.Driver != "mysql" {
		t.Fatalf("Database.Driver = %q, want legacy fallback mysql", cfg.Database.Driver)
	}
	if cfg.Database.Host != "legacy.example.com" {
		t.Fatalf("Database.Host = %q, want legacy.example.com", cfg.Database.Host)
	}
	if cfg.Database.Port != 3306 {
		t.Fatalf("Database.Port = %d, want 3306", cfg.Database.Port)
	}
	if cfg.Database.User != "legacy" {
		t.Fatalf("Database.User = %q, want legacy", cfg.Database.User)
	}
	if cfg.Database.Password != "legacy-secret" {
		t.Fatalf("Database.Password = %q, want legacy-secret", cfg.Database.Password)
	}
	if cfg.Database.DBName != "legacydb" {
		t.Fatalf("Database.DBName = %q, want legacydb", cfg.Database.DBName)
	}
	if cfg.Database.MaxOpenConns != 17 {
		t.Fatalf("Database.MaxOpenConns = %d, want 17", cfg.Database.MaxOpenConns)
	}
	if cfg.Database.MaxIdleConns != 9 {
		t.Fatalf("Database.MaxIdleConns = %d, want 9", cfg.Database.MaxIdleConns)
	}
}

func TestOverrideWithEnvUsesDocumentedNonDatabaseNames(t *testing.T) {
	cfg := testCompleteConfig()

	t.Setenv(EnvRedisEnabled, "false")
	t.Setenv(EnvRedisHost, "redis.example.com")
	t.Setenv(EnvRedisPort, "6380")
	t.Setenv(EnvRedisPassword, "redis-secret-2")
	t.Setenv(EnvRedisDB, "2")
	t.Setenv(EnvRedisPoolSize, "30")
	t.Setenv(EnvRedisMinIdleConns, "6")
	t.Setenv(EnvRedisMaxRetries, "4")
	t.Setenv(EnvRedisDialTimeout, "7")
	t.Setenv(EnvRedisReadTimeout, "8")
	t.Setenv(EnvRedisWriteTimeout, "9")
	t.Setenv(EnvServerPort, "9090")
	t.Setenv(EnvServerMode, "release")
	t.Setenv(EnvServerReadTimeout, "11")
	t.Setenv(EnvServerWriteTimeout, "12")
	t.Setenv(EnvLogLevel, "warn")
	t.Setenv(EnvLogFormat, "console")
	t.Setenv(EnvLogOutput, "stdout")
	t.Setenv(EnvI18nDefault, "en-US")
	t.Setenv(EnvI18nSupported, "zh-CN,en-US")

	OverrideWithEnv(cfg)

	if cfg.Redis.Enabled {
		t.Fatal("Redis.Enabled = true, want false")
	}
	if cfg.Redis.Host != "redis.example.com" || cfg.Redis.Port != 6380 || cfg.Redis.Password != "redis-secret-2" {
		t.Fatalf("Redis override mismatch: %#v", cfg.Redis)
	}
	if cfg.Redis.DB != 2 || cfg.Redis.PoolSize != 30 || cfg.Redis.MinIdleConns != 6 ||
		cfg.Redis.MaxRetries != 4 || cfg.Redis.DialTimeout != 7 || cfg.Redis.ReadTimeout != 8 || cfg.Redis.WriteTimeout != 9 {
		t.Fatalf("Redis numeric override mismatch: %#v", cfg.Redis)
	}
	if cfg.Server.Port != 9090 || cfg.Server.Mode != "release" || cfg.Server.ReadTimeout != 11 || cfg.Server.WriteTimeout != 12 {
		t.Fatalf("Server override mismatch: %#v", cfg.Server)
	}
	if cfg.Logger.Level != "warn" || cfg.Logger.Format != "console" || cfg.Logger.Output != "stdout" {
		t.Fatalf("Logger override mismatch: %#v", cfg.Logger)
	}
	if !reflect.DeepEqual(cfg.I18n.Supported, []string{"zh-CN", "en-US"}) || cfg.I18n.Default != "en-US" {
		t.Fatalf("I18n override mismatch: %#v", cfg.I18n)
	}
}

func testCompleteConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Host:         "127.0.0.1",
			Port:         8080,
			Mode:         "test",
			ReadTimeout:  5,
			WriteTimeout: 10,
			IdleTimeout:  60,
		},
		Database: DatabaseConfig{
			Driver:       "sqlite",
			Host:         "localhost",
			Port:         5432,
			User:         "user",
			Password:     "secret",
			DBName:       ":memory:",
			MaxOpenConns: 10,
			MaxIdleConns: 5,
		},
		Redis: RedisConfig{
			Enabled:      true,
			Host:         "127.0.0.1",
			Port:         6379,
			Password:     "redis-secret",
			DB:           1,
			PoolSize:     20,
			MinIdleConns: 5,
			MaxRetries:   3,
			DialTimeout:  5,
			ReadTimeout:  3,
			WriteTimeout: 3,
		},
		Logger: LoggerConfig{
			Level:         "debug",
			Format:        "json",
			ConsoleFormat: "console",
			FileFormat:    "json",
			Output:        "both",
			FilePath:      "./logs/app.log",
			MaxSize:       100,
			MaxBackups:    7,
			MaxAge:        30,
		},
		I18n: I18nConfig{
			Default:     "zh-CN",
			Supported:   []string{"zh-CN", "en-US"},
			MessagesDir: "./configs/locales",
		},
		InitDB: InitDBConfig{
			ScriptDir:        "./scripts",
			LockFile:         ".initdb.lock",
			ScriptFilePrefix: "init_",
		},
		Executor: ExecutorConfig{
			Enabled: true,
			Pools: []ExecutorPoolConfig{
				{
					Name:        "default",
					Size:        10,
					Expiry:      30,
					NonBlocking: true,
				},
			},
		},
		Storage: StorageConfig{
			Enabled:         true,
			FSType:          "memory",
			BasePath:        "./data",
			EnableWatch:     true,
			WatchBufferSize: 16,
		},
		CORS: CORSConfig{
			Enabled:          true,
			AllowOrigins:     []string{"https://example.com"},
			AllowMethods:     []string{"GET", "POST"},
			AllowHeaders:     []string{"Origin", "Content-Type"},
			ExposeHeaders:    []string{"X-Request-ID"},
			AllowCredentials: true,
			MaxAge:           3600,
		},
	}
}
