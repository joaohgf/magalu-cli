package task

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
	"gopkg.in/yaml.v3"
)

type DetailView struct {
	table *tablewriter.Table
}

func NewDetailView() *DetailView {
	detail := &DetailView{
		table: tablewriter.NewTable(
			os.Stdout,
			tablewriter.WithHeaderAlignment(tw.AlignLeft),
			tablewriter.WithRowAutoWrap(tw.WrapNormal),
		),
	}
	return detail
}

func (d *DetailView) Table(task *domain.Task) error {
	defer d.table.Close()
	d.table.Header([]string{"ID", "Title", "Priority", "Status", "Tags", "Created At", "Estimated Done At", "Description", "Done At"})
	var err error
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
	err = d.table.Append(row)
	if err != nil {
		return fmt.Errorf("failed to append task to table: %w", err)
	}
	err = d.table.Render()
	if err != nil {
		return fmt.Errorf("failed to render table: %w", err)
	}
	return nil
}

func (d *DetailView) JSON(task *domain.Task) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(task)
	if err != nil {
		return fmt.Errorf("failed to encode tasks to JSON: %w", err)
	}
	return nil
}

func (d *DetailView) YAML(task *domain.Task) error {
	encoder := yaml.NewEncoder(os.Stdout)
	err := encoder.Encode(task)
	if err != nil {
		return fmt.Errorf("failed to encode tasks to YAML: %w", err)
	}
	return nil
}
