package task

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/joaohgf/magalu-cli/internal/port"
	"gopkg.in/yaml.v3"
)

// List is responsible for rendering a list of tasks in a tabular format,
// as well as providing options to render the tasks in JSON and YAML formats.
type List struct {
	table port.TableWriter
}

func NewList(table port.TableWriter) *List {
	list := &List{table: table}
	return list
}

// Table is responsible for rendering a list of tasks in a tabular format.
func (l *List) Table(filter *domain.TaskFilter) error {
	defer l.table.Close()
	if len(filter.Data) == 0 {
		l.table.Header([]string{"No tasks found"})
		err := l.table.Render()
		if err != nil {
			return fmt.Errorf("failed to render table: %w", err)
		}
		return nil
	}
	l.table.Header([]string{"ID", "Title", "Priority", "Status"})
	var err error
	for _, task := range filter.Data {
		err = l.table.Append(
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
	firstElements := filter.GetSize() * (filter.GetPage() - 1)
	l.table.Footer([]string{"", "", "", fmt.Sprintf("rows %d - %d", firstElements, filter.GetTotal()+firstElements)})
	err = l.table.Render()
	if err != nil {
		return fmt.Errorf("failed to render table: %w", err)
	}
	return nil
}

// JSON is responsible for rendering a list of tasks in JSON format.
func (l *List) JSON(filter *domain.TaskFilter) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(filter.Data)
	if err != nil {
		return fmt.Errorf("failed to encode tasks to JSON: %w", err)
	}
	return nil
}

// YAML is responsible for rendering a list of tasks in YAML format.
func (l *List) YAML(filter *domain.TaskFilter) error {
	encoder := yaml.NewEncoder(os.Stdout)
	defer encoder.Close()
	err := encoder.Encode(filter.Data)
	if err != nil {
		return fmt.Errorf("failed to encode tasks to YAML: %w", err)
	}
	return nil
}
