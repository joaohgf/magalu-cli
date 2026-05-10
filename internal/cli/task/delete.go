package task

import (
	"github.com/joaohgf/magalu-cli/internal/core/domain"
	usecase "github.com/joaohgf/magalu-cli/internal/core/usecase/task"
	"github.com/joaohgf/magalu-cli/internal/persistence"
	"github.com/joaohgf/magalu-cli/internal/render"
	taskrender "github.com/joaohgf/magalu-cli/internal/render/task"
	handler "github.com/joaohgf/magalu-cli/internal/runner/task"
	"github.com/nanobox-io/scribble"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

const (
	taskDeleteUse   = `delete "<task_id>"`
	taskDeleteShort = "Delete a task by its ID"
	taskDeleteLong  = "Delete a task by its ID. This action is irreversible."
)

// BuildDeleteCommandHandler builds the delete subcommand runner for the task command.
func BuildDeleteCommandHandler(db *scribble.Driver, tableWriter *tablewriter.Table) *cobra.Command {
	repository := persistence.NewServiceDeleter[*domain.Task](db)
	rule := usecase.NewDelete(repository)
	view := taskrender.NewDeleteView(tableWriter)
	renderer := render.NewRender[*domain.Task](view)
	runner := handler.NewDeleteRunner(rule, renderer)
	cmd := &cobra.Command{
		Use:   taskDeleteUse,
		Short: taskDeleteShort,
		Long:  taskDeleteLong,
		Args:  cobra.ExactArgs(1),
		RunE:  runner.Run,
	}
	return cmd
}
