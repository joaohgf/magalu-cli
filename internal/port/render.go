package port

type (
	Render[T any] interface {
		Render(target ...T) error
	}
)
