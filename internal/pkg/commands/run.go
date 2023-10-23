// Package commands provides the commands for the application.
package commands

import (
	"io"

	"github.com/haapjari/caliber/internal/pkg/log"
	"github.com/haapjari/caliber/internal/pkg/utils"
	"github.com/spf13/cobra"
)

// NewRunCommand returns a new run command.
//
//nolint:gomnd
//nolint:godox
func NewRunCommand(out io.Writer) *cobra.Command {
	use := "run <projectPath>"
	short := "run"

	// TODO caliber -run path=<path_to_project> load=<path_to_config> output=<output_file_path>.json
	// TODO caliber -run path=<path_to_project> output=<output_file_path>.json load=<path_to_config>

	logic := func(cmd *cobra.Command, args []string) {
		// Case: caliber run path=<path_to_project>
		if len(args) == 1 {
			path := args[0]

			absolutePath, err := utils.ReturnAbsolutePath(path)
			if err != nil {
				log.Errorf("unable to read absolute path: %w", err)
				return
			}

			log.Info(absolutePath)

			// TODO: Call analyzer for "absolutePath" variable.
		}
	}

	return &cobra.Command{
		Use:   use,
		Short: short,
		Args:  cobra.RangeArgs(1, 3),
		Run:   logic,
	}
}
