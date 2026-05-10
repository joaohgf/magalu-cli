package task

import (
	"github.com/joaohgf/magalu-cli/internal/core/domain"
	usecase "github.com/joaohgf/magalu-cli/internal/core/usecase/task"
	"github.com/joaohgf/magalu-cli/internal/persistence"
	"github.com/joaohgf/magalu-cli/internal/render"
	taskrender "github.com/joaohgf/magalu-cli/internal/render/task"
	handler "github.com/joaohgf/magalu-cli/internal/runner/task"
	"github.com/nanobox-io/scribble"
	"github.com/spf13/cobra"
)

const (
	taskDoneUse              = "complete"
	taskDoneShortDescription = "Mark a task as done"
	taskDoneLongDescription  = "Mark a task as done by providing its ID. \n" +
		"This will update the task's status to 'done' and set the completion timestamp."
)

// BuildTaskDoneCommandHandler builds the done subcommand runner for the task command.
func BuildTaskDoneCommandHandler(db *scribble.Driver) *cobra.Command {
	finder := persistence.NewServiceFinder[*domain.Task](db)
	saver := persistence.NewServiceSaver[*domain.Task](db)
	rule := usecase.NewDone(saver, finder)
	renderer := render.NewRender[*domain.Task](taskrender.NewDoneView())
	runner := handler.NewDoneRunner(rule, renderer)
	cmd := &cobra.Command{
		Use:   taskDoneUse,
		Short: taskDoneShortDescription,
		Long:  taskDoneLongDescription,
		RunE:  runner.Run,
	}
	return cmd
}
