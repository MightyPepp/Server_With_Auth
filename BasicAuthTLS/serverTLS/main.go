package main

import (
	"fmt"
	"log"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

// Пока сделаем просто - мапа user:hashPassword
var userPasswords = map[string]string{
	"Alex": "",
	"Bob": "",
	"Sara": "",
}

// Тут мы, следуя занятному примеру от Gemini,
// реализуем базовый процесс регистрации и 
// авторизации c использованием хэша пароля

func main() {
	http.HandleFunc("/api/basicauthTLS", func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`) // По стандартам надо так
			return
		}
		if user == "Alex" && pass == "secret" {
			w.Write([]byte("Auth with TLS OK\nBababa bebebe\nBobobo bibibi"))
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
	http.HandleFunc("/api/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About page"))
	})

	fmt.Println("Запускаем сервер...")
	if err := http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil); err != nil {
		log.Fatal(err)
	}
}
