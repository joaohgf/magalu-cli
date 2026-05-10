package task

import (
	"github.com/joaohgf/magalu-cli/internal/core/domain"
	usecase "github.com/joaohgf/magalu-cli/internal/core/usecase/task"
	"github.com/joaohgf/magalu-cli/internal/persistence"
	"github.com/joaohgf/magalu-cli/internal/render"
	rendertask "github.com/joaohgf/magalu-cli/internal/render/task"
	handler "github.com/joaohgf/magalu-cli/internal/runner/task"
	"github.com/nanobox-io/scribble"
	"github.com/spf13/cobra"
)

const (
	taskListUse              = "list"
	taskListShortDescription = "List all tasks"
	taskListLongDescription  = "List all tasks with their details. \n" +
		"You can also filter tasks by status and priority."
)

// BuildTaskListCommandHandler builds the list subcommand runner for the task command.
func BuildTaskListCommandHandler(db *scribble.Driver) *cobra.Command {
	repository := persistence.NewServiceLister[*domain.Task, *domain.TaskFilter](db)
	rule := usecase.NewList(repository)
	view := rendertask.NewList()
	renderer := render.NewRender[*domain.TaskFilter](view)
	runner := handler.NewListRunner(rule, renderer)
	cmd := &cobra.Command{
		Use:   taskListUse,
		Short: taskListShortDescription,
		Long:  taskListLongDescription,
		RunE:  runner.Run,
	}
	cmd.Flags().StringP("title", "t", "", "Title filter (contains)")
	cmd.Flags().StringP("description", "d", "", "Description filter (contains)")
	cmd.Flags().StringP("status", "s", "", "Status of the task (in_progress, done)")
	cmd.Flags().StringP("priority", "p", "", "Priority of the task (low, medium, high)")
	cmd.Flags().StringP("size", "S", "", "Number of tasks to return")
	cmd.Flags().StringP("page", "P", "", "Page number for pagination")
	return cmd
}
