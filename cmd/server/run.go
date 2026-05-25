package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rei0721/go-scaffold/internal/app"
	"github.com/rei0721/go-scaffold/types/constants"
)

func runApp(configPath string) {
	application, err := app.New(app.Options{
		ConfigPath: configPath,
	})
	if err != nil {
		os.Stderr.WriteString("failed to initialize application: " + err.Error() + "\n")
		os.Exit(1)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	errChan := make(chan error, 1)
	go func() {
		if err := application.Run(); err != nil {
			errChan <- err
		}
	}()

	select {
	case sig := <-quit:
		application.Core.Logger.Info("received shutdown signal", "signal", sig.String())
	case err := <-errChan:
		application.Core.Logger.Error("server error", "error", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), constants.AppShutdownTimeout)
	defer cancel()

	if err := application.Shutdown(ctx); err != nil {
		application.Core.Logger.Error("shutdown error", "error", err)
		os.Exit(1)
	}

	application.Core.Logger.Info("application exited gracefully")
}
