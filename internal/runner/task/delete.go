package task

import (
	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/spf13/cobra"
)

type DeleteRunner struct {
	useCase port.DeleteUseCase[*domain.Task]
	render  port.Render[*domain.Task]
}

func NewDeleteRunner(
	useCase port.DeleteUseCase[*domain.Task],
	render port.Render[*domain.Task],
) *DeleteRunner {
	return &DeleteRunner{useCase: useCase, render: render}
}

func (dr *DeleteRunner) Run(cmd *cobra.Command, args []string) error {
	var id string
	if len(args) > 0 {
		id = args[0]
	}
	if id != "" {
	}
	task := &domain.Task{ID: id}
	ctx := cmd.Context()
	err := dr.useCase.Delete(ctx, task)
	if err != nil {
		return err
	}
	err = dr.render.Render(ctx, task)
	return err
}
