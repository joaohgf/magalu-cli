//go:build unit

package task

import (
	"context"
	"testing"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/errors"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestListRunner(t *testing.T) {
	t.Run("successfully", runListSuccessfully)
	t.Run("with error", func(t *testing.T) {
		t.Run("saving", runListWithSavingError)
		t.Run("rendering", runListWithRenderingError)
	})
}

func runListSuccessfully(t *testing.T) {
	cmd := getListCommand()
	runner := NewListRunner(new(mockFindAllUseCase), new(mockFilterRender))
	err := runner.Run(cmd, []string{})
	assert.NoError(t, err)
}

func runListWithSavingError(t *testing.T) {
	cmd := getListCommand()
	runner := NewListRunner(new(mockFindAllUseCaseWithError), new(mockFilterRender))
	err := runner.Run(cmd, []string{})
	assert.Error(t, err)
}

func runListWithRenderingError(t *testing.T) {
	cmd := getListCommand()
	runner := NewListRunner(new(mockFindAllUseCase), new(mockFilterRenderWithError))
	err := runner.Run(cmd, []string{})
	assert.Error(t, err)
}

/*
Mocks below
*/

func getListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
	}
	cmd.Flags().String("title", "Test Task", "")
	cmd.Flags().String("description", "This is a test task", "")
	cmd.Flags().String("status", "pending", "")
	cmd.Flags().String("priority", "medium", "")
	cmd.Flags().String("size", "10", "")
	cmd.Flags().String("page", "1", "")
	return cmd
}

type (
	mockFindAllUseCase          struct{}
	mockFindAllUseCaseWithError struct{}
	mockFilterRender            struct{}
	mockFilterRenderWithError   struct{}
)

func (m *mockFindAllUseCase) All(_ context.Context, _ *domain.TaskFilter) (*domain.TaskFilter, error) {
	return &domain.TaskFilter{}, nil
}

func (m *mockFindAllUseCaseWithError) All(_ context.Context, _ *domain.TaskFilter) (*domain.TaskFilter, error) {
	return nil, errors.New("error finding tasks")
}

func (m *mockFilterRender) Render(_ context.Context, _ *domain.TaskFilter) error {
	return nil
}

func (m *mockFilterRenderWithError) Render(_ context.Context, _ *domain.TaskFilter) error {
	return errors.New("error rendering tasks")
}
