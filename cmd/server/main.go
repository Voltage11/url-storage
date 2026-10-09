package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	

	"github.com/Voltage11/url-storage/internal/config"
	"github.com/Voltage11/url-storage/internal/transport/rest"
)

func main() {
	//Получим конфигурацию
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Ошибка получения конфигурации: %v", err)
	}

	// Создание HTTP сервера
	server := rest.NewHTTPServer(&cfg.Server)

	go func() {
		log.Printf("Запуск сервера на порту: %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Сервер остановлен: %v", err)
		}
	}()

	shutdownCh := make(chan os.Signal, 1)

	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	<-shutdownCh
	log.Println("Получен сигнал остановки сервиса")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		if err := server.Close(); err != nil {
			log.Printf("Ошибка остановки сервера: %v", err)
		}
		log.Printf("Ошибка остановки сервиса: %v", err)
	}

}
