package config

import (
	"net/http"
	"log"
	"os"
	"github.com/ilyakaznacheev/cleanenv"
)

const ConfigPath = "/app/cmd/api/configs/config.yml"

type AppConfig struct {
	ServerConfig	`yaml:"server"`
	LoggerConfig	`yaml:"logger"`
	DBConfig		`yaml:"db"`
}

type ServerConfig struct {
	Port           	string `yaml:"port"`
	PathToCertfile 	string `yaml:"path_to_certfile"`
	PathToKeyfile  	string `yaml:"path_to_keyfile"`
}

type LoggerConfig struct {
	LoggerPrefix	string `yaml:"logger_prefix"`
	PathToLogfile 	string `yaml:"path_to_logfile"`
}

// TODO: Сделат подлючение к БД и саму БД
type DBConfig struct {} 

func GetConfig() (*AppConfig, error) {
	const op = "GetConfig"

	var cfg AppConfig

	err := cleanenv.ReadConfig(ConfigPath, &cfg); 
	if err != nil {
		log.Printf("\nFROM: %s\nОшибка парсинга файла конфигурации: %s", op, err)
		return nil, err
	}
	return &cfg, nil
}

func ConfigureApp() (appLogger *log.Logger, logFile *os.File, appServer *http.Server, mux *http.ServeMux, err error) {
	const op = "ConfigureApp"

	// Один раз вызвать GetConfig
	cfg, err := GetConfig()
	if err != nil {
		log.Printf("\nFROM: %s\nОшибка парсинга файла конфигурации: %s", op, err)
		return nil, nil, nil, nil, err
	}

	// Настроить логгер, передав указатель на структуру конфига
	appLogger, logFile, err = GetAndConfigLoggerFunc(cfg)
	if err != nil {
		log.Printf("\nFROM: %s\nОшибка настройки логгера: %s", op, err)	
		return nil, nil, nil, nil, err
	}

	// Настроить сервер, передав указатель на структуру конфига
	appServer, mux, err = GetAndConfigServerFunc(cfg)
	if err != nil {
		log.Printf("\nFROM: %s\nОшибка настройки сервера: %s", op, err)
		return nil, nil, nil, nil, err
	}

	// TODO: настроить подключение к БД, передав указатель на структуру конфига

	return appLogger, logFile, appServer, mux, nil
}