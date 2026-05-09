package task

import (
	"fmt"
	"time"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/core/enum"
	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/oklog/ulid/v2"
	"github.com/spf13/cobra"
)

type AddRunner struct {
	useCase port.SaveUseCase[*domain.Task]
}

func NewAddRunner(useCase port.SaveUseCase[*domain.Task]) *AddRunner {
	return &AddRunner{useCase: useCase}
}

// Run is the method that executes the logic for adding a new task.
// It generates a unique ID for the task, captures the current timestamp,
// and constructs a new Task object using the provided command-line arguments and flags.
func (cr *AddRunner) Run(cmd *cobra.Command, args []string) error {
	id := ulid.Make()
	ts := id.Time()
	createdAt := ulid.Time(ts)
	tags, _ := cmd.Flags().GetStringSlice("tags")
	var title string
	if len(args) > 0 {
		title = args[0]
	}
	task := domain.NewTask(
		id.String(),
		title,
		cmd.Flag("description").Value.String(),
		enum.PriorityOf(cmd.Flag("priority").Value.String()),
		createdAt,
		tags...,
	)
	estimatedDoneAt := cmd.Flag("estimated_done_at").Value.String()
	if estimatedDoneAt != "" {
		parsedTime, err := time.Parse(time.DateTime, estimatedDoneAt)
		if err != nil {
			return fmt.Errorf("invalid date format: %w", err)
		}
		task.EstimatedDoneAt = &parsedTime
	}
	saved, err := cr.useCase.Save(task)
	if err != nil {
		return err
	}
	cmd.Printf("Task created with ID: %s\n", saved.ID)
	return nil
}
