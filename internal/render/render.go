package render

import (
	"context"

	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/joaohgf/magalu-cli/internal/port"
)

type Render[T any] struct {
	view port.View[T]
}

func NewRender[T any](view port.View[T]) *Render[T] {
	return &Render[T]{view: view}
}

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
