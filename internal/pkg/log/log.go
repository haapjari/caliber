// Package log provides a simple interface for logging.
package log

import "io"

// Logger is the interface for the logger.
type Logger interface {
	Debug(message string)
	Debugf(format string, args ...interface{})
	Info(message string)
	Infof(format string, args ...interface{})
	Error(message string)
	Errorf(format string, args ...interface{})
	Fatal(message string)
	Fatalf(format string, args ...interface{})
	SetOutput(io.Writer)
}

// Debug logs a debug message.
func Debug(message string) {
	Log.Debug(message)
}

// Debugf logs a formatted debug message.
func Debugf(format string, args ...interface{}) {
	Log.Debugf(format, args...)
}

// Info logs an info message.
func Info(message string) {
	Log.Info(message)
}

// Infof logs a formatted info message.
func Infof(format string, args ...interface{}) {
	Log.Infof(format, args...)
}

// Error logs an error message.
func Error(message string) {
	Log.Error(message)
}

// Errorf logs a formatted error message.
func Errorf(format string, args ...interface{}) {
	Log.Errorf(format, args...)
}

// Fatal logs a fatal message.
func Fatal(message string) {
	Log.Fatal(message)
}

// Fatalf logs a formatted fatal message.
func Fatalf(format string, args ...interface{}) {
	Log.Fatalf(format, args...)
}

// SetOutput sets the output destination for the logger.
func SetOutput(out io.Writer) {
	Log.SetOutput(out)
}
