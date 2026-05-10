package persistence

import (
	"context"
	"fmt"

	"github.com/joaohgf/magalu-cli/internal/errors"
	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/nanobox-io/scribble"
)

type ServiceSaver[T port.Domain] struct {
	*scribble.Driver
}

func NewServiceSaver[T port.Domain](db *scribble.Driver) *ServiceSaver[T] {
	return &ServiceSaver[T]{Driver: db}
}

func (s *ServiceSaver[T]) Save(_ context.Context, target T) (T, error) {
	if err := s.Driver.Write(target.GetCollection(), target.GetID(), target); err != nil {
		return target, errors.NotSaved(fmt.Sprintf("%v with ID %s", target.GetCollection(), target.GetID()))
	}
	return target, nil
}
