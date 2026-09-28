package config

import (
	"log"
	"os"
)

func GetAndConfigLoggerFunc(cfg *AppConfig) (appLogger *log.Logger, file *os.File, err error) {
	const op = "ConfigLoggerFunc"

	logFile, err := os.OpenFile(cfg.LoggerConfig.PathToLogfile, os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("\nFROM: %s\nОшибка открытия файла логов: %s", op, err)
		return nil, nil, err
	}

	appLogger = &log.Logger{}

	appLogger.SetPrefix(cfg.LoggerConfig.LoggerPrefix)
	appLogger.SetOutput(logFile)

	return appLogger, logFile, nil
}