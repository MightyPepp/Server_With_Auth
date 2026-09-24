package config

import (
	"net/http"
	"crypto/tls"
	"log"
)

func GetAndConfigServerFunc() (server *http.Server, mux *http.ServeMux, err error) {
	const op = "ConfigServerFunc"

	appConfig, err := GetConfig()
	if err != nil {
		log.Printf("\nFROM: %s\nОшибка получения конфигурации сервера: %s", op, err)
		return nil, nil, err
	}

	cert, err := tls.LoadX509KeyPair(appConfig.ServerConfig.PathToCertfile, appConfig.ServerConfig.PathToKeyfile)
	if err != nil {
		log.Printf("\nFROM: %s\nОшибка чтения файла сертификата или ключа: %s", op, err)
		return nil, nil, err
	}

	mux = &http.ServeMux{}

	server = &http.Server{
		Addr:    appConfig.ServerConfig.Port,
		Handler: mux,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	return server, mux, nil
}