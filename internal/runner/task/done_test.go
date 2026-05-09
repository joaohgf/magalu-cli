//go:build unit

package task

import (
	"fmt"
	"testing"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/spf13/cobra"
)

func TestDoneRunner(t *testing.T) {
	t.Run("run successfully", runDoneSuccessfully)
	t.Run("run with error", runDoneWithError)
}

func runDoneSuccessfully(t *testing.T) {
	cmd := getCompleteCommand()
	runner := NewDoneRunner(new(mockSaveUseCase))
	err := runner.Run(cmd, []string{"123"})
	if err != nil {
		t.Fatalf("expected no error but got: %v", err)
	}
}

func runDoneWithError(t *testing.T) {
	cmd := getCompleteCommand()
	runner := NewDoneRunner(new(mockSaveUseCaseWithError))
	err := runner.Run(cmd, []string{"123"})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

/*
Mocks below
*/

func getCompleteCommand() *cobra.Command {
	return &cobra.Command{
		Use: "complete",
	}
}

type (
	mockSaveUseCase          struct{}
	mockSaveUseCaseWithError struct{}
)

func (m *mockSaveUseCase) Save(task *domain.Task) (*domain.Task, error) {
	return task, nil
}

func (m *mockSaveUseCaseWithError) Save(_ *domain.Task) (*domain.Task, error) {
	return nil, fmt.Errorf("failed to mark task as done")
}
