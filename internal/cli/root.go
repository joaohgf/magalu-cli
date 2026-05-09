package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/joaohgf/magalu-cli/internal/cli/start"
	"github.com/joaohgf/magalu-cli/internal/cli/task"
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

var (
	// skipInitCheck lists commands that are allowed before start is run.
	skipInitCheck = map[string]bool{
		"start":      true,
		"help":       true,
		"completion": true,
	}
)

// buildRootCommandHandler creates the root command for the CLI application.
// The root command does not perform any action when executed;
// instead, it displays the help message to guide users on how to use the CLI.
func buildRootCommandHandler() *cobra.Command {
	cmd := &cobra.Command{
		Use:     useRoot,
		Short:   shortDescriptionRoot,
		Long:    longDescriptionRoot,
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				os.Exit(1)
			}
		},
	}
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if skipInitCheck[cmd.Name()] {
			return nil
		}
		if _, err := os.Stat("./tarefeiro/"); errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("configuration not found. Please run '%s start' to set up the CLI", cmd.Root().Name())
		}
		return nil
	}
	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(db *scribble.Driver) {
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
	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}
