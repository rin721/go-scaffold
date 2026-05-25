package plugin

import (
	"context"
	"fmt"
	"sync"
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
	Close(ctx context.Context) error
}

type manager struct {
	mu      sync.RWMutex
	plugins map[string]Plugin
}

// NewManager creates a plugin manager.
func NewManager() Manager {
	return &manager{
		plugins: make(map[string]Plugin),
	}
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

	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.plugins[metadata.Name]; exists {
		return &Error{Op: "register", Plugin: metadata.Name, Err: ErrPluginExists}
	}
	m.plugins[metadata.Name] = plugin
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
	return plugin.Invoke(ctx, req)
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

// Close closes all registered plugins.
func (m *manager) Close(ctx context.Context) error {
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
	return closeErr
}
