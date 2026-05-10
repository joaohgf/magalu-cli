package task

type UpdateView struct {
	*DetailView
}

func NewUpdateView() *UpdateView {
	update := &UpdateView{}
	update.DetailView = NewDetailView()
	return update
}
