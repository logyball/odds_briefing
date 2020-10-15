package main

import (
	logger "github.com/loganballard/odds_briefing/logger"
)

// var (
// 	WarningLogger *log.Logger
// 	InfoLogger    *log.Logger
// 	ErrorLogger   *log.Logger
// )

func init() {
	// appLogFilePath := path.Join("logs", "app.log")
	// errorLogFilePath := path.Join("logs", "error.log")
	// if os.Getenv("CI") == "true" {
	// 	appLogFilePath += ".example"
	// 	errorLogFilePath += ".example"
	// }

	// file, err := os.OpenFile(appLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// errorLogFile, err := os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	// WarningLogger = log.New(errorLogFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	// ErrorLogger = log.New(errorLogFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	logger.Info("starting the app")
}
