package config

import (
	"net/http"
	"crypto/tls"
	"log"
)

func ConfigServerFunc() (err error, mux *http.ServeMux, server *http.Server) {
	const op = "ConfigServerFunc"

	appConfig, err := GetConfig()
	if err != nil {
		log.Printf("\nFROM: %s\nОшибка получения конфигурации сервера: %s", op, err)
		return err, nil, nil
	}

	cert, err := tls.LoadX509KeyPair(appConfig.ServerConfig.PathToCertfile, appConfig.ServerConfig.PathToKeyfile)
	if err != nil {
		log.Printf("\nFROM: %s\nОшибка чтения файла сертификата или ключа: %s", op, err)
		return err, nil, nil
	}

	mux = &http.ServeMux{}

	server = &http.Server{
		Addr:    appConfig.ServerConfig.Port,
		Handler: mux,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	return nil, mux, server
}