package httpserver

import (
	"context"
	"errors"
	"net"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/rei0721/go-scaffold/pkg/logger"
)

func TestNewAppliesDefaultsAndRejectsInvalidInput(t *testing.T) {
	if _, err := New(nil, &Config{}, noopLogger{}); err == nil {
		t.Fatal("New(nil handler) error = nil, want error")
	}

	srv, err := New(http.NewServeMux(), &Config{}, noopLogger{})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	got := srv.(*httpServer)
	if got.config.Host != DefaultHost {
		t.Fatalf("Host = %q, want %q", got.config.Host, DefaultHost)
	}
	if got.config.Port != DefaultPort {
		t.Fatalf("Port = %d, want %d", got.config.Port, DefaultPort)
	}
	if got.config.ReadTimeout != DefaultReadTimeout {
		t.Fatalf("ReadTimeout = %s, want %s", got.config.ReadTimeout, DefaultReadTimeout)
	}

	_, err = New(http.NewServeMux(), &Config{Port: -1}, noopLogger{})
	var serverErr *ServerError
	if !errors.As(err, &serverErr) {
		t.Fatalf("New(invalid config) error = %T, want *ServerError", err)
	}
	var configErr *ConfigError
	if !errors.As(err, &configErr) {
		t.Fatalf("New(invalid config) unwrap = %v, want *ConfigError", err)
	}
}

func TestReloadUpdatesConfigWhenStopped(t *testing.T) {
	srv, err := New(http.NewServeMux(), &Config{Host: "127.0.0.1", Port: 18080}, noopLogger{})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	next := &Config{
		Host:         "0.0.0.0",
		Port:         18081,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
	if err := srv.Reload(context.Background(), next); err != nil {
		t.Fatalf("Reload(stopped) error = %v", err)
	}
	got := srv.(*httpServer)
	if got.config.Host != "0.0.0.0" || got.config.Port != 18081 {
		t.Fatalf("config = %s:%d, want 0.0.0.0:18081", got.config.Host, got.config.Port)
	}
	if got.config.IdleTimeout != DefaultIdleTimeout {
		t.Fatalf("IdleTimeout = %s, want %s", got.config.IdleTimeout, DefaultIdleTimeout)
	}

	if err := srv.Reload(context.Background(), nil); err == nil {
		t.Fatal("Reload(nil) error = nil, want error")
	}
	if err := srv.Shutdown(context.Background()); err != nil {
		t.Fatalf("Shutdown(stopped) error = %v", err)
	}
}

func TestStartRejectsAlreadyRunningServer(t *testing.T) {
	srv, err := New(http.NewServeMux(), &Config{}, noopLogger{})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	got := srv.(*httpServer)
	got.state.Store(int32(stateRunning))
	if err := srv.Start(context.Background()); err == nil {
		t.Fatal("Start(already running) error = nil, want error")
	}
}

func TestStartReturnsBindError(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen occupied port: %v", err)
	}
	defer listener.Close()

	_, portValue, err := net.SplitHostPort(listener.Addr().String())
	if err != nil {
		t.Fatalf("split listener addr: %v", err)
	}
	port, err := strconv.Atoi(portValue)
	if err != nil {
		t.Fatalf("parse listener port: %v", err)
	}

	srv, err := New(http.NewServeMux(), &Config{Host: "127.0.0.1", Port: port}, noopLogger{})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if err := srv.Start(context.Background()); err == nil {
		t.Fatal("Start(occupied port) error = nil, want bind error")
	}
}

type noopLogger struct{}

func (noopLogger) Debug(string, ...interface{}) {}
func (noopLogger) Info(string, ...interface{})  {}
func (noopLogger) Warn(string, ...interface{})  {}
func (noopLogger) Error(string, ...interface{}) {}
func (noopLogger) Fatal(string, ...interface{}) {}
func (l noopLogger) With(...interface{}) logger.Logger {
	return l
}
func (noopLogger) Sync() error {
	return nil
}
func (noopLogger) Reload(*logger.Config) error {
	return nil
}
