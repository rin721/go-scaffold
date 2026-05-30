package app

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
	"time"

	"github.com/rei0721/go-scaffold/internal/config"
	"github.com/rei0721/go-scaffold/internal/modules/demo/model"
	usermodel "github.com/rei0721/go-scaffold/internal/modules/user/model"
)

func TestNewServerModeBuildsMinimalApplication(t *testing.T) {
	clearAppIntegrationEnv(t)

	configPath := writeAppIntegrationConfig(t, filepath.Join(t.TempDir(), "server-mode.db"))

	application, err := New(Options{ConfigPath: configPath, Mode: ModeServer})
	if err != nil {
		t.Fatalf("new server app: %v", err)
	}
	defer shutdownApp(t, application)

	if application.Core.Config == nil {
		t.Fatal("expected core config")
	}
	if application.Core.ConfigManager == nil {
		t.Fatal("expected core config manager")
	}
	if application.Core.Logger == nil {
		t.Fatal("expected core logger")
	}
	if application.Core.I18n == nil {
		t.Fatal("expected core i18n")
	}
	if application.Core.I18nUtils == nil {
		t.Fatal("expected core i18n utils")
	}
	if application.Core.IDGenerator == nil {
		t.Fatal("expected core id generator")
	}

	if application.Infra.Database == nil {
		t.Fatal("expected database infrastructure")
	}
	if application.Infra.Cache != nil {
		t.Fatal("expected redis cache to be disabled")
	}
	if application.Infra.Executor != nil {
		t.Fatal("expected executor to be disabled")
	}
	if application.Infra.Storage != nil {
		t.Fatal("expected storage to be disabled")
	}
	if application.Infra.IAM != nil {
		t.Fatal("expected iam to be disabled")
	}

	if application.Modules.Demo.TodoRepository == nil {
		t.Fatal("expected demo repository")
	}
	if application.Modules.Demo.TodoService == nil {
		t.Fatal("expected demo service")
	}
	if application.Modules.Demo.TodoHandler == nil {
		t.Fatal("expected demo handler")
	}
	if application.Modules.User.Repository == nil {
		t.Fatal("expected user repository")
	}
	if application.Modules.User.Service == nil {
		t.Fatal("expected user service")
	}
	if application.Modules.User.Handler == nil {
		t.Fatal("expected user handler")
	}
	if application.Modules.User.Tokens == nil {
		t.Fatal("expected user token service")
	}

	if application.Transport.Router == nil {
		t.Fatal("expected HTTP router")
	}
	if application.Transport.HTTPServer == nil {
		t.Fatal("expected HTTP server wrapper")
	}

	if !application.Infra.Database.DB().Migrator().HasTable(&model.Todo{}) {
		t.Fatal("expected demo todo schema to be created in server mode")
	}
	for _, table := range []interface{}{
		&usermodel.User{},
		&usermodel.Role{},
		&usermodel.Permission{},
		&usermodel.UserRole{},
		&usermodel.RolePermission{},
	} {
		if !application.Infra.Database.DB().Migrator().HasTable(table) {
			t.Fatalf("expected user schema table for %#v to be created in server mode", table)
		}
	}

	if err := application.Core.ConfigManager.Update(func(cfg *config.Config) {
		cfg.Server.Port = 18082
		cfg.Server.ReadTimeout = 2
	}); err != nil {
		t.Fatalf("update config through manager: %v", err)
	}
	if got := application.Core.Config.Server.Port; got != 18082 {
		t.Fatalf("expected app hook to update core config port to 18082, got %d", got)
	}
	if got := application.Core.ConfigManager.Get().Server.Port; got != 18082 {
		t.Fatalf("expected manager config port to be 18082, got %d", got)
	}
}

