package port

type (
	// SaveUseCase defines use case for saving an entity of type T.
	SaveUseCase[T any] interface {
		Save(target T) (T, error)
	}
	// FindUseCase defines use case for finding an entity of type T.
	FindUseCase[T any] interface {
		Find(target T) (T, error)
	}
	// FindAllUseCase defines use case for finding all entities of type T.
	FindAllUseCase[T any] interface {
		All(target T) ([]T, error)
	}
	// DeleteUseCase defines use case for deleting an entity of type T.
	DeleteUseCase[T any] interface {
		Delete(target T) error
	}
)
