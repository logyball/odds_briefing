package main

import (
	"log"
	"os"
	"path"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile(path.Join("logs", "app.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	errorLogFile, err := os.OpenFile(path.Join("logs", "error.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(errorLogFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(errorLogFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	InfoLogger.Println("starting the app")
}
