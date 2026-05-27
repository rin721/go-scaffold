package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/rei0721/go-scaffold/pkg/plugin/hooks"
)

// Manager is a passive plugin registry and runtime.
//
// Plugin services or the host composition layer create plugin instances and
// call Register. The manager does not discover, load, or create services.
type Manager interface {
	Register(plugin Plugin) error
	Invoke(ctx context.Context, name string, req Request) (*Response, error)
	Get(name string) (Plugin, bool)
	List() []Metadata
	Hooks() hooks.Registry
	RegisterHook(point hooks.Point, handler hooks.Handler, opts ...hooks.RegisterOption) error
	Close(ctx context.Context) error
}

type manager struct {
	mu                sync.RWMutex
	plugins           map[string]Plugin
	hooks             hooks.Registry
	hookEventEnricher HookEventEnricher
}

// ManagerOption configures a plugin manager.
type ManagerOption func(*manager)

// HookEventEnricher can add host-owned context to hook events.
type HookEventEnricher func(ctx context.Context, event hooks.Event) hooks.Event

// WithHooks sets the hook registry used by a manager.
func WithHooks(registry hooks.Registry) ManagerOption {
	return func(m *manager) {
		if registry != nil {
			m.hooks = registry
		}
	}
}

// WithHookEventEnricher configures a function that enriches every hook event
// before handlers receive it.
func WithHookEventEnricher(enricher HookEventEnricher) ManagerOption {
	return func(m *manager) {
		m.hookEventEnricher = enricher
	}
}

// NewManager creates a plugin manager.
func NewManager(opts ...ManagerOption) Manager {
	mgr := &manager{
		plugins: make(map[string]Plugin),
		hooks:   hooks.NewRegistry(),
	}
	for _, opt := range opts {
		opt(mgr)
	}
	return mgr
}

// Register adds a plugin instance.
func (m *manager) Register(plugin Plugin) error {
	if plugin == nil {
		return ErrNilPlugin
	}
	metadata := plugin.Metadata()
	if metadata.Name == "" {
		return fmt.Errorf("%w: name is required", ErrInvalidDefinition)
	}

	if _, err := m.emitHook(context.Background(), HookBeforeRegister, metadata.Name, "", registerHookPayload{Metadata: metadata}); err != nil {
		return &Error{Op: "before_register", Plugin: metadata.Name, Err: err}
	}

	m.mu.Lock()
	if _, exists := m.plugins[metadata.Name]; exists {
		m.mu.Unlock()
		return &Error{Op: "register", Plugin: metadata.Name, Err: ErrPluginExists}
	}
	m.plugins[metadata.Name] = plugin
	m.mu.Unlock()

	if _, err := m.emitHook(context.Background(), HookAfterRegister, metadata.Name, "", registerHookPayload{Metadata: metadata}); err != nil {
		return &Error{Op: "after_register", Plugin: metadata.Name, Err: err}
	}
	return nil
}

// Invoke calls a registered plugin.
func (m *manager) Invoke(ctx context.Context, name string, req Request) (*Response, error) {
	plugin, ok := m.Get(name)
	if !ok {
		return nil, &Error{Op: "invoke", Plugin: name, Err: ErrPluginNotFound}
	}
	if req.Plugin == "" {
		req.Plugin = name
	}
	skipHooks := req.Operation == OperationHooksExecute
	if !skipHooks {
		if _, err := m.emitHook(ctx, HookIAMAuthorize, name, req.Operation, invokeHookPayload{Plugin: name, Request: req}); err != nil {
			return nil, &Error{Op: "iam_authorize", Plugin: name, Err: err}
		}
		if _, err := m.emitHook(ctx, HookBeforeInvoke, name, req.Operation, invokeHookPayload{Plugin: name, Request: req}); err != nil {
			return nil, &Error{Op: "before_invoke", Plugin: name, Err: err}
		}
	}

	resp, err := plugin.Invoke(ctx, req)
	if err != nil {
		if !skipHooks {
			_, _ = m.emitHook(ctx, HookInvokeError, name, req.Operation, invokeHookPayload{Plugin: name, Request: req, Response: resp, Error: err.Error()})
		}
		return resp, err
	}

	if !skipHooks {
		if _, err := m.emitHook(ctx, HookAfterInvoke, name, req.Operation, invokeHookPayload{Plugin: name, Request: req, Response: resp}); err != nil {
			return resp, &Error{Op: "after_invoke", Plugin: name, Err: err}
		}
	}
	return resp, nil
}

// Get returns a registered plugin by name.
func (m *manager) Get(name string) (Plugin, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	p, ok := m.plugins[name]
	return p, ok
}

// List returns metadata for all registered plugins.
func (m *manager) List() []Metadata {
	m.mu.RLock()
	defer m.mu.RUnlock()
	list := make([]Metadata, 0, len(m.plugins))
	for _, plugin := range m.plugins {
		list = append(list, plugin.Metadata())
	}
	return list
}

// Hooks returns the manager hook registry.
func (m *manager) Hooks() hooks.Registry {
	return m.hooks
}

// RegisterHook registers a hook handler on the manager hook registry.
func (m *manager) RegisterHook(point hooks.Point, handler hooks.Handler, opts ...hooks.RegisterOption) error {
	return m.hooks.Register(point, handler, opts...)
}

// Close closes all registered plugins.
func (m *manager) Close(ctx context.Context) error {
	if _, err := m.emitHook(ctx, HookBeforeClose, "", "", nil); err != nil {
		return &Error{Op: "before_close", Err: err}
	}

	m.mu.RLock()
	plugins := make([]Plugin, 0, len(m.plugins))
	for _, plugin := range m.plugins {
		plugins = append(plugins, plugin)
	}
	m.mu.RUnlock()

	var closeErr error
	for _, plugin := range plugins {
		if err := plugin.Close(ctx); err != nil && closeErr == nil {
			closeErr = err
		}
	}
	if _, err := m.emitHook(ctx, HookAfterClose, "", "", closeHookPayload{Error: errorString(closeErr)}); err != nil && closeErr == nil {
		closeErr = &Error{Op: "after_close", Err: err}
	}
	return closeErr
}

type registerHookPayload struct {
	Metadata Metadata `json:"metadata"`
}

type invokeHookPayload struct {
	Plugin   string    `json:"plugin"`
	Request  Request   `json:"request"`
	Response *Response `json:"response,omitempty"`
	Error    string    `json:"error,omitempty"`
}

type closeHookPayload struct {
	Error string `json:"error,omitempty"`
}

func (m *manager) emitHook(ctx context.Context, point hooks.Point, pluginName, operation string, payload interface{}) (hooks.Result, error) {
	if m.hooks == nil {
		return hooks.Result{}, nil
	}
	event := hooks.Event{
		Point:     point,
		Plugin:    pluginName,
		Operation: operation,
		CreatedAt: time.Now().UTC(),
	}
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return hooks.Result{}, err
		}
		event.Payload = data
	}
	if m.hookEventEnricher != nil {
		event = m.hookEventEnricher(ctx, event)
	}
	return m.hooks.Emit(ctx, event)
}

func errorString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
