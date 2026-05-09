package cli

import (
	"log"
	"os"

	"github.com/joaohgf/magalu-cli/internal/cli/start"
	"github.com/joaohgf/magalu-cli/internal/cli/task"
	handler "github.com/joaohgf/magalu-cli/internal/runner/root"
	"github.com/nanobox-io/scribble"
	"github.com/spf13/cobra"
)

const (
	useRoot              = "tarefeiro"
	shortDescriptionRoot = "It's a CLI application to manage your tasks"
	longDescriptionRoot  = "It's a CLI application to manage your tasks.\n" +
		"You can create, list, update and delete your tasks with this application.\n\n" +
		"It's a CLI created to resolve the challenge of the Magalu's selection process."
	version = "0.1.0"
)

// buildRootCommandHandler creates the root command for the CLI application.
// The root command does not perform any action when executed;
// instead, it displays the help message to guide users on how to use the CLI.
func buildRootCommandHandler() *cobra.Command {
	runner := handler.NewRunner()
	cmd := &cobra.Command{
		Use:               useRoot,
		Short:             shortDescriptionRoot,
		Long:              longDescriptionRoot,
		Version:           version,
		RunE:              runner.Run,
		PersistentPreRunE: runner.PreRun,
	}
	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	db, err := scribble.New("./tarefeiro", nil)
	if err != nil {
		logger := log.Default()
		logger.SetOutput(os.Stderr)
		logger.Printf("Failed to initialize database: %v", err)
		os.Exit(1)
	}
	root := buildRootCommandHandler()
	commands := []*cobra.Command{
		start.BuildStartCommandHandler(db),
		task.BuildTaskAddCommandHandler(db),
		task.BuildTaskListCommandHandler(db),
		task.BuildTaskShowCommandHandler(db),
		task.BuildTaskDoneCommandHandler(db),
		task.BuildUpdateCommandHandler(db),
		task.BuildDeleteCommandHandler(db),
	}
	root.AddCommand(commands...)
	err = root.Execute()
	if err != nil {
		os.Exit(1)
	}
}
