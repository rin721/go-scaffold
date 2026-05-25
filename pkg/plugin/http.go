package plugin

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type httpPlugin struct {
	metadata         Metadata
	endpoint         string
	headers          map[string]string
	client           *http.Client
	maxResponseBytes int64
}

func newHTTP(def Definition, client *http.Client, defaultTimeout time.Duration, maxResponseBytes int64) (Plugin, error) {
	if err := def.Validate(); err != nil {
		return nil, err
	}
	if client == nil {
		client = http.DefaultClient
	}
	httpClient := *client
	if httpClient.Timeout == 0 {
		httpClient.Timeout = def.timeout(defaultTimeout)
	}
	return &httpPlugin{
		metadata:         def.metadata(),
		endpoint:         def.Endpoint,
		headers:          copyStringMap(def.Headers),
		client:           &httpClient,
		maxResponseBytes: maxResponseBytes,
	}, nil
}

func (p *httpPlugin) Metadata() Metadata {
	return p.metadata
}

func (p *httpPlugin) Invoke(ctx context.Context, req Request) (*Response, error) {
	if err := req.Validate(); err != nil {
		return nil, &Error{Op: "invoke", Plugin: p.metadata.Name, Err: err}
	}
	if req.Plugin == "" {
		req.Plugin = p.metadata.Name
	}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, &Error{Op: "marshal", Plugin: p.metadata.Name, Err: err}
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, p.endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, &Error{Op: "request", Plugin: p.metadata.Name, Err: err}
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")
	for k, v := range p.headers {
		httpReq.Header.Set(k, v)
	}
	for k, v := range req.Headers {
		httpReq.Header.Set(k, v)
	}

	httpResp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, &Error{Op: "http", Plugin: p.metadata.Name, Err: err}
	}
	defer httpResp.Body.Close()

	reader := io.Reader(httpResp.Body)
	if p.maxResponseBytes > 0 {
		reader = io.LimitReader(httpResp.Body, p.maxResponseBytes)
	}
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, &Error{Op: "read", Plugin: p.metadata.Name, Err: err}
	}

	resp := &Response{}
	if len(data) > 0 {
		if err := json.Unmarshal(data, resp); err != nil {
			return nil, &Error{Op: "decode", Plugin: p.metadata.Name, Err: err}
		}
	}

	if httpResp.StatusCode < http.StatusOK || httpResp.StatusCode >= http.StatusMultipleChoices {
		msg := http.StatusText(httpResp.StatusCode)
		if resp.Error != "" {
			msg = resp.Error
		}
		return resp, &Error{
			Op:     "http",
			Plugin: p.metadata.Name,
			Err:    fmt.Errorf("%w: %d %s", ErrHTTPStatus, httpResp.StatusCode, msg),
		}
	}
	if resp.Error != "" {
		return resp, &Error{Op: "invoke", Plugin: p.metadata.Name, Err: errors.New(resp.Error)}
	}
	return resp, nil
}

func (p *httpPlugin) Close(ctx context.Context) error {
	return nil
}
