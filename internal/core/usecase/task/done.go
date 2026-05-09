package task

import (
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

func (d *Done) Save(target *domain.Task) (*domain.Task, error) {
	existing, err := d.finder.Find(target)
	if err != nil {
		return nil, err
	}
	existing.MarkAsDone()
	updatedTask, err := d.saver.Save(existing)
	if err != nil {
		return nil, err
	}
	return updatedTask, nil
}
