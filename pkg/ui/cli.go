package ui

import (
	"os"

	"github.com/haapjari/caliber/pkg/log"
)

type UserInterface interface {
	Close() error
	Listener()
}

// CommandLineInterface defines a new interface.
type CommandLineInterface struct {
	version int
}

func New() *CommandLineInterface {
	return &CommandLineInterface{
		version: 1,
	}
}

func init() {}

func (c *CommandLineInterface) Listen() {
	for _, arg := range os.Args {
		if arg == "-h" || arg == "--help" {
			log.Info("help")
		}
	}
}
