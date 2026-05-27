package plugin

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rei0721/go-scaffold/pkg/plugin/hooks"
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

func TestHTTPServerHelperInvoke(t *testing.T) {
	echo, err := NewLocal(Metadata{Name: "echo", Protocol: ProtocolLocal}, func(ctx context.Context, req Request) (*Response, error) {
		return NewResponse(map[string]string{"plugin": req.Plugin, "operation": req.Operation})
	})
	if err != nil {
		t.Fatalf("NewLocal() error = %v", err)
	}
	server := httptest.NewServer(NewHTTPServer(echo))
	defer server.Close()

	remote, err := NewHTTP(Definition{Name: "echo", Protocol: ProtocolHTTP, Endpoint: server.URL + HTTPInvokePath})
	if err != nil {
		t.Fatalf("NewHTTP() error = %v", err)
	}
	resp, err := remote.Invoke(context.Background(), MustNewRequest("status", nil))
	if err != nil {
		t.Fatalf("Invoke() error = %v", err)
	}
	var got map[string]string
	if err := resp.DecodePayload(&got); err != nil {
		t.Fatalf("DecodePayload() error = %v", err)
	}
	if got["plugin"] != "echo" || got["operation"] != "status" {
		t.Fatalf("unexpected response: %#v", got)
	}
}

