package ui_test

import (
	"os"
	"testing"

	"github.com/haapjari/caliber/pkg/ui"
)

func TestListener(t *testing.T) {
	ui.New().Listen()
	os.Args = append(os.Args, "--help")

	// Check that the buffer is not empty
	// assert.NotEmpty(t, buf.String(), "Log should not be empty")

	// (Optional) Check that the buffer contains a specific log message
	// assert.Contains(t, buf.String(), "This is a log message.")
}
