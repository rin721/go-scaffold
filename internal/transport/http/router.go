package httptransport

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rei0721/go-scaffold/internal/middleware"
	demohandler "github.com/rei0721/go-scaffold/internal/modules/demo/handler"
	userhandler "github.com/rei0721/go-scaffold/internal/modules/user/handler"
	userservice "github.com/rei0721/go-scaffold/internal/modules/user/service"
	"github.com/rei0721/go-scaffold/pkg/database"
	"github.com/rei0721/go-scaffold/pkg/i18n"
	"github.com/rei0721/go-scaffold/pkg/logger"
	"github.com/rei0721/go-scaffold/pkg/plugin"
	"github.com/rei0721/go-scaffold/types/result"
)

type RouterDeps struct {
	Logger             logger.Logger
	I18n               i18n.I18n
	Database           database.Database
	Middleware         middleware.MiddlewareConfig
	TodoHandler        *demohandler.TodoHandler
	UserHandler        *userhandler.UserHandler
	PluginRegistration http.Handler
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
	if deps.PluginRegistration != nil {
		r.Any(plugin.HTTPRegisterPath, gin.WrapH(deps.PluginRegistration))
	}

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
	if deps.UserHandler != nil {
		auth := v1.Group("/auth")
		auth.POST("/register", deps.UserHandler.Register)
		auth.POST("/login", deps.UserHandler.Login)
		auth.GET("/me", deps.UserHandler.Authenticate(), deps.UserHandler.Me)

		users := v1.Group("/users", deps.UserHandler.Authenticate())
		users.POST("", deps.UserHandler.RequirePermission(userservice.PermissionUsersCreate), deps.UserHandler.CreateUser)
		users.GET("", deps.UserHandler.RequirePermission(userservice.PermissionUsersRead), deps.UserHandler.ListUsers)
		users.GET("/:id", deps.UserHandler.RequirePermission(userservice.PermissionUsersRead), deps.UserHandler.GetUser)
		users.PUT("/:id", deps.UserHandler.RequirePermission(userservice.PermissionUsersUpdate), deps.UserHandler.UpdateUser)
		users.DELETE("/:id", deps.UserHandler.RequirePermission(userservice.PermissionUsersDelete), deps.UserHandler.DeleteUser)
		users.POST("/:id/roles/:roleID", deps.UserHandler.RequirePermission(userservice.PermissionUsersAssignRoles), deps.UserHandler.AssignRoleToUser)
		users.DELETE("/:id/roles/:roleID", deps.UserHandler.RequirePermission(userservice.PermissionUsersAssignRoles), deps.UserHandler.RemoveRoleFromUser)

		roles := v1.Group("/roles", deps.UserHandler.Authenticate())
		roles.POST("", deps.UserHandler.RequirePermission(userservice.PermissionRolesCreate), deps.UserHandler.CreateRole)
		roles.GET("", deps.UserHandler.RequirePermission(userservice.PermissionRolesRead), deps.UserHandler.ListRoles)
		roles.GET("/:id", deps.UserHandler.RequirePermission(userservice.PermissionRolesRead), deps.UserHandler.GetRole)
		roles.PUT("/:id", deps.UserHandler.RequirePermission(userservice.PermissionRolesUpdate), deps.UserHandler.UpdateRole)
		roles.DELETE("/:id", deps.UserHandler.RequirePermission(userservice.PermissionRolesDelete), deps.UserHandler.DeleteRole)
		roles.POST("/:id/permissions/:permissionID", deps.UserHandler.RequirePermission(userservice.PermissionRolesAssignPermission), deps.UserHandler.AssignPermissionToRole)
		roles.DELETE("/:id/permissions/:permissionID", deps.UserHandler.RequirePermission(userservice.PermissionRolesAssignPermission), deps.UserHandler.RemovePermissionFromRole)

		permissions := v1.Group("/permissions", deps.UserHandler.Authenticate())
		permissions.POST("", deps.UserHandler.RequirePermission(userservice.PermissionPermissionsCreate), deps.UserHandler.CreatePermission)
		permissions.GET("", deps.UserHandler.RequirePermission(userservice.PermissionPermissionsRead), deps.UserHandler.ListPermissions)
		permissions.GET("/:id", deps.UserHandler.RequirePermission(userservice.PermissionPermissionsRead), deps.UserHandler.GetPermission)
		permissions.PUT("/:id", deps.UserHandler.RequirePermission(userservice.PermissionPermissionsUpdate), deps.UserHandler.UpdatePermission)
		permissions.DELETE("/:id", deps.UserHandler.RequirePermission(userservice.PermissionPermissionsDelete), deps.UserHandler.DeletePermission)
	}

	return r
}

func NewPluginRouter(registration http.Handler) http.Handler {
	mux := http.NewServeMux()
	if registration != nil {
		mux.Handle(plugin.HTTPRegisterPath, registration)
	}
	return mux
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, result.Success(gin.H{"status": "ok"}))
}

func ready(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		if db == nil {
			c.JSON(http.StatusServiceUnavailable, result.Success(gin.H{
				"status": "not_ready",
				"checks": gin.H{"database": "missing"},
			}))
			return
		}
		if err := db.Ping(c.Request.Context()); err != nil {
			c.JSON(http.StatusServiceUnavailable, result.Success(gin.H{
				"status": "not_ready",
				"checks": gin.H{"database": err.Error()},
			}))
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
