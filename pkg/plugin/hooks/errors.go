package hooks

import "errors"

var (
	ErrInvalidPoint   = errors.New("invalid hook point")
	ErrNilHandler     = errors.New("hook handler is nil")
	ErrStopped        = errors.New("hook chain stopped")
	ErrInvalidResult  = errors.New("invalid hook result")
	ErrInvalidService = errors.New("invalid hook service")
)
