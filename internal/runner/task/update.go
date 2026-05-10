package task

import (
	"fmt"
	"time"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/spf13/cobra"
)

// UpdateRunner is responsible for executing the logic to update a task based on the provided command-line arguments.
type UpdateRunner struct {
	useCase port.SaveUseCase[*domain.Task]
	render  port.Render[*domain.Task]
}

func NewUpdateRunner(
	useCase port.SaveUseCase[*domain.Task],
	render port.Render[*domain.Task],
) *UpdateRunner {
	return &UpdateRunner{useCase: useCase, render: render}
}

// Run is the method that executes the logic for updating a task.
// It retrieves the task ID from the command-line arguments, constructs a Task object with that ID,
// and then calls the save use case to perform the update. Finally, it renders the result.
func (ur *UpdateRunner) Run(cmd *cobra.Command, args []string) error {
	var id string
	if len(args) > 0 {
		id = args[0]
	}
	task := &domain.Task{ID: id}
	task.Title = cmd.Flag("title").Value.String()
	task.Description = cmd.Flag("description").Value.String()
	tags, _ := cmd.Flags().GetStringSlice("tags")
	task.Tags = tags
	priority := cmd.Flag("priority").Value.String()
	if priority != "" {
		task.Priority = enum.PriorityOf(priority)
	}
	estimatedAt := cmd.Flag("estimated_done_at").Value.String()
	if estimatedAt != "" {
		parsedTime, err := time.Parse(time.DateTime, estimatedAt)
		if err != nil {
			return fmt.Errorf("errors parsing estimated done at, please use the format: %s", time.DateTime)
		}
		task.EstimatedDoneAt = &parsedTime
	}
	ctx := cmd.Context()
	updated, err := ur.useCase.Save(ctx, task)
	if err != nil {
		return err
	}
	err = ur.render.Render(ctx, updated)
	return err
}
