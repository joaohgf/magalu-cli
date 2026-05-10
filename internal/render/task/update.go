package task

import "github.com/joaohgf/magalu-cli/internal/port"

// UpdateView is responsible for rendering the details of an updated task.
// It embeds the DetailView, which provides the necessary methods for rendering
// the task details in different formats (table, JSON, YAML).
type UpdateView struct {
	*DetailView
}

func NewUpdateView(tableWriter port.TableWriter) *UpdateView {
	update := &UpdateView{}
	update.DetailView = NewDetailView(tableWriter)
	return update
}
