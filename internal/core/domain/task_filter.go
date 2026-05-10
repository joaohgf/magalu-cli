package domain

import "github.com/joaohgf/magalu-cli/internal/enum"

// TaskFilter is a struct that represents a filter for tasks, containing the filter criteria, the resulting data, pagination information (page and size), and the total number of matching tasks.
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

// GetID returns the ID of the filter if it exists, otherwise it returns an empty string.
func (tf *TaskFilter) GetID() string {
	if tf.Filter == nil {
		return ""
	}
	return tf.Filter.ID
}

// GetCollection returns the name of the collection where tasks are stored, which is "tasks".
func (tf *TaskFilter) GetCollection() string {
	return "tasks"
}

// IsEmpty checks if the TaskFilter struct is empty, meaning the Filter field is nil or the Filter itself is empty.
func (tf *TaskFilter) IsEmpty() bool {
	return tf == nil || tf.Filter.IsEmpty()
}

// IsEqual checks if the TaskFilter matches another Task based on the criteria defined in the Filter field.
// If the TaskFilter is empty, it returns true. Otherwise, it uses the Matches method of the Filter to compare with the other Task.
func (tf *TaskFilter) IsEqual(other *Task) bool {
	if tf.IsEmpty() {
		return true
	}
	condition := tf.Filter.Matches(other)
	return condition
}

// GetSize returns the size of the page for pagination.
func (tf *TaskFilter) GetSize() int {
	return tf.Size
}

// SetSize sets the size of the page for pagination.
func (tf *TaskFilter) SetSize(size int) {
	tf.Size = size
}

// GetPage returns the current page number for pagination.
func (tf *TaskFilter) GetPage() int {
	return tf.Page
}

// SetPage sets the current page number for pagination.
func (tf *TaskFilter) SetPage(page int) {
	tf.Page = page
}

// SetContent sets the content of the TaskFilter with a variable number of Task pointers.
func (tf *TaskFilter) SetContent(content ...*Task) {
	tf.Data = content
}

// GetContent returns the content of the TaskFilter, which is a slice of Task pointers.
func (tf *TaskFilter) GetContent() []*Task {
	return tf.Data
}

// SetTotal sets the total number of matching tasks for pagination.
func (tf *TaskFilter) SetTotal(total int) {
	tf.Total = total
}

// GetTotal returns the total number of matching tasks for pagination.
func (tf *TaskFilter) GetTotal() int {
	return tf.Total
}
