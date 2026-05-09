package task

import (
	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/spf13/cobra"
)

type FindRunner struct {
	useCase port.FindUseCase[*domain.Task]
	render  port.Render[*domain.Task]
}

func NewFindRunner(
	useCase port.FindUseCase[*domain.Task],
	render port.Render[*domain.Task],
) *FindRunner {
	return &FindRunner{useCase: useCase, render: render}
}

func (fr *FindRunner) Run(_ *cobra.Command, args []string) error {
	task := &domain.Task{}
	if len(args) > 0 {
		task.ID = args[0]
	}
	found, err := fr.useCase.Find(task)
	if err != nil {
		return err
	}
	err = fr.render.Render(found)
	return err
}
