package main

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rei0721/go-scaffold/pkg/plugin"
)

type Config struct {
	Name              string
	Version           string
	ListenAddr        string
	PublicHTTPURL     string
	HostHTTPURL       string
	HostWSURL         string
	RegistrationToken string
	SharedSecret      string
	RegisterTimeout   time.Duration
	HookPoints        []string
}

func LoadConfigFromEnv() Config {
	hostHTTPURL := envString("BLOG_PLUGIN_HOST_HTTP_URL", envString("BLOG_PLUGIN_MAIN_HTTP_URL", ""))
	hostWSURL := envString("BLOG_PLUGIN_HOST_WS_URL", envString("BLOG_PLUGIN_MAIN_WS_URL", ""))
	return Config{
		Name:              envString("BLOG_PLUGIN_NAME", "blog"),
		Version:           envString("BLOG_PLUGIN_VERSION", "0.1.0"),
		ListenAddr:        envString("BLOG_PLUGIN_LISTEN_ADDR", "127.0.0.1:18081"),
		PublicHTTPURL:     envString("BLOG_PLUGIN_PUBLIC_HTTP_URL", "http://127.0.0.1:18081"),
		HostHTTPURL:       hostHTTPURL,
		HostWSURL:         hostWSURL,
		RegistrationToken: envString("BLOG_PLUGIN_REGISTRATION_TOKEN", ""),
		SharedSecret:      envString("BLOG_PLUGIN_SHARED_SECRET", ""),
		RegisterTimeout:   envDuration("BLOG_PLUGIN_REGISTER_TIMEOUT_SECONDS", 10*time.Second),
		HookPoints:        envList("BLOG_PLUGIN_HOOK_POINTS", []string{string(plugin.HookAfterInvoke)}),
	}
}

func (c Config) InvokeEndpoint() string {
	return strings.TrimRight(c.PublicHTTPURL, "/") + plugin.HTTPInvokePath
}

func envString(name, fallback string) string {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return fallback
	}
	return value
}

func envDuration(name string, fallback time.Duration) time.Duration {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return fallback
	}
	seconds, err := strconv.Atoi(value)
	if err != nil || seconds <= 0 {
		return fallback
	}
	return time.Duration(seconds) * time.Second
}

func envList(name string, fallback []string) []string {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return append([]string(nil), fallback...)
	}
	parts := strings.Split(value, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			out = append(out, part)
		}
	}
	if len(out) == 0 {
		return append([]string(nil), fallback...)
	}
	return out
}
