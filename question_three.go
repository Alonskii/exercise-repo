package main

import (
	"fmt"
	"os"
)

// Logger defines the interface for logging messages.
type Logger interface {
	Log(message string) // Log logs the provided message.
}

// FileLogger logs messages to a file.
type FileLogger struct {
	FilePath string // FilePath is the path to the log file.
}

// Log implements the Log method of the Logger interface for FileLogger.
func (f *FileLogger) Log(message string) {
	file, err := os.OpenFile(f.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, FileMode)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%s\n", message)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

// ConsoleLogger logs messages to the console.
type ConsoleLogger struct{}

// Log implements the Log method of the Logger interface for ConsoleLogger.
func (c *ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

// FileMode is the file mode used when creating log files.
const (
	FileMode = 0644
)

// Example usage of logging to both file and console.
func ExampleLogger() {
	fileLogger := &FileLogger{FilePath: "log.txt"}
	fileLogger.Log("This is a message logged to a file.")

	consoleLogger := &ConsoleLogger{}
	consoleLogger.Log("This is a message logged to the console.")
}

func main() {
	ExampleLogger()
}
