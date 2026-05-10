package task

import (
	"context"
	"fmt"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/port"
)

type Delete struct {
	persistence port.PersistenceDeleter[*domain.Task]
}

func NewDelete(persistence port.PersistenceDeleter[*domain.Task]) *Delete {
	return &Delete{persistence: persistence}
}

func (d *Delete) Delete(ctx context.Context, target *domain.Task) error {
	if target == nil || target.ID == "" {
		return fmt.Errorf("task ID is required for deletion")
	}
	err := d.persistence.Delete(ctx, target)
	if err != nil {
		return fmt.Errorf("deleting task: %w", err)
	}
	return nil
}
