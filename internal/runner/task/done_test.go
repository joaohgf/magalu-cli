//go:build unit

package task

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestDoneRunner(t *testing.T) {
	t.Run("run successfully", runDoneSuccessfully)
	t.Run("run with errors", func(t *testing.T) {
		t.Run("when saving task", runDoneWithError)
		t.Run("when rendering task", runDoneRenderWithError)
	})
}

func runDoneSuccessfully(t *testing.T) {
	cmd := getCompleteCommand()
	runner := NewDoneRunner(new(mockSaveUseCase), new(mockRender))
	err := runner.Run(cmd, []string{"123"})
	if err != nil {
		t.Fatalf("expected no errors but got: %v", err)
	}
}

func runDoneWithError(t *testing.T) {
	cmd := getCompleteCommand()
	runner := NewDoneRunner(new(mockSaveUseCaseWithError), new(mockRender))
	err := runner.Run(cmd, []string{"123"})
	if err == nil {
		t.Fatal("expected an errors but got nil")
	}
}

func runDoneRenderWithError(t *testing.T) {
	cmd := getCompleteCommand()
	runner := NewDoneRunner(new(mockSaveUseCase), new(mockRenderWithError))
	err := runner.Run(cmd, []string{"123"})
	if err == nil {
		t.Fatal("expected an errors but got nil")
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
