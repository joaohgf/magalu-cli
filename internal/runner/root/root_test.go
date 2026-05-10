//go:build unit

package root

import (
	"testing"

	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/spf13/cobra"
)

func TestRootRunner(t *testing.T) {
	t.Run("run", func(t *testing.T) {
		t.Run("successfully", runSuccessfully)
	})
	t.Run("pre-run", func(t *testing.T) {
		t.Run("successfully", preRunSuccessfully)
		t.Run("with error", preRunWithError)
	})
}

func runSuccessfully(t *testing.T) {
	command := getRootCommand()
	runner := NewRunner()
	err := runner.Run(command, []string{})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func preRunSuccessfully(t *testing.T) {
	command := getRootCommand()
	runner := NewRunner()
	err := runner.PreRun(command, []string{})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func preRunWithError(t *testing.T) {
	command := getRootCommand()
	runner := NewRunner()
	err := runner.PreRun(command, []string{})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

/*
Mocks below
*/

func getRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "tarefeiro",
	}
	cmd.Flags().StringP(enum.OutputDefaultKey.ToLower(), "o", "", "output format")
	return cmd
}
