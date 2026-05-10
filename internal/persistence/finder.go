package persistence

import (
	"context"
	"fmt"

	"github.com/joaohgf/magalu-cli/internal/errors"
	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/nanobox-io/scribble"
)

// ServiceFinder is responsible for finding a domain entity in the persistence layer using the scribble driver.
type ServiceFinder[T port.Domain] struct {
	*scribble.Driver
}

func NewServiceFinder[T port.Domain](db *scribble.Driver) *ServiceFinder[T] {
	return &ServiceFinder[T]{Driver: db}
}

// Find retrieves an entity from the persistence layer based on the collection and ID provided by the target.
func (s *ServiceFinder[T]) Find(_ context.Context, target T) (T, error) {
	var out T
	if err := s.Driver.Read(target.GetCollection(), target.GetID(), &out); err != nil {
		return target, errors.NotFound(fmt.Sprintf("%s with ID %s", target.GetCollection(), target.GetID()))
	}
	return out, nil
}
