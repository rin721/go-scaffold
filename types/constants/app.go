package constants

import "time"

const (
	// AppDefaultConfigPath 是配置文件的默认路径
	// 当环境变量 CONFIG_PATH 未设置时使用此路径
	AppDefaultConfigPath = "configs/config.yaml"

	// AppShutdownTimeout 是优雅关闭的最大等待时间
	// 设置为 30 秒以确保所有正在处理的请求能够完成
	// 超过此时间后将强制关闭,避免进程无限期等待
	AppShutdownTimeout = 30 * time.Second

	// EnvConfigPathName 是配置文件路径的环境变量名称
	// Deprecated: 应用运行时应通过 internal/config.EnvConfigPathName 基于 AppPrefix 动态生成。
	EnvConfigPathName = "RIN_CONFIG_PATH"

	AppPrefix            = "Rin"                           // AppPrefix 是应用前缀
	AppName              = "go-scaffold"                   // AppName 是应用名称
	AppDescription       = "This is a go backend scaffold" // AppDescription 是应用描述
	AppServerCommandName = "server"                        // AppServerCommandName 是应用命令名称
	AppInitDBCommandName = "initdb"                        // AppInitDBCommandName 是初始化数据库命令名称
	AppVersion           = "0.1.2"                         // AppVersion 是应用版本号
)
