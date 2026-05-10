package task

import (
	"context"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/errors"
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

func (u *Update) Save(ctx context.Context, target *domain.Task) (*domain.Task, error) {
	if target == nil || target.ID == "" {
		return nil, errors.Invalid("task ID is required for update")
	}
	existing, err := u.finder.Find(ctx, target)
	if err != nil {
		return nil, err
	}
	if !target.HasAnyUpdate() {
		return nil, errors.Invalid("no fields to update")
	}
	target = u.setFields(existing, target)
	updatedTask, err := u.saver.Save(ctx, target)
	if err != nil {
		return nil, err
	}
	return updatedTask, nil
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
