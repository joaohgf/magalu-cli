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
	// PersistenceLister defines the interface for listing entities of type T based on a filter of type F in a persistence layer.
	PersistenceLister[T Domain, F FilterDomain[T]] interface {
		List(ctx context.Context, target F) (F, error)
	}
	// PersistenceDeleter defines the interface for deleting an entity of type T from a persistence layer.
	PersistenceDeleter[T Domain] interface {
		Delete(ctx context.Context, target T) error
	}
)
