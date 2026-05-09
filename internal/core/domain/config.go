package domain

import "github.com/joaohgf/magalu-cli/internal/core/enum"

type ConfigCommand struct {
	ID     string
	Output string
}

func NewConfigCommand() *ConfigCommand {
	return &ConfigCommand{
		ID:     "default",
		Output: enum.OutputTable.String(),
	}
}

func (c *ConfigCommand) GetID() string {
	return c.ID
}

func (c *ConfigCommand) GetCollection() string {
	return "config"
}
