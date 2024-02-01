package sslogger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger struct{}

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
	GreenColor   = "\033[1;32m%s\033[0m"
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

func (l Logger) Success(message string) {
	LogMessage("success", message)
}

func LogMessage(level, message string) {

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	formattedMessage := fmt.Sprintf("[%s] » %s", currentTime, message)

	// Create a custom logger output
	log.SetFlags(0) // Removes the default timestamp
	log.SetOutput(os.Stdout)


	switch level {
	case "info":
		log.Printf(InfoColor, formattedMessage)
	case "warning":
		log.Printf(WarningColor, formattedMessage)
	case "error":
		log.Printf(ErrorColor, formattedMessage)
	case "debug":
		log.Printf(DebugColor, formattedMessage)
	case "success":
		log.Printf(GreenColor, formattedMessage)
	default:
		log.Printf(formattedMessage)
	}
}
