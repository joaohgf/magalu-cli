package start

import (
	"github.com/joaohgf/magalu-cli/internal/core/domain"
	usecase "github.com/joaohgf/magalu-cli/internal/core/usecase/start"
	"github.com/joaohgf/magalu-cli/internal/persistence"
	handler "github.com/joaohgf/magalu-cli/internal/runner/start"
	"github.com/nanobox-io/scribble"
	"github.com/spf13/cobra"
)

const (
	startUse              = "start"
	startShortDescription = "Configures the CLI default settings"
	startLongDescription  = "Initial setup for the CLI. \nYou can define the tasks database path."
)

func BuildStartCommandHandler(db *scribble.Driver) *cobra.Command {
	repository := persistence.NewServiceSaver[*domain.ConfigCommand](db)
	rule := usecase.NewCreate(repository)
	runner := handler.NewRunner(rule)
	cmd := &cobra.Command{
		Use:   startUse,
		Short: startShortDescription,
		Long:  startLongDescription,
		RunE:  runner.Run,
	}

	return cmd
}
