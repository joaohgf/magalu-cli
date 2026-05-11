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
	taskUpdateUse   = `edit "<task_id>"`
	taskUpdateShort = "Update a task by its ID"
	taskUpdateLong  = "Update a task by its ID. You can update the title, description, priority, estimated done date, and tags."
)

// BuildTaskUpdateCommandHandler builds the update subcommand runner for the task command.
func BuildTaskUpdateCommandHandler(db *scribble.Driver, tableWriter *tablewriter.Table) (*cobra.Command, error) {
	saver := persistence.NewServiceSaver[*domain.Task](db)
	finder := persistence.NewServiceFinder[*domain.Task](db)
	rule := usecase.NewUpdate(saver, finder)
	view := taskrender.NewUpdateView(tableWriter)
	renderer := render.NewRender[*domain.Task](view)
	runner := handler.NewUpdateRunner(rule, renderer)
	cmd, err := buildTaskCommandHandler(
		taskUpdateUse, taskUpdateShort, taskUpdateLong, runner,
		withArgs(cobra.ExactArgs(1)),
		withPriorityFlag(), withDescriptionFlag(), withEstimatedDoneDateFlag(), withTagsFlag(),
	)
	return cmd, err
}
