package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func pingFoodService() {
	resp, err := http.Get("http://localhost:8081/food-service/api/ping")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Response is:\n%s", body)
}

func createOrder() {
	jsonData := `{
		"user_id": 1,
		"items": [
			{"menu_item_id": 1, "quantity": 3},
			{"menu_item_id": 1, "quantity": 3},
			{"menu_item_id": 1, "quantity": 3},
			{"menu_item_id": 1, "quantity": 3},
			{"menu_item_id": 1, "quantity": 3},
			{"menu_item_id": 1, "quantity": 3}
		]
	}`
	body := strings.NewReader(jsonData)

	resp, err := http.Post("http://localhost:8080/aggregator/api/orders", "application/json", body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	fmt.Printf("Response is:\n%s", respBody)
}

func main() {
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About page"))
	})
	http.HandleFunc("/auth/{user}/{passwd}", func(w http.ResponseWriter, r *http.Request) {
		user, passwd := r.PathValue("user"), r.PathValue("passwd")
		if user == "Alex" && passwd == "secret" {
			w.Write([]byte("Auth OK"))
		} else {
			http.NotFound(w, r)
		}
	})
	http.HandleFunc("/auth/basicauth", func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if user == "Alex" && pass == "secret" {
			w.Write([]byte("Auth OK"))
		}
	})
	
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}