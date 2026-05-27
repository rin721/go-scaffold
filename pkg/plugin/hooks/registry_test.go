package hooks

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestRegistryEmitPriorityAndList(t *testing.T) {
	registry := NewRegistry()
	var order []string

	if err := registry.Register("plugin.invoke", HandlerFunc(func(ctx context.Context, event Event) (Result, error) {
		order = append(order, "low")
		return Result{}, nil
	}), WithName("low"), WithPriority(1)); err != nil {
		t.Fatalf("register low: %v", err)
	}
	if err := registry.Register("plugin.invoke", HandlerFunc(func(ctx context.Context, event Event) (Result, error) {
		order = append(order, "high")
		return Result{}, nil
	}), WithName("high"), WithPriority(10)); err != nil {
		t.Fatalf("register high: %v", err)
	}

	event := Event{Point: "plugin.invoke"}
	if _, err := registry.Emit(context.Background(), event); err != nil {
		t.Fatalf("emit: %v", err)
	}

	if !reflect.DeepEqual(order, []string{"high", "low"}) {
		t.Fatalf("order = %#v, want high then low", order)
	}
	list := registry.List("plugin.invoke")
	if len(list) != 2 || list[0].Name != "high" || list[1].Name != "low" {
		t.Fatalf("list = %#v, want execution order", list)
	}
}

func TestRegistryStopAndError(t *testing.T) {
	registry := NewRegistry()
	called := false
	if err := registry.Register("plugin.invoke", HandlerFunc(func(ctx context.Context, event Event) (Result, error) {
		return Stop("blocked"), nil
	}), WithPriority(10)); err != nil {
		t.Fatalf("register stopper: %v", err)
	}
	if err := registry.Register("plugin.invoke", HandlerFunc(func(ctx context.Context, event Event) (Result, error) {
		called = true
		return Result{}, nil
	}), WithPriority(1)); err != nil {
		t.Fatalf("register second: %v", err)
	}

	result, err := registry.Emit(context.Background(), Event{Point: "plugin.invoke"})
	if !errors.Is(err, ErrStopped) {
		t.Fatalf("emit error = %v, want ErrStopped", err)
	}
	if !result.Stopped || result.Reason != "blocked" {
		t.Fatalf("result = %#v, want stopped blocked", result)
	}
	if called {
		t.Fatal("handler after stop was called")
	}
}

func TestRegistryContextAndValidation(t *testing.T) {
	registry := NewRegistry()
	if err := registry.Register("", HandlerFunc(func(ctx context.Context, event Event) (Result, error) {
		return Result{}, nil
	})); !errors.Is(err, ErrInvalidPoint) {
		t.Fatalf("Register(empty) error = %v, want ErrInvalidPoint", err)
	}
	if err := registry.Register("point", nil); !errors.Is(err, ErrNilHandler) {
		t.Fatalf("Register(nil) error = %v, want ErrNilHandler", err)
	}

	if err := registry.Register("point", HandlerFunc(func(ctx context.Context, event Event) (Result, error) {
		t.Fatal("handler should not run after cancellation")
		return Result{}, nil
	})); err != nil {
		t.Fatalf("register: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := registry.Emit(ctx, Event{Point: "point"}); !errors.Is(err, context.Canceled) {
		t.Fatalf("Emit(canceled) error = %v, want context.Canceled", err)
	}
}

func TestEventResultAndServices(t *testing.T) {
	event, err := NewEvent("point", map[string]string{"hello": "world"})
	if err != nil {
		t.Fatalf("NewEvent: %v", err)
	}
	if event.CreatedAt.IsZero() || time.Since(event.CreatedAt) > time.Second {
		t.Fatalf("unexpected CreatedAt: %v", event.CreatedAt)
	}
	var eventPayload map[string]string
	if err := event.DecodePayload(&eventPayload); err != nil {
		t.Fatalf("event DecodePayload: %v", err)
	}
	if eventPayload["hello"] != "world" {
		t.Fatalf("event payload = %#v", eventPayload)
	}

	result, err := NewResult(map[string]string{"ok": "true"})
	if err != nil {
		t.Fatalf("NewResult: %v", err)
	}
	var resultPayload map[string]string
	if err := result.DecodePayload(&resultPayload); err != nil {
		t.Fatalf("result DecodePayload: %v", err)
	}
	if resultPayload["ok"] != "true" {
		t.Fatalf("result payload = %#v", resultPayload)
	}

	services := NewServices()
	if err := services.Set("", "bad"); !errors.Is(err, ErrInvalidService) {
		t.Fatalf("Set(empty) error = %v, want ErrInvalidService", err)
	}
	if err := services.Set("logger", "stub"); err != nil {
		t.Fatalf("Set(logger): %v", err)
	}
	if got, ok := services.Get("logger"); !ok || got != "stub" {
		t.Fatalf("Get(logger) = %#v, %v", got, ok)
	}
	if !reflect.DeepEqual(services.List(), []string{"logger"}) {
		t.Fatalf("List() = %#v", services.List())
	}
}
