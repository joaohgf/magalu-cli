package port

// Domain represents a generic entity in the application.
type Domain interface {
	// GetCollection returns the name of the collection(folder) where this domain entity is stored.
	GetCollection() string
	// GetID returns the unique identifier of this domain entity.
	GetID() string
}

// FilterDomain represents a domain entity that can be used for filtering operations.
type FilterDomain[T any] interface {
	Domain
	// IsEqual compares the current filter with another filter and returns true if they are considered equal.
	IsEqual(other T) bool
	GetPage() int
	GetSize() int
	GetTotal() int
	GetContent() []T
	SetSize(size int)
	SetTotal(total int)
	SetPage(page int)
	SetContent(content ...T)
}
