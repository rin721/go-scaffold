package config

import (
	"strings"
	"unicode"

	app "github.com/rei0721/go-scaffold/types/constants"
)

// EnvPrefix returns the dynamic prefix used by configuration env overrides.
//
// Rules:
//   - derive the application token from app.AppPrefix
//   - normalize it to uppercase underscore-separated env-token format
//   - configuration override names use <APP_PREFIX>_APP_<MODULE>_<FIELD>
//
// Example: app.AppPrefix = "Rin" returns "RIN_APP".
func EnvPrefix() string {
	prefix := normalizeEnvToken(app.AppPrefix)
	if prefix == "" {
		return "APP"
	}
	return prefix + "_APP"
}

// EnvConfigPathName returns the dynamic env name for the config file path.
//
// Example: app.AppPrefix = "Rin" returns "RIN_CONFIG_PATH".
func EnvConfigPathName() string {
	prefix := normalizeEnvToken(app.AppPrefix)
	if prefix == "" {
		return "CONFIG_PATH"
	}
	return prefix + "_CONFIG_PATH"
}

// EnvPrefixJoin joins an envname tag value with the dynamic application prefix.
func EnvPrefixJoin(field string) string {
	field = strings.TrimSpace(field)
	prefix := EnvPrefix()
	if field == "" {
		return prefix
	}
	if prefix == "" || strings.HasPrefix(field, prefix+"_") {
		return field
	}
	return prefix + "_" + field
}

func normalizeEnvToken(value string) string {
	var builder strings.Builder
	lastUnderscore := false
	for _, r := range value {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			builder.WriteRune(unicode.ToUpper(r))
			lastUnderscore = false
			continue
		}
		if !lastUnderscore && builder.Len() > 0 {
			builder.WriteByte('_')
			lastUnderscore = true
		}
	}
	return strings.Trim(builder.String(), "_")
}

const (
	EnvFilePath        = ".env"
	EnvFilePathExample = ".env.example"
)

const (
	DefaultSeparator = ","
)

const (
	AppServerName   = "server"
	AppDatabaseName = "database"
	AppRedisName    = "redis"
	AppLoggerName   = "logger"
	AppI18nName     = "i18n"
	AppExecutorName = "executor"
	AppStorageName  = "storage"
	AppDemoName     = "demo"
	AppPluginName   = "plugin"
	AppIAMName      = "iam"
	AppAuthName     = "auth"
	AppRBACName     = "rbac"
	AppCORSName     = "cors"
)
