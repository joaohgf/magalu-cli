package render

import (
	"context"

	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/joaohgf/magalu-cli/internal/port"
)

// Render is responsible for rendering the output of the CLI application in different formats (table, JSON, YAML) based on the context value for the output type.
type Render[T any] struct {
	view port.View[T]
}

func NewRender[T any](view port.View[T]) *Render[T] {
	return &Render[T]{view: view}
}

// Render checks the context for the output type and renders the data accordingly using the provided view.
// It supports rendering in JSON, YAML, or table format based on the output type specified in the context.
func (r *Render[T]) Render(ctx context.Context, data T) error {
	output, ok := ctx.Value(enum.OutputDefaultKey.String()).(enum.Type)
	if !ok {
		output = enum.OutputTable
	}
	switch output {
	case enum.OutputJSON:
		return r.view.JSON(data)
	case enum.OutputYAML:
		return r.view.YAML(data)
	default:
		return r.view.Table(data)
	}
}
