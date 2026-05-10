package task

import (
	"context"
	"fmt"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/joaohgf/magalu-cli/internal/errors"
	"github.com/joaohgf/magalu-cli/internal/port"
)

// Add is a use case for adding a new task to the system.
// It interacts with the persistence layer to save the task and returns the saved task or an errors if the operation fails.
type Add struct {
	persistence port.PersistenceSaver[*domain.Task]
}

func NewAdd(persistence port.PersistenceSaver[*domain.Task]) *Add {
	return &Add{persistence: persistence}
}

// Save takes a task as input and attempts to save it using the persistence layer.
// If the save operation is successful, it returns the saved task.
// If there is an errors during saving, it wraps the errors with additional context and returns it.
func (a *Add) Save(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	// Validate the task before saving
	if err := a.validate(task); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	saved, err := a.persistence.Save(ctx, task)
	if err != nil {
		return nil, err
	}
	return saved, nil
}

func (a *Add) validate(task *domain.Task) error {
	if task == nil {
		return errors.Invalid("task cannot be nil")
	}
	if task.ID == "" {
		return errors.Invalid("task ID cannot be empty")
	}
	if task.Title == "" {
		return errors.Invalid("task title cannot be empty")
	}
	if task.Priority == enum.PriorityUnknown {
		return errors.Invalid("task priority")
	}
	if task.Status == enum.StatusUnknown {
		return errors.Invalid("task status")
	}
	return nil
}
