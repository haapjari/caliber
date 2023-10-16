package commands_test

import (
	"bytes"
	"testing"

	"github.com/haapjari/caliber/internal/pkg/commands"
	"github.com/haapjari/caliber/internal/pkg/log"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestRunCommand(t *testing.T) {
	t.Parallel()

	buf := new(bytes.Buffer)

	log.NewLogger()
	log.SetOutput(buf)

	root := commands.NewRootCommand(buf)

	t.Run("Run: Without Arguments", func(t *testing.T) {
		output, err := executeCobraCommand(root, buf, "run", "/path/to/project")
		assert.NoError(t, err)
		assert.Contains(t, output, "/path/to/project")

		buf.Reset()
	})
}

// executeCobraCommand executes a cobra command and returns the output and error.
func executeCobraCommand(cmd *cobra.Command, buf *bytes.Buffer, args ...string) (output string, err error) {
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)

	err = cmd.Execute()

	return buf.String(), err
}
