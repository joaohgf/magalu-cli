package domain

import (
	"strings"
	"time"

	"github.com/joaohgf/magalu-cli/internal/enum"
)

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

func (t *Task) GetID() string {
	if t == nil {
		return ""
	}
	return t.ID
}

func (t *Task) GetCollection() string {
	return "tasks"
}

func (t *Task) MarkAsDone() {
	if t == nil {
		return
	}
	t.Status = enum.StatusDone
	t.DoneAt = new(time.Now())
}

func (t *Task) IsEmpty() bool {
	return t == nil || (t.ID == "" && t.Title == "" && t.Description == "" &&
		t.CreatedAt.IsZero() && t.Priority == "" && t.Status == "")
}

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
