package root

import (
	"context"
	"fmt"

	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/spf13/cobra"
)

// Runner is responsible for executing the root command of the CLI application.
type Runner struct{}

func NewRunner() *Runner {
	return &Runner{}
}

// Run executes the root command, which displays the help message for the CLI application.
func (r *Runner) Run(cmd *cobra.Command, _ []string) (err error) {
	err = cmd.Help()
	if err != nil {
		return fmt.Errorf("failed to display help: %w", err)
	}
	return nil
}

// PreRun is a hook that runs before every command execution
// It checks for the presence of the output flag and updates the command's context accordingly.
func (r *Runner) PreRun(cmd *cobra.Command, _ []string) error {
	output := cmd.Flag(enum.OutputDefaultKey.ToLower()).Value.String()
	ctx := cmd.Context()
	if output != "" {
		ctx = context.WithValue(ctx, enum.OutputDefaultKey.String(), enum.OutputOf(output))
	}
	cmd.SetContext(ctx)
	return nil
}
