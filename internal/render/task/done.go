package task

import "github.com/joaohgf/magalu-cli/internal/port"

// DoneView is responsible for rendering the details of a completed task.
// It embeds the DetailView, which provides the necessary methods for rendering
// the task details in different formats (table, JSON, YAML).
type DoneView struct {
	*DetailView
}

func NewDoneView(tableWriter port.TableWriter) *DoneView {
	done := &DoneView{}
	done.DetailView = NewDetailView(tableWriter)
	return done
}
