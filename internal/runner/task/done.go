package task

import (
	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/spf13/cobra"
)

// DoneRunner is responsible for executing the logic to mark a task as done based on the provided command-line arguments.
type DoneRunner struct {
	useCase port.SaveUseCase[*domain.Task]
	render  port.Render[*domain.Task]
}

func NewDoneRunner(
	useCase port.SaveUseCase[*domain.Task],
	render port.Render[*domain.Task],
) *DoneRunner {
	return &DoneRunner{useCase: useCase, render: render}
}

// Run is the method that executes the logic for marking a task as done.
// It retrieves the task ID from the command-line arguments, constructs a Task object with that ID,
// and then calls the save use case to update the task's status to done. Finally, it renders the result.
func (dr *DoneRunner) Run(cmd *cobra.Command, args []string) error {
	var id string
	if len(args) > 0 {
		id = args[0]
	}
	task := &domain.Task{ID: id}
	ctx := cmd.Context()
	updated, err := dr.useCase.Save(ctx, task)
	if err != nil {
		return err
	}
	err = dr.render.Render(ctx, updated)
	return err
}
