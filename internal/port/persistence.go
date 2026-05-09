package port

type (
	// PersistenceSaver defines the interface for saving an entity of type T to a persistence layer.
	PersistenceSaver[T Domain] interface {
		Save(target T) (T, error)
	}
	// PersistenceFinder defines the interface for finding an entity of type T in a persistence layer.
	PersistenceFinder[T Domain] interface {
		Find(target T) (T, error)
		FindAll(target T) ([]T, error)
	}
	PersistenceDeleter[T Domain] interface {
		Delete(target T) error
	}
)
