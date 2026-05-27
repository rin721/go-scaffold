package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/rei0721/go-scaffold/pkg/plugin"
)

func main() {
	cfg := LoadConfigFromEnv()
	if err := run(context.Background(), cfg); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, cfg Config) error {
	blog := NewBlogPlugin(cfg)
	handler := protectWithSharedSecret(plugin.NewHTTPServer(blog), cfg.SharedSecret)
	server := &http.Server{
		Addr:    cfg.ListenAddr,
		Handler: handler,
	}

	listener, err := net.Listen("tcp", cfg.ListenAddr)
	if err != nil {
		return err
	}
	if strings.TrimSpace(cfg.PublicHTTPURL) == "" {
		cfg.PublicHTTPURL = "http://" + listener.Addr().String()
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- server.Serve(listener)
	}()

	if cfg.MainHTTPURL != "" {
		registerCtx, cancel := context.WithTimeout(ctx, cfg.RegisterTimeout)
		if err := RegisterWithHost(registerCtx, cfg, http.DefaultClient); err != nil {
			cancel()
			_ = server.Close()
			return err
		}
		cancel()
		log.Printf("registered blog plugin with host %s; ws address reserved as %s", cfg.MainHTTPURL, cfg.MainWSURL)
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	select {
	case <-ctx.Done():
	case <-signalCh:
	case err := <-errCh:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.RegisterTimeout)
	defer cancel()
	return server.Shutdown(shutdownCtx)
}

func protectWithSharedSecret(next http.Handler, secret string) http.Handler {
	if secret == "" {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Blog-Plugin-Secret") != secret {
			http.Error(w, "blog plugin secret mismatch", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func printConfig(cfg Config) string {
	return fmt.Sprintf("listen=%s public_http=%s main_http=%s main_ws=%s", cfg.ListenAddr, cfg.PublicHTTPURL, cfg.MainHTTPURL, cfg.MainWSURL)
}
