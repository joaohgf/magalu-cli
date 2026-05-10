package task

type AddView struct {
	*DetailView
}

func NewAddView() *AddView {
	add := &AddView{}
	add.DetailView = NewDetailView()
	return add
}
