package rest

import (
	"net/http"

	"github.com/Voltage11/url-storage/internal/config"
	"github.com/Voltage11/url-storage/internal/transport/rest/handlers"
)

// NewHTTPServer конструктор создания http сервера
func NewHTTPServer(cfg *config.ServerConfig) *http.Server {

	router := http.NewServeMux()
	registerHandlers(router)

	server := &http.Server{
		Addr:              cfg.Addr,
		Handler:           router,
		ReadTimeout:       cfg.ReadTimeout,       //Таймаут на чтение
		WriteTimeout:      cfg.WriteTimeout,      //Таймаут на запись ответа
		IdleTimeout:       cfg.IdleTimeout,       //Таймаут соединение на переиспользование тем же клиентом
		ReadHeaderTimeout: cfg.ReadHeaderTimeout, //Таймаут на чтение заголовков
	}

	return server
}

// registerHandlers регистрация всех роутеров
func registerHandlers(router *http.ServeMux) {
	router.HandleFunc("GET /health", handlers.Health)
}
