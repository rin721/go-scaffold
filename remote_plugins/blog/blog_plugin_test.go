package main

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/rei0721/go-scaffold/pkg/plugin"
	"github.com/rei0721/go-scaffold/pkg/plugin/hooks"
)

func TestBlogPluginOperationsAndHookIdentity(t *testing.T) {
	blog := NewBlogPlugin(Config{Name: "blog", Version: "test"})

	createResp, err := blog.Invoke(context.Background(), plugin.MustNewRequest(operationBlogCreate, createPostRequest{
		Title: "First post",
		Body:  "Hello from remote plugin",
	}))
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	var post Post
	if err := createResp.DecodePayload(&post); err != nil {
		t.Fatalf("decode post: %v", err)
	}
	if post.ID != 1 || post.Title != "First post" {
		t.Fatalf("post = %#v", post)
	}

	event := hooks.Event{
		Point:     plugin.HookAfterInvoke,
		Plugin:    "blog",
		Operation: operationBlogList,
		Identity: &hooks.Identity{Principal: hooks.Principal{
			ID:    "admin",
			Roles: []string{"maintainer"},
		}},
	}
	hookResp, err := blog.Invoke(context.Background(), plugin.MustNewRequest(plugin.OperationHooksExecute, event))
	if err != nil {
		t.Fatalf("hook execute: %v", err)
	}
	var result hooks.Result
	if err := hookResp.DecodePayload(&result); err != nil {
		t.Fatalf("decode hook result: %v", err)
	}
	var audit HookAudit
	if err := result.DecodePayload(&audit); err != nil {
		t.Fatalf("decode audit: %v", err)
	}
	if audit.PrincipalID != "admin" || audit.Point != string(plugin.HookAfterInvoke) {
		t.Fatalf("audit = %#v", audit)
	}
}

func TestRegisterWithHost(t *testing.T) {
	blog := NewBlogPlugin(Config{Name: "blog", Version: "test"})
	blogServer := httptest.NewServer(protectWithSharedSecret(plugin.NewHTTPServer(blog), "plugin-secret"))
	defer blogServer.Close()

	hostManager := plugin.NewManager()
	hostServer := httptest.NewServer(plugin.NewHTTPRegistrationHandler(hostManager, plugin.WithRegistrationToken("register-secret")))
	defer hostServer.Close()

	cfg := Config{
		Name:              "blog",
		Version:           "test",
		PublicHTTPURL:     blogServer.URL,
		HostHTTPURL:       hostServer.URL,
		RegistrationToken: "register-secret",
		SharedSecret:      "plugin-secret",
		HookPoints:        []string{string(plugin.HookAfterInvoke)},
	}
	if err := RegisterWithHost(context.Background(), cfg, hostServer.Client()); err != nil {
		t.Fatalf("RegisterWithHost() error = %v", err)
	}
	if _, ok := hostManager.Get("blog"); !ok {
		t.Fatal("host manager does not contain registered blog plugin")
	}

	resp, err := hostManager.Invoke(context.Background(), "blog", plugin.MustNewRequest(plugin.OperationHealth, nil))
	if err != nil {
		t.Fatalf("invoke registered blog: %v", err)
	}
	var got map[string]string
	if err := resp.DecodePayload(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if got["status"] != "ok" {
		t.Fatalf("response = %#v", got)
	}
}
