package port

import "context"

type (
	// SaveUseCase defines use case for saving an entity of type T.
	SaveUseCase[T any] interface {
		Save(ctx context.Context, target T) (T, error)
	}
	// FindUseCase defines use case for finding an entity of type T.
	FindUseCase[T any] interface {
		Find(ctx context.Context, target T) (T, error)
	}
	// FindAllUseCase defines use case for finding all entities of type T.
	FindAllUseCase[T any] interface {
		All(ctx context.Context, target T) (T, error)
	}
	// DeleteUseCase defines use case for deleting an entity of type T.
	DeleteUseCase[T any] interface {
		Delete(ctx context.Context, target T) error
	}
)
