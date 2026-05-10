package port

import "context"

type (
	Render[T any] interface {
		Render(ctx context.Context, target T) error
	}
	View[T any] interface {
		JSON(T) error
		YAML(T) error
		Table(T) error
	}
)
