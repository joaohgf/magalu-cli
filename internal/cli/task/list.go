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
	taskListUse              = "list"
	taskListShortDescription = "List all tasks"
	taskListLongDescription  = "List all tasks with their details. \n" +
		"You can also filter tasks by status and priority."
)

// BuildTaskListCommandHandler builds the list subcommand runner for the task command.
func BuildTaskListCommandHandler(db *scribble.Driver, tableWriter *tablewriter.Table) *cobra.Command {
	repository := persistence.NewServiceLister[*domain.Task, *domain.TaskFilter](db)
	rule := usecase.NewList(repository)
	view := rendertask.NewList(tableWriter)
	renderer := render.NewRender[*domain.TaskFilter](view)
	runner := handler.NewListRunner(rule, renderer)
	cmd := buildTaskCommandHandler(
		taskListUse, taskListShortDescription, taskListLongDescription, runner,
		withDescriptionFlag(), withStatusFlag(), withPriorityFlag(),
		withStringField("size", "S", "Number of tasks to return"),
		withStringField("page", "P", "Page number for pagination"),
		withStringField("title", "t", "Title filter (contains)"),
		withBoolField("interactive", "i", false, "Enable interactive selection after table rendering"),
	)
	return cmd
}
