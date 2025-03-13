package main

import "fmt"

type LegacyLogger struct{}

func (l *LegacyLogger) LogMessage(message string) {
	fmt.Println("Legacy logger: " + message)
}

type NewLogger interface {
	Log(message string)
}

type NewSystemLogger struct{}

func (l *NewSystemLogger) Log(message string) {
	fmt.Println("New system logger: " + message)
}

type LoggerAdapter struct {
	legacyLogger *LegacyLogger
}

func (a *LoggerAdapter) Log(message string) {
	// Adapting the interface here:
	a.legacyLogger.LogMessage(message)
}

func main() {
	legacyLogger := &LegacyLogger{}
	loggerAdapter := &LoggerAdapter{legacyLogger: legacyLogger}
	newLogger := &NewSystemLogger{}

	loggerAdapter.Log("Hello, world!")
	newLogger.Log("Hello, world!")
}
