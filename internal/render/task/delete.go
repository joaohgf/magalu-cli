package task

import (
	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/errors"
)

type DeleteView struct {
	*DetailView
}

func NewDeleteView() *DeleteView {
	del := &DeleteView{}
	del.DetailView = NewDetailView()
	return del
}

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
