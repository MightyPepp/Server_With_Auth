package config

import (
	"log"
	"github.com/ilyakaznacheev/cleanenv"
)

const ConfigPath = "/home/mighty-pepe/Desktop/Server_With_Auth/cmd/api/configs/config.yml"

type AppConfig struct {
	ServerConfig
	LoggerConfig
	DBConfig
}

type ServerConfig struct {
	Port           	string `yaml:"port"`
	PathToCertfile 	string `yaml:"path_to_certfile"`
	PathToKeyfile  	string `yaml:"path_to_keyfile"`
}

// TODO:
type LoggerConfig struct {
	LoggerPrefix	string `yaml:"logger_prefix"`
	PathToLogfile 	string `yaml:"path_to_logfile"`
}

// TODO:
type DBConfig struct {} 

func GetConfig() (cfg *AppConfig, err error) {
	const op = "GetConfig"
	if err = cleanenv.ReadConfig(ConfigPath, cfg); err != nil {
		log.Printf("\nFROM: %s\nОшибка парсинга файла конфигурации: %s", op, err)
		return nil, err
	}
	return cfg, nil
}