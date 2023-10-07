package ui

import "fmt"

// CommandLineInterface defines a new interface.
type CommandLineInterface struct {
	version int
}

// New is a constructor.
func NewCommandLineInterface() *CommandLineInterface {
	return &CommandLineInterface{
		version: 1,
	}
}

// Create returns a new command-line interface.
func Create() {
	fmt.Println("Command-Line Interface")
	c := NewCommandLineInterface()
	fmt.Println(c.version)
}
