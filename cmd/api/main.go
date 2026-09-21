package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	handlers "Server_With_Auth/internal/handlers"
)

func configServerTLS() (*http.ServeMux, *http.Server, error) {
	mux := http.NewServeMux()

	cert, err := tls.LoadX509KeyPair("self_signed/certs/server.crt", "self_signed/private/server.key")
	if err != nil {
		log.Fatal("Ошибка загрузки сертификата:", err)
		return nil, nil, err
	}

	server := &http.Server{
		Addr:    ":8443",
		Handler: mux,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}
	return mux, server, nil
}

func getLogger(outputPath string) (*log.Logger, *os.File) {
	file, err := os.OpenFile(outputPath, os.O_WRONLY, 0664)
	if err != nil {
		log.Printf("Не удалось открыть файл логов: %s\n", err)
	}
	myLogger := log.New(file, "SERVER: ", log.Lmsgprefix|log.Ldate|log.Ltime)
	return myLogger, file
}

func main() {
	psswds := make(map[string]string)
	psswds["User1"] = "Psswd1"
	psswds["User2"] = "Psswd2"
	myHandler := handlers.NewMyHandler(psswds)

	myLogger, outputFile := getLogger("logs/log.txt")
	defer outputFile.Close()

	mux, server, err := configServerTLS()
	if err != nil {
		myLogger.Printf("Ошибка конфигурации сервера: %s\n", err)
	}
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		myLogger.Printf("Запрос на ручку: %s", r.URL)
		w.Write([]byte("pong"))
	})
	mux.HandleFunc("/api/auth", myHandler.AuthHandler)

	go func() {
		myLogger.Printf("Запуск сервера на порту %s", server.Addr)
		if err := server.ListenAndServeTLS("", ""); err != http.ErrServerClosed {
			myLogger.Fatal("Сервер упал с ошибкой:", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)
	<-sigChan
	myLogger.Println("Получен сигнал SIGTERM, начинаем завершение работы сервера.")
	if err := server.Shutdown(context.Background()); err != nil {
		myLogger.Printf("Ошибка остановки работы сервера: %s\n", err)
	}
	myLogger.Println("Работа сервера завершена.")
}
