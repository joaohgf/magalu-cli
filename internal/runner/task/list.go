package task

import (
	"fmt"
	"strconv"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/joaohgf/magalu-cli/internal/errors"
	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/spf13/cobra"
)

// ListRunner is responsible for executing the logic to list tasks based on the provided command-line arguments and flags.
type ListRunner struct {
	useCase port.FindAllUseCase[*domain.TaskFilter]
	render  port.Render[*domain.TaskFilter]
}

func NewListRunner(
	useCase port.FindAllUseCase[*domain.TaskFilter],
	render port.Render[*domain.TaskFilter],
) *ListRunner {
	return &ListRunner{useCase: useCase, render: render}
}

// Run is the method that executes the logic for listing tasks.
// It constructs a TaskFilter object using the provided command-line arguments and flags,
// and then calls the find all use case to perform the search. Finally, it renders the result.
func (lr *ListRunner) Run(cmd *cobra.Command, _ []string) error {
	task := &domain.Task{}
	task.Title = cmd.Flag("title").Value.String()
	task.Description = cmd.Flag("description").Value.String()
	if status := enum.StatusOf(cmd.Flag("status").Value.String()); status != enum.StatusUnknown {
		task.Status = status
	}
	if priority := enum.PriorityOf(cmd.Flag("priority").Value.String()); priority != enum.PriorityUnknown {
		task.Priority = priority
	}
	filter := domain.NewTaskFilter(task)
	size := cmd.Flag("size").Value.String()
	if size != "" {
		parsedSize, err := strconv.ParseInt(size, 10, 64)
		if err != nil {
			return errors.Invalid(fmt.Sprintf("size: %v", size))
		}
		filter.SetSize(int(parsedSize))
	}
	page := cmd.Flag("page").Value.String()
	if page != "" {
		parsedPage, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			return errors.Invalid(fmt.Sprintf("page: %v", page))
		}
		filter.SetPage(int(parsedPage))
	}
	ctx := cmd.Context()
	found, err := lr.useCase.All(ctx, filter)
	if err != nil {
		return err
	}
	err = lr.render.Render(ctx, found)
	return err
}
