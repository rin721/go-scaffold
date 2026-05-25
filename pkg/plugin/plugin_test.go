package plugin

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLocalPluginInvoke(t *testing.T) {
	mgr := NewManager()
	echo, err := NewLocal(Metadata{Name: "echo", Protocol: ProtocolLocal}, func(ctx context.Context, req Request) (*Response, error) {
		var input map[string]string
		if err := req.DecodePayload(&input); err != nil {
			return nil, err
		}
		input["operation"] = req.Operation
		return NewResponse(input)
	})
	if err != nil {
		t.Fatalf("NewLocal() error = %v", err)
	}
	if err := mgr.Register(echo); err != nil {
		t.Fatalf("Register() error = %v", err)
	}

	req := MustNewRequest("run", map[string]string{"hello": "world"})
	resp, err := mgr.Invoke(context.Background(), "echo", req)
	if err != nil {
		t.Fatalf("Invoke() error = %v", err)
	}

	var got map[string]string
	if err := resp.DecodePayload(&got); err != nil {
		t.Fatalf("DecodePayload() error = %v", err)
	}
	if got["hello"] != "world" || got["operation"] != "run" {
		t.Fatalf("unexpected response: %#v", got)
	}
}

func TestHTTPPluginInvoke(t *testing.T) {
	var received Request
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s, want POST", r.Method)
		}
		if r.Header.Get("X-Plugin-Token") != "secret" {
			t.Fatalf("missing configured header")
		}
		if err := json.NewDecoder(r.Body).Decode(&received); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		_ = json.NewEncoder(w).Encode(MustNewResponse(map[string]string{
			"plugin":    received.Plugin,
			"operation": received.Operation,
		}))
	}))
	defer server.Close()

	mgr := NewManager()
	remote, err := NewHTTP(Definition{
		Name:     "remote",
		Protocol: ProtocolHTTP,
		Endpoint: server.URL,
		Headers:  map[string]string{"X-Plugin-Token": "secret"},
	})
	if err != nil {
		t.Fatalf("NewHTTP() error = %v", err)
	}
	if err := mgr.Register(remote); err != nil {
		t.Fatalf("Register() error = %v", err)
	}

	resp, err := mgr.Invoke(context.Background(), "remote", MustNewRequest("status", nil))
	if err != nil {
		t.Fatalf("Invoke() error = %v", err)
	}
	var got map[string]string
	if err := resp.DecodePayload(&got); err != nil {
		t.Fatalf("DecodePayload() error = %v", err)
	}
	if got["plugin"] != "remote" || got["operation"] != "status" {
		t.Fatalf("unexpected response: %#v", got)
	}
}

func TestHTTPPluginStatusError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		_ = json.NewEncoder(w).Encode(&Response{Error: "upstream failed"})
	}))
	defer server.Close()

	mgr := NewManager()
	remote, err := NewHTTP(Definition{Name: "remote", Protocol: ProtocolHTTP, Endpoint: server.URL})
	if err != nil {
		t.Fatalf("NewHTTP() error = %v", err)
	}
	if err := mgr.Register(remote); err != nil {
		t.Fatalf("Register() error = %v", err)
	}

	_, err = mgr.Invoke(context.Background(), "remote", MustNewRequest("run", nil))
	if !errors.Is(err, ErrHTTPStatus) {
		t.Fatalf("Invoke() error = %v, want ErrHTTPStatus", err)
	}
}

func TestManagerUnsupportedProtocol(t *testing.T) {
	_, err := NewHTTP(Definition{Name: "future", Protocol: ProtocolWS, Endpoint: "ws://example.com"})
	if !errors.Is(err, ErrUnsupportedProtocol) {
		t.Fatalf("NewHTTP() error = %v, want ErrUnsupportedProtocol", err)
	}
}

func TestManagerPluginNotFound(t *testing.T) {
	mgr := NewManager()
	_, err := mgr.Invoke(context.Background(), "missing", MustNewRequest("run", nil))
	if !errors.Is(err, ErrPluginNotFound) {
		t.Fatalf("Invoke() error = %v, want ErrPluginNotFound", err)
	}
}
