package task

import (
	"encoding/json"
	"os"
	"strings"
	"time"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/joaohgf/magalu-cli/internal/errors"
	"github.com/joaohgf/magalu-cli/internal/port"
	"gopkg.in/yaml.v3"
)

// DetailView is responsible for rendering the details of a task in a tabular format,
// as well as providing options to render the task in JSON and YAML formats.
type DetailView struct {
	table port.TableWriter
}

func NewDetailView(t port.TableWriter) *DetailView {
	return &DetailView{
		table: t,
	}
}

// Table is responsible for rendering the details of a task in a tabular format.
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
		return errors.Invalid("failed to append task to table")
	}
	err = d.table.Render()
	if err != nil {
		return errors.Invalid("failed to render table")
	}
	return nil
}

// JSON is responsible for rendering the details of a task in JSON format.
func (d *DetailView) JSON(task *domain.Task) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(task)
	if err != nil {
		return errors.Invalid("failed to encode tasks to JSON")
	}
	return nil
}

// YAML is responsible for rendering the details of a task in YAML format.
func (d *DetailView) YAML(task *domain.Task) error {
	encoder := yaml.NewEncoder(os.Stdout)
	err := encoder.Encode(task)
	if err != nil {
		return errors.Invalid("failed to encode tasks to YAML")
	}
	return nil
}
