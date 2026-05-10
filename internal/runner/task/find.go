package task

import (
	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/spf13/cobra"
)

// FindRunner is responsible for executing the logic to find a task based on the provided command-line arguments.
type FindRunner struct {
	useCase port.FindUseCase[*domain.Task]
	render  port.Render[*domain.Task]
}

func NewFindRunner(
	useCase port.FindUseCase[*domain.Task],
	render port.Render[*domain.Task],
) *FindRunner {
	return &FindRunner{useCase: useCase, render: render}
}

// Run is the method that executes the logic for finding a task.
// It retrieves the task ID from the command-line arguments, constructs a Task object with that ID,
// and then calls the find use case to perform the search. Finally, it renders the result.
func (fr *FindRunner) Run(cmd *cobra.Command, args []string) error {
	task := &domain.Task{}
	if len(args) > 0 {
		task.ID = args[0]
	}
	ctx := cmd.Context()
	found, err := fr.useCase.Find(ctx, task)
	if err != nil {
		return err
	}
	err = fr.render.Render(ctx, found)
	return err
}
