package task

import (
	"fmt"
	"time"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/core/enum"
	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/spf13/cobra"
)

type UpdateRunner struct {
	useCase port.SaveUseCase[*domain.Task]
}

func NewUpdateRunner(useCase port.SaveUseCase[*domain.Task]) *UpdateRunner {
	return &UpdateRunner{useCase: useCase}
}

func (ur *UpdateRunner) Run(cmd *cobra.Command, args []string) error {
	var id string
	if len(args) > 0 {
		id = args[0]
	}
	task := &domain.Task{ID: id}
	task.Title = cmd.Flag("title").Value.String()
	task.Description = cmd.Flag("description").Value.String()
	tags, err := cmd.Flags().GetStringSlice("tags")
	if err != nil {
		return fmt.Errorf("failed to get tags: %w", err)
	}
	task.Tags = tags
	priority := cmd.Flag("priority").Value.String()
	if priority != "" {
		task.Priority = enum.PriorityOf(priority)
	}
	estimatedAt := cmd.Flag("estimated_done_at").Value.String()
	if estimatedAt != "" {
		parsedTime, err := time.Parse(time.DateTime, estimatedAt)
		if err != nil {
			return fmt.Errorf("error parsing estimated done at, please use the format: %s", time.DateTime)
		}
		task.EstimatedDoneAt = &parsedTime
	}
	updated, err := ur.useCase.Save(task)
	if err != nil {
		return err
	}
	cmd.Printf("Task with ID %s updated successfully\n", updated.ID)
	return nil
}
