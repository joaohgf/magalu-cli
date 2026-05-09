package start

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/spf13/cobra"
)

type CreateRunner struct {
	useCase port.SaveUseCase[*domain.ConfigCommand]
}

func NewRunner(useCase port.SaveUseCase[*domain.ConfigCommand]) *CreateRunner {
	return &CreateRunner{useCase: useCase}
}

func (cr *CreateRunner) Run(cmd *cobra.Command, _ []string) error {
	config := domain.NewConfigCommand()
	cmd.Printf("Enter the path of the database file(default is %s):", config.DatabasePath)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}
	config.DatabasePath = input
	config, err = cr.useCase.Save(config)
	if err != nil {
		return fmt.Errorf("failed to save configuration: %w", err)
	}
	cmd.Printf("Configuration saved in %s\n", config.DatabasePath)
	return nil
}
