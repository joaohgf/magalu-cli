package task

import (
	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/spf13/cobra"
)

type DoneRunner struct {
	useCase port.SaveUseCase[*domain.Task]
}

func NewDoneRunner(useCase port.SaveUseCase[*domain.Task]) *DoneRunner {
	return &DoneRunner{useCase: useCase}
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
	cmd.Printf("Task with ID %s marked as done successfully\n", updated.ID)
	return nil
}
