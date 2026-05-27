package initapp

import (
	"context"
	"fmt"
	"time"

	"github.com/rei0721/go-scaffold/internal/config"
	"github.com/rei0721/go-scaffold/pkg/iam"
	"github.com/rei0721/go-scaffold/pkg/logger"
	"github.com/rei0721/go-scaffold/pkg/plugin"
	"github.com/rei0721/go-scaffold/pkg/plugin/hooks"
)

func NewPluginManager(cfg *config.Config, log logger.Logger, iamService iam.Service) (plugin.Manager, error) {
	if !cfg.Plugin.Enabled {
		if log != nil {
			log.Info("plugin manager disabled")
		}
		return nil, nil
	}

	mgr := plugin.NewManager()
	if err := RegisterPluginLogHooks(mgr, log); err != nil {
		return nil, err
	}
	if err := RegisterIAMHooks(mgr, iamService); err != nil {
		return nil, err
	}

	for _, def := range cfg.Plugin.Plugins {
		remote, err := plugin.NewHTTP(PluginDefinition(def), pluginHTTPOptions(cfg)...)
		if err != nil {
			return nil, fmt.Errorf("create plugin %q: %w", def.Name, err)
		}
		if err := mgr.Register(remote); err != nil {
			return nil, err
		}
	}
	for _, binding := range cfg.Plugin.Hooks {
		if err := mgr.RegisterHook(
			hooks.Point(binding.Point),
			plugin.NewRemoteHook(mgr, binding.Plugin),
			hooks.WithName(binding.Name),
			hooks.WithPriority(binding.Priority),
		); err != nil {
			return nil, err
		}
	}

	if log != nil {
		log.Info("plugin manager initialized", "plugins", len(cfg.Plugin.Plugins), "hooks", len(cfg.Plugin.Hooks))
	}
	return mgr, nil
}

func PluginDefinition(def config.PluginDefinitionConfig) plugin.Definition {
	return plugin.Definition{
		Name:         def.Name,
		Version:      def.Version,
		Protocol:     plugin.Protocol(def.Protocol),
		Endpoint:     def.Endpoint,
		Timeout:      time.Duration(def.Timeout) * time.Second,
		Headers:      copyStringMap(def.Headers),
		Description:  def.Description,
		Capabilities: append([]string(nil), def.Capabilities...),
		Labels:       copyStringMap(def.Labels),
	}
}

func RegisterPluginLogHooks(mgr plugin.Manager, log logger.Logger) error {
	if mgr == nil || log == nil {
		return nil
	}
	return mgr.RegisterHook(plugin.HookAfterInvoke, hooks.HandlerFunc(func(ctx context.Context, event hooks.Event) (hooks.Result, error) {
		log.Debug("plugin invoked", "plugin", event.Plugin, "operation", event.Operation)
		return hooks.Result{}, nil
	}), hooks.WithName("app.plugin.audit"), hooks.WithPriority(-100))
}

func RegisterIAMHooks(mgr plugin.Manager, iamService iam.Service) error {
	if mgr == nil || iamService == nil {
		return nil
	}
	return mgr.RegisterHook(plugin.HookIAMAuthorize, hooks.HandlerFunc(func(ctx context.Context, event hooks.Event) (hooks.Result, error) {
		principal, err := iam.RequirePrincipal(ctx)
		if err != nil {
			return hooks.Result{}, err
		}
		resource := iam.Resource(event.Plugin)
		if resource == "" {
			resource = iam.Resource("*")
		}
		_, err = iamService.Authorize(ctx, principal, iam.Action(event.Operation), resource)
		if err != nil {
			return hooks.Stop(err.Error()), err
		}
		return hooks.Result{}, nil
	}), hooks.WithName("app.iam.authorize"), hooks.WithPriority(1000))
}

func pluginHTTPOptions(cfg *config.Config) []plugin.HTTPOption {
	var opts []plugin.HTTPOption
	if timeout := cfg.Plugin.DefaultTimeoutDuration(); timeout > 0 {
		opts = append(opts, plugin.WithHTTPDefaultTimeout(timeout))
	}
	if cfg.Plugin.MaxResponseBytes > 0 {
		opts = append(opts, plugin.WithHTTPMaxResponseBytes(cfg.Plugin.MaxResponseBytes))
	}
	return opts
}

func copyStringMap(src map[string]string) map[string]string {
	if len(src) == 0 {
		return nil
	}
	dst := make(map[string]string, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
