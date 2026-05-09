package domain

import (
	"time"

	"github.com/joaohgf/magalu-cli/internal/core/enum"
)

type Task struct {
	ID              string        `json:"id,omitempty"`
	Title           string        `json:"title,omitempty"`
	Description     string        `json:"description,omitempty"`
	Status          enum.Status   `json:"status,omitempty"`
	Priority        enum.Priority `json:"priority,omitempty"`
	CreatedAt       time.Time     `json:"created_at,omitempty"`
	DoneAt          *time.Time    `json:"done_at,omitempty"`
	EstimatedDoneAt *time.Time    `json:"estimated_done_at,omitempty"`
	Tags            []string      `json:"tags,omitempty"`
}

func NewTask(id, title, description string, priority enum.Priority, createdAt time.Time, tags ...string) *Task {
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
	return t.ID
}

func (t *Task) GetCollection() string {
	return "tasks"
}

func (t *Task) MarkAsDone() {
	t.Status = enum.StatusDone
	t.DoneAt = new(time.Now())
}

func (t *Task) IsOverdue() bool {
	if t.EstimatedDoneAt == nil {
		return false
	}
	condition := time.Now().After(*t.EstimatedDoneAt) && t.Status != enum.StatusDone
	return condition
}

func (t *Task) IsEmpty() bool {
	return t == nil || (t.ID == "" && t.Title == "" && t.Description == "" && t.CreatedAt.IsZero() && t.Priority == "" && t.Status == "")
}

func (t *Task) IsEqual(other *Task) bool {
	if t.IsEmpty() {
		return true
	}
	if t.ID != "" && t.ID == other.ID {
		return true
	}
	if t.Title != "" && t.Title == other.Title {
		return true
	}
	if t.Description != "" && t.Description == other.Description {
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
