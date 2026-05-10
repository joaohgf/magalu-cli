package task

import (
	"context"
	"fmt"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/port"
)

type Done struct {
	saver  port.PersistenceSaver[*domain.Task]
	finder port.PersistenceFinder[*domain.Task]
}

func NewDone(
	saver port.PersistenceSaver[*domain.Task],
	finder port.PersistenceFinder[*domain.Task],
) *Done {
	return &Done{
		saver:  saver,
		finder: finder,
	}
}

func (d *Done) Save(ctx context.Context, target *domain.Task) (*domain.Task, error) {
	if target == nil || target.ID == "" {
		return nil, fmt.Errorf("task ID is required to mark as done")
	}
	existing, err := d.finder.Find(ctx, target)
	if err != nil {
		return nil, err
	}
	existing.MarkAsDone()
	updatedTask, err := d.saver.Save(ctx, existing)
	if err != nil {
		return nil, err
	}
	return updatedTask, nil
}
