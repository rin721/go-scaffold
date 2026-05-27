package config

import (
	"os"
	"path/filepath"
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
	got.Plugin.Plugins[0].Headers["X-Test"] = "changed"
	got.Plugin.Plugins[0].Capabilities[0] = "changed"
	got.Plugin.Hooks[0].Point = "changed"
	got.IAM.Tokens[0].Principal.Roles[0] = "changed"
	got.IAM.Tokens[0].Principal.Attributes["team"] = "changed"
	got.IAM.Policies[0].Action = "changed"
	*got.IAM.DefaultDeny = false
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
	if src.Plugin.Plugins[0].Headers["X-Test"] == got.Plugin.Plugins[0].Headers["X-Test"] {
		t.Fatal("copyConfig() shares Plugin.Plugins headers map with source")
	}
	if src.Plugin.Plugins[0].Capabilities[0] == got.Plugin.Plugins[0].Capabilities[0] {
		t.Fatal("copyConfig() shares Plugin.Plugins capabilities slice with source")
	}
	if src.Plugin.Hooks[0].Point == got.Plugin.Hooks[0].Point {
		t.Fatal("copyConfig() shares Plugin.Hooks slice with source")
	}
	if src.IAM.Tokens[0].Principal.Roles[0] == got.IAM.Tokens[0].Principal.Roles[0] {
		t.Fatal("copyConfig() shares IAM token roles slice with source")
	}
	if src.IAM.Tokens[0].Principal.Attributes["team"] == got.IAM.Tokens[0].Principal.Attributes["team"] {
		t.Fatal("copyConfig() shares IAM token attributes map with source")
	}
	if src.IAM.Policies[0].Action == got.IAM.Policies[0].Action {
		t.Fatal("copyConfig() shares IAM.Policies slice with source")
	}
	if *src.IAM.DefaultDeny == *got.IAM.DefaultDeny {
		t.Fatal("copyConfig() shares IAM.DefaultDeny pointer with source")
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

func taggedEnvName(t *testing.T, target any, fieldName string) string {
	t.Helper()

	typ := reflect.TypeOf(target)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok {
		t.Fatalf("%s has no field %s", typ.Name(), fieldName)
	}
	name := field.Tag.Get("envname")
	if name == "" || name == "-" {
		t.Fatalf("%s.%s has no envname tag", typ.Name(), fieldName)
	}
	return name
}

func setTaggedEnv(t *testing.T, target any, fieldName, value string) {
	t.Helper()

	t.Setenv(EnvPrefixJoin(taggedEnvName(t, target, fieldName)), value)
}

func setUnprefixedTaggedEnv(t *testing.T, target any, fieldName, value string) {
	t.Helper()

	t.Setenv(taggedEnvName(t, target, fieldName), value)
}

func TestEnvNamesUseDynamicAppPrefix(t *testing.T) {
	if EnvPrefix() != "RIN_APP" {
		t.Fatalf("EnvPrefix() = %q, want RIN_APP", EnvPrefix())
	}
	dbHostEnvName := taggedEnvName(t, DatabaseConfig{}, "Host")
	if EnvPrefixJoin(dbHostEnvName) != "RIN_APP_DB_HOST" {
		t.Fatalf("EnvPrefixJoin(%q) = %q, want RIN_APP_DB_HOST", dbHostEnvName, EnvPrefixJoin(dbHostEnvName))
	}
	if EnvConfigPathName() != "RIN_CONFIG_PATH" {
		t.Fatalf("EnvConfigPathName() = %q, want RIN_CONFIG_PATH", EnvConfigPathName())
	}
}

func TestOverrideWithEnvUsesDynamicPrefixFromAppPrefix(t *testing.T) {
	cfg := testCompleteConfig()

	setUnprefixedTaggedEnv(t, DatabaseConfig{}, "Driver", "mysql")
	setUnprefixedTaggedEnv(t, DatabaseConfig{}, "Host", "fallback.example.com")
	setUnprefixedTaggedEnv(t, DatabaseConfig{}, "Port", "3306")
	setUnprefixedTaggedEnv(t, DatabaseConfig{}, "User", "fallback")
	setUnprefixedTaggedEnv(t, DatabaseConfig{}, "Password", "fallback-secret")
	setUnprefixedTaggedEnv(t, DatabaseConfig{}, "DBName", "fallbackdb")
	setUnprefixedTaggedEnv(t, DatabaseConfig{}, "MaxOpenConns", "17")
	setUnprefixedTaggedEnv(t, DatabaseConfig{}, "MaxIdleConns", "9")
	setTaggedEnv(t, DatabaseConfig{}, "Driver", "postgres")
	setTaggedEnv(t, DatabaseConfig{}, "Host", "db.example.com")
	setTaggedEnv(t, DatabaseConfig{}, "Port", "15432")
	setTaggedEnv(t, DatabaseConfig{}, "User", "app")
	setTaggedEnv(t, DatabaseConfig{}, "Password", "secret")
	setTaggedEnv(t, DatabaseConfig{}, "DBName", "appdb")
	setTaggedEnv(t, DatabaseConfig{}, "MaxOpenConns", "42")
	setTaggedEnv(t, DatabaseConfig{}, "MaxIdleConns", "21")

	OverrideWithEnv(cfg)

	if cfg.Database.Driver != "postgres" {
		t.Fatalf("Database.Driver = %q, want dynamic prefixed variable to win", cfg.Database.Driver)
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

func TestOverrideWithEnvKeepsUnprefixedFallback(t *testing.T) {
	cfg := testCompleteConfig()

	setUnprefixedTaggedEnv(t, DatabaseConfig{}, "Driver", "mysql")
	setUnprefixedTaggedEnv(t, DatabaseConfig{}, "Host", "fallback.example.com")
	setUnprefixedTaggedEnv(t, DatabaseConfig{}, "Port", "3306")
	setUnprefixedTaggedEnv(t, DatabaseConfig{}, "User", "fallback")
	setUnprefixedTaggedEnv(t, DatabaseConfig{}, "Password", "fallback-secret")
	setUnprefixedTaggedEnv(t, DatabaseConfig{}, "DBName", "fallbackdb")
	setUnprefixedTaggedEnv(t, DatabaseConfig{}, "MaxOpenConns", "17")
	setUnprefixedTaggedEnv(t, DatabaseConfig{}, "MaxIdleConns", "9")

	OverrideWithEnv(cfg)

	if cfg.Database.Driver != "mysql" {
		t.Fatalf("Database.Driver = %q, want unprefixed fallback mysql", cfg.Database.Driver)
	}
	if cfg.Database.Host != "fallback.example.com" {
		t.Fatalf("Database.Host = %q, want fallback.example.com", cfg.Database.Host)
	}
	if cfg.Database.Port != 3306 {
		t.Fatalf("Database.Port = %d, want 3306", cfg.Database.Port)
	}
	if cfg.Database.User != "fallback" {
		t.Fatalf("Database.User = %q, want fallback", cfg.Database.User)
	}
	if cfg.Database.Password != "fallback-secret" {
		t.Fatalf("Database.Password = %q, want fallback-secret", cfg.Database.Password)
	}
	if cfg.Database.DBName != "fallbackdb" {
		t.Fatalf("Database.DBName = %q, want fallbackdb", cfg.Database.DBName)
	}
	if cfg.Database.MaxOpenConns != 17 {
		t.Fatalf("Database.MaxOpenConns = %d, want 17", cfg.Database.MaxOpenConns)
	}
	if cfg.Database.MaxIdleConns != 9 {
		t.Fatalf("Database.MaxIdleConns = %d, want 9", cfg.Database.MaxIdleConns)
	}
}

func TestOverrideWithEnvUsesEnvnameTagsForNonDatabaseConfigs(t *testing.T) {
	cfg := testCompleteConfig()

	setTaggedEnv(t, RedisConfig{}, "Enabled", "false")
	setTaggedEnv(t, RedisConfig{}, "Host", "redis.example.com")
	setTaggedEnv(t, RedisConfig{}, "Port", "6380")
	setTaggedEnv(t, RedisConfig{}, "Password", "redis-secret-2")
	setTaggedEnv(t, RedisConfig{}, "DB", "2")
	setTaggedEnv(t, RedisConfig{}, "PoolSize", "30")
	setTaggedEnv(t, RedisConfig{}, "MinIdleConns", "6")
	setTaggedEnv(t, RedisConfig{}, "MaxRetries", "4")
	setTaggedEnv(t, RedisConfig{}, "DialTimeout", "7")
	setTaggedEnv(t, RedisConfig{}, "ReadTimeout", "8")
	setTaggedEnv(t, RedisConfig{}, "WriteTimeout", "9")
	setTaggedEnv(t, ServerConfig{}, "Host", "0.0.0.0")
	setTaggedEnv(t, ServerConfig{}, "Port", "9090")
	setTaggedEnv(t, ServerConfig{}, "Mode", "release")
	setTaggedEnv(t, ServerConfig{}, "ReadTimeout", "11")
	setTaggedEnv(t, ServerConfig{}, "WriteTimeout", "12")
	setTaggedEnv(t, ServerConfig{}, "IdleTimeout", "13")
	setTaggedEnv(t, LoggerConfig{}, "Level", "warn")
	setTaggedEnv(t, LoggerConfig{}, "Format", "console")
	setTaggedEnv(t, LoggerConfig{}, "ConsoleFormat", "console")
	setTaggedEnv(t, LoggerConfig{}, "FileFormat", "json")
	setTaggedEnv(t, LoggerConfig{}, "Output", "stdout")
	setTaggedEnv(t, LoggerConfig{}, "FilePath", "./logs/env.log")
	setTaggedEnv(t, LoggerConfig{}, "MaxSize", "64")
	setTaggedEnv(t, LoggerConfig{}, "MaxBackups", "3")
	setTaggedEnv(t, LoggerConfig{}, "MaxAge", "14")
	setTaggedEnv(t, I18nConfig{}, "Default", "en-US")
	setTaggedEnv(t, I18nConfig{}, "Supported", "zh-CN,en-US")
	setTaggedEnv(t, I18nConfig{}, "MessagesDir", "./locales")
	setTaggedEnv(t, InitDBConfig{}, "ScriptDir", "./init")
	setTaggedEnv(t, InitDBConfig{}, "LockFile", ".env-init.lock")
	setTaggedEnv(t, InitDBConfig{}, "ScriptFilePrefix", "seed_")
	setTaggedEnv(t, ExecutorConfig{}, "Enabled", "false")
	setTaggedEnv(t, StorageConfig{}, "Enabled", "false")
	setTaggedEnv(t, StorageConfig{}, "FSType", "basepath")
	setTaggedEnv(t, StorageConfig{}, "BasePath", "./env-data")
	setTaggedEnv(t, StorageConfig{}, "EnableWatch", "false")
	setTaggedEnv(t, StorageConfig{}, "WatchBufferSize", "32")
	setTaggedEnv(t, PluginConfig{}, "Enabled", "true")
	setTaggedEnv(t, PluginConfig{}, "DefaultTimeout", "15")
	setTaggedEnv(t, PluginConfig{}, "MaxResponseBytes", "2048")
	setTaggedEnv(t, IAMConfig{}, "Enabled", "true")
	setTaggedEnv(t, IAMConfig{}, "Mode", "memory")
	setTaggedEnv(t, IAMConfig{}, "DefaultDeny", "false")
	setTaggedEnv(t, CORSConfig{}, "Enabled", "false")
	setTaggedEnv(t, CORSConfig{}, "AllowOrigins", "https://app.example.com, https://admin.example.com")
	setTaggedEnv(t, CORSConfig{}, "AllowMethods", "GET,POST")
	setTaggedEnv(t, CORSConfig{}, "AllowHeaders", "Origin,Authorization")
	setTaggedEnv(t, CORSConfig{}, "ExposeHeaders", "X-Request-ID,X-Total-Count")
	setTaggedEnv(t, CORSConfig{}, "AllowCredentials", "false")
	setTaggedEnv(t, CORSConfig{}, "MaxAge", "7200")

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
	if cfg.Server.Host != "0.0.0.0" || cfg.Server.Port != 9090 || cfg.Server.Mode != "release" ||
		cfg.Server.ReadTimeout != 11 || cfg.Server.WriteTimeout != 12 || cfg.Server.IdleTimeout != 13 {
		t.Fatalf("Server override mismatch: %#v", cfg.Server)
	}
	if cfg.Logger.Level != "warn" || cfg.Logger.Format != "console" || cfg.Logger.ConsoleFormat != "console" ||
		cfg.Logger.FileFormat != "json" || cfg.Logger.Output != "stdout" || cfg.Logger.FilePath != "./logs/env.log" ||
		cfg.Logger.MaxSize != 64 || cfg.Logger.MaxBackups != 3 || cfg.Logger.MaxAge != 14 {
		t.Fatalf("Logger override mismatch: %#v", cfg.Logger)
	}
	if !reflect.DeepEqual(cfg.I18n.Supported, []string{"zh-CN", "en-US"}) ||
		cfg.I18n.Default != "en-US" || cfg.I18n.MessagesDir != "./locales" {
		t.Fatalf("I18n override mismatch: %#v", cfg.I18n)
	}
	if cfg.InitDB.ScriptDir != "./init" || cfg.InitDB.LockFile != ".env-init.lock" || cfg.InitDB.ScriptFilePrefix != "seed_" {
		t.Fatalf("InitDB override mismatch: %#v", cfg.InitDB)
	}
	if cfg.Executor.Enabled {
		t.Fatalf("Executor.Enabled = true, want false")
	}
	if cfg.Storage.Enabled || cfg.Storage.FSType != "basepath" || cfg.Storage.BasePath != "./env-data" ||
		cfg.Storage.EnableWatch || cfg.Storage.WatchBufferSize != 32 {
		t.Fatalf("Storage override mismatch: %#v", cfg.Storage)
	}
	if !cfg.Plugin.Enabled || cfg.Plugin.DefaultTimeout != 15 || cfg.Plugin.MaxResponseBytes != 2048 {
		t.Fatalf("Plugin override mismatch: %#v", cfg.Plugin)
	}
	if !cfg.IAM.Enabled || cfg.IAM.Mode != "memory" || cfg.IAM.DefaultDenyEnabled() {
		t.Fatalf("IAM override mismatch: %#v", cfg.IAM)
	}
	if cfg.CORS.Enabled || !reflect.DeepEqual(cfg.CORS.AllowOrigins, []string{"https://app.example.com", "https://admin.example.com"}) ||
		!reflect.DeepEqual(cfg.CORS.AllowMethods, []string{"GET", "POST"}) ||
		!reflect.DeepEqual(cfg.CORS.AllowHeaders, []string{"Origin", "Authorization"}) ||
		!reflect.DeepEqual(cfg.CORS.ExposeHeaders, []string{"X-Request-ID", "X-Total-Count"}) ||
		cfg.CORS.AllowCredentials || cfg.CORS.MaxAge != 7200 {
		t.Fatalf("CORS override mismatch: %#v", cfg.CORS)
	}
}

func TestDirectOverrideConfigUsesDynamicPrefix(t *testing.T) {
	storageCfg := StorageConfig{}
	setTaggedEnv(t, StorageConfig{}, "FSType", "memory")
	setTaggedEnv(t, StorageConfig{}, "WatchBufferSize", "64")
	storageCfg.OverrideConfig()
	if storageCfg.FSType != "memory" || storageCfg.WatchBufferSize != 64 {
		t.Fatalf("Storage OverrideConfig mismatch: %#v", storageCfg)
	}

	corsCfg := CORSConfig{}
	setTaggedEnv(t, CORSConfig{}, "AllowOrigins", "https://app.example.com,https://admin.example.com")
	setTaggedEnv(t, CORSConfig{}, "MaxAge", "1800")
	corsCfg.OverrideConfig()
	if !reflect.DeepEqual(corsCfg.AllowOrigins, []string{"https://app.example.com", "https://admin.example.com"}) ||
		corsCfg.MaxAge != 1800 {
		t.Fatalf("CORS OverrideConfig mismatch: %#v", corsCfg)
	}
}

func TestManagerLoadAutoLoadsDotEnvWithDynamicPrefix(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.yaml")
	dotEnvPath := filepath.Join(tempDir, EnvFilePath)
	serverPortName := taggedEnvName(t, ServerConfig{}, "Port")
	serverPortEnv := EnvPrefixJoin(serverPortName)

	unsetEnvForTest(t, serverPortEnv, serverPortName)
	writeTestConfig(t, configPath)
	if err := os.WriteFile(dotEnvPath, []byte(serverPortEnv+"=19090\n"), 0600); err != nil {
		t.Fatalf("write .env: %v", err)
	}

	oldWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("get wd: %v", err)
	}
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("chdir temp dir: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(oldWD); err != nil {
			t.Errorf("restore wd: %v", err)
		}
	})

	m := NewManager()
	if err := m.Load(configPath); err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	got := m.Get()
	if got.Server.Port != 19090 {
		t.Fatalf("Server.Port = %d, want 19090 from .env", got.Server.Port)
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
		Plugin: PluginConfig{
			Enabled:          true,
			DefaultTimeout:   10,
			MaxResponseBytes: 1024,
			Plugins: []PluginDefinitionConfig{
				{
					Name:         "remote",
					Protocol:     "http",
					Endpoint:     "http://127.0.0.1:18090/plugin/v1/invoke",
					Timeout:      5,
					Headers:      map[string]string{"X-Test": "original"},
					Capabilities: []string{"hooks"},
					Labels:       map[string]string{"env": "test"},
				},
			},
			Hooks: []PluginHookBindingConfig{
				{Point: "plugin.after_invoke", Plugin: "remote", Name: "audit", Priority: 1},
			},
		},
		IAM: IAMConfig{
			Enabled:     true,
			Mode:        "memory",
			DefaultDeny: boolPtr(true),
			Tokens: []IAMTokenConfig{
				{
					Token: "admin-token",
					Principal: IAMPrincipalConfig{
						ID:         "admin",
						Roles:      []string{"admin"},
						Attributes: map[string]string{"team": "platform"},
					},
				},
			},
			Policies: []IAMPolicyConfig{
				{Subject: "admin", Action: "read", Resource: "*", Effect: "allow"},
			},
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

func boolPtr(value bool) *bool {
	return &value
}

func writeTestConfig(t *testing.T, path string) {
	t.Helper()

	const content = `
server:
  host: 127.0.0.1
  port: 8080
  mode: test
  read_timeout: 5
  write_timeout: 10
  idle_timeout: 60
database:
  driver: sqlite
  host: localhost
  port: 5432
  user: user
  password: secret
  dbname: ":memory:"
  max_open_conns: 10
  max_idle_conns: 5
redis:
  enabled: false
logger:
  level: debug
  format: json
  output: stdout
i18n:
  default: zh-CN
  supported:
    - zh-CN
  messages_dir: ./configs/locales
initdb:
  script_dir: ./scripts
  lock_file: .initdb.lock
  script_file_prefix: init_
executor:
  enabled: false
storage:
  enabled: false
plugin:
  enabled: false
iam:
  enabled: false
cors:
  enabled: false
`
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		t.Fatalf("write config: %v", err)
	}
}

func unsetEnvForTest(t *testing.T, keys ...string) {
	t.Helper()

	for _, key := range keys {
		key := key
		oldValue, existed := os.LookupEnv(key)
		if err := os.Unsetenv(key); err != nil {
			t.Fatalf("unset %s: %v", key, err)
		}
		t.Cleanup(func() {
			if existed {
				if err := os.Setenv(key, oldValue); err != nil {
					t.Errorf("restore %s: %v", key, err)
				}
				return
			}
			if err := os.Unsetenv(key); err != nil {
				t.Errorf("restore unset %s: %v", key, err)
			}
		})
	}
}
