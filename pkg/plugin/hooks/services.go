package hooks

import (
	"sort"
	"sync"
)

// Services stores optional shared services for hook handlers.
type Services struct {
	mu     sync.RWMutex
	values map[string]interface{}
}

// NewServices creates an empty service registry.
func NewServices() *Services {
	return &Services{values: make(map[string]interface{})}
}

// Set stores a service by name.
func (s *Services) Set(name string, service interface{}) error {
	if name == "" || service == nil {
		return ErrInvalidService
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[name] = service
	return nil
}

// Get returns a service by name.
func (s *Services) Get(name string) (interface{}, bool) {
	if s == nil {
		return nil, false
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.values[name]
	return value, ok
}

// List returns service names in stable order.
func (s *Services) List() []string {
	if s == nil {
		return nil
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	names := make([]string, 0, len(s.values))
	for name := range s.values {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
