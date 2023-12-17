package sslogger

import (
	"fmt"
	"log"
	"time"
)

type Logger struct{}

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

func (l Logger) Info(message string) {
	LogMessage("info", message)
}

func (l Logger) Warn(message string) {
	LogMessage("warning", message)
}

func (l Logger) Error(message string) {
	LogMessage("error", message)
}

func (l Logger) Debug(message string) {
	LogMessage("debug", message)
}

func LogMessage(level, message string) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	formattedMessage := fmt.Sprintf("[%s] » %s", currentTime, message)
	switch level {
	case "info":
		log.Printf(InfoColor, formattedMessage)
	case "warning":
		log.Printf(WarningColor, formattedMessage)
	case "error":
		log.Printf(ErrorColor, formattedMessage)
	case "debug":
		log.Printf(DebugColor, formattedMessage)
	default:
		log.Printf(formattedMessage)
	}
}
