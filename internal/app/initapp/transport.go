package initapp

import (
	"fmt"

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
		modules.Demo.TodoHandler,
		modules.User.Handler,
	)
	if err != nil {
		return Transport{}, err
	}

	return Transport{
		Router:     router,
		HTTPServer: server,
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
	})

	server, err := httpserver.New(router, HTTPServerConfig(cfg), log)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create http server: %w", err)
	}

	return router, server, nil
}
