//go:build unit

package task

import (
	"context"
	"fmt"
	"testing"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/spf13/cobra"
)

func TestUpdateRunner(t *testing.T) {
	t.Run("run successfully", runUpdateSuccessfully)
	t.Run("run with errors", func(t *testing.T) {
		t.Run("run with valid estimated_done_at", runUpdateErrorParsingEstimatedDoneAt)
		t.Run("errors saving task", runUpdateErrorSavingTask)
	})
}

func runUpdateSuccessfully(t *testing.T) {
	cmd := getUpdateCommand()
	cmd.Flags().String("title", "Updated Task", "")
	cmd.Flags().String("description", "Updated Description", "")
	cmd.Flags().StringSlice("tags", []string{"tag1", "tag2"}, "")
	cmd.Flags().String("priority", "high", "")
	cmd.Flags().String("estimated_done_at", "2024-12-31 15:00:00", "")
	useCase := &mockUpdateUseCase{}
	runner := NewUpdateRunner(useCase)
	err := runner.Run(cmd, []string{"123"})
	if err != nil {
		t.Fatalf("expected no errors, got %v", err)
	}
}

func runUpdateErrorParsingEstimatedDoneAt(t *testing.T) {
	cmd := getUpdateCommand()
	cmd.Flags().String("title", "Updated Task", "")
	cmd.Flags().String("description", "Updated Description", "")
	cmd.Flags().StringSlice("tags", []string{"tag1", "tag2"}, "")
	cmd.Flags().String("priority", "high", "")
	cmd.Flags().String("estimated_done_at", "invalid-date", "")
	useCase := &mockUpdateUseCase{}
	runner := NewUpdateRunner(useCase)
	err := runner.Run(cmd, []string{"123"})
	if err == nil {
		t.Fatal("expected errors, got nil")
	}
}

func runUpdateErrorSavingTask(t *testing.T) {
	cmd := getUpdateCommand()
	useCase := &mockUpdateUseCaseWithError{}
	runner := NewUpdateRunner(useCase)
	cmd.Flags().String("title", "Updated Task", "")
	cmd.Flags().String("description", "Updated Description", "")
	cmd.Flags().StringSlice("tags", []string{"tag1", "tag2"}, "")
	cmd.Flags().String("priority", "high", "")
	cmd.Flags().String("estimated_done_at", "2024-12-31 15:00:00", "")
	err := runner.Run(cmd, []string{"123"})
	if err == nil {
		t.Fatal("expected errors, got nil")
	}
}

/*
Mocks below
*/

func getUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "edit",
	}
	return cmd
}

type (
	mockUpdateUseCase          struct{}
	mockUpdateUseCaseWithError struct{}
)

func (m *mockUpdateUseCase) Save(_ context.Context, task *domain.Task) (*domain.Task, error) {
	task.Title = "Updated Task"
	return task, nil
}

func (m *mockUpdateUseCaseWithError) Save(_ context.Context, task *domain.Task) (*domain.Task, error) {
	return nil, fmt.Errorf("errors saving task")
}
