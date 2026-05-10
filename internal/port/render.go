package port

import "context"

type (
	// Render defines the interface for rendering an entity of type T
	Render[T any] interface {
		Render(ctx context.Context, target T) error
	}
	// View defines the interface for rendering an entity of type T in specific formats (JSON, YAML, Table).
	View[T any] interface {
		JSON(T) error
		YAML(T) error
		Table(T) error
	}
	// TableWriter defines the interface for table operations.
	TableWriter interface {
		Header(...any)
		Append(...any) error
		Render() error
		Footer(...any)
		Close() error
	}
)
