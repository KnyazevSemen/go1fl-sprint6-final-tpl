package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags)

	srv := server.New(logger)

	if err := srv.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("ошибка запуска сервера: %v", err)
	}
}
