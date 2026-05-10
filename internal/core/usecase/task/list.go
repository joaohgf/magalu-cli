package task

import (
	"context"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/joaohgf/magalu-cli/internal/port"
)

// List is a use case for retrieving all tasks that match certain criteria from the system.
// It interacts with the persistence layer to fetch the tasks and returns a slice of tasks or an errors if the operation fails.
type List struct {
	persistence port.PersistenceLister[*domain.Task, *domain.TaskFilter]
}

func NewList(persistence port.PersistenceLister[*domain.Task, *domain.TaskFilter]) *List {
	return &List{persistence: persistence}
}

// All takes a target task as input, which serves as a filter for the search criteria.
func (l *List) All(ctx context.Context, target *domain.TaskFilter) (*domain.TaskFilter, error) {
	l.setDefault(target)
	found, err := l.persistence.List(ctx, target)
	return found, err
}

func (l *List) setDefault(target *domain.TaskFilter) {
	if target.GetSize() <= 0 {
		target.SetSize(enum.DefaultSize)
	}
	if target.GetPage() <= 0 {
		target.SetPage(enum.DefaultPage)
	}
}
