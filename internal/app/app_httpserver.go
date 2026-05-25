package app

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rei0721/go-scaffold/internal/middleware"
	demohandler "github.com/rei0721/go-scaffold/internal/modules/demo/handler"
	demorepository "github.com/rei0721/go-scaffold/internal/modules/demo/repository"
	demoservice "github.com/rei0721/go-scaffold/internal/modules/demo/service"
	httptransport "github.com/rei0721/go-scaffold/internal/transport/http"
	"github.com/rei0721/go-scaffold/pkg/httpserver"
)

// initHTTPServer 初始化 HTTP 服务器
// 使用 pkg/httpserver 封装，替代原来的直接创建 http.Server
// 这个函数应该在 Router 初始化之后调用
func (app *App) initHTTPServer() error {
	gin.SetMode(app.Config.Server.Mode)

	todoRepo := demorepository.NewTodoRepository()
	todoService := demoservice.NewTodoService(app.DB, todoRepo)
	todoHandler := demohandler.NewTodoHandler(todoService, app.Logger)

	middlewareCfg := middleware.DefaultMiddlewareConfig()
	middlewareCfg.CORS = app.getCORSMiddlewareConfig()
	router := httptransport.NewRouter(httptransport.RouterDeps{
		Logger:      app.Logger,
		I18n:        app.I18n,
		Database:    app.DB,
		Middleware:  middlewareCfg,
		TodoHandler: todoHandler,
	})

	// 创建 HTTP 服务器配置
	cfg := &httpserver.Config{
		Host:         app.Config.Server.Host,
		Port:         app.Config.Server.Port,
		ReadTimeout:  time.Duration(app.Config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(app.Config.Server.WriteTimeout) * time.Second,
	}

	server, err := httpserver.New(router, cfg, app.Logger)
	if err != nil {
		return fmt.Errorf("failed to create http server: %w", err)
	}

	app.HTTPServer = server
	app.Router = router

	return nil
}
