package persistence

import (
	"fmt"

	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/nanobox-io/scribble"
)

type ServiceSaver[T port.Domain] struct {
	*scribble.Driver
}

func NewServiceSaver[T port.Domain](db *scribble.Driver) *ServiceSaver[T] {
	return &ServiceSaver[T]{Driver: db}
}

func (s *ServiceSaver[T]) Save(target T) (T, error) {
	if err := s.Driver.Write(target.GetCollection(), target.GetID(), target); err != nil {
		return target, fmt.Errorf("error saving: %w", err)
	}
	return target, nil
}
