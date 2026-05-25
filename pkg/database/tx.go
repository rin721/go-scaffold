package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

const defaultTxTimeout = 30 * time.Second

var (
	ErrNilTxFunc                 = errors.New("transaction function is nil")
	ErrInvalidTxOptions          = errors.New("invalid transaction options")
	ErrNestedTransactionDisabled = errors.New("nested transaction is disabled")
)

type txContextKey struct{}

func DefaultTxOptions() *TxOptions {
	return &TxOptions{
		Isolation: sql.LevelDefault,
		Timeout:   defaultTxTimeout,
	}
}

func (opts *TxOptions) normalize() error {
	if opts == nil {
		return ErrInvalidTxOptions
	}
	if opts.Timeout < 0 {
		return ErrInvalidTxOptions
	}
	if opts.Timeout == 0 {
		opts.Timeout = defaultTxTimeout
	}
	return nil
}

func txFromContext(ctx context.Context) *gorm.DB {
	if ctx == nil {
		return nil
	}
	tx, _ := ctx.Value(txContextKey{}).(*gorm.DB)
	return tx
}

func contextWithTx(ctx context.Context, tx *gorm.DB) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, txContextKey{}, tx)
}

func (d *database) WithTx(ctx context.Context, fn TxFunc) error {
	return d.WithTxOptions(ctx, nil, fn)
}

func (d *database) WithTxOptions(ctx context.Context, opts *TxOptions, fn TxFunc) error {
	if fn == nil {
		return ErrNilTxFunc
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if opts == nil {
		opts = DefaultTxOptions()
	} else if err := opts.normalize(); err != nil {
		return err
	}

	txCtx := ctx
	var cancel context.CancelFunc
	if opts.Timeout > 0 {
		txCtx, cancel = context.WithTimeout(ctx, opts.Timeout)
		defer cancel()
	}

	if parentTx := txFromContext(txCtx); parentTx != nil {
		if opts.DisableNestedTransaction {
			return ErrNestedTransactionDisabled
		}
		return runGormTransaction(txCtx, parentTx.WithContext(txCtx), opts, fn)
	}

	return runGormTransaction(txCtx, d.DB().WithContext(txCtx), opts, fn)
}

func runGormTransaction(ctx context.Context, db *gorm.DB, opts *TxOptions, fn TxFunc) error {
	txOpts := &sql.TxOptions{
		Isolation: opts.Isolation,
		ReadOnly:  opts.ReadOnly,
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		txCtx := contextWithTx(ctx, tx)
		if err := fn(txCtx, tx); err != nil {
			return err
		}
		if err := txCtx.Err(); err != nil {
			return err
		}
		return nil
	}, txOpts)
	if err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}
	return nil
}
