// Package main provides the entry point for the application.
package main

import (
	"os"

	"github.com/haapjari/caliber/internal/pkg/commands"
	"github.com/haapjari/caliber/internal/pkg/log"
	"github.com/spf13/cobra"
)

func main() {
	log.NewLogger()

	root := &cobra.Command{Use: "caliber"}

	run := commands.NewRunCommand(os.Stdout)

	root.AddCommand(run)

	if err := root.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
