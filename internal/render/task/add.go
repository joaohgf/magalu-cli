package task

import "github.com/joaohgf/magalu-cli/internal/port"

// AddView is responsible for rendering the details of a task when adding a new task.
// It embeds the DetailView, which provides the necessary methods for rendering
// the task details in different formats (table, JSON, YAML).
type AddView struct {
	*DetailView
}

func NewAddView(tableWriter port.TableWriter) *AddView {
	add := &AddView{}
	add.DetailView = NewDetailView(tableWriter)
	return add
}
