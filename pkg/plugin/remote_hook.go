package plugin

import (
	"context"
	"encoding/json"

	"github.com/rei0721/go-scaffold/pkg/plugin/hooks"
)

// RemoteHook adapts a remote plugin operation to a local hook handler.
type RemoteHook struct {
	manager Manager
	plugin  string
}

// NewRemoteHook creates a hook handler that invokes a remote plugin.
func NewRemoteHook(manager Manager, pluginName string) hooks.Handler {
	return &RemoteHook{manager: manager, plugin: pluginName}
}

// HandleHook sends the hook event through the standard plugin invocation path.
func (h *RemoteHook) HandleHook(ctx context.Context, event hooks.Event) (hooks.Result, error) {
	if h == nil || h.manager == nil || h.plugin == "" {
		return hooks.Result{}, ErrInvalidDefinition
	}
	payload, err := json.Marshal(event)
	if err != nil {
		return hooks.Result{}, err
	}
	resp, err := h.manager.Invoke(ctx, h.plugin, Request{
		Plugin:    h.plugin,
		Operation: OperationHooksExecute,
		Payload:   payload,
	})
	if err != nil {
		return hooks.Result{}, err
	}
	var result hooks.Result
	if err := resp.DecodePayload(&result); err != nil {
		return hooks.Result{}, err
	}
	return result, nil
}
