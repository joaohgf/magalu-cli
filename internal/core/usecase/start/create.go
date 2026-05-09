package start

import (
	"fmt"
	"os"
	"strings"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/port"
)

type Create struct {
	persistence port.PersistenceSaver[*domain.ConfigCommand]
}

func NewCreate(persistence port.PersistenceSaver[*domain.ConfigCommand]) *Create {
	return &Create{persistence: persistence}
}

func (c *Create) Save(target *domain.ConfigCommand) (*domain.ConfigCommand, error) {
	// todo set by envs
	err := os.MkdirAll("./tarefeiro/", os.ModeDir)
	if err != nil {
		return nil, fmt.Errorf("failed to create base directory: %w", err)
	}
	if strings.EqualFold(target.DatabasePath, "\n") {
		defaultConfig := domain.NewConfigCommand()
		target = defaultConfig
	}
	if target.DatabasePath != "" {
		target.DatabasePath = strings.TrimSpace(target.DatabasePath)
		target.DatabasePath = strings.TrimSuffix(target.DatabasePath, "\n")
	}
	target, err = c.persistence.Save(target)
	if err != nil {
		return nil, fmt.Errorf("failed to save configuration: %w", err)
	}
	return target, nil
}
