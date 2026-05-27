package app

import "github.com/rei0721/go-scaffold/internal/app/modeapp"

type AppMode = modeapp.Mode

const (
	ModeServer = modeapp.ModeServer
)

const (
	ConstantsI18nMessagesDir     = "./configs/locales"
	ConstantsI18nDefaultLanguage = "zh-CN"
	ConstantsDefaultHost         = "localhost"
)

var ConstantsI18nSupportedLanguages = []string{"zh-CN", "en-US"}
