package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"crypto/x509"
)

func main() {
	// Чтение сертификата CA из файла
	caCert, err := os.ReadFile("ca.crt")
	if err != nil {
		log.Fatal("Failed to read CA cert:", err)
	}
	// Пул доверенных CA-сертификатов
	caCertPool := x509.NewCertPool() 
	if !caCertPool.AppendCertsFromPEM(caCert) {
		log.Fatal("Failed to parse CA cert")
	}

	// Настройка транспорта (low-level клиент) для клиента (high-level)
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: caCertPool,
			MinVersion: tls.VersionTLS12,
		},
		ForceAttemptHTTP2: true,
	}
	client := &http.Client{Transport: transport}

	// Настроим так же клиента для тестов
	testTransport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // Отключили проверку сертификатов
		},
		ForceAttemptHTTP2: true,
	}
	testClient := &http.Client{Transport: testTransport}

	// Подготовим запрос
	req, err := http.NewRequest("GET", "https://localhost:8443/api/basicauthTLS", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth("Alex", "secret")

	// Выполним запрос с обоих клиентов
	resp, err := testClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal("Request failed:", err)
		return 
	}
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Status: %s\nBody: %s\n", resp.Status, body)

	resp, err = client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal("Request failed:", err)
		return
	}
	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("Status: %s\nBody: %s\n", resp.Status, body)
}