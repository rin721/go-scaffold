package plugin

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Manager manages plugin registration, loading, invocation, and shutdown.
type Manager interface {
	Load(config *Config) error
	Register(plugin Plugin) error
	RegisterLocalFactory(name string, factory LocalFactory)
	Invoke(ctx context.Context, name string, req Request) (*Response, error)
	Get(name string) (Plugin, bool)
	List() []Metadata
	Close(ctx context.Context) error
}

// Option configures a plugin manager.
type Option func(*manager)

// WithHTTPClient configures the HTTP client used by HTTP plugins.
func WithHTTPClient(client *http.Client) Option {
	return func(m *manager) {
		if client != nil {
			m.httpClient = client
		}
	}
}

// WithLocalFactory registers a local factory during manager creation.
func WithLocalFactory(name string, factory LocalFactory) Option {
	return func(m *manager) {
		m.RegisterLocalFactory(name, factory)
	}
}

type manager struct {
	mu               sync.RWMutex
	plugins          map[string]Plugin
	localFactories   map[string]LocalFactory
	httpClient       *http.Client
	defaultTimeout   time.Duration
	maxResponseBytes int64
}

// NewManager creates a plugin manager.
func NewManager(opts ...Option) Manager {
	m := &manager{
		plugins:          make(map[string]Plugin),
		localFactories:   make(map[string]LocalFactory),
		httpClient:       http.DefaultClient,
		defaultTimeout:   DefaultTimeout,
		maxResponseBytes: DefaultMaxResponseBytes,
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// Load creates and registers plugins from config.
func (m *manager) Load(config *Config) error {
	if config == nil {
		config = DefaultConfig()
	}
	if config.DefaultTimeout > 0 {
		m.defaultTimeout = config.DefaultTimeout
	}
	if config.MaxResponseBytes > 0 {
		m.maxResponseBytes = config.MaxResponseBytes
	}
	if err := config.Validate(); err != nil {
		return err
	}
	for _, def := range config.Plugins {
		p, err := m.build(def)
		if err != nil {
			return err
		}
		if err := m.Register(p); err != nil {
			return err
		}
	}
	return nil
}

func (m *manager) build(def Definition) (Plugin, error) {
	switch def.Protocol {
	case ProtocolLocal:
		key := def.Endpoint
		if key == "" {
			key = def.Name
		}
		m.mu.RLock()
		factory, ok := m.localFactories[key]
		m.mu.RUnlock()
		if !ok {
			return nil, &Error{Op: "load", Plugin: def.Name, Err: ErrLocalFactoryNotFound}
		}
		return factory(def)
	case ProtocolHTTP:
		return newHTTP(def, m.httpClient, m.defaultTimeout, m.maxResponseBytes)
	default:
		return nil, &Error{Op: "load", Plugin: def.Name, Err: fmt.Errorf("%w: %s", ErrUnsupportedProtocol, def.Protocol)}
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

// RegisterLocalFactory registers a factory for in-process plugins.
func (m *manager) RegisterLocalFactory(name string, factory LocalFactory) {
	if name == "" || factory == nil {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.localFactories[name] = factory
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
