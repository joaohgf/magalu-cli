package persistence

import (
	"context"
	"fmt"

	"github.com/joaohgf/magalu-cli/internal/errors"
	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/nanobox-io/scribble"
)

type ServiceFinder[T port.Domain] struct {
	*scribble.Driver
}

func NewServiceFinder[T port.Domain](db *scribble.Driver) *ServiceFinder[T] {
	return &ServiceFinder[T]{Driver: db}
}

func (s *ServiceFinder[T]) Find(_ context.Context, target T) (T, error) {
	var out T
	if err := s.Driver.Read(target.GetCollection(), target.GetID(), &out); err != nil {
		return target, errors.NotFound(fmt.Sprintf("%s with ID %s", target.GetCollection(), target.GetID()))
	}
	return out, nil
}
