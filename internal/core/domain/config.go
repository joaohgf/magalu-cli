package domain

import (
	"github.com/joaohgf/magalu-cli/internal/enum"
)

type ConfigCommand struct {
	ID     string
	Output enum.Output
}

func NewConfigCommand() *ConfigCommand {
	return &ConfigCommand{
		ID:     "default",
		Output: enum.OutputTable,
	}
}

func (c *ConfigCommand) GetID() string {
	return c.ID
}

func (c *ConfigCommand) GetCollection() string {
	return "config"
}
