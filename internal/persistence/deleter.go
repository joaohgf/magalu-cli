package persistence

import (
	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/nanobox-io/scribble"
)

type ServiceDeleter[T port.Domain] struct {
	*scribble.Driver
}

func NewServiceDeleter[T port.Domain](db *scribble.Driver) *ServiceDeleter[T] {
	return &ServiceDeleter[T]{Driver: db}
}

func (s *ServiceDeleter[T]) Delete(target T) error {
	if err := s.Driver.Delete(target.GetCollection(), target.GetID()); err != nil {
		return ErrNotFound
	}
	return nil
}
