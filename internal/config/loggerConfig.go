package config

import (
	"log"
	"os"
)

// TODO:
func GetAndConfigLoggerFunc() (appLogger *log.Logger, file *os.File, err error) {
	const op = "ConfigLoggerFunc"

	аppConfig, err := GetConfig()
	if err != nil {
		log.Printf("\nFROM: %s\nОшибка получения конфигурации сервера: %s", op, err)
		return nil, nil, err
	}

	logFile, err := os.OpenFile(аppConfig.LoggerConfig.PathToLogfile, os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("\nFROM: %s\nОшибка открытия файла логов: %s", op, err)
		return nil, nil, err
	}

	appLogger = &log.Logger{}

	appLogger.SetPrefix(аppConfig.LoggerConfig.LoggerPrefix)
	appLogger.SetOutput(logFile)

	return appLogger, logFile, nil
}