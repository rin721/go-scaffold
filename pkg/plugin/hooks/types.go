package hooks

import (
	"context"
	"encoding/json"
	"time"
)

// Point identifies where a hook is emitted.
type Point string

// Event is passed to every registered hook handler.
type Event struct {
	Point     Point             `json:"point"`
	Plugin    string            `json:"plugin,omitempty"`
	Operation string            `json:"operation,omitempty"`
	Payload   json.RawMessage   `json:"payload,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
	CreatedAt time.Time         `json:"created_at"`
}

// NewEvent creates an Event and marshals payload into JSON.
func NewEvent(point Point, payload interface{}) (Event, error) {
	event := Event{Point: point, CreatedAt: time.Now().UTC()}
	if payload == nil {
		return event, nil
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return Event{}, err
	}
	event.Payload = data
	return event, nil
}

// DecodePayload decodes the event payload.
func (e Event) DecodePayload(v interface{}) error {
	if len(e.Payload) == 0 {
		return nil
	}
	return json.Unmarshal(e.Payload, v)
}

// Result describes the outcome of a hook handler or hook chain.
type Result struct {
	Stopped  bool              `json:"stopped,omitempty"`
	Reason   string            `json:"reason,omitempty"`
	Payload  json.RawMessage   `json:"payload,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// NewResult creates a Result and marshals payload into JSON.
func NewResult(payload interface{}) (Result, error) {
	result := Result{}
	if payload == nil {
		return result, nil
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return Result{}, err
	}
	result.Payload = data
	return result, nil
}

// DecodePayload decodes the result payload.
func (r Result) DecodePayload(v interface{}) error {
	if len(r.Payload) == 0 {
		return nil
	}
	return json.Unmarshal(r.Payload, v)
}

// Stop returns a Result that stops the hook chain.
func Stop(reason string) Result {
	return Result{Stopped: true, Reason: reason}
}

// Handler handles one hook event.
type Handler interface {
	HandleHook(ctx context.Context, event Event) (Result, error)
}

// HandlerFunc adapts a function to Handler.
type HandlerFunc func(ctx context.Context, event Event) (Result, error)

// HandleHook implements Handler.
func (f HandlerFunc) HandleHook(ctx context.Context, event Event) (Result, error) {
	if f == nil {
		return Result{}, ErrNilHandler
	}
	return f(ctx, event)
}

// Registration describes a registered hook without exposing the handler.
type Registration struct {
	Point    Point  `json:"point"`
	Name     string `json:"name,omitempty"`
	Priority int    `json:"priority"`
}
