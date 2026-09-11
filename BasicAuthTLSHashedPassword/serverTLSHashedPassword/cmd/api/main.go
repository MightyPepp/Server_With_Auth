package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	muxNoTLS := http.NewServeMux()
	muxNoTLS.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	go http.ListenAndServe(":8080", muxNoTLS)

	fmt.Println("Запускаем сервер...")
	if err := http.ListenAndServeTLS(":8443", "../server.crt", "../server.key", mux); err != nil {
		log.Fatal(err)
	}
}