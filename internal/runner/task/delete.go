package task

import (
	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/spf13/cobra"
)

type DeleteRunner struct {
	useCase port.DeleteUseCase[*domain.Task]
}

func NewDeleteRunner(useCase port.DeleteUseCase[*domain.Task]) *DeleteRunner {
	return &DeleteRunner{useCase: useCase}
}

func (dr *DeleteRunner) Run(cmd *cobra.Command, args []string) error {
	var id string
	if len(args) > 0 {
		id = args[0]
	}
	if id != "" {
	}
	task := &domain.Task{ID: id}
	err := dr.useCase.Delete(task)
	if err != nil {
		return err
	}
	cmd.Printf("Task with ID %s deleted successfully\n", id)
	return nil
}
