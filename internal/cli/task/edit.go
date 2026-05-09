package task

import (
	"github.com/joaohgf/magalu-cli/internal/core/domain"
	usecase "github.com/joaohgf/magalu-cli/internal/core/usecase/task"
	"github.com/joaohgf/magalu-cli/internal/persistence"
	handler "github.com/joaohgf/magalu-cli/internal/runner/task"
	"github.com/nanobox-io/scribble"
	"github.com/spf13/cobra"
)

const (
	taskUpdateUse   = `edit "<task_id>"`
	taskUpdateShort = "Update a task by its ID"
	taskUpdateLong  = "Update a task by its ID. You can update the title, description, priority, estimated done date, and tags."
)

// BuildUpdateCommandHandler builds the update subcommand runner for the task command.
func BuildUpdateCommandHandler(db *scribble.Driver) *cobra.Command {
	saver := persistence.NewServiceSaver[*domain.Task](db)
	finder := persistence.NewServiceFinder[*domain.Task](db)
	rule := usecase.NewUpdate(saver, finder)
	runner := handler.NewUpdateRunner(rule)
	cmd := &cobra.Command{
		Use:   taskUpdateUse,
		Short: taskUpdateShort,
		Long:  taskUpdateLong,
		Args:  cobra.ExactArgs(1),
		RunE:  runner.Run,
	}
	cmd.Flags().StringP("title", "t", "", "Title of the task")
	cmd.Flags().StringP("description", "d", "", "Description of the task")
	cmd.Flags().StringP("priority", "p", "", "Priority of the task (low, medium, high)")
	cmd.Flags().StringP("estimated_done_at", "e", "", "Estimated done date for the task (YYYY-MM-DD HH:MM:SS format)")
	cmd.Flags().StringSliceP("tags", "T", []string{}, "Comma-separated list of tags for the task")
	return cmd
}
