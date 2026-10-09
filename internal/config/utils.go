package config

import (
	"os"
	"strconv"
	"strings"
)

// getEnv получение значение переменной окружения с дефолтным значением при отсутствии
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	value = strings.TrimSpace(value)
	if value == "" {
		return defaultValue
	}

	return value
}

// getEnvint получение значение переменной окружения с дефолтным значением при отсутствии
func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	value = strings.TrimSpace(value)
	if value == "" {
		return defaultValue
	}

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return valueInt
}
