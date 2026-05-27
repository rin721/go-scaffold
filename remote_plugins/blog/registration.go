package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/rei0721/go-scaffold/pkg/plugin"
	"github.com/rei0721/go-scaffold/pkg/plugin/hooks"
)

func RegisterWithHost(ctx context.Context, cfg Config, client *http.Client) error {
	if strings.TrimSpace(cfg.HostHTTPURL) == "" {
		return nil
	}
	if client == nil {
		client = http.DefaultClient
	}

	reqBody := plugin.RegistrationRequest{
		Plugin: plugin.Definition{
			Name:         cfg.Name,
			Version:      cfg.Version,
			Protocol:     plugin.ProtocolHTTP,
			Endpoint:     cfg.InvokeEndpoint(),
			Headers:      pluginHeaders(cfg),
			Description:  "Remote Blog plugin sample",
			Capabilities: []string{operationBlogCreate, operationBlogList, plugin.OperationHooksExecute},
			Labels:       map[string]string{"example": "remote-blog"},
		},
		Hooks: hookBindings(cfg),
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	url := strings.TrimRight(cfg.HostHTTPURL, "/") + plugin.HTTPRegisterPath
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if cfg.RegistrationToken != "" {
		req.Header.Set("Authorization", "Bearer "+cfg.RegistrationToken)
		req.Header.Set("X-Plugin-Registration-Token", cfg.RegistrationToken)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		var pluginResp plugin.Response
		_ = json.NewDecoder(resp.Body).Decode(&pluginResp)
		if pluginResp.Error != "" {
			return fmt.Errorf("host registration failed: %s", pluginResp.Error)
		}
		return fmt.Errorf("host registration failed: %s", resp.Status)
	}
	return nil
}

func hookBindings(cfg Config) []plugin.HookBinding {
	bindings := make([]plugin.HookBinding, 0, len(cfg.HookPoints))
	for _, point := range cfg.HookPoints {
		point = strings.TrimSpace(point)
		if point == "" {
			continue
		}
		bindings = append(bindings, plugin.HookBinding{
			Point:    hooks.Point(point),
			Plugin:   cfg.Name,
			Name:     cfg.Name + "-" + strings.ReplaceAll(point, ".", "-"),
			Priority: 0,
		})
	}
	return bindings
}

func pluginHeaders(cfg Config) map[string]string {
	if cfg.SharedSecret == "" {
		return nil
	}
	return map[string]string{"X-Blog-Plugin-Secret": cfg.SharedSecret}
}
