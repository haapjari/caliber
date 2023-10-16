// Package commands provides the commands for the application.
package commands

import (
	"io"

	"github.com/haapjari/caliber/internal/pkg/log"
	"github.com/spf13/cobra"
)

// NewRunCommand returns a new run command.
//
//nolint:gomnd
func NewRunCommand(out io.Writer) *cobra.Command {
	use := "run <projectPath>"
	short := "Run"

	// TODO caliber -run path=<path_to_project>
	// TODO caliber -run path=<path_to_project> load=<path_to_config>
	// TODO caliber -run path=<path_to_project> output=<output_file_path>.json
	// TODO caliber -run path=<path_to_project> load=<path_to_config> output=<output_file_path>.json
	// TODO caliber -run path=<path_to_project> output=<output_file_path>.json load=<path_to_config>

	// TODO
	logic := func(cmd *cobra.Command, args []string) {
		projectPath := args[0]
		log.Infof("Project Path: %v", projectPath)
	}

	return &cobra.Command{
		Use:   use,
		Short: short,
		Args:  cobra.RangeArgs(1, 3),
		Run:   logic,
	}
}
