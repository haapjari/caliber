package log

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormattedLogging(t *testing.T) {
	tests := []struct {
		name     string
		logFunc  func(string, ...interface{})
		logLevel string
		args     []interface{}
	}{
		{"Debugf", Debugf, "[DEBUG]", []interface{}{"%s", "formatted"}},
		{"Infof", Infof, "[INFO]", []interface{}{"%s", "formatted"}},
		{"Errorf", Errorf, "[ERROR]", []interface{}{"%s", "formatted"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			NewTestLogger(&buf)

			if tt.args != nil {
				tt.logFunc(tt.args[0].(string), tt.args[1:]...)
			} else {
				tt.logFunc("test message")
			}

			expectedMessage := "test message"
			if tt.args != nil {
				expectedMessage = "formatted"
			}

			assert.Contains(t, buf.String(), tt.logLevel+" "+expectedMessage)
		})
	}
}

func TestLogging(t *testing.T) {
	tests := []struct {
		name     string
		logFunc  func(string)
		logLevel string
	}{
		{"Debug", Debug, "[DEBUG]"},
		{"Info", Info, "[INFO]"},
		{"Error", Error, "[ERROR]"},
	}

	message := "test message"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			NewTestLogger(&buf)

			tt.logFunc(message)

			assert.Contains(t, buf.String(), tt.logLevel+" "+message)
		})
	}
}
