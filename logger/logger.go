package logger

import (
	"log"
	"os"
	"path"
	"runtime"
)

var l *log.Logger

func getProjectRootDir() string {
	// shamelessly stolen from https://brandur.org/fragments/testing-go-project-root
	_, filename, _, _ := runtime.Caller(0)
	projectRootDir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(projectRootDir)
	if err != nil {
		log.Fatal(err)
	}
	return projectRootDir
}

func init() {
	appLogFilePath := path.Join(getProjectRootDir(), "logs", "app.log")
	if os.Getenv("CI") == "true" {
		appLogFilePath += ".example"
	}

	file, err := os.OpenFile(appLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	l = log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs an info message
func Info(msg string) {
	l.SetPrefix("INFO: ")
	l.Println(msg)
}

// Warn logs a warning message
func Warn(msg string) {
	l.SetPrefix("WARN: ")
	l.Println(msg)
}

// Error logs an error message and panics
func Error(msg string) {
	l.SetPrefix("ERROR: ")
	l.Fatal(msg)
}
