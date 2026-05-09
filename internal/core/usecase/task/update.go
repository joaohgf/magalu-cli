package task

import (
	"fmt"
	"strings"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/core/enum"
	"github.com/joaohgf/magalu-cli/internal/port"
)

type Update struct {
	saver  port.PersistenceSaver[*domain.Task]
	finder port.PersistenceFinder[*domain.Task]
}

func NewUpdate(
	saver port.PersistenceSaver[*domain.Task],
	finder port.PersistenceFinder[*domain.Task],
) *Update {
	return &Update{
		saver:  saver,
		finder: finder,
	}
}

func (u *Update) Save(target *domain.Task) (*domain.Task, error) {
	existing, err := u.finder.Find(target)
	if err != nil {
		return nil, fmt.Errorf("task with ID %s not found: %w", target.ID, err)
	}
	target = u.setFields(existing, target)
	if err = u.validate(target); err != nil {
		return nil, err
	}
	updatedTask, err := u.saver.Save(target)
	if err != nil {
		return nil, err
	}
	return updatedTask, nil
}

// validate checks if the target task has valid fields for an update operation.
func (u *Update) validate(target *domain.Task) error {
	if target.IsEmpty() {
		return fmt.Errorf("task cannot be empty")
	}
	if target.ID == "" {
		return fmt.Errorf("task ID is required for update")
	}
	if target.Priority == enum.PriorityUnknown {
		return fmt.Errorf("invalid priority value, must be one of: %s",
			strings.Join([]string{
				enum.PriorityLow.String(), enum.PriorityMedium.String(), enum.PriorityHigh.String(),
			}, ", "))
	}
	return nil
}

// setFields updates the existing task with non-empty fields from the target task.
func (u *Update) setFields(existing, target *domain.Task) *domain.Task {
	if target.Title != "" {
		existing.Title = target.Title
	}
	if target.Description != "" {
		existing.Description = target.Description
	}
	if target.Priority.String() != "" {
		existing.Priority = target.Priority
	}
	if len(target.Tags) > 0 {
		existing.Tags = target.Tags
	}
	if target.EstimatedDoneAt != nil && !target.EstimatedDoneAt.IsZero() {
		existing.EstimatedDoneAt = target.EstimatedDoneAt
	}
	return existing
}
