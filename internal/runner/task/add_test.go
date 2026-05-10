//go:build unit

package task

import (
	"context"
	"fmt"
	"testing"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/spf13/cobra"
)

func TestAddRunner_Run(t *testing.T) {
	t.Run("run successfully", runAddSuccessfully)
	t.Run("run with errors", func(t *testing.T) {
		t.Run("invalid estimated_done_at", runAddErrorParsingEstimatedDoneAt)
		t.Run("saving", runAddErrorSavingTask)
		t.Run("rendering", runAddErrorRender)
	})
}

func runAddSuccessfully(t *testing.T) {
	cmd := getAddCommand()
	cmd.Flags().String("priority", "medium", "")
	cmd.Flags().String("description", "This is a test task", "")
	cmd.Flags().StringSlice("tags", []string{
		"work",
	}, "Tags for the task")
	cmd.Flags().String("estimated_done_at", "2024-12-31 15:00:00", "")
	runner := NewAddRunner(new(mockSaveUseCase), new(mockRender))
	err := runner.Run(cmd, []string{"Test Task"})
	if err != nil {
		t.Fatalf("expected no errors but got: %v", err)
	}
}

func runAddErrorParsingEstimatedDoneAt(t *testing.T) {
	cmd := getAddCommand()
	cmd.Flags().String("priority", "", "")
	cmd.Flags().String("description", "", "")
	cmd.Flags().StringSlice("tags", []string{}, "")
	cmd.Flags().String("estimated_done_at", "invalid-date", "")
	runner := NewAddRunner(new(mockSaveUseCase), new(mockRender))
	err := runner.Run(cmd, []string{"Test Task"})
	if err == nil {
		t.Fatal("expected an errors but got nil")
	}
}

func runAddErrorSavingTask(t *testing.T) {
	cmd := getAddCommand()
	cmd.Flags().String("priority", "", "")
	cmd.Flags().String("description", "", "")
	cmd.Flags().StringSlice("tags", []string{}, "Tags for the task")
	cmd.Flags().String("estimated_done_at", "", "")
	runner := NewAddRunner(new(mockSaveUseCaseWithError), new(mockRender))
	err := runner.Run(cmd, []string{""})
	if err == nil {
		t.Fatal("expected an errors but got nil")
	}
}

func runAddErrorRender(t *testing.T) {
	cmd := getAddCommand()
	cmd.Flags().String("priority", "", "")
	cmd.Flags().String("description", "", "")
	cmd.Flags().StringSlice("tags", []string{}, "Tags for the task")
	cmd.Flags().String("estimated_done_at", "", "")
	runner := NewAddRunner(new(mockSaveUseCase), new(mockRenderWithError))
	err := runner.Run(cmd, []string{""})
	if err == nil {
		t.Fatal("expected an errors but got nil")
	}
}

/*
Mocks below
*/

func getAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add",
	}
	return cmd
}

type (
	mockSaveUseCase          struct{}
	mockSaveUseCaseWithError struct{}
)

func (m *mockSaveUseCase) Save(_ context.Context, task *domain.Task) (*domain.Task, error) {
	task.ID = "mock-id"
	return task, nil
}

func (m *mockSaveUseCaseWithError) Save(_ context.Context, _ *domain.Task) (*domain.Task, error) {
	return nil, fmt.Errorf("failed to save task")
}
