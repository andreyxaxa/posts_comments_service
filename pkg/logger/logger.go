package logger

import (
	"log"
	"os"
)

type Logger struct {
	Error *log.Logger
	Info  *log.Logger
}

func NewLogger() *Logger {
	logger := Logger{}
	logger.Error = log.New(os.Stderr, "ERROR  ", log.Ldate|log.Ltime)
	logger.Info = log.New(os.Stdout, "INFO  ", log.Ldate|log.Ltime)

	return &logger
}
