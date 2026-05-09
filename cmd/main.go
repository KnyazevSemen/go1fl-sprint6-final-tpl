package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags)

	srv := server.New(logger)

	logger.Println("Запуск сервера")

	err := srv.Start()
	if err != nil {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
	}

	logger.Println("Сервер успешно запущен")
}
