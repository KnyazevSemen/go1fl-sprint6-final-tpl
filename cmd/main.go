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

	// Запускаем сервер — это единственная операция после создания
	// ListenAndServe() блокирует выполнение, поэтому сервер работает до прерывания
	if err := srv.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("ошибка запуска сервера: %v", err)
	}

	logger.Println("Сервер успешно запущен")
}
