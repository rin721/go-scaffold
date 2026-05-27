package plugin

import (
	"time"

	"github.com/rei0721/go-scaffold/pkg/plugin/hooks"
)

// Protocol identifies how a plugin is invoked.
type Protocol string

const (
	// ProtocolLocal invokes an in-process plugin registered by a plugin service.
	ProtocolLocal Protocol = "local"

	// ProtocolHTTP invokes a remote plugin over HTTP.
	ProtocolHTTP Protocol = "http"

	// ProtocolRPC is reserved for a future RPC adapter.
	ProtocolRPC Protocol = "rpc"

	// ProtocolWS is reserved for a future WebSocket adapter.
	ProtocolWS Protocol = "ws"
)

const (
	DefaultTimeout          = 10 * time.Second
	DefaultMaxResponseBytes = 10 << 20
)

const (
	OperationManifest     = "manifest"
	OperationHealth       = "health"
	OperationHooksExecute = "hooks.execute"
)

const HTTPInvokePath = "/plugin/v1/invoke"
const HTTPRegisterPath = "/plugin/v1/register"

const (
	HookBeforeRegister hooks.Point = "plugin.before_register"
	HookAfterRegister  hooks.Point = "plugin.after_register"
	HookIAMAuthorize   hooks.Point = "plugin.iam_authorize"
	HookBeforeInvoke   hooks.Point = "plugin.before_invoke"
	HookAfterInvoke    hooks.Point = "plugin.after_invoke"
	HookInvokeError    hooks.Point = "plugin.invoke_error"
	HookBeforeClose    hooks.Point = "plugin.before_close"
	HookAfterClose     hooks.Point = "plugin.after_close"
	HookConfigChanged  hooks.Point = "plugin.config_changed"
	HookLoggerReady    hooks.Point = "plugin.logger_ready"
)
