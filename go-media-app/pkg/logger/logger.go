package logger

import (
	"log"
)


type Logger struct{}


func New() *Logger {
	return &Logger{}
}


func (l *Logger) Info(msg string) {
	log.Printf("INFO: %s", msg)
}


func (l *Logger) Errorf(format string, args ...interface{}) {
	log.Printf("ERROR: "+format, args...)
}
