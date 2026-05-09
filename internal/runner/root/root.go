package root

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Runner struct {
}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) Run(cmd *cobra.Command, _ []string) error {
	err := cmd.Help()
	if err != nil {
		return fmt.Errorf("failed to display help: %w", err)
	}
	return nil
}

func (r *Runner) PreRun(cmd *cobra.Command, _ []string) error {
	skip := map[string]bool{
		"start":      true,
		"help":       true,
		"completion": true,
	}
	if cmd == cmd.Root() {
		return nil
	}
	if skip[cmd.Name()] {
		return nil
	}
	const bootstrapFile = "./tarefeiro/config/default.json"
	if _, err := os.Stat(bootstrapFile); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("CLI not initialized. Run `tarefeiro start` first")
	}
	return nil
}
