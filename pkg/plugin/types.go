package plugin

import (
	"context"
	"encoding/json"
	"errors"
)

// Metadata describes a plugin instance.
type Metadata struct {
	Name         string            `json:"name" yaml:"name" mapstructure:"name"`
	Version      string            `json:"version,omitempty" yaml:"version,omitempty" mapstructure:"version"`
	Protocol     Protocol          `json:"protocol" yaml:"protocol" mapstructure:"protocol"`
	Description  string            `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description"`
	Capabilities []string          `json:"capabilities,omitempty" yaml:"capabilities,omitempty" mapstructure:"capabilities"`
	Labels       map[string]string `json:"labels,omitempty" yaml:"labels,omitempty" mapstructure:"labels"`
}

// Plugin is the common interface implemented by local and remote plugins.
type Plugin interface {
	Metadata() Metadata
	Invoke(ctx context.Context, req Request) (*Response, error)
	Close(ctx context.Context) error
}

// Request is the common invocation payload sent to all plugin protocols.
type Request struct {
	Plugin    string            `json:"plugin,omitempty"`
	Operation string            `json:"operation"`
	Payload   json.RawMessage   `json:"payload,omitempty"`
	Headers   map[string]string `json:"headers,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

// Validate verifies the request has enough information to be dispatched.
func (r Request) Validate() error {
	if r.Operation == "" {
		return ErrInvalidRequest
	}
	return nil
}

// NewRequest creates a Request and marshals payload into JSON.
func NewRequest(operation string, payload interface{}) (Request, error) {
	req := Request{Operation: operation}
	if payload == nil {
		return req, nil
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return Request{}, err
	}
	req.Payload = data
	return req, nil
}

// MustNewRequest creates a Request and panics on marshal errors.
func MustNewRequest(operation string, payload interface{}) Request {
	req, err := NewRequest(operation, payload)
	if err != nil {
		panic(err)
	}
	return req
}

// DecodePayload decodes the request payload.
func (r Request) DecodePayload(v interface{}) error {
	if len(r.Payload) == 0 {
		return nil
	}
	return json.Unmarshal(r.Payload, v)
}

// Response is the common plugin invocation response.
type Response struct {
	Payload  json.RawMessage   `json:"payload,omitempty"`
	Headers  map[string]string `json:"headers,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Error    string            `json:"error,omitempty"`
}

// NewResponse creates a Response and marshals payload into JSON.
func NewResponse(payload interface{}) (*Response, error) {
	resp := &Response{}
	if payload == nil {
		return resp, nil
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	resp.Payload = data
	return resp, nil
}

// MustNewResponse creates a Response and panics on marshal errors.
func MustNewResponse(payload interface{}) *Response {
	resp, err := NewResponse(payload)
	if err != nil {
		panic(err)
	}
	return resp
}

// DecodePayload decodes the response payload.
func (r *Response) DecodePayload(v interface{}) error {
	if r == nil {
		return ErrInvalidResponse
	}
	if r.Error != "" {
		return errors.New(r.Error)
	}
	if len(r.Payload) == 0 {
		return nil
	}
	return json.Unmarshal(r.Payload, v)
}
