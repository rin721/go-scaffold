package initapp

import (
	"context"
	"fmt"

	"github.com/rei0721/go-scaffold/internal/app/dbapp"
	"github.com/rei0721/go-scaffold/internal/config"
	demohandler "github.com/rei0721/go-scaffold/internal/modules/demo/handler"
	demorepository "github.com/rei0721/go-scaffold/internal/modules/demo/repository"
	demoservice "github.com/rei0721/go-scaffold/internal/modules/demo/service"
	userhandler "github.com/rei0721/go-scaffold/internal/modules/user/handler"
	userrepository "github.com/rei0721/go-scaffold/internal/modules/user/repository"
	userservice "github.com/rei0721/go-scaffold/internal/modules/user/service"
	authapi "github.com/rei0721/go-scaffold/pkg/auth"
	passwordcrypto "github.com/rei0721/go-scaffold/pkg/crypto"
	"github.com/rei0721/go-scaffold/pkg/database"
	"github.com/rei0721/go-scaffold/pkg/logger"
	rbacapi "github.com/rei0721/go-scaffold/pkg/rbac"
)

func NewModules(core Core, infra Infrastructure) (Modules, error) {
	var demoModule DemoModule
	if core.Config.Demo.EnabledValue() {
		if core.Config.Demo.ApplySchemaOnStartValue() {
			if _, err := ApplyDemoSchemaForTrigger(infra.Database, core.Config.Database.Driver, core.Logger, DemoSchemaTriggerServerStart); err != nil {
				return Modules{}, err
			}
		}
		demoModule = NewDemoModule(infra.Database, core.Logger)
	} else if core.Logger != nil {
		core.Logger.Info("demo module disabled")
	}
	if err := ApplyUserSchema(infra.Database, core.Config.Database.Driver, core.Logger); err != nil {
		return Modules{}, err
	}
	userModule, err := NewUserModule(infra.Database, core.Logger, core.Config.Auth, core.Config.RBAC)
	if err != nil {
		return Modules{}, err
	}
	if err := ApplyConfiguredRBAC(context.Background(), userModule.Service, core.Config.RBAC, core.Logger); err != nil {
		return Modules{}, err
	}
	return Modules{
		Demo: demoModule,
		User: userModule,
	}, nil
}

type DemoSchemaTrigger string

const (
	DemoSchemaTriggerServerStart DemoSchemaTrigger = "server-start"
	DemoSchemaTriggerReload      DemoSchemaTrigger = "reload"
)

type DemoSchemaPolicy struct {
	Trigger DemoSchemaTrigger
	Apply   bool
	Reason  string
}

func DemoSchemaPolicyFor(trigger DemoSchemaTrigger) DemoSchemaPolicy {
	switch trigger {
	case DemoSchemaTriggerServerStart:
		return DemoSchemaPolicy{
			Trigger: trigger,
			Apply:   true,
			Reason:  "demo server startup keeps the local development schema ready through sqlgen",
		}
	case DemoSchemaTriggerReload:
		return DemoSchemaPolicy{
			Trigger: trigger,
			Apply:   false,
			Reason:  "database reload must not perform implicit schema changes",
		}
	default:
		return DemoSchemaPolicy{
			Trigger: trigger,
			Apply:   false,
			Reason:  "unknown schema trigger requires an explicit policy",
		}
	}
}

func ApplyDemoSchema(db database.Database, driver string, log logger.Logger) error {
	_, err := ApplyDemoSchemaForTrigger(db, driver, log, DemoSchemaTriggerServerStart)
	return err
}

func ApplyDemoSchemaForTrigger(db database.Database, driver string, log logger.Logger, trigger DemoSchemaTrigger) (DemoSchemaPolicy, error) {
	policy := DemoSchemaPolicyFor(trigger)
	if !policy.Apply {
		logDemoSchemaSkipped(log, policy)
		return policy, nil
	}
	if db == nil {
		return policy, nil
	}
	if _, err := dbapp.ApplyDemoSchema(context.Background(), db, driver); err != nil {
		return policy, fmt.Errorf("apply demo schema: %w", err)
	}
	logDemoSchemaApplied(log, policy)
	return policy, nil
}

func logDemoSchemaApplied(log logger.Logger, policy DemoSchemaPolicy) {
	if log == nil {
		return
	}
	log.Info("demo schema applied", "trigger", policy.Trigger, "reason", policy.Reason)
}

