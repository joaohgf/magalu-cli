package cli

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/joaohgf/magalu-cli/internal/cli/task"
	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/joaohgf/magalu-cli/internal/port"
	handler "github.com/joaohgf/magalu-cli/internal/runner/root"
	"github.com/nanobox-io/scribble"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
	"github.com/spf13/cobra"
)

const (
	useRoot              = "tarefeiro"
	shortDescriptionRoot = "It's a CLI application to manage your tasks"
	longDescriptionRoot  = "It's a CLI application to manage your tasks.\n" +
		"You can create, list, update and delete your tasks with this application.\n" +
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
	cmd.PersistentFlags().StringP("output", "o", "",
		fmt.Sprintf("Output format (%s)", strings.Join([]string{
			enum.OutputTable.String(), enum.OutputJSON.String(), enum.OutputYAML.String()}, ", "),
		))
	return cmd
}

func buildCommands(db *scribble.Driver, tableWriter *tablewriter.Table) ([]*cobra.Command, error) {
	commands := []port.Command{
		task.BuildTaskAddCommandHandler,
		task.BuildTaskDoneCommandHandler,
		task.BuildTaskUpdateCommandHandler,
		task.BuildTaskDeleteCommandHandler,
		task.BuildTaskListCommandHandler,
		task.BuildTaskShowCommandHandler,
	}
	var cobraCommands []*cobra.Command
	for _, command := range commands {
		cmd, err := command(db, tableWriter)
		if err != nil {
			return nil, err
		}
		cobraCommands = append(cobraCommands, cmd)
	}
	return cobraCommands, nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ctx context.Context) error {
	db, err := scribble.New("./data/", nil)
	if err != nil {
		return err
	}
	tableWriter := tablewriter.NewTable(
		os.Stdout,
		tablewriter.WithHeaderAlignment(tw.AlignLeft),
		tablewriter.WithRowAutoWrap(tw.WrapNormal),
	)
	root := buildRootCommandHandler()
	commands, err := buildCommands(db, tableWriter)
	if err != nil {
		return err
	}
	root.AddCommand(commands...)
	err = root.Execute()
	return err
}
