package start

import (
	"fmt"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/joaohgf/magalu-cli/internal/port"
)

type Create struct {
	persistence port.PersistenceSaver[*domain.ConfigCommand]
}

func NewCreate(persistence port.PersistenceSaver[*domain.ConfigCommand]) *Create {
	return &Create{persistence: persistence}
}

func (c *Create) Save(target *domain.ConfigCommand) (*domain.ConfigCommand, error) {
	if target.Output == enum.OutputUnknown {
		target.Output = enum.OutputTable
	}
	target, err := c.persistence.Save(target)
	if err != nil {
		return nil, fmt.Errorf("failed to save configuration: %w", err)
	}
	return target, nil
}
