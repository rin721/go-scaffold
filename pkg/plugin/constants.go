package plugin

import "time"

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
