package task

import (
	"github.com/joaohgf/magalu-cli/internal/core/domain"
	enum "github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/spf13/cobra"
)

type ListRunner struct {
	useCase port.FindAllUseCase[*domain.Task]
	render  port.Render[*domain.Task]
}

func NewListRunner(
	useCase port.FindAllUseCase[*domain.Task],
	render port.Render[*domain.Task],
) *ListRunner {
	return &ListRunner{useCase: useCase, render: render}
}

func (lr *ListRunner) Run(cmd *cobra.Command, _ []string) error {
	filter := &domain.Task{}
	filter.Title = cmd.Flag("title").Value.String()
	filter.Description = cmd.Flag("description").Value.String()
	if status := enum.StatusOf(cmd.Flag("status").Value.String()); status != enum.StatusUnknown {
		filter.Status = status
	}
	if priority := enum.PriorityOf(cmd.Flag("priority").Value.String()); priority != enum.PriorityUnknown {
		filter.Priority = priority
	}
	tasks, err := lr.useCase.All(filter)
	if err != nil {
		return err
	}
	err = lr.render.Render(tasks...)
	return err
}
