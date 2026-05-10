package task

import (
	"github.com/joaohgf/magalu-cli/internal/core/domain"
	usecase "github.com/joaohgf/magalu-cli/internal/core/usecase/task"
	"github.com/joaohgf/magalu-cli/internal/persistence"
	"github.com/joaohgf/magalu-cli/internal/render"
	rendertask "github.com/joaohgf/magalu-cli/internal/render/task"
	handler "github.com/joaohgf/magalu-cli/internal/runner/task"
	"github.com/nanobox-io/scribble"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

const (
	taskShowUse              = "show"
	taskShowShortDescription = "Show details of a task"
	taskShowLongDescription  = "Show details of a task by its ID. \n" +
		"You can also filter tasks by status and priority."
)

// BuildTaskShowCommandHandler builds the show subcommand runner for the task command.
func BuildTaskShowCommandHandler(db *scribble.Driver, tableWriter *tablewriter.Table) *cobra.Command {
	repository := persistence.NewServiceFinder[*domain.Task](db)
	rule := usecase.NewFind(repository)
	view := rendertask.NewDetailView(tableWriter)
	renderer := render.NewRender[*domain.Task](view)
	runner := handler.NewFindRunner(rule, renderer)
	cmd := &cobra.Command{
		Use:   taskShowUse,
		Short: taskShowShortDescription,
		Long:  taskShowLongDescription,
		Args:  cobra.ExactArgs(1),
		RunE:  runner.Run,
	}
	return cmd
}
