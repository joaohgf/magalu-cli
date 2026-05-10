//go:build unit

package task

import (
	"context"
	"fmt"
	"testing"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestDeleteRunner(t *testing.T) {
	t.Run("run successfully", runDeleteSuccessfully)
	t.Run("run with errors", func(t *testing.T) {
		t.Run("deleting", runDeleteWithError)
		t.Run("rendering", runDeleteRenderWithError)
	})
}

func runDeleteSuccessfully(t *testing.T) {
	cmd := getDeleteCommand()
	runner := NewDeleteRunner(new(mockDeleteUseCase), new(mockRender))
	err := runner.Run(cmd, []string{"123"})
	assert.NoError(t, err)
}

func runDeleteWithError(t *testing.T) {
	cmd := getDeleteCommand()
	runner := NewDeleteRunner(new(mockDeleteUseCaseWithError), new(mockRender))
	err := runner.Run(cmd, []string{"123"})
	assert.Error(t, err)
}

func runDeleteRenderWithError(t *testing.T) {
	cmd := getDeleteCommand()
	runner := NewDeleteRunner(new(mockDeleteUseCase), new(mockRenderWithError))
	err := runner.Run(cmd, []string{"123"})
	assert.Error(t, err)
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
