package server

import (
	"net/http"
	"log"
	"os"

	cfg 	"Server_With_Auth/internal/config"
)

// TODO:
func registerHandlers(mux *http.ServeMux) {

}

// TODO:
func GetConfiguredServerFunc() (appLogger *log.Logger, logFile *os.File, server *http.Server, mux *http.ServeMux, err error) {
	const op = "StartServerFunc"

	appLogger, logFile, err = cfg.GetAndConfigLoggerFunc()
	if err != nil {
		log.Printf("\nFROM: %s\nОшибка получения или конфигурации логгера: %s", op, err)
		return nil, nil, nil, nil, err
	}
	
	server, mux, err = cfg.GetAndConfigServerFunc()

	// TODO: регистрация обработчиков доделать

	return appLogger, logFile, server, mux, nil
}