package task

import (
	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/spf13/cobra"
)

type DoneRunner struct {
	useCase port.SaveUseCase[*domain.Task]
	render  port.Render[*domain.Task]
}

func NewDoneRunner(
	useCase port.SaveUseCase[*domain.Task],
	render port.Render[*domain.Task],
) *DoneRunner {
	return &DoneRunner{useCase: useCase, render: render}
}

func (dr *DoneRunner) Run(cmd *cobra.Command, args []string) error {
	var id string
	if len(args) > 0 {
		id = args[0]
	}
	task := &domain.Task{ID: id}
	ctx := cmd.Context()
	updated, err := dr.useCase.Save(ctx, task)
	if err != nil {
		return err
	}
	err = dr.render.Render(ctx, updated)
	return err
}
