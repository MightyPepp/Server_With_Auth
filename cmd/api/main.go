package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	srv 		"Server_With_Auth/internal/server"
	handlers 	"Server_With_Auth/internal/handlers"
)

// TODO: запуск сервера написать
func main() {
	const op = "main"

	appLogger, logFile, appServer, mux, err := srv.GetConfiguredServerFunc()
	defer logFile.Close()
	if err != nil {
		appLogger.Printf("\nFROM: %s\nОшибка получения сконфигурированного сервера: %s", op, err)
	}

	hh := handlers.NewHealthHandler()
	hh.RegisterHealthHandler(mux)


	go func() {
		appLogger.Printf("\nFROM: %s\nЗапуск сервера на %s...", op, appServer.Addr)
		if err := appServer.ListenAndServeTLS("", ""); err != nil {
			appLogger.Printf("\nFROM: %s\nОшибка запуска сервера: %s", op, err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)
	<-sigChan
	appLogger.Printf("\nFROM: %s\nПолучен SIGTERM, работа сервера завершается...", op)
	if err := appServer.Shutdown(context.Background()); err != nil {
		appLogger.Printf("\nFROM: %s\nОшибка остановки работы сервера: %s", op, err)
	}
	appLogger.Printf("\nFROM: %s\nСервер остановлен!", op)
}
