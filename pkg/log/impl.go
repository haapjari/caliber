package log

import (
	"io"
	"log"
	"os"
)

var Log Logger

type LoggerImpl struct {
	log *log.Logger
}

func NewLogger() {
	Log = &LoggerImpl{
		log: log.New(os.Stderr, "", log.LstdFlags),
	}
}

func NewTestLogger(output io.Writer) {
	if output == nil {
		output = os.Stderr
	}
	Log = &LoggerImpl{
		log: log.New(output, "", log.LstdFlags),
	}
}

func (l *LoggerImpl) Debug(message string) {
	l.log.Printf("[DEBUG] " + message)
}

func (l *LoggerImpl) Debugf(format string, args ...interface{}) {
	l.log.Printf("[DEBUG] "+format, args...)
}

func (l *LoggerImpl) Info(message string) {
	l.log.Println("[INFO] " + message)
}

func (l *LoggerImpl) Infof(format string, args ...interface{}) {
	l.log.Printf("[INFO] "+format, args...)
}

func (l *LoggerImpl) Error(message string) {
	l.log.Println("[ERROR] " + message)
}

func (l *LoggerImpl) Errorf(format string, args ...interface{}) {
	l.log.Printf("[ERROR] "+format, args...)
}

func (l *LoggerImpl) Fatal(message string) {
	l.log.Fatalln("[FATAL] " + message)
}

func (l *LoggerImpl) Fatalf(format string, args ...interface{}) {
	l.log.Fatalf("[FATAL] "+format, args...)
}
