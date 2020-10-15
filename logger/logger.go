package logger

import (
	"log"
	"os"
	"path"
)

var l *log.Logger

func init() {
	appLogFilePath := path.Join("logs", "app.log")
	if os.Getenv("CI") == "true" {
		appLogFilePath += ".example"
	}
	file, err := os.OpenFile(appLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	l = log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(msg string) {
	l.SetPrefix("INFO: ")
	l.Println(msg)
}

func Warn(msg string) {
	l.SetPrefix("WARN: ")
	l.Println(msg)
}

func Error(msg string) {
	l.SetPrefix("ERROR: ")
	l.Fatal(msg)
}
