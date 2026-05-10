package port

import "context"

type (
	// PersistenceSaver defines the interface for saving an entity of type T to a persistence layer.
	PersistenceSaver[T Domain] interface {
		Save(ctx context.Context, target T) (T, error)
	}
	// PersistenceFinder defines the interface for finding an entity of type T in a persistence layer.
	PersistenceFinder[T Domain] interface {
		Find(ctx context.Context, target T) (T, error)
	}
	PersistenceLister[T Domain, F FilterDomain[T]] interface {
		List(ctx context.Context, target F) (F, error)
	}
	PersistenceDeleter[T Domain] interface {
		Delete(ctx context.Context, target T) error
	}
)
