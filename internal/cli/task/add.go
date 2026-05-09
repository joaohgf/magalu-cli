package task

import (
	"log"
	"os"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	usecase "github.com/joaohgf/magalu-cli/internal/core/usecase/task"
	"github.com/joaohgf/magalu-cli/internal/persistence"
	handler "github.com/joaohgf/magalu-cli/internal/runner/task"
	"github.com/nanobox-io/scribble"
	"github.com/spf13/cobra"
)

const (
	taskAddUse              = `add "Title"`
	taskAddShortDescription = "Add a new task"
	taskAddLongDescription  = "Add a new task with a title, description, priority and status. \n" +
		"You can also specify the due date for the task."
)

// BuildTaskAddCommandHandler builds the add subcommand runner for the task command.
func BuildTaskAddCommandHandler(db *scribble.Driver) *cobra.Command {
	repository := persistence.NewServiceSaver[*domain.Task](db)
	rule := usecase.NewAdd(repository)
	runner := handler.NewAddRunner(rule)
	cmd := &cobra.Command{
		Use:   taskAddUse,
		Short: taskAddShortDescription,
		Long:  taskAddLongDescription,
		Args:  cobra.ExactArgs(1),
		RunE:  runner.Run,
	}
	cmd.Flags().StringP("description", "d", "", "Description of the task")
	cmd.Flags().StringP("priority", "p", "medium", "Priority of the task (low, medium, high)")
	cmd.Flags().StringSliceP("tags", "T", []string{}, "Comma-separated list of tags for the task")
	cmd.Flags().StringP("estimated_done_at", "e", "", "Estimated done date for the task (YYYY-MM-DD HH:MM:SS format)")
	logger := log.Default()
	logger.SetOutput(os.Stdout)
	if err := cmd.MarkFlagRequired("priority"); err != nil {
		logger.Printf("failed to mark flag as required: %v", err)
		return nil
	}
	return cmd
}
