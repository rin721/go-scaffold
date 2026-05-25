package plugin

import (
	"context"
	"fmt"
)

// Handler handles local plugin invocations.
type Handler func(ctx context.Context, req Request) (*Response, error)

type localPlugin struct {
	metadata Metadata
	handler  Handler
	close    func(context.Context) error
}

// LocalOption configures a local plugin.
type LocalOption func(*localPlugin)

// WithLocalClose sets an optional close hook for a local plugin.
func WithLocalClose(close func(context.Context) error) LocalOption {
	return func(p *localPlugin) {
		p.close = close
	}
}

// NewLocal creates an in-process plugin.
func NewLocal(metadata Metadata, handler Handler, opts ...LocalOption) (Plugin, error) {
	if metadata.Name == "" {
		return nil, fmt.Errorf("%w: name is required", ErrInvalidDefinition)
	}
	if handler == nil {
		return nil, ErrNilHandler
	}
	metadata.Protocol = ProtocolLocal
	p := &localPlugin{
		metadata: metadata,
		handler:  handler,
	}
	for _, opt := range opts {
		opt(p)
	}
	return p, nil
}

func (p *localPlugin) Metadata() Metadata {
	return p.metadata
}

func (p *localPlugin) Invoke(ctx context.Context, req Request) (*Response, error) {
	if err := req.Validate(); err != nil {
		return nil, &Error{Op: "invoke", Plugin: p.metadata.Name, Err: err}
	}
	if req.Plugin == "" {
		req.Plugin = p.metadata.Name
	}
	resp, err := p.handler(ctx, req)
	if err != nil {
		return nil, &Error{Op: "invoke", Plugin: p.metadata.Name, Err: err}
	}
	if resp == nil {
		return &Response{}, nil
	}
	return resp, nil
}

func (p *localPlugin) Close(ctx context.Context) error {
	if p.close == nil {
		return nil
	}
	if err := p.close(ctx); err != nil {
		return &Error{Op: "close", Plugin: p.metadata.Name, Err: err}
	}
	return nil
}
