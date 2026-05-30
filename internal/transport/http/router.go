package httptransport

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

type RouterDeps struct {
	Logger      logger.Logger
	I18n        i18n.I18n
	Database    database.Database
	Middleware  middleware.MiddlewareConfig
	TodoHandler *demohandler.TodoHandler
}

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

func health(c *gin.Context) {
	c.JSON(http.StatusOK, result.Success(gin.H{"status": "ok"}))
}

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

func ReadyCheck(db database.Database) func(context.Context) error {
	return func(ctx context.Context) error {
		if db == nil {
			return http.ErrServerClosed
		}
		return db.Ping(ctx)
	}
}
