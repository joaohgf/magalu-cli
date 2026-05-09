package task

import (
	"fmt"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/port"
)

type Find struct {
	persistence port.FindUseCase[*domain.Task]
}

func NewFind(persistence port.FindUseCase[*domain.Task]) *Find {
	return &Find{persistence: persistence}
}

func (f *Find) Find(target *domain.Task) (*domain.Task, error) {
	if target.ID == "" {
		return nil, fmt.Errorf("task ID is required")
	}
	task, err := f.persistence.Find(target)
	if err != nil {
		return nil, fmt.Errorf("error finding task: %w", err)
	}
	return task, nil
}
