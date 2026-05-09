package task

import (
	"fmt"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/core/enum"
	"github.com/joaohgf/magalu-cli/internal/port"
)

// Add is a use case for adding a new task to the system.
// It interacts with the persistence layer to save the task and returns the saved task or an error if the operation fails.
type Add struct {
	persistence port.PersistenceSaver[*domain.Task]
}

func NewAdd(persistence port.PersistenceSaver[*domain.Task]) *Add {
	return &Add{persistence: persistence}
}

// Save takes a task as input and attempts to save it using the persistence layer.
// If the save operation is successful, it returns the saved task.
// If there is an error during saving, it wraps the error with additional context and returns it.
func (a *Add) Save(task *domain.Task) (*domain.Task, error) {
	// Validate the task before saving
	if err := a.validate(task); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}
	saved, err := a.persistence.Save(task)
	if err != nil {
		return nil, fmt.Errorf("error saving task: %w", err)
	}
	return saved, nil
}

func (a *Add) validate(task *domain.Task) error {
	if task.Title == "" {
		return fmt.Errorf("task title cannot be empty")
	}
	if task.Priority == enum.PriorityUnknown {
		return fmt.Errorf("invalid task priority")
	}
	return nil
}
