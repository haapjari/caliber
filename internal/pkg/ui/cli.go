// Package ui provides the user interface for the application.
package ui

import (
	"flag"
	"sync"

	"github.com/haapjari/caliber/internal/pkg/log"
)

// UserInterface is the interface for the user interface.
type UserInterface interface {
	Close() error
	Listener()
}

// CommandLineInterface is the command line interface for the application.
type CommandLineInterface struct {
	version string
	mu      sync.Mutex
}

// New returns a new CommandLineInterface.
func New() *CommandLineInterface {
	return &CommandLineInterface{
		version: "0.0.1", // TODO: Get version from git tag
		mu:      sync.Mutex{},
	}
}

func init() {}

// Listen listens for commands from the user.
func (c *CommandLineInterface) Listen() {
	var (
		help    = flag.Bool("help", false, "Display Help")
		version = flag.Bool("version", false, "Display Version")
		load    = flag.Bool("load", false, "Load Config")
		update  = flag.Bool("update", false, "Update Definitions")
	)

	flag.Parse()

	// TODO: Add support for commands
	if *help {
		log.Info("TODO")

		flag.PrintDefaults()

		return
	}

	// TODO: Add support for commands
	if *version {
		log.Infof("Caliber Version: %v", c.GetVersion())

		return
	}

	// TODO: Add support for commands
	if *load {
		log.Info("TODO")

		return
	}

	// TODO: Add support for commands
	if *update {
		log.Info("TODO")

		return
	}

	log.Info("Command Not Supported")
}

// GetVersion returns the version of the application.
func (c *CommandLineInterface) GetVersion() string {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.version
}
