package initapp

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rei0721/go-scaffold/internal/config"
	"github.com/rei0721/go-scaffold/internal/middleware"
	demohandler "github.com/rei0721/go-scaffold/internal/modules/demo/handler"
	userhandler "github.com/rei0721/go-scaffold/internal/modules/user/handler"
	httptransport "github.com/rei0721/go-scaffold/internal/transport/http"
	"github.com/rei0721/go-scaffold/pkg/database"
	"github.com/rei0721/go-scaffold/pkg/httpserver"
	"github.com/rei0721/go-scaffold/pkg/i18n"
	"github.com/rei0721/go-scaffold/pkg/logger"
	"github.com/rei0721/go-scaffold/pkg/plugin"
)

func NewTransport(core Core, infra Infrastructure, modules Modules) (Transport, error) {
	corsConfig, err := NewCORS(core.Config, core.Logger)
	if err != nil {
		return Transport{}, err
	}

	router, server, err := NewHTTPServer(
		core.Config,
		core.Logger,
		core.I18n,
		infra.Database,
		corsConfig,
		infra.Plugins,
		modules.Demo.TodoHandler,
		modules.User.Handler,
	)
	if err != nil {
		return Transport{}, err
	}

	pluginHTTPServer, err := NewPluginHTTPServer(core.Config, core.Logger, infra.Plugins)
	if err != nil {
		return Transport{}, err
	}

	return Transport{
		Router:           router,
		HTTPServer:       server,
		PluginHTTPServer: pluginHTTPServer,
	}, nil
}

func NewCORS(cfg *config.Config, log logger.Logger) (middleware.CORSConfig, error) {
	corsCfg := cfg.CORS
	corsCfg.DefaultConfig()
	corsCfg.OverrideConfig()

	if err := corsCfg.Validate(); err != nil {
		return middleware.CORSConfig{}, err
	}

	if corsCfg.Enabled {
		log.Info(
			"CORS middleware enabled",
			"allow_origins", corsCfg.AllowOrigins,
			"allow_credentials", corsCfg.AllowCredentials,
			"max_age", corsCfg.MaxAge,
		)
	} else {
		log.Info("CORS middleware disabled")
	}

	return middleware.CORSConfig{
		Enabled:          corsCfg.Enabled,
		AllowOrigins:     corsCfg.AllowOrigins,
		AllowMethods:     corsCfg.AllowMethods,
		AllowHeaders:     corsCfg.AllowHeaders,
		ExposeHeaders:    corsCfg.ExposeHeaders,
		AllowCredentials: corsCfg.AllowCredentials,
		MaxAge:           corsCfg.MaxAge,
	}, nil
}

func NewHTTPServer(
	cfg *config.Config,
	log logger.Logger,
	i18nApp i18n.I18n,
	db database.Database,
	corsConfig middleware.CORSConfig,
	plugins plugin.Manager,
	todoHandler *demohandler.TodoHandler,
	userHandler *userhandler.UserHandler,
) (*gin.Engine, httpserver.HTTPServer, error) {
	gin.SetMode(cfg.Server.Mode)

	middlewareCfg := middleware.DefaultMiddlewareConfig()
	middlewareCfg.CORS = corsConfig

	router := httptransport.NewRouter(httptransport.RouterDeps{
		Logger:      log,
		I18n:        i18nApp,
		Database:    db,
		Middleware:  middlewareCfg,
		TodoHandler: todoHandler,
		UserHandler: userHandler,
		PluginRegistration: NewMainPluginRegistrationHandler(
			cfg,
			plugins,
		),
	})

	server, err := httpserver.New(router, HTTPServerConfig(cfg), log)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create http server: %w", err)
	}

	return router, server, nil
}

func NewPluginHTTPServer(cfg *config.Config, log logger.Logger, plugins plugin.Manager) (httpserver.HTTPServer, error) {
	if cfg == nil || !cfg.Plugin.Interface.HTTP.Enabled {
		return nil, nil
	}
	handler := NewPluginRegistrationHandler(cfg, plugins)
	if handler == nil {
		return nil, fmt.Errorf("plugin http interface requires plugin registration handler")
	}
	server, err := httpserver.New(httptransport.NewPluginRouter(handler), PluginHTTPServerConfig(cfg), log)
	if err != nil {
		return nil, fmt.Errorf("failed to create plugin http server: %w", err)
	}
	return server, nil
}

func NewMainPluginRegistrationHandler(cfg *config.Config, plugins plugin.Manager) http.Handler {
	if cfg == nil || !cfg.Plugin.Registration.ExposeOnMainHTTP {
		return nil
	}
	return NewPluginRegistrationHandler(cfg, plugins)
}

func NewPluginRegistrationHandler(cfg *config.Config, plugins plugin.Manager) http.Handler {
	if cfg == nil || plugins == nil || !cfg.Plugin.Enabled || !cfg.Plugin.Registration.Enabled {
		return nil
	}
	return plugin.NewHTTPRegistrationHandler(
		plugins,
		plugin.WithRegistrationToken(cfg.Plugin.Registration.Token),
		plugin.WithRegistrationHTTPOptions(pluginHTTPOptions(cfg)...),
	)
}
