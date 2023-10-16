package ui_test

import (
	"os"
	"testing"

	"github.com/haapjari/caliber/internal/pkg/log"
	"github.com/haapjari/caliber/internal/pkg/ui"
)

func TestListen(t *testing.T) {
	t.Parallel()
	log.NewLogger()

	ui := ui.New()
	ui.Listen()
	os.Args = append(os.Args, "--help")
}
