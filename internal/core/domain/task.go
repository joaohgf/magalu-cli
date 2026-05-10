package domain

import (
	"strings"
	"time"

	"github.com/joaohgf/magalu-cli/internal/enum"
)

// Task represents a task in the task management system.
type Task struct {
	ID              string     `json:"id,omitempty" yaml:"id"`
	Title           string     `json:"title,omitempty" yaml:"title"`
	Description     string     `json:"description,omitempty" yaml:"description"`
	Status          enum.Type  `json:"status,omitempty" yaml:"status"`
	Priority        enum.Type  `json:"priority,omitempty" yaml:"priority"`
	CreatedAt       time.Time  `json:"created_at,omitempty" yaml:"created_at"`
	DoneAt          *time.Time `json:"done_at,omitempty" yaml:"done_at"`
	EstimatedDoneAt *time.Time `json:"estimated_done_at,omitempty" yaml:"estimated_done_at"`
	Tags            []string   `json:"tags,omitempty" yaml:"tags"`
}

func NewTask(id, title, description string, priority enum.Type, createdAt time.Time, tags ...string) *Task {
	return &Task{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      enum.StatusInProgress,
		Priority:    priority,
		CreatedAt:   createdAt,
		Tags:        tags,
	}
}

// GetID returns the ID of the task. If the Task struct is nil, it returns an empty string.
func (t *Task) GetID() string {
	if t == nil {
		return ""
	}
	return t.ID
}

// GetCollection returns the name of the collection where tasks are stored, which is "tasks".
func (t *Task) GetCollection() string {
	return "tasks"
}

// MarkAsDone updates the status of the task to "done" and sets the DoneAt timestamp to the current time.
func (t *Task) MarkAsDone() {
	if t == nil {
		return
	}
	t.Status = enum.StatusDone
	t.DoneAt = new(time.Now())
}

// IsEmpty checks if the Task struct is empty, meaning all fields are either zero values or empty strings.
func (t *Task) IsEmpty() bool {
	return t == nil || (t.ID == "" && t.Title == "" && t.Description == "" &&
		t.CreatedAt.IsZero() && t.Priority == "" && t.Status == "")
}

// Matches checks if the current Task matches another Task based on non-empty fields.
func (t *Task) Matches(other *Task) bool {
	if t.IsEmpty() {
		return true
	}
	if t.ID != "" && t.ID == other.ID {
		return true
	}
	if t.Title != "" && strings.Contains(strings.ToLower(other.Title), strings.ToLower(t.Title)) {
		return true
	}
	if t.Description != "" && strings.Contains(strings.ToLower(other.Description), strings.ToLower(t.Description)) {
		return true
	}
	if t.Status != enum.StatusUnknown && t.Status == other.Status {
		return true
	}
	if t.Priority != enum.PriorityUnknown && t.Priority == other.Priority {
		return true
	}
	return false
}

// HasAnyUpdate checks if any of the fields in the Task struct have been updated (i.e., are not empty or zero values).
func (t *Task) HasAnyUpdate() bool {
	if t == nil {
		return false
	}
	return t.Title != "" ||
		t.Description != "" ||
		(t.Priority != "" && t.Priority != enum.PriorityUnknown) ||
		len(t.Tags) > 0 ||
		(t.EstimatedDoneAt != nil && !t.EstimatedDoneAt.IsZero())
}
