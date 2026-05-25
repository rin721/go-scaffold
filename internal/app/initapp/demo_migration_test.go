package initapp

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/rei0721/go-scaffold/internal/modules/demo/model"
	"github.com/rei0721/go-scaffold/pkg/database"
	"github.com/rei0721/go-scaffold/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestDemoMigrationPolicyFor(t *testing.T) {
	tests := []struct {
		name        string
		trigger     DemoMigrationTrigger
		autoMigrate bool
	}{
		{name: "server start migrates demo schema", trigger: DemoMigrationTriggerServerStart, autoMigrate: true},
		{name: "initdb migrates demo schema", trigger: DemoMigrationTriggerInitDB, autoMigrate: true},
		{name: "reload skips demo schema migration", trigger: DemoMigrationTriggerReload, autoMigrate: false},
		{name: "unknown trigger skips migration", trigger: DemoMigrationTrigger("unknown"), autoMigrate: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			policy := DemoMigrationPolicyFor(tt.trigger)
			if policy.Trigger != tt.trigger {
				t.Fatalf("trigger = %q, want %q", policy.Trigger, tt.trigger)
			}
			if policy.AutoMigrate != tt.autoMigrate {
				t.Fatalf("AutoMigrate = %v, want %v", policy.AutoMigrate, tt.autoMigrate)
			}
			if policy.Reason == "" {
				t.Fatal("Reason must describe why the policy exists")
			}
		})
	}
}

func TestMigrateDemoSchemaForTrigger(t *testing.T) {
	tests := []struct {
		name     string
		trigger  DemoMigrationTrigger
		hasTable bool
	}{
		{name: "server start creates todo table", trigger: DemoMigrationTriggerServerStart, hasTable: true},
		{name: "initdb creates todo table", trigger: DemoMigrationTriggerInitDB, hasTable: true},
		{name: "reload leaves schema untouched", trigger: DemoMigrationTriggerReload, hasTable: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDatabase(t)

			policy, err := MigrateDemoSchemaForTrigger(db, testLogger{}, tt.trigger)
			if err != nil {
				t.Fatalf("MigrateDemoSchemaForTrigger() error = %v", err)
			}
			if policy.AutoMigrate != tt.hasTable {
				t.Fatalf("AutoMigrate = %v, want %v", policy.AutoMigrate, tt.hasTable)
			}
			if got := db.DB().Migrator().HasTable(&model.Todo{}); got != tt.hasTable {
				t.Fatalf("HasTable(Todo) = %v, want %v", got, tt.hasTable)
			}
		})
	}
}

func TestMigrateDemoSchemaKeepsServerStartDefault(t *testing.T) {
	db := newTestDatabase(t)

	if err := MigrateDemoSchema(db, testLogger{}); err != nil {
		t.Fatalf("MigrateDemoSchema() error = %v", err)
	}
	if !db.DB().Migrator().HasTable(&model.Todo{}) {
		t.Fatal("MigrateDemoSchema() should create the demo todo table")
	}
}

func TestMigrateDemoSchemaForTriggerAllowsMissingDatabase(t *testing.T) {
	policy, err := MigrateDemoSchemaForTrigger(nil, testLogger{}, DemoMigrationTriggerInitDB)
	if err != nil {
		t.Fatalf("MigrateDemoSchemaForTrigger(nil) error = %v", err)
	}
	if !policy.AutoMigrate {
		t.Fatal("initdb policy should still describe an auto-migration trigger")
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
