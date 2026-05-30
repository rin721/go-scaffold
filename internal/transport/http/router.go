package httptransport

// 本文件定义 HTTP 传输层装配，把中间件顺序、健康检查和业务路由注册为 Gin Engine。

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rei0721/go-scaffold/internal/middleware"
	demohandler "github.com/rei0721/go-scaffold/internal/modules/demo/handler"
	"github.com/rei0721/go-scaffold/pkg/database"
	"github.com/rei0721/go-scaffold/pkg/i18n"
	"github.com/rei0721/go-scaffold/pkg/logger"
	apperrors "github.com/rei0721/go-scaffold/types/errors"
	"github.com/rei0721/go-scaffold/types/result"
)

// RouterDeps 聚合 HTTP 路由装配所需依赖，允许测试或可选模块传入 nil 以裁剪路由。
type RouterDeps struct {
	Logger      logger.Logger
	I18n        i18n.I18n
	Database    database.Database
	Middleware  middleware.MiddlewareConfig
	TodoHandler *demohandler.TodoHandler
}

// NewRouter 按固定顺序注册中间件、健康检查和业务路由，返回可直接交给 HTTPServer 的 Gin Engine。
func NewRouter(deps RouterDeps) *gin.Engine {
	r := gin.New()

	if deps.I18n != nil {
		r.Use(middleware.I18n(deps.I18n))
	}
	r.Use(middleware.TraceID(deps.Middleware.TraceID))
	r.Use(middleware.CORSMiddleware(deps.Middleware.CORS))
	if deps.Logger != nil {
		r.Use(middleware.Logger(deps.Middleware.Logger, deps.Logger))
		r.Use(middleware.Recovery(deps.Middleware.Recovery, deps.Logger))
	} else {
		r.Use(gin.Recovery())
	}

	r.GET("/health", health)
	r.GET("/ready", ready(deps.Database))

	v1 := r.Group("/api/v1")
	demo := v1.Group("/demo")
	if deps.TodoHandler != nil {
		todos := demo.Group("/todos")
		todos.POST("", deps.TodoHandler.Create)
		todos.GET("", deps.TodoHandler.List)
		todos.GET("/:id", deps.TodoHandler.Get)
		todos.PUT("/:id", deps.TodoHandler.Update)
		todos.DELETE("/:id", deps.TodoHandler.Delete)
	}

	return r
}

// health 返回轻量存活探针响应，只证明进程与路由栈仍可处理请求。
func health(c *gin.Context) {
	c.JSON(http.StatusOK, result.Success(gin.H{"status": "ok"}))
}

// ready 执行数据库就绪检查，并把失败原因转化为 503 响应。
func ready(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		if db == nil {
			c.JSON(http.StatusServiceUnavailable, &result.Result[gin.H]{
				Code:    apperrors.ErrDatabaseError,
				Message: "not ready",
				Data: gin.H{
					"status": "not_ready",
					"checks": gin.H{"database": "missing"},
				},
				ServerTime: time.Now().Unix(),
			})
			return
		}
		if err := db.Ping(c.Request.Context()); err != nil {
			c.JSON(http.StatusServiceUnavailable, &result.Result[gin.H]{
				Code:    apperrors.ErrDatabaseError,
				Message: "not ready",
				Data: gin.H{
					"status": "not_ready",
					"checks": gin.H{"database": err.Error()},
				},
				ServerTime: time.Now().Unix(),
			})
			return
		}
		c.JSON(http.StatusOK, result.Success(gin.H{
			"status": "ready",
			"checks": gin.H{"database": "ok"},
		}))
	}
}

// ReadyCheck 构造就绪探针回调，通过数据库健康检查表达服务是否可以承接流量。
func ReadyCheck(db database.Database) func(context.Context) error {
	return func(ctx context.Context) error {
		if db == nil {
			return http.ErrServerClosed
		}
		return db.Ping(ctx)
	}
}
