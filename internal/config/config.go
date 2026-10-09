package config

import (
	"time"

	"github.com/joho/godotenv"
)

// ServerConfig Конфигурация сервера
type ServerConfig struct {
	Addr              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	ShutdownTimeout   time.Duration
}

// Config конфигурация приложения, будем дорабатывать по мере потребности
type Config struct {
	Server ServerConfig
}

// New конструктор создания конфигурации
func New() (*Config, error) {

	_ = godotenv.Load()

	serverConfig := ServerConfig{
		Addr:              getEnv("SERVER_ADDR", ":8080"),
		ReadTimeout:       intToTimeDurationSecond(getEnvInt("READ_TIMEOUT_SEC", 10)),
		WriteTimeout:      intToTimeDurationSecond(getEnvInt("WRITE_TIMEOUT_SEC", 30)),
		IdleTimeout:       intToTimeDurationSecond(getEnvInt("IDLE_TIMEOUT_SEC", 60)),
		ReadHeaderTimeout: intToTimeDurationSecond(getEnvInt("READ_HEADER_TIMEOUT_SEC", 5)),
		ShutdownTimeout:   intToTimeDurationSecond(getEnvInt("SHUTDOWN_TIMEOUT_SEC", 5)),
	}

	config := &Config{
		Server: serverConfig,
	}

	return config, nil
}

func intToTimeDurationSecond(value int) time.Duration {
	return time.Duration(value) * time.Second
}
