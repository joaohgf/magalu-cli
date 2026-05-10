package persistence

import (
	"context"
	"fmt"

	"github.com/joaohgf/magalu-cli/internal/errors"

	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/nanobox-io/scribble"
)

// ServiceDeleter is responsible for deleting a domain entity from the persistence layer using the scribble driver.
type ServiceDeleter[T port.Domain] struct {
	*scribble.Driver
}

func NewServiceDeleter[T port.Domain](db *scribble.Driver) *ServiceDeleter[T] {
	return &ServiceDeleter[T]{Driver: db}
}

// Delete removes the target entity from the persistence layer using the scribble driver.
func (s *ServiceDeleter[T]) Delete(_ context.Context, target T) error {
	if err := s.Driver.Delete(target.GetCollection(), target.GetID()); err != nil {
		return errors.NotDeleted(fmt.Sprintf("%s with ID %s", target.GetCollection(), target.GetID()))
	}
	return nil
}
