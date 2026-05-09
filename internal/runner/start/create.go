package start

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/enum"
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
	cmd.Printf("Enter the output[%s] (default is %s):",
		strings.Join([]string{enum.OutputJSON.String(), enum.OutputYAML.String(), enum.OutputTable.String()}, ", "),
		enum.OutputTable)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}
	config := domain.NewConfigCommand()
	config.Output = enum.OutputOf(input)
	config, err = cr.useCase.Save(config)
	if err != nil {
		return fmt.Errorf("failed to save configuration: %w", err)
	}
	cmd.Printf("Configuration saved in %s\n", config.ID)
	return nil
}
