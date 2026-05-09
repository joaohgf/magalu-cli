package task

import (
	"github.com/joaohgf/magalu-cli/internal/core/domain"
	usecase "github.com/joaohgf/magalu-cli/internal/core/usecase/task"
	"github.com/joaohgf/magalu-cli/internal/persistence"
	render "github.com/joaohgf/magalu-cli/internal/render/task"
	handler "github.com/joaohgf/magalu-cli/internal/runner/task"
	"github.com/nanobox-io/scribble"
	"github.com/spf13/cobra"
)

const (
	taskShowUse              = "show"
	taskShowShortDescription = "Show details of a task"
	taskShowLongDescription  = "Show details of a task by its ID. \n" +
		"You can also filter tasks by status and priority."
)

// BuildTaskShowCommandHandler builds the show subcommand runner for the task command.
func BuildTaskShowCommandHandler(db *scribble.Driver) *cobra.Command {
	repository := persistence.NewServiceFinder[*domain.Task](db)
	rule := usecase.NewFind(repository)
	renderer := render.NewDetail(nil)
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
