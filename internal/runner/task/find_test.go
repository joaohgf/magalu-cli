//go:build unit

package task

import (
	"fmt"
	"testing"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/spf13/cobra"
)

func TestFindRunner(t *testing.T) {
	t.Run("successfully", runFindSuccessfully)
	t.Run("with error", func(t *testing.T) {
		t.Run("when finding task", runFindWithError)
		t.Run("when rendering task", runRenderWithError)
	})
}

func runFindSuccessfully(t *testing.T) {
	cmd := getFindCommand()
	runner := NewFindRunner(new(mockFindUseCase), new(mockRender))
	err := runner.Run(cmd, []string{"123"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func runFindWithError(t *testing.T) {
	cmd := getFindCommand()
	runner := NewFindRunner(new(mockFindUseCaseWithError), new(mockRender))
	err := runner.Run(cmd, []string{"123"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func runRenderWithError(t *testing.T) {
	cmd := getFindCommand()
	runner := NewFindRunner(new(mockFindUseCase), new(mockRenderWithError))
	err := runner.Run(cmd, []string{"123"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

/*
Mocks below
*/

func getFindCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "show",
	}
	return cmd
}

type (
	mockFindUseCase          struct{}
	mockFindUseCaseWithError struct{}
	mockRender               struct{}
	mockRenderWithError      struct{}
)

func (m *mockFindUseCase) Find(_ *domain.Task) (*domain.Task, error) {
	return &domain.Task{ID: "123"}, nil
}

func (m *mockFindUseCaseWithError) Find(_ *domain.Task) (*domain.Task, error) {
	return nil, fmt.Errorf("error finding task")
}

func (m *mockRender) Render(_ ...*domain.Task) error {
	return nil
}

func (m *mockRenderWithError) Render(_ ...*domain.Task) error {
	return fmt.Errorf("error rendering task")
}
