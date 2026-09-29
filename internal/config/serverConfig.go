package config

import (
	"net/http"
	"crypto/tls"
	"log"
)

func GetAndConfigServerFunc(cfg *AppConfig) (server *http.Server, mux *http.ServeMux, err error) {
	const op = "ConfigServerFunc"

	cert, err := tls.LoadX509KeyPair(cfg.ServerConfig.PathToCertfile, cfg.ServerConfig.PathToKeyfile)
	if err != nil {
		log.Printf("\nFROM: %s\nОшибка чтения файла сертификата или ключа: %s", op, err)
		return nil, nil, err
	}

	mux = &http.ServeMux{}

	server = &http.Server{
		Addr:    cfg.ServerConfig.Port,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	return server, mux, nil
}