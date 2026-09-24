package logging

import (
	"os"
	"log"
)

// TODO:
func GetLogger(outputPath string) (*log.Logger, *os.File) {
	file, err := os.OpenFile(outputPath, os.O_WRONLY, 0664)
	if err != nil {
		log.Printf("Не удалось открыть файл логов: %s\n", err)
	}
	myLogger := log.New(file, "SERVER: ", log.Lmsgprefix|log.Ldate|log.Ltime)
	return myLogger, file
}