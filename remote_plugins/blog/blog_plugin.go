package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/rei0721/go-scaffold/pkg/plugin"
	"github.com/rei0721/go-scaffold/pkg/plugin/hooks"
)

const (
	operationBlogCreate = "blog.create"
	operationBlogList   = "blog.list"
)

type BlogPlugin struct {
	mu       sync.Mutex
	metadata plugin.Metadata
	posts    []Post
	audits   []HookAudit
	nextID   int
}

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type HookAudit struct {
	Point       string `json:"point"`
	Plugin      string `json:"plugin,omitempty"`
	Operation   string `json:"operation,omitempty"`
	PrincipalID string `json:"principal_id,omitempty"`
}

type createPostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body,omitempty"`
}

func NewBlogPlugin(cfg Config) *BlogPlugin {
	return &BlogPlugin{
		metadata: plugin.Metadata{
			Name:        cfg.Name,
			Version:     cfg.Version,
			Protocol:    plugin.ProtocolHTTP,
			Description: "Remote Blog plugin sample",
			Capabilities: []string{
				plugin.OperationManifest,
				plugin.OperationHealth,
				plugin.OperationHooksExecute,
				operationBlogCreate,
				operationBlogList,
			},
			Labels: map[string]string{
				"example": "remote-blog",
			},
		},
		nextID: 1,
	}
}

func (p *BlogPlugin) Metadata() plugin.Metadata {
	return p.metadata
}

func (p *BlogPlugin) Invoke(ctx context.Context, req plugin.Request) (*plugin.Response, error) {
	switch req.Operation {
	case plugin.OperationManifest:
		return plugin.NewResponse(p.Metadata())
	case plugin.OperationHealth:
		return plugin.NewResponse(map[string]string{"status": "ok"})
	case operationBlogCreate:
		var input createPostRequest
		if err := req.DecodePayload(&input); err != nil {
			return nil, err
		}
		return plugin.NewResponse(p.create(input))
	case operationBlogList:
		return plugin.NewResponse(p.list())
	case plugin.OperationHooksExecute:
		var event hooks.Event
		if err := req.DecodePayload(&event); err != nil {
			return nil, err
		}
		return p.handleHook(event)
	default:
		return nil, fmt.Errorf("unsupported blog operation %q", req.Operation)
	}
}

func (p *BlogPlugin) Close(ctx context.Context) error {
	return nil
}

func (p *BlogPlugin) create(input createPostRequest) Post {
	p.mu.Lock()
	defer p.mu.Unlock()

	post := Post{
		ID:        p.nextID,
		Title:     input.Title,
		Body:      input.Body,
		CreatedAt: time.Now().UTC(),
	}
	p.nextID++
	p.posts = append(p.posts, post)
	return post
}

func (p *BlogPlugin) list() []Post {
	p.mu.Lock()
	defer p.mu.Unlock()

	out := make([]Post, len(p.posts))
	copy(out, p.posts)
	return out
}

func (p *BlogPlugin) handleHook(event hooks.Event) (*plugin.Response, error) {
	audit := HookAudit{
		Point:     string(event.Point),
		Plugin:    event.Plugin,
		Operation: event.Operation,
	}
	if event.Identity != nil {
		audit.PrincipalID = event.Identity.Principal.ID
	}

	p.mu.Lock()
	p.audits = append(p.audits, audit)
	p.mu.Unlock()

	result, err := hooks.NewResult(audit)
	if err != nil {
		return nil, err
	}
	return plugin.NewResponse(result)
}

func (p *BlogPlugin) auditsSnapshot() []HookAudit {
	p.mu.Lock()
	defer p.mu.Unlock()

	out := make([]HookAudit, len(p.audits))
	copy(out, p.audits)
	return out
}
