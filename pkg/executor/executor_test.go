package executor

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestConfigValidateAppliesBoundsAndDefaults(t *testing.T) {
	cfg := Config{Name: "jobs", Size: 0}
	if err := cfg.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
	if cfg.Size != MinPoolSize {
		t.Fatalf("Size = %d, want %d", cfg.Size, MinPoolSize)
	}
	if cfg.Expiry != DefaultWorkerExpiry {
		t.Fatalf("Expiry = %s, want %s", cfg.Expiry, DefaultWorkerExpiry)
	}

	cfg = Config{Name: "jobs", Size: MaxPoolSize + 1, Expiry: time.Second}
	if err := cfg.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
	if cfg.Size != MaxPoolSize {
		t.Fatalf("Size = %d, want %d", cfg.Size, MaxPoolSize)
	}

	cfg = Config{Size: 1}
	if err := cfg.Validate(); !errors.Is(err, ErrInvalidConfig) {
		t.Fatalf("Validate() error = %v, want ErrInvalidConfig", err)
	}
}

func TestManagerExecuteAndShutdown(t *testing.T) {
	mgr, err := NewManager([]Config{{Name: "jobs", Size: 1, Expiry: time.Millisecond, NonBlocking: true}})
	if err != nil {
		t.Fatalf("NewManager() error = %v", err)
	}
	defer mgr.Shutdown()

	done := make(chan struct{})
	if err := mgr.Execute("jobs", func() { close(done) }); err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("task did not run")
	}

	if err := mgr.Execute("missing", func() {}); !errors.Is(err, ErrPoolNotFound) {
		t.Fatalf("Execute(missing) error = %v, want ErrPoolNotFound", err)
	}

	mgr.Shutdown()
	if err := mgr.Execute("jobs", func() {}); !errors.Is(err, ErrManagerClosed) {
		t.Fatalf("Execute(after shutdown) error = %v, want ErrManagerClosed", err)
	}
	if err := mgr.Reload([]Config{{Name: "jobs", Size: 1}}); !errors.Is(err, ErrManagerClosed) {
		t.Fatalf("Reload(after shutdown) error = %v, want ErrManagerClosed", err)
	}
}

func TestManagerReportsOverloadAndKeepsRunning(t *testing.T) {
	mgr, err := NewManager([]Config{{Name: "jobs", Size: 1, Expiry: time.Second, NonBlocking: true}})
	if err != nil {
		t.Fatalf("NewManager() error = %v", err)
	}
	defer mgr.Shutdown()

	started := make(chan struct{})
	release := make(chan struct{})
	finished := make(chan struct{})
	if err := mgr.Execute("jobs", func() {
		close(started)
		<-release
		close(finished)
	}); err != nil {
		t.Fatalf("Execute(blocking task) error = %v", err)
	}
	<-started

	if err := mgr.Execute("jobs", func() {}); !errors.Is(err, ErrPoolOverload) {
		close(release)
		t.Fatalf("Execute(overload) error = %v, want ErrPoolOverload", err)
	}
	close(release)
	select {
	case <-finished:
	case <-time.After(time.Second):
		t.Fatal("blocking task did not finish")
	}

	done := make(chan struct{})
	if err := mgr.Execute("jobs", func() { close(done) }); err != nil {
		t.Fatalf("Execute(after overload) error = %v", err)
	}
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("task did not run after overload")
	}
}

func TestManagerReloadFailureKeepsExistingPools(t *testing.T) {
	mgr, err := NewManager([]Config{{Name: "old", Size: 1, Expiry: time.Millisecond, NonBlocking: true}})
	if err != nil {
		t.Fatalf("NewManager() error = %v", err)
	}
	defer mgr.Shutdown()

	if err := mgr.Reload([]Config{{Name: "new", Size: 1}, {Name: "new", Size: 1}}); !errors.Is(err, ErrInvalidConfig) {
		t.Fatalf("Reload(duplicate) error = %v, want ErrInvalidConfig", err)
	}

	done := make(chan struct{})
	if err := mgr.Execute("old", func() { close(done) }); err != nil {
		t.Fatalf("Execute(old) after failed reload error = %v", err)
	}
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("old pool did not survive failed reload")
	}
	if err := mgr.Execute("new", func() {}); !errors.Is(err, ErrPoolNotFound) {
		t.Fatalf("Execute(new) error = %v, want ErrPoolNotFound", err)
	}
}

func TestPanicHandlerObservesRecoveredTaskPanic(t *testing.T) {
	handler := &recordingPanicHandler{seen: make(chan interface{}, 1)}
	SetPanicHandler(handler)
	defer SetPanicHandler(nil)

	mgr, err := NewManager([]Config{{Name: "jobs", Size: 1, Expiry: time.Millisecond, NonBlocking: true}})
	if err != nil {
		t.Fatalf("NewManager() error = %v", err)
	}
	defer mgr.Shutdown()

	if err := mgr.Execute("jobs", func() { panic("boom") }); err != nil {
		t.Fatalf("Execute(panic task) error = %v", err)
	}
	select {
	case recovered := <-handler.seen:
		if recovered != "boom" {
			t.Fatalf("recovered panic = %v, want boom", recovered)
		}
	case <-time.After(time.Second):
		t.Fatal("panic handler was not called")
	}

	done := make(chan struct{})
	if err := mgr.Execute("jobs", func() { close(done) }); err != nil {
		t.Fatalf("Execute(after panic) error = %v", err)
	}
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("pool did not continue after recovered panic")
	}
}

type recordingPanicHandler struct {
	mu   sync.Mutex
	seen chan interface{}
}

func (h *recordingPanicHandler) HandlePanic(_ PoolName, recovered interface{}) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.seen <- recovered
}
