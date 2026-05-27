package plugin

import (
	"crypto/subtle"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/rei0721/go-scaffold/pkg/plugin/hooks"
)

const registrationTokenHeader = "X-Plugin-Registration-Token"

// HookBinding describes a remote plugin hook registration request.
type HookBinding struct {
	Point    hooks.Point `json:"point"`
	Plugin   string      `json:"plugin,omitempty"`
	Name     string      `json:"name,omitempty"`
	Priority int         `json:"priority,omitempty"`
}

// RegistrationRequest is posted by a remote plugin service to the host.
type RegistrationRequest struct {
	Plugin Definition    `json:"plugin"`
	Hooks  []HookBinding `json:"hooks,omitempty"`
}

// RegistrationResponse describes the host-side registration result.
type RegistrationResponse struct {
	Metadata Metadata             `json:"metadata"`
	Hooks    []hooks.Registration `json:"hooks,omitempty"`
}

// RegistrationOption configures the HTTP registration handler.
type RegistrationOption func(*registrationOptions)

type registrationOptions struct {
	token       string
	httpOptions []HTTPOption
}

// WithRegistrationToken requires a bearer token or X-Plugin-Registration-Token.
func WithRegistrationToken(token string) RegistrationOption {
	return func(opts *registrationOptions) {
		opts.token = strings.TrimSpace(token)
	}
}

// WithRegistrationHTTPOptions configures HTTP plugin instances created from registrations.
func WithRegistrationHTTPOptions(httpOptions ...HTTPOption) RegistrationOption {
	return func(opts *registrationOptions) {
		opts.httpOptions = append(opts.httpOptions, httpOptions...)
	}
}

// NewHTTPRegistrationHandler accepts explicit remote plugin registrations.
func NewHTTPRegistrationHandler(manager Manager, opts ...RegistrationOption) http.Handler {
	options := registrationOptions{}
	for _, opt := range opts {
		opt(&options)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != HTTPRegisterPath {
			writeHTTPPluginError(w, http.StatusNotFound, "plugin registration endpoint not found")
			return
		}
		if r.Method != http.MethodPost {
			writeHTTPPluginError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}
		if manager == nil {
			writeHTTPPluginError(w, http.StatusInternalServerError, "plugin manager is nil")
			return
		}
		if !registrationAuthorized(r, options.token) {
			writeHTTPPluginError(w, http.StatusUnauthorized, "plugin registration unauthorized")
			return
		}

		var req RegistrationRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeHTTPPluginError(w, http.StatusBadRequest, err.Error())
			return
		}
		resp, err := RegisterRemoteHTTP(manager, req, options.httpOptions...)
		if err != nil {
			status := http.StatusBadRequest
			if errors.Is(err, ErrPluginExists) {
				status = http.StatusConflict
			}
			writeHTTPPluginResponse(w, status, &Response{Error: err.Error()})
			return
		}
		writeHTTPPluginResponse(w, http.StatusCreated, MustNewResponse(resp))
	})
}

// RegisterRemoteHTTP registers a remote HTTP plugin and optional remote hooks.
func RegisterRemoteHTTP(manager Manager, req RegistrationRequest, httpOptions ...HTTPOption) (RegistrationResponse, error) {
	if manager == nil {
		return RegistrationResponse{}, ErrInvalidConfig
	}
	def := req.Plugin
	if def.Protocol == "" {
		def.Protocol = ProtocolHTTP
	}
	remote, err := NewHTTP(def, httpOptions...)
	if err != nil {
		return RegistrationResponse{}, err
	}
	if err := manager.Register(remote); err != nil {
		return RegistrationResponse{}, err
	}

	registrations := make([]hooks.Registration, 0, len(req.Hooks))
	for _, binding := range req.Hooks {
		pluginName := binding.Plugin
		if pluginName == "" {
			pluginName = def.Name
		}
		if pluginName == "" {
			return RegistrationResponse{}, fmt.Errorf("%w: hook plugin is required", ErrInvalidDefinition)
		}
		if err := manager.RegisterHook(
			binding.Point,
			NewRemoteHook(manager, pluginName),
			hooks.WithName(binding.Name),
			hooks.WithPriority(binding.Priority),
		); err != nil {
			return RegistrationResponse{}, err
		}
		registrations = append(registrations, hooks.Registration{
			Point:    binding.Point,
			Name:     binding.Name,
			Priority: binding.Priority,
		})
	}

	return RegistrationResponse{
		Metadata: remote.Metadata(),
		Hooks:    registrations,
	}, nil
}

func registrationAuthorized(r *http.Request, token string) bool {
	if token == "" {
		return true
	}
	if subtleTokenEqual(r.Header.Get(registrationTokenHeader), token) {
		return true
	}
	auth := strings.TrimSpace(r.Header.Get("Authorization"))
	if strings.HasPrefix(strings.ToLower(auth), "bearer ") {
		return subtleTokenEqual(strings.TrimSpace(auth[len("bearer "):]), token)
	}
	return false
}

func subtleTokenEqual(got, want string) bool {
	return got != "" && subtle.ConstantTimeCompare([]byte(got), []byte(want)) == 1
}
