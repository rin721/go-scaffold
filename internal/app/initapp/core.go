package initapp

import (
	"fmt"

	"github.com/rei0721/go-scaffold/internal/config"
	"github.com/rei0721/go-scaffold/pkg/i18n"
	"github.com/rei0721/go-scaffold/pkg/logger"
	"github.com/rei0721/go-scaffold/pkg/utils"
)

func NewCore(configPath string) (Core, error) {
	configManager, cfg, err := LoadConfig(configPath)
	if err != nil {
		return Core{}, err
	}

	log, err := NewLogger(cfg)
	if err != nil {
		return Core{}, err
	}

	i18nApp, i18nUtils, err := NewI18n(cfg)
	if err != nil {
		return Core{}, err
	}

	return Core{
		Config:        cfg,
		ConfigManager: configManager,
		Logger:        log,
		I18n:          i18nApp,
		I18nUtils:     i18nUtils,
		IDGenerator:   NewIDGenerator(),
	}, nil
}

func LoadConfig(configPath string) (config.Manager, *config.Config, error) {
	configManager := config.NewManager()
	if err := configManager.Load(configPath); err != nil {
		return nil, nil, fmt.Errorf("failed to load config: %w", err)
	}
	return configManager, configManager.Get(), nil
}

func NewLogger(cfg *config.Config) (logger.Logger, error) {
	log, err := logger.New(LoggerConfig(cfg))
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}
	log.Info("logger initialized successfully")
	return log, nil
}

func NewI18n(cfg *config.Config) (i18n.I18n, *utils.I18nUtils, error) {
	i18nCfg := &i18n.Config{
		DefaultLanguage:    cfg.I18n.Default,
		SupportedLanguages: cfg.I18n.Supported,
		MessagesDir:        cfg.I18n.MessagesDir,
	}

	i18nApp, err := i18n.New(i18nCfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create i18n: %w", err)
	}
	return i18nApp, utils.NewI18nUtils(i18nApp, cfg.I18n.Default), nil
}

func NewIDGenerator() utils.IDGenerator {
	return utils.DefaultSnowflake()
}

func DebugConfig(core Core, configPath string) {
	core.Logger.Debug("configuration loaded successfully",
		"config_file", configPath,
		"env_support", "enabled")

	core.Logger.Debug("server configuration",
		"port", core.Config.Server.Port,
		"mode", core.Config.Server.Mode)

	core.Logger.Debug("database configuration",
		"driver", core.Config.Database.Driver,
		"host", core.Config.Database.Host,
		"db", core.Config.Database.DBName)

	if core.Config.Redis.Enabled {
		core.Logger.Debug("redis configuration",
			"enabled", true,
			"host", core.Config.Redis.Host,
			"db", core.Config.Redis.DB)
		return
	}
	core.Logger.Debug("redis configuration", "enabled", false)
}
