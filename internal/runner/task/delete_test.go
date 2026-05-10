//go:build unit

package task

import (
	"context"
	"fmt"
	"testing"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/spf13/cobra"
)

func TestDeleteRunner(t *testing.T) {
	t.Run("run successfully", runDeleteSuccessfully)
	t.Run("run with errors", runDeleteWithError)
}

func runDeleteSuccessfully(t *testing.T) {
	cmd := getDeleteCommand()
	runner := NewDeleteRunner(new(mockDeleteUseCase))
	err := runner.Run(cmd, []string{"123"})
	if err != nil {
		t.Fatalf("expected no errors but got: %v", err)
	}
}

func runDeleteWithError(t *testing.T) {
	cmd := getDeleteCommand()
	runner := NewDeleteRunner(new(mockDeleteUseCaseWithError))
	err := runner.Run(cmd, []string{"123"})
	if err == nil {
		t.Fatal("expected an errors but got nil")
	}
}

/*
Mocks below
*/

func getDeleteCommand() *cobra.Command {
	return &cobra.Command{
		Use: "delete",
	}
}

type (
	mockDeleteUseCase          struct{}
	mockDeleteUseCaseWithError struct{}
)

func (m *mockDeleteUseCase) Delete(_ context.Context, _ *domain.Task) error {
	return nil
}

func (m *mockDeleteUseCaseWithError) Delete(_ context.Context, _ *domain.Task) error {
	return fmt.Errorf("failed to delete task")
}
