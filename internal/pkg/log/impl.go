package log

import (
	"io"
	"log"
	"os"
)

// Log is the logger.
var Log Logger

// LoggerImpl is the implementation of the logger.
type LoggerImpl struct {
	log *log.Logger
}

// NewLogger creates a new logger.
func NewLogger() {
	Log = &LoggerImpl{
		log: log.New(os.Stderr, "", log.LstdFlags),
	}
}

// NewTestLogger creates a new logger for testing. Output is written to the
// provided io.Writer. If no io.Writer is provided, output is written to
// os.Stderr.
func NewTestLogger(output io.Writer) {
	if output == nil {
		output = os.Stderr
	}
	Log = &LoggerImpl{
		log: log.New(output, "", log.LstdFlags),
	}
}

// Debug logs a debug message.
func (l *LoggerImpl) Debug(message string) {
	l.log.Printf("[DEBUG] " + message)
}

// Debugf logs a formatted debug message.
func (l *LoggerImpl) Debugf(format string, args ...interface{}) {
	l.log.Printf("[DEBUG] "+format, args...)
}

// Info logs an info message.
func (l *LoggerImpl) Info(message string) {
	l.log.Println("[INFO] " + message)
}

// Infof logs a formatted info message.
func (l *LoggerImpl) Infof(format string, args ...interface{}) {
	l.log.Printf("[INFO] "+format, args...)
}

// Error logs an error message.
func (l *LoggerImpl) Error(message string) {
	l.log.Println("[ERROR] " + message)
}

// Errorf logs a formatted error message.
func (l *LoggerImpl) Errorf(format string, args ...interface{}) {
	l.log.Printf("[ERROR] "+format, args...)
}

// Fatal logs a fatal message.
func (l *LoggerImpl) Fatal(message string) {
	l.log.Fatalln("[FATAL] " + message)
}

// Fatalf logs a formatted fatal message.
func (l *LoggerImpl) Fatalf(format string, args ...interface{}) {
	l.log.Fatalf("[FATAL] "+format, args...)
}

// SetOutput sets the output destination for the logger.
func (l *LoggerImpl) SetOutput(out io.Writer) {
	l.log.SetOutput(out)
}
