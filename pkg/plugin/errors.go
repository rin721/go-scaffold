package plugin

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidConfig        = errors.New("invalid plugin config")
	ErrInvalidDefinition    = errors.New("invalid plugin definition")
	ErrInvalidRequest       = errors.New("invalid plugin request")
	ErrInvalidResponse      = errors.New("invalid plugin response")
	ErrUnsupportedProtocol  = errors.New("unsupported plugin protocol")
	ErrPluginExists         = errors.New("plugin already exists")
	ErrPluginNotFound       = errors.New("plugin not found")
	ErrNilPlugin            = errors.New("plugin is nil")
	ErrNilHandler           = errors.New("plugin handler is nil")
	ErrLocalFactoryNotFound = errors.New("local plugin factory not found")
	ErrHTTPStatus           = errors.New("unexpected plugin http status")
)

// Error wraps plugin operations with plugin name and operation context.
type Error struct {
	Op     string
	Plugin string
	Err    error
}

func (e *Error) Error() string {
	if e.Plugin == "" {
		return fmt.Sprintf("plugin: %s: %v", e.Op, e.Err)
	}
	return fmt.Sprintf("plugin: %s %q: %v", e.Op, e.Plugin, e.Err)
}

func (e *Error) Unwrap() error {
	return e.Err
}
