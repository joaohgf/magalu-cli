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
	taskDeleteUse   = `delete "<task_id>"`
	taskDeleteShort = "Delete a task by its ID"
	taskDeleteLong  = "Delete a task by its ID. This action is irreversible."
)

// BuildDeleteCommandHandler builds the delete subcommand runner for the task command.
func BuildDeleteCommandHandler(db *scribble.Driver) *cobra.Command {
	repository := persistence.NewServiceDeleter[*domain.Task](db)
	rule := usecase.NewDelete(repository)
	runner := handler.NewDeleteRunner(rule)
	cmd := &cobra.Command{
		Use:   taskDeleteUse,
		Short: taskDeleteShort,
		Long:  taskDeleteLong,
		Args:  cobra.ExactArgs(1),
		RunE:  runner.Run,
	}
	return cmd
}
