package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	cfg "Server_With_Auth/internal/config"
	handlers "Server_With_Auth/internal/handlers"
	"Server_With_Auth/internal/middleware"
)

// TODO: golangci настроить
func main() {
	const op = "main"

	// _____Настройка сервера_____

	// Получаем сконфигурированные компоненты приложения
	appLogger, logFile, appServer, mux, err := cfg.ConfigureApp()
	defer logFile.Close()
	if err != nil {
		appLogger.Printf("\nFROM: %s\nОшибка получения сконфигурированного приложения: %s", op, err)
		return
	}

	// I guess здесь оборачиваем mux в logging middleware
	loggingMux := middleware.LoggingMiddleware(mux, appLogger)
	appServer.Handler = loggingMux

	// Создаём и регистрируем обработчики
	healthHandler 	:= handlers.NewHealthHandler()
	authHandler 	:= handlers.NewAuthHandler()

	mux.Handle("/health", healthHandler)
	mux.Handle("/auth", authHandler)

	// _____Запуск сервера_____

	// В отдельной горутине запускаем сервер
	go func() {
		appLogger.Printf("\nFROM: %s\nЗапуск сервера на %s...", op, appServer.Addr)
		if err := appServer.ListenAndServeTLS("", ""); err != nil {
			appLogger.Printf("\nFROM: %s\nОшибка запуска сервера: %s", op, err)
		}
	}()

	// Слушаем сигналы Ctrl + C и kill по PID для graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	<-sigChan
	appLogger.Printf("\nFROM: %s\nПолучен SIGTERM, работа сервера завершается...", op)
	if err := appServer.Shutdown(context.Background()); err != nil {
		appLogger.Printf("\nFROM: %s\nОшибка остановки работы сервера: %s", op, err)
	}
	appLogger.Printf("\nFROM: %s\nСервер остановлен!", op)
}