func logDemoSchemaSkipped(log logger.Logger, policy DemoSchemaPolicy) {
	if log == nil {
		return
	}
	log.Debug("demo schema apply skipped", "trigger", policy.Trigger, "reason", policy.Reason)
}

func ApplyUserSchema(db database.Database, driver string, log logger.Logger) error {
	if db == nil {
		return nil
	}
	if _, err := dbapp.ApplyUserSchema(context.Background(), db, driver); err != nil {
		return fmt.Errorf("apply user schema: %w", err)
	}
	if log != nil {
		log.Info("user schema applied", "trigger", DemoSchemaTriggerServerStart)
	}
	return nil
}

func NewDemoModule(db database.Database, log logger.Logger) DemoModule {
	todoRepo := demorepository.NewTodoRepository()
	todoService := demoservice.NewTodoService(db, todoRepo)
	todoHandler := demohandler.NewTodoHandler(todoService, log)

	return DemoModule{
		TodoRepository: todoRepo,
		TodoService:    todoService,
		TodoHandler:    todoHandler,
	}
}

func NewUserModule(db database.Database, log logger.Logger, authCfg config.AuthConfig, rbacCfg config.RBACConfig) (UserModule, error) {
	repo := userrepository.NewRepository()
	hasher, err := passwordcrypto.NewBcrypt()
	if err != nil {
		return UserModule{}, fmt.Errorf("create password hasher: %w", err)
	}
	tokens, err := NewAuthTokenService(authCfg)
	if err != nil {
		return UserModule{}, fmt.Errorf("create auth token service: %w", err)
	}
	authorizer, err := rbacapi.NewCasbinAuthorizer(rbacCfg.ModelPath)
	if err != nil {
		return UserModule{}, fmt.Errorf("create rbac authorizer: %w", err)
	}
	userService := userservice.NewUserService(db, repo, hasher, tokens, authorizer)
	userHandler := userhandler.NewUserHandler(userService, log,
		userhandler.WithPublicRegistration(authCfg.PublicRegistrationEnabled()))
	return UserModule{
		Repository: repo,
		Service:    userService,
		Handler:    userHandler,
		Tokens:     tokens,
	}, nil
}

func NewAuthTokenService(authCfg config.AuthConfig) (authapi.TokenService, error) {
	secret := authCfg.TokenSecretBytes()
	var err error
	if len(secret) == 0 {
		secret, err = authapi.GenerateSecret(authapi.MinTokenSecretBytes)
		if err != nil {
			return nil, err
		}
	}
	return authapi.NewJWTService(authapi.JWTConfig{
		Secret: secret,
		TTL:    authCfg.TokenTTLDuration(),
		Issuer: "go-scaffold",
	})
}

func ApplyConfiguredRBAC(ctx context.Context, service userservice.UserService, cfg config.RBACConfig, log logger.Logger) error {
	if !cfg.Enabled || !cfg.ApplyOnStart {
		return nil
	}
	if service == nil {
		return fmt.Errorf("apply rbac config: user service is nil")
	}
	seed := userservice.RBACSeed{
		Roles:           make([]userservice.RBACRoleSeed, 0, len(cfg.Roles)),
		Permissions:     make([]userservice.RBACPermissionSeed, 0, len(cfg.Permissions)),
		RolePermissions: make([]userservice.RBACRolePermissionSeed, 0, len(cfg.RolePermissions)),
	}
	for _, role := range cfg.Roles {
		seed.Roles = append(seed.Roles, userservice.RBACRoleSeed{
			Name:        role.Name,
			Description: role.Description,
		})
	}
	for _, permission := range cfg.Permissions {
		seed.Permissions = append(seed.Permissions, userservice.RBACPermissionSeed{
			Code:        permission.Code,
			Description: permission.Description,
		})
	}
	for _, grant := range cfg.RolePermissions {
		seed.RolePermissions = append(seed.RolePermissions, userservice.RBACRolePermissionSeed{
			Role:        grant.Role,
			Permissions: append([]string(nil), grant.Permissions...),
		})
	}
	if err := service.ApplyRBACSeed(ctx, seed); err != nil {
		return fmt.Errorf("apply rbac config: %w", err)
	}
	if log != nil {
		log.Info("rbac config seed applied",
			"roles", len(seed.Roles),
			"permissions", len(seed.Permissions),
			"role_permissions", len(seed.RolePermissions))
	}
	return nil
}