func TestHTTPServerHelperRejectsInvalidRequests(t *testing.T) {
	echo, err := NewLocal(Metadata{Name: "echo", Protocol: ProtocolLocal}, func(ctx context.Context, req Request) (*Response, error) {
		return NewResponse(nil)
	})
	if err != nil {
		t.Fatalf("NewLocal() error = %v", err)
	}
	server := httptest.NewServer(NewHTTPServer(echo))
	defer server.Close()

	methodResp, err := http.Get(server.URL + HTTPInvokePath)
	if err != nil {
		t.Fatalf("GET invoke path: %v", err)
	}
	defer methodResp.Body.Close()
	if methodResp.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("GET status = %d, want 405", methodResp.StatusCode)
	}

	missingResp, err := http.Post(server.URL+"/missing", "application/json", nil)
	if err != nil {
		t.Fatalf("POST missing path: %v", err)
	}
	defer missingResp.Body.Close()
	if missingResp.StatusCode != http.StatusNotFound {
		t.Fatalf("missing status = %d, want 404", missingResp.StatusCode)
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

func TestManagerInvokeHooks(t *testing.T) {
	mgr := NewManager()
	var order []hooks.Point
	for _, point := range []hooks.Point{HookIAMAuthorize, HookBeforeInvoke, HookAfterInvoke} {
		p := point
		if err := mgr.RegisterHook(p, hooks.HandlerFunc(func(ctx context.Context, event hooks.Event) (hooks.Result, error) {
			order = append(order, p)
			if event.Plugin != "echo" || event.Operation != "run" {
				t.Fatalf("event = %#v, want echo/run", event)
			}
			return hooks.Result{}, nil
		})); err != nil {
			t.Fatalf("RegisterHook(%s): %v", p, err)
		}
	}
	echo, err := NewLocal(Metadata{Name: "echo", Protocol: ProtocolLocal}, func(ctx context.Context, req Request) (*Response, error) {
		return NewResponse(map[string]string{"ok": "true"})
	})
	if err != nil {
		t.Fatalf("NewLocal() error = %v", err)
	}
	if err := mgr.Register(echo); err != nil {
		t.Fatalf("Register() error = %v", err)
	}

	if _, err := mgr.Invoke(context.Background(), "echo", MustNewRequest("run", nil)); err != nil {
		t.Fatalf("Invoke() error = %v", err)
	}

	want := []hooks.Point{HookIAMAuthorize, HookBeforeInvoke, HookAfterInvoke}
	if len(order) != len(want) {
		t.Fatalf("order = %#v, want %#v", order, want)
	}
	for i := range want {
		if order[i] != want[i] {
			t.Fatalf("order = %#v, want %#v", order, want)
		}
	}
}

func TestManagerBeforeInvokeHookCanStop(t *testing.T) {
	mgr := NewManager()
	called := false
	if err := mgr.RegisterHook(HookBeforeInvoke, hooks.HandlerFunc(func(ctx context.Context, event hooks.Event) (hooks.Result, error) {
		return hooks.Stop("blocked"), nil
	})); err != nil {
		t.Fatalf("RegisterHook() error = %v", err)
	}
	echo, err := NewLocal(Metadata{Name: "echo", Protocol: ProtocolLocal}, func(ctx context.Context, req Request) (*Response, error) {
		called = true
		return NewResponse(nil)
	})
	if err != nil {
		t.Fatalf("NewLocal() error = %v", err)
	}
	if err := mgr.Register(echo); err != nil {
		t.Fatalf("Register() error = %v", err)
	}

	_, err = mgr.Invoke(context.Background(), "echo", MustNewRequest("run", nil))
	if !errors.Is(err, hooks.ErrStopped) {
		t.Fatalf("Invoke() error = %v, want hooks.ErrStopped", err)
	}
	if called {
		t.Fatal("plugin handler was called after before hook stopped")
	}
}

func TestManagerInvokeErrorHookIsBestEffort(t *testing.T) {
	mgr := NewManager()
	hookCalled := false
	pluginErr := errors.New("plugin failed")
	if err := mgr.RegisterHook(HookInvokeError, hooks.HandlerFunc(func(ctx context.Context, event hooks.Event) (hooks.Result, error) {
		hookCalled = true
		return hooks.Result{}, errors.New("hook failed")
	})); err != nil {
		t.Fatalf("RegisterHook() error = %v", err)
	}
	failing, err := NewLocal(Metadata{Name: "failing", Protocol: ProtocolLocal}, func(ctx context.Context, req Request) (*Response, error) {
		return nil, pluginErr
	})
	if err != nil {
		t.Fatalf("NewLocal() error = %v", err)
	}
	if err := mgr.Register(failing); err != nil {
		t.Fatalf("Register() error = %v", err)
	}

	_, err = mgr.Invoke(context.Background(), "failing", MustNewRequest("run", nil))
	if !errors.Is(err, pluginErr) {
		t.Fatalf("Invoke() error = %v, want plugin error", err)
	}
	if !hookCalled {
		t.Fatal("invoke error hook was not called")
	}
}

func TestManagerSkipsHooksExecuteOperation(t *testing.T) {
	mgr := NewManager()
	hookCalled := false
	if err := mgr.RegisterHook(HookBeforeInvoke, hooks.HandlerFunc(func(ctx context.Context, event hooks.Event) (hooks.Result, error) {
		hookCalled = true
		return hooks.Result{}, nil
	})); err != nil {
		t.Fatalf("RegisterHook() error = %v", err)
	}
	echo, err := NewLocal(Metadata{Name: "echo", Protocol: ProtocolLocal}, func(ctx context.Context, req Request) (*Response, error) {
		return NewResponse(map[string]string{"operation": req.Operation})
	})
	if err != nil {
		t.Fatalf("NewLocal() error = %v", err)
	}
	if err := mgr.Register(echo); err != nil {
		t.Fatalf("Register() error = %v", err)
	}

	if _, err := mgr.Invoke(context.Background(), "echo", MustNewRequest(OperationHooksExecute, nil)); err != nil {
		t.Fatalf("Invoke() error = %v", err)
	}
	if hookCalled {
		t.Fatal("hooks.execute should not re-enter manager hooks")
	}
}

func TestRemoteHookInvokesHooksExecute(t *testing.T) {
	mgr := NewManager()
	remote, err := NewLocal(Metadata{Name: "remote", Protocol: ProtocolLocal}, func(ctx context.Context, req Request) (*Response, error) {
		if req.Operation != OperationHooksExecute {
			t.Fatalf("operation = %q, want hooks.execute", req.Operation)
		}
		var event hooks.Event
		if err := req.DecodePayload(&event); err != nil {
			t.Fatalf("DecodePayload(event): %v", err)
		}
		result, err := hooks.NewResult(map[string]string{
			"point": string(event.Point),
		})
		if err != nil {
			return nil, err
		}
		return NewResponse(result)
	})
	if err != nil {
		t.Fatalf("NewLocal() error = %v", err)
	}
	if err := mgr.Register(remote); err != nil {
		t.Fatalf("Register() error = %v", err)
	}

	handler := NewRemoteHook(mgr, "remote")
	result, err := handler.HandleHook(context.Background(), hooks.Event{Point: HookBeforeInvoke, Plugin: "echo"})
	if err != nil {
		t.Fatalf("HandleHook() error = %v", err)
	}
	var got map[string]string
	if err := result.DecodePayload(&got); err != nil {
		t.Fatalf("DecodePayload(result): %v", err)
	}
	if got["point"] != string(HookBeforeInvoke) {
		t.Fatalf("result payload = %#v", got)
	}
}
