package database

import (
	"context"
	"errors"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type txTestUser struct {
	ID    int64  `gorm:"primaryKey"`
	Name  string `gorm:"size:100"`
	Email string `gorm:"size:100;uniqueIndex"`
}

func setupTxTestDB(t *testing.T) Database {
	t.Helper()
	gdb, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	if err := gdb.AutoMigrate(&txTestUser{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	sqlDB, err := gdb.DB()
	if err != nil {
		t.Fatalf("failed to get sql db: %v", err)
	}
	t.Cleanup(func() { _ = sqlDB.Close() })
	return &database{db: gdb, sqlDB: sqlDB}
}

func TestWithTxCommit(t *testing.T) {
	db := setupTxTestDB(t)
	err := db.WithTx(context.Background(), func(ctx context.Context, tx *gorm.DB) error {
		return tx.WithContext(ctx).Create(&txTestUser{Name: "Alice", Email: "alice@example.com"}).Error
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var count int64
	db.DB().Model(&txTestUser{}).Count(&count)
	if count != 1 {
		t.Fatalf("expected 1 user, got %d", count)
	}
}

func TestWithTxRollback(t *testing.T) {
	db := setupTxTestDB(t)
	expected := errors.New("stop")
	err := db.WithTx(context.Background(), func(ctx context.Context, tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&txTestUser{Name: "Bob", Email: "bob@example.com"}).Error; err != nil {
			return err
		}
		return expected
	})
	if !errors.Is(err, expected) {
		t.Fatalf("expected wrapped stop error, got %v", err)
	}

	var count int64
	db.DB().Model(&txTestUser{}).Count(&count)
	if count != 0 {
		t.Fatalf("expected rollback, got %d users", count)
	}
}

func TestWithTxNested(t *testing.T) {
	db := setupTxTestDB(t)
	err := db.WithTx(context.Background(), func(ctx context.Context, tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&txTestUser{Name: "Outer", Email: "outer@example.com"}).Error; err != nil {
			return err
		}
		return db.WithTx(ctx, func(ctx context.Context, tx *gorm.DB) error {
			return tx.WithContext(ctx).Create(&txTestUser{Name: "Inner", Email: "inner@example.com"}).Error
		})
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var count int64
	db.DB().Model(&txTestUser{}).Count(&count)
	if count != 2 {
		t.Fatalf("expected 2 users, got %d", count)
	}
}

func TestWithTxNestedRollback(t *testing.T) {
	db := setupTxTestDB(t)
	innerErr := errors.New("inner failed")
	err := db.WithTx(context.Background(), func(ctx context.Context, tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&txTestUser{Name: "Outer", Email: "outer@example.com"}).Error; err != nil {
			return err
		}
		err := db.WithTx(ctx, func(ctx context.Context, tx *gorm.DB) error {
			if err := tx.WithContext(ctx).Create(&txTestUser{Name: "Inner", Email: "inner@example.com"}).Error; err != nil {
				return err
			}
			return innerErr
		})
		if !errors.Is(err, innerErr) {
			t.Fatalf("expected inner error, got %v", err)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var count int64
	db.DB().Model(&txTestUser{}).Count(&count)
	if count != 1 {
		t.Fatalf("expected only outer user, got %d", count)
	}
}

func TestWithTxDisableNested(t *testing.T) {
	db := setupTxTestDB(t)
	err := db.WithTx(context.Background(), func(ctx context.Context, tx *gorm.DB) error {
		return db.WithTxOptions(ctx, &TxOptions{DisableNestedTransaction: true, Timeout: time.Second}, func(ctx context.Context, tx *gorm.DB) error {
			return nil
		})
	})
	if !errors.Is(err, ErrNestedTransactionDisabled) {
		t.Fatalf("expected nested transaction disabled, got %v", err)
	}
}
