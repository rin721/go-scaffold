package constants

import "time"

const (
	AppDefaultConfigPath = "configs/config.yaml"
	AppShutdownTimeout   = 30 * time.Second
	EnvConfigPathName    = "RIN_CONFIG_PATH"

	AppPrefix            = "Rin"
	AppName              = "go-scaffold"
	AppDescription       = "This is a go backend scaffold"
	AppServerCommandName = "server"
	AppVersion           = "0.1.2"
)
