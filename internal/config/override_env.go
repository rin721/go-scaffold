package config

import (
	"encoding/json"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const envNameTag = "envname"

func overrideConfigFromEnv(target any) {
	value := reflect.ValueOf(target)
	if !value.IsValid() || value.Kind() != reflect.Pointer || value.IsNil() {
		return
	}
	overrideStructFields(value.Elem())
}

func overrideStructFields(value reflect.Value) {
	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return
		}
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return
	}

	valueType := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		structField := valueType.Field(i)
		if structField.PkgPath != "" {
			continue
		}

		if envName := strings.TrimSpace(structField.Tag.Get(envNameTag)); envName != "" && envName != "-" {
			if raw, ok := lookupTaggedEnv(envName); ok {
				setValueFromEnv(field, raw)
			}
		}

		if shouldRecurseEnvField(field) {
			overrideStructFields(field)
		}
	}
}

func shouldRecurseEnvField(field reflect.Value) bool {
	switch field.Kind() {
	case reflect.Struct:
		return true
	case reflect.Pointer:
		return !field.IsNil() && field.Elem().Kind() == reflect.Struct
	default:
		return false
	}
}

func lookupTaggedEnv(envName string) (string, bool) {
	for _, candidate := range envNameCandidates(envName) {
		if value, ok := os.LookupEnv(candidate); ok && value != "" {
			return value, true
		}
	}
	return "", false
}

func envNameCandidates(envName string) []string {
	envName = strings.TrimSpace(envName)
	if envName == "" {
		return nil
	}

	prefixed := EnvPrefixJoin(envName)
	if prefixed == envName {
		return []string{envName}
	}
	return []string{prefixed, envName}
}

func setValueFromEnv(field reflect.Value, raw string) bool {
	if !field.CanSet() {
		return false
	}

	if field.Kind() == reflect.Pointer {
		return setPointerValueFromEnv(field, raw)
	}

	return setConcreteValueFromEnv(field, raw)
}

func setPointerValueFromEnv(field reflect.Value, raw string) bool {
	elem := reflect.New(field.Type().Elem()).Elem()
	if !setConcreteValueFromEnv(elem, raw) {
		return false
	}
	field.Set(elem.Addr())
	return true
}

func setConcreteValueFromEnv(field reflect.Value, raw string) bool {
	switch field.Kind() {
	case reflect.String:
		field.SetString(raw)
		return true
	case reflect.Bool:
		value, err := strconv.ParseBool(raw)
		if err != nil {
			return false
		}
		field.SetBool(value)
		return true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value, err := strconv.ParseInt(raw, 10, field.Type().Bits())
		if err != nil {
			return false
		}
		field.SetInt(value)
		return true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value, err := strconv.ParseUint(raw, 10, field.Type().Bits())
		if err != nil {
			return false
		}
		field.SetUint(value)
		return true
	case reflect.Float32, reflect.Float64:
		value, err := strconv.ParseFloat(raw, field.Type().Bits())
		if err != nil {
			return false
		}
		field.SetFloat(value)
		return true
	case reflect.Slice:
		return setSliceValueFromEnv(field, raw)
	default:
		return false
	}
}

func setSliceValueFromEnv(field reflect.Value, raw string) bool {
	elemKind := field.Type().Elem().Kind()
	switch elemKind {
	case reflect.String:
		values := splitEnvList(raw)
		if len(values) == 0 {
			return false
		}
		field.Set(reflect.ValueOf(values).Convert(field.Type()))
		return true
	default:
		value := reflect.New(field.Type()).Interface()
		if err := json.Unmarshal([]byte(raw), value); err != nil {
			return false
		}
		field.Set(reflect.ValueOf(value).Elem())
		return true
	}
}

func splitEnvList(raw string) []string {
	parts := strings.Split(raw, DefaultSeparator)
	values := make([]string, 0, len(parts))
	for _, part := range parts {
		value := strings.TrimSpace(part)
		if value != "" {
			values = append(values, value)
		}
	}
	return values
}
