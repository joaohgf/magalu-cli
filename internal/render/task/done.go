package task

type DoneView struct {
	*DetailView
}

func NewDoneView() *DoneView {
	done := &DoneView{}
	done.DetailView = NewDetailView()
	return done
}