func writeAppIntegrationConfig(t *testing.T, dbPath string) string {
	t.Helper()

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("locate test file")
	}
	repoRoot := filepath.Clean(filepath.Join(filepath.Dir(file), "..", ".."))
	messagesDir := filepath.Join(repoRoot, "configs", "locales")
	configPath := filepath.Join(t.TempDir(), "config.yaml")
	storagePath := filepath.Join(t.TempDir(), "storage")

	content := fmt.Sprintf(`server:
  host: "127.0.0.1"
  port: 18081
  mode: "test"
  read_timeout: 1
  write_timeout: 1
  idle_timeout: 1
database:
  driver: "sqlite"
  host: ""
  port: 0
  user: ""
  password: ""
  dbname: %s
  max_open_conns: 1
  max_idle_conns: 1
redis:
  enabled: false
  host: "127.0.0.1"
  port: 6379
  password: ""
  db: 0
  pool_size: 1
  min_idle_conns: 0
  max_retries: 0
  dial_timeout: 1
  read_timeout: 1
  write_timeout: 1
logger:
  level: "error"
  format: "console"
  output: "stdout"
  file_path: ""
  max_size: 1
  max_backups: 1
  max_age: 1
i18n:
  default: "zh-CN"
  supported:
    - "zh-CN"
    - "en-US"
  messages_dir: %s
executor:
  enabled: false
  pools: []
storage:
  enabled: false
  fs_type: "memory"
  base_path: %s
  enable_watch: false
  watch_buffer_size: 1
auth:
  token_secret: "0123456789abcdef0123456789abcdef"
  token_ttl: 60
cors:
  enabled: true
  allow_origins:
    - "*"
  allow_methods:
    - "GET"
    - "POST"
    - "PUT"
    - "DELETE"
    - "PATCH"
    - "OPTIONS"
  allow_headers:
    - "Origin"
    - "Content-Type"
    - "Authorization"
    - "X-Request-ID"
  expose_headers:
    - "X-Request-ID"
  allow_credentials: false
  max_age: 60
`, yamlString(dbPath), yamlString(messagesDir), yamlString(storagePath))

	if err := os.WriteFile(configPath, []byte(content), 0600); err != nil {
		t.Fatalf("write test config: %v", err)
	}
	return configPath
}

func shutdownApp(t *testing.T, application *App) {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := application.Shutdown(ctx); err != nil {
		t.Fatalf("shutdown app: %v", err)
	}
}

func yamlString(value string) string {
	return strconv.Quote(filepath.ToSlash(value))
}

func clearAppIntegrationEnv(t *testing.T) {
	t.Helper()

	keys := []string{
		"DB_DRIVER",
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"DB_MAX_OPEN_CONNS",
		"DB_MAX_IDLE_CONNS",
		"REI_APP_DB_DRIVER",
		"REI_APP_DB_HOST",
		"REI_APP_DB_PORT",
		"REI_APP_DB_USER",
		"REI_APP_DB_PASSWORD",
		"REI_APP_DB_NAME",
		"REI_APP_DB_MAX_OPEN_CONNS",
		"REI_APP_DB_MAX_IDLE_CONNS",
		"REDIS_ENABLED",
		"REDIS_HOST",
		"REDIS_PORT",
		"REDIS_PASSWORD",
		"REDIS_DB",
		"REDIS_POOL_SIZE",
		"REDIS_MIN_IDLE_CONNS",
		"REDIS_MAX_RETRIES",
		"REDIS_DIAL_TIMEOUT",
		"REDIS_READ_TIMEOUT",
		"REDIS_WRITE_TIMEOUT",
		"SERVER_PORT",
		"SERVER_MODE",
		"SERVER_READ_TIMEOUT",
		"SERVER_WRITE_TIMEOUT",
		"LOG_LEVEL",
		"LOG_FORMAT",
		"LOG_OUTPUT",
		"I18N_DEFAULT",
		"I18N_SUPPORTED",
		"STORAGE_ENABLED",
		"STORAGE_FS_TYPE",
		"STORAGE_BASE_PATH",
		"STORAGE_ENABLE_WATCH",
		"STORAGE_WATCH_BUFFER_SIZE",
		"IAM_ENABLED",
		"IAM_MODE",
		"IAM_DEFAULT_DENY",
		"AUTH_TOKEN_SECRET",
		"AUTH_TOKEN_TTL",
		"RBAC_ENABLED",
		"RBAC_APPLY_ON_START",
		"CORS_ENABLED",
		"CORS_ALLOW_ORIGINS",
		"CORS_ALLOW_METHODS",
		"CORS_ALLOW_HEADERS",
		"CORS_EXPOSE_HEADERS",
		"CORS_ALLOW_CREDENTIALS",
		"CORS_MAX_AGE",
	}
	for _, key := range keys {
		t.Setenv(key, "")
		t.Setenv(config.EnvPrefixJoin(key), "")
	}
}
