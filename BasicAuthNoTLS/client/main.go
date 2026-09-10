package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8081/auth/basicauth", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth("Alex", "secret")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	responseBody, _ := io.ReadAll(resp.Body)
	fmt.Println(string(responseBody))
}