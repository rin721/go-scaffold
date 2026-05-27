package hooks

import (
	"context"
	"errors"
	"sort"
	"sync"
	"time"
)

// Registry stores and emits hook handlers.
type Registry interface {
	Register(point Point, handler Handler, opts ...RegisterOption) error
	Emit(ctx context.Context, event Event) (Result, error)
	List(point Point) []Registration
	Services() *Services
}

// RegisterOption configures one hook registration.
type RegisterOption func(*registration)

type registration struct {
	Registration
	handler Handler
}

type registry struct {
	mu       sync.RWMutex
	handlers map[Point][]registration
	services *Services
}

// WithPriority sets handler priority. Higher priority runs earlier.
func WithPriority(priority int) RegisterOption {
	return func(reg *registration) {
		reg.Priority = priority
	}
}

// WithName sets a human-readable handler name.
func WithName(name string) RegisterOption {
	return func(reg *registration) {
		reg.Name = name
	}
}

// NewRegistry creates an empty hook registry.
func NewRegistry() Registry {
	return &registry{
		handlers: make(map[Point][]registration),
		services: NewServices(),
	}
}

// Register adds a handler for one point.
func (r *registry) Register(point Point, handler Handler, opts ...RegisterOption) error {
	if point == "" {
		return ErrInvalidPoint
	}
	if handler == nil {
		return ErrNilHandler
	}
	reg := registration{
		Registration: Registration{Point: point},
		handler:      handler,
	}
	for _, opt := range opts {
		opt(&reg)
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.handlers[point] = append(r.handlers[point], reg)
	sort.SliceStable(r.handlers[point], func(i, j int) bool {
		return r.handlers[point][i].Priority > r.handlers[point][j].Priority
	})
	return nil
}

// Emit executes all handlers for event.Point.
func (r *registry) Emit(ctx context.Context, event Event) (Result, error) {
	if event.Point == "" {
		return Result{}, ErrInvalidPoint
	}
	if event.CreatedAt.IsZero() {
		event.CreatedAt = time.Now().UTC()
	}

	handlers := r.handlersFor(event.Point)
	var last Result
	for _, reg := range handlers {
		if err := ctx.Err(); err != nil {
			return last, err
		}
		result, err := reg.handler.HandleHook(ctx, event)
		if err != nil {
			return result, err
		}
		last = result
		if result.Stopped {
			return result, ErrStopped
		}
	}
	return last, nil
}

// List returns registrations for a point in execution order.
func (r *registry) List(point Point) []Registration {
	handlers := r.handlersFor(point)
	list := make([]Registration, 0, len(handlers))
	for _, reg := range handlers {
		list = append(list, reg.Registration)
	}
	return list
}

// Services returns the shared hook service registry.
func (r *registry) Services() *Services {
	return r.services
}

func (r *registry) handlersFor(point Point) []registration {
	r.mu.RLock()
	defer r.mu.RUnlock()
	src := r.handlers[point]
	dst := make([]registration, len(src))
	copy(dst, src)
	return dst
}

// IsStopped reports whether an error means the hook chain stopped.
func IsStopped(err error) bool {
	return errors.Is(err, ErrStopped)
}
