package task

import (
	"fmt"
	"os"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
)

type List struct {
	*tablewriter.Table
	config *domain.ConfigCommand
}

func NewList(config *domain.ConfigCommand) *List {
	table := &List{
		Table: tablewriter.NewTable(
			os.Stdout,
			tablewriter.WithHeaderAlignment(tw.AlignLeft),
			tablewriter.WithRowAutoWrap(tw.WrapNormal),
		),
		config: config,
	}
	return table
}

func (l *List) Render(tasks ...*domain.Task) error {
	defer l.Table.Close()
	if len(tasks) == 0 {
		l.Table.Header([]string{"No tasks found"})
		err := l.Table.Render()
		if err != nil {
			return fmt.Errorf("failed to render table: %w", err)
		}
		return nil
	}
	l.Table.Header([]string{"ID", "Title", "Priority", "Status"})
	var err error
	for _, task := range tasks {
		err = l.Table.Append(
			[]string{
				task.ID,
				task.Title,
				enum.ColorByPriority(task.Priority),
				enum.ColorByStatus(task.Status),
			},
		)
		if err != nil {
			return fmt.Errorf("failed to append task to table: %w", err)
		}
	}
	err = l.Table.Render()
	if err != nil {
		return fmt.Errorf("failed to render table: %w", err)
	}
	return nil
}
