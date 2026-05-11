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
	taskAddUse              = `add "Title"`
	taskAddShortDescription = "Add a new task"
	taskAddLongDescription  = "Add a new task with a title, description, priority and status. \n" +
		"You can also specify the due date for the task."
)

// BuildTaskAddCommandHandler builds the add subcommand runner for the task command.
func BuildTaskAddCommandHandler(db *scribble.Driver, table *tablewriter.Table) (*cobra.Command, error) {
	repository := persistence.NewServiceSaver[*domain.Task](db)
	rule := usecase.NewAdd(repository)
	view := taskrender.NewAddView(table)
	renderer := render.NewRender[*domain.Task](view)
	runner := handler.NewAddRunner(rule, renderer)
	cmd, err := buildTaskCommandHandler(
		taskAddUse, taskAddShortDescription, taskAddLongDescription, runner,
		withArgs(cobra.ExactArgs(1)),
		withPriorityFlag(), withDescriptionFlag(), withEstimatedDoneDateFlag(), withTagsFlag(),
	)
	return cmd, err
}
