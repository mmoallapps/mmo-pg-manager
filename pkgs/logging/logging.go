package logging

import (
	"log"
	"os"
	"path/filepath"
)

// create a basic function that accepts errors and writes to logs/error.log
func LogError(err error) {
	// open the logs/error.log file for appending
	// check if the file exists, if not create it
	// create the logs directory if it doesn't exist

	logFilePath := filepath.Join("logs", "error.log")

	// open the log file for appending
	logFile, fileErr := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	defer logFile.Close()
	// create a new logger that writes to the log file
	logger := log.New(logFile, "", log.LstdFlags)
	logger.Println(err.Error())
}
