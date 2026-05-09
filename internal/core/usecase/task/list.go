package task

import (
	"errors"
	"os"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/port"
)

// List is a use case for retrieving all tasks that match certain criteria from the system.
// It interacts with the persistence layer to fetch the tasks and returns a slice of tasks or an error if the operation fails.
type List struct {
	persistence port.PersistenceFinder[*domain.Task]
}

func NewList(persistence port.PersistenceFinder[*domain.Task]) *List {
	return &List{persistence: persistence}
}

// All takes a target task as input, which serves as a filter for the search criteria.
func (l *List) All(target *domain.Task) ([]*domain.Task, error) {
	tasks, err := l.persistence.FindAll(target)
	if err != nil {
		if _, ok := errors.AsType[*os.PathError](err); ok {
			return nil, nil
		}
		return nil, err
	}
	return tasks, nil
}
