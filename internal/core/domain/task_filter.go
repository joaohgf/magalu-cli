package domain

import "github.com/joaohgf/magalu-cli/internal/enum"

type TaskFilter struct {
	Filter *Task
	Data   []*Task
	Page   int
	Size   int
	Total  int
}

func NewTaskFilter(filter *Task) *TaskFilter {
	return &TaskFilter{
		Filter: filter,
		Page:   enum.DefaultPage,
		Size:   enum.DefaultSize,
	}
}

func (tf *TaskFilter) GetID() string {
	if tf.Filter == nil {
		return ""
	}
	return tf.Filter.ID
}

func (tf *TaskFilter) GetCollection() string {
	return "tasks"
}

func (tf *TaskFilter) IsEmpty() bool {
	return tf == nil || tf.Filter.IsEmpty()
}

func (tf *TaskFilter) IsEqual(other *Task) bool {
	if tf.IsEmpty() {
		return true
	}
	condition := tf.Filter.Matches(other)
	return condition
}

func (tf *TaskFilter) GetSize() int {
	return tf.Size
}

func (tf *TaskFilter) SetSize(size int) {
	tf.Size = size
}

func (tf *TaskFilter) GetPage() int {
	return tf.Page
}

func (tf *TaskFilter) SetPage(page int) {
	tf.Page = page
}

func (tf *TaskFilter) SetContent(content ...*Task) {
	tf.Data = content
}

func (tf *TaskFilter) GetContent() []*Task {
	return tf.Data
}

func (tf *TaskFilter) SetTotal(total int) {
	tf.Total = total
}

func (tf *TaskFilter) GetTotal() int {
	return tf.Total
}
