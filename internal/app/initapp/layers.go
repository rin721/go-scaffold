package initapp

import (
	"github.com/gin-gonic/gin"
	"github.com/rei0721/go-scaffold/internal/config"
	demohandler "github.com/rei0721/go-scaffold/internal/modules/demo/handler"
	demorepository "github.com/rei0721/go-scaffold/internal/modules/demo/repository"
	demoservice "github.com/rei0721/go-scaffold/internal/modules/demo/service"
	"github.com/rei0721/go-scaffold/pkg/cache"
	"github.com/rei0721/go-scaffold/pkg/database"
	"github.com/rei0721/go-scaffold/pkg/executor"
	"github.com/rei0721/go-scaffold/pkg/httpserver"
	"github.com/rei0721/go-scaffold/pkg/i18n"
	"github.com/rei0721/go-scaffold/pkg/iam"
	"github.com/rei0721/go-scaffold/pkg/logger"
	"github.com/rei0721/go-scaffold/pkg/plugin"
	"github.com/rei0721/go-scaffold/pkg/storage"
	"github.com/rei0721/go-scaffold/pkg/utils"
)

type Core struct {
	Config        *config.Config
	ConfigManager config.Manager
	Logger        logger.Logger
	I18n          i18n.I18n
	I18nUtils     *utils.I18nUtils
	IDGenerator   utils.IDGenerator
}

type Infrastructure struct {
	Database database.Database
	Cache    cache.Cache
	Executor executor.Manager
	Storage  storage.Storage
	IAM      iam.Service
	Plugins  plugin.Manager
}

type Modules struct {
	Demo DemoModule
}

type DemoModule struct {
	TodoRepository demorepository.TodoRepository
	TodoService    demoservice.TodoService
	TodoHandler    *demohandler.TodoHandler
}

type Transport struct {
	Router           *gin.Engine
	HTTPServer       httpserver.HTTPServer
	PluginHTTPServer httpserver.HTTPServer
}
