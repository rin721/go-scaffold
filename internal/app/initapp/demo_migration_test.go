package initapp

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/rei0721/go-scaffold/internal/config"
	"github.com/rei0721/go-scaffold/internal/modules/demo/model"
	"github.com/rei0721/go-scaffold/pkg/database"
	"github.com/rei0721/go-scaffold/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestDemoSchemaPolicyFor(t *testing.T) {
	tests := []struct {
		name    string
		trigger DemoSchemaTrigger
		apply   bool
	}{
		{name: "server start applies demo schema", trigger: DemoSchemaTriggerServerStart, apply: true},
		{name: "reload skips demo schema apply", trigger: DemoSchemaTriggerReload, apply: false},
		{name: "unknown trigger skips schema apply", trigger: DemoSchemaTrigger("unknown"), apply: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			policy := DemoSchemaPolicyFor(tt.trigger)
			if policy.Trigger != tt.trigger {
				t.Fatalf("trigger = %q, want %q", policy.Trigger, tt.trigger)
			}
			if policy.Apply != tt.apply {
				t.Fatalf("Apply = %v, want %v", policy.Apply, tt.apply)
			}
			if policy.Reason == "" {
				t.Fatal("Reason must describe why the policy exists")
			}
		})
	}
}

func TestApplyDemoSchemaForTrigger(t *testing.T) {
	tests := []struct {
		name     string
		trigger  DemoSchemaTrigger
		hasTable bool
	}{
		{name: "server start creates todo table", trigger: DemoSchemaTriggerServerStart, hasTable: true},
		{name: "reload leaves schema untouched", trigger: DemoSchemaTriggerReload, hasTable: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDatabase(t)

			policy, err := ApplyDemoSchemaForTrigger(db, string(database.DriverSQLite), testLogger{}, tt.trigger)
			if err != nil {
				t.Fatalf("ApplyDemoSchemaForTrigger() error = %v", err)
			}
			if policy.Apply != tt.hasTable {
				t.Fatalf("Apply = %v, want %v", policy.Apply, tt.hasTable)
			}
			if got := db.DB().Migrator().HasTable(&model.Todo{}); got != tt.hasTable {
				t.Fatalf("HasTable(Todo) = %v, want %v", got, tt.hasTable)
			}
		})
	}
}

func TestApplyDemoSchemaKeepsServerStartDefault(t *testing.T) {
	db := newTestDatabase(t)

	if err := ApplyDemoSchema(db, string(database.DriverSQLite), testLogger{}); err != nil {
		t.Fatalf("ApplyDemoSchema() error = %v", err)
	}
	if !db.DB().Migrator().HasTable(&model.Todo{}) {
		t.Fatal("ApplyDemoSchema() should create the demo todo table")
	}
}

func TestNewModulesRespectsDemoConfig(t *testing.T) {
	tests := []struct {
		name        string
		enabled     bool
		applySchema bool
		wantHandler bool
		wantTable   bool
	}{
		{name: "demo disabled", enabled: false, applySchema: true, wantHandler: false, wantTable: false},
		{name: "demo enabled without schema apply", enabled: true, applySchema: false, wantHandler: true, wantTable: false},
		{name: "demo enabled with schema apply", enabled: true, applySchema: true, wantHandler: true, wantTable: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDatabase(t)
			enabled := tt.enabled
			applySchema := tt.applySchema
			core := Core{
				Config: &config.Config{
					Database: config.DatabaseConfig{Driver: string(database.DriverSQLite)},
					Demo: config.DemoConfig{
						Enabled:            &enabled,
						ApplySchemaOnStart: &applySchema,
					},
					Auth: config.AuthConfig{
						TokenSecret: "0123456789abcdef0123456789abcdef",
						TokenTTL:    60,
					},
				},
				Logger: testLogger{},
			}

			modules, err := NewModules(core, Infrastructure{Database: db})
			if err != nil {
				t.Fatalf("NewModules() error = %v", err)
			}
			if got := modules.Demo.TodoHandler != nil; got != tt.wantHandler {
				t.Fatalf("TodoHandler present = %v, want %v", got, tt.wantHandler)
			}
			if got := db.DB().Migrator().HasTable(&model.Todo{}); got != tt.wantTable {
				t.Fatalf("HasTable(Todo) = %v, want %v", got, tt.wantTable)
			}
		})
	}
}

func TestApplyDemoSchemaForTriggerAllowsMissingDatabase(t *testing.T) {
	policy, err := ApplyDemoSchemaForTrigger(nil, string(database.DriverSQLite), testLogger{}, DemoSchemaTriggerServerStart)
	if err != nil {
		t.Fatalf("ApplyDemoSchemaForTrigger(nil) error = %v", err)
	}
	if !policy.Apply {
		t.Fatal("server-start policy should still describe a schema apply trigger")
	}
}

type testDatabase struct {
	db *gorm.DB
}

func newTestDatabase(t *testing.T) *testDatabase {
	t.Helper()

	name := strings.NewReplacer("/", "_", " ", "_").Replace(t.Name())
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("file:%s?mode=memory&cache=shared", name)), &gorm.Config{})
	if err != nil {
		t.Fatalf("open sqlite memory database: %v", err)
	}

	testDB := &testDatabase{db: db}
	t.Cleanup(func() {
		if err := testDB.Close(); err != nil {
			t.Fatalf("close sqlite memory database: %v", err)
		}
	})
	return testDB
}

func (d *testDatabase) DB() *gorm.DB {
	return d.db
}

func (d *testDatabase) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (d *testDatabase) Ping(ctx context.Context) error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.PingContext(ctx)
}

func (d *testDatabase) WithTx(ctx context.Context, fn database.TxFunc) error {
	return d.WithTxOptions(ctx, nil, fn)
}

func (d *testDatabase) WithTxOptions(ctx context.Context, _ *database.TxOptions, fn database.TxFunc) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}

func (d *testDatabase) Reload(_ *database.Config) error {
	return nil
}

type testLogger struct{}

func (testLogger) Debug(string, ...interface{}) {}
func (testLogger) Info(string, ...interface{})  {}
func (testLogger) Warn(string, ...interface{})  {}
func (testLogger) Error(string, ...interface{}) {}
func (testLogger) Fatal(string, ...interface{}) {}

func (l testLogger) With(...interface{}) logger.Logger {
	return l
}

func (testLogger) Sync() error {
	return nil
}

func (testLogger) Reload(*logger.Config) error {
	return nil
}
