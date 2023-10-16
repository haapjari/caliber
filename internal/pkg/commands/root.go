// Package commands provides the commands for the application.
package commands

import (
	"io"

	"github.com/spf13/cobra"
)

// NewRootCommand returns a new root command.
func NewRootCommand(out io.Writer) *cobra.Command {
	rootCmd := &cobra.Command{Use: "caliber"}
	rootCmd.AddCommand(NewRunCommand(out))

	return rootCmd
}
