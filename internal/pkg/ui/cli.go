// Package ui provides the user interface for the application.
package ui

import (
	"flag"
	"strings"
	"sync"

	"github.com/haapjari/caliber/internal/pkg/config"
	"github.com/haapjari/caliber/internal/pkg/log"
)

// UserInterface is the interface for the user interface.
type UserInterface interface {
	Close() error
	Listener()
}

// CommandLineInterface is the command line interface for the application.
type CommandLineInterface struct {
	sync.Mutex

	cfg *config.Config
}

// New returns a new CommandLineInterface.
func New() *CommandLineInterface {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	commandLineInterface := &CommandLineInterface{
		Mutex: sync.Mutex{},
		cfg:   cfg,
	}

	return commandLineInterface
}

func init() {}

// Listen listens for commands from the user.
func (c *CommandLineInterface) Listen() {
	var (
		help    = flag.Bool("help", false, "Display Help")
		version = flag.Bool("version", false, "Display Version")
		run     = flag.Bool("run", false, "Run the command")
	)

	flag.Parse()

	if *help {
		log.Info("Displaying Help...")
		flag.PrintDefaults()
		return
	}

	if *version {
		log.Infof("Caliber Version: %v", c.GetVersion())
		return
	}

	if *run {
		args := flag.Args()
		if len(args) < 2 { // We ignore the 'run' command, so at least 2 arguments are needed
			log.Error("Error: Project path is required for the run command.")
			return
		}

		projectPath := args[1]
		var loadPath, outputPath string

		if len(args) > 2 {
			secondArg := args[2]
			if strings.HasSuffix(secondArg, ".json") {
				outputPath = secondArg
			} else {
				loadPath = secondArg
			}
		}

		if len(args) > 3 {
			outputPath = args[3]
		}

		// Handle the paths
		log.Infof("Project Path: %v", projectPath)
		if loadPath != "" {
			log.Infof("Load Path: %v", loadPath)
		}
		if outputPath != "" {
			log.Infof("Output Path: %v", outputPath)
		}
		return
	}

	log.Error("Command Not Supported")
}

// GetVersion returns the version of the application.
func (c *CommandLineInterface) GetVersion() string {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	return c.cfg.Version
}
