package lifecycleapp

import (
	"context"
	"fmt"

	"github.com/rei0721/go-scaffold/internal/app/initapp"
)

func Start(ctx context.Context, transport initapp.Transport) error {
	if transport.HTTPServer == nil {
		return fmt.Errorf("http server is not initialized")
	}
	if err := transport.HTTPServer.Start(ctx); err != nil {
		return fmt.Errorf("start HTTP server: %w", err)
	}
	return nil
}

func Shutdown(ctx context.Context, core initapp.Core, infra initapp.Infrastructure, transport initapp.Transport) error {
	log := core.Logger
	if log != nil {
		log.Info("shutting down application...")
	}

	var errs []error

	if transport.HTTPServer != nil {
		if err := transport.HTTPServer.Shutdown(ctx); err != nil {
			errs = append(errs, fmt.Errorf("http server shutdown: %w", err))
			if log != nil {
				log.Error("failed to shutdown HTTP server", "error", err)
			}
		} else if log != nil {
			log.Info("HTTP server stopped")
		}
	}

	if infra.Storage != nil {
		if err := infra.Storage.Close(); err != nil {
			errs = append(errs, fmt.Errorf("storage close: %w", err))
			if log != nil {
				log.Error("failed to close storage", "error", err)
			}
		} else if log != nil {
			log.Info("storage closed")
		}
	}

	if infra.Executor != nil {
		infra.Executor.Shutdown()
		if log != nil {
			log.Info("executor stopped")
		}
	}

	if infra.Cache != nil {
		if err := infra.Cache.Close(); err != nil {
			errs = append(errs, fmt.Errorf("cache close: %w", err))
			if log != nil {
				log.Error("failed to close cache", "error", err)
			}
		} else if log != nil {
			log.Info("cache closed")
		}
	}

	if infra.Database != nil {
		if err := infra.Database.Close(); err != nil {
			errs = append(errs, fmt.Errorf("database close: %w", err))
			if log != nil {
				log.Error("failed to close database", "error", err)
			}
		} else if log != nil {
			log.Info("database connection closed")
		}
	}

	if log != nil {
		_ = log.Sync()
	}

	if len(errs) > 0 {
		return fmt.Errorf("shutdown completed with %d errors", len(errs))
	}
	if log != nil {
		log.Info("application shutdown complete")
	}
	return nil
}
