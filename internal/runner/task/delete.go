package task

import (
	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/spf13/cobra"
)

// DeleteRunner is responsible for executing the logic to delete a task based on the provided command-line arguments.
type DeleteRunner struct {
	useCase port.DeleteUseCase[*domain.Task]
	render  port.Render[*domain.Task]
}

func NewDeleteRunner(
	useCase port.DeleteUseCase[*domain.Task],
	render port.Render[*domain.Task],
) *DeleteRunner {
	return &DeleteRunner{useCase: useCase, render: render}
}

// Run is the method that executes the logic for deleting a task.
// It retrieves the task ID from the command-line arguments, constructs a Task object with that ID,
// and then calls the delete use case to perform the deletion. Finally, it renders the result.
func (dr *DeleteRunner) Run(cmd *cobra.Command, args []string) error {
	var id string
	if len(args) > 0 {
		id = args[0]
	}
	task := &domain.Task{ID: id}
	ctx := cmd.Context()
	err := dr.useCase.Delete(ctx, task)
	if err != nil {
		return err
	}
	err = dr.render.Render(ctx, task)
	return err
}
