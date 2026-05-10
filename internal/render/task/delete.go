package task

import (
	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/errors"
	"github.com/joaohgf/magalu-cli/internal/port"
)

// DeleteView is responsible for rendering the details of a deleted task in a tabular format.
// It embeds the DetailView to reuse its functionality for rendering task details.
type DeleteView struct {
	*DetailView
}

func NewDeleteView(tableWriter port.TableWriter) *DeleteView {
	del := &DeleteView{}
	del.DetailView = NewDetailView(tableWriter)
	return del
}

// Table is responsible for rendering the details of a deleted task in a tabular format.
func (dv *DeleteView) Table(task *domain.Task) error {
	defer dv.table.Close()
	dv.table.Header([]string{"", "ID"})
	row := []string{
		"Deleted successfully task with ID:",
		task.ID,
	}
	err := dv.table.Append(row)
	if err != nil {
		return errors.Invalid("failed to append task to table")
	}
	err = dv.table.Render()
	if err != nil {
		return errors.Invalid("failed to render table")
	}
	return nil
}
