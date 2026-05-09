package task

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/core/enum"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
)

type Detail struct {
	*tablewriter.Table
}

func NewDetail() *Detail {
	table := &Detail{
		Table: tablewriter.NewTable(
			os.Stdout,
			tablewriter.WithHeaderAlignment(tw.AlignLeft),
			tablewriter.WithRowAutoWrap(tw.WrapNormal),
		),
	}
	return table
}

func (d *Detail) Render(tasks ...*domain.Task) error {
	defer d.Table.Close()
	d.Table.Header([]string{"ID", "Title", "Priority", "Status", "Tags", "Created At", "Estimated Done At", "Description", "Done At"})
	var err error
	for _, task := range tasks {
		row := []string{
			task.ID,
			task.Title,
			enum.ColorByPriority(task.Priority),
			enum.ColorByStatus(task.Status),
			strings.Join(task.Tags, ","),
			task.CreatedAt.Format(time.DateTime),
		}
		estimatedDoneAt := "N/A"
		if task.EstimatedDoneAt != nil && !task.EstimatedDoneAt.IsZero() {
			estimatedDoneAt = task.EstimatedDoneAt.Format(time.DateTime)
		}
		row = append(row, estimatedDoneAt)
		row = append(row, task.Description)
		doneAt := "N/A"
		if task.DoneAt != nil && !task.DoneAt.IsZero() {
			doneAt = task.DoneAt.Format(time.DateTime)
		}
		row = append(row, doneAt)
		err = d.Table.Append(row)
		if err != nil {
			return fmt.Errorf("failed to append task to table: %w", err)
		}
	}
	err = d.Table.Render()
	if err != nil {
		return fmt.Errorf("failed to render table: %w", err)
	}
	return nil
}
