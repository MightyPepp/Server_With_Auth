package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func configServerTLS() (*http.ServeMux, *http.Server, error) {
	mux := http.NewServeMux()

	cert, err := tls.LoadX509KeyPair("/etc/ssl/self_signed/certs/server.crt", "/etc/ssl/self_signed/private/server.key")
	if err != nil {
		log.Fatal("Ошибка загрузки сертификата:", err)
		return nil, nil, err
	}

	server := &http.Server{
		Addr: ":8443",
		Handler: mux,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}
	return mux, server, nil
}

func main() {
	file, err := os.OpenFile("/home/mighty-pepe/Desktop/Server_With_Auth/BasicAuthTLSHashedPassword/serverTLSHashedPassword/logs/log.txt", os.O_WRONLY, 0664)
	if err != nil {
		log.Printf("Не удалось открыть файл логов: %s\n", err)
	}
	defer file.Close()
	log.SetOutput(file)
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGTERM)
	
	mux, server, err := configServerTLS()
	if err != nil {
		log.Printf("Ошибка конфигурации сервера: %s\n", err)
	}

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Запрос на ручку: %s", r.URL)
		w.Write([]byte("pong"))
	})

	go func() {
		if err := server.ListenAndServeTLS("", ""); err != http.ErrServerClosed {
			log.Fatal("Сервер упал с ошибкой:", err)
		}
	}()

	<-sigChan
	log.Println("Получен сигнал SIGTERM, начинаем завершение работы сервера.")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("Ошибка остановки работы сервера: %s\n", err)
	}
	log.Println("Работа сервера завершена.")
}