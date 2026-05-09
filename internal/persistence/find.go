package persistence

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/nanobox-io/scribble"
)

var (
	ErrNotFound = errors.New("not found")
)

type ServiceFinder[T port.FilterDomain[T]] struct {
	*scribble.Driver
}

func NewServiceFinder[T port.FilterDomain[T]](db *scribble.Driver) *ServiceFinder[T] {
	return &ServiceFinder[T]{Driver: db}
}

func (s *ServiceFinder[T]) Find(target T) (T, error) {
	var out T
	if err := s.Driver.Read(target.GetCollection(), target.GetID(), &out); err != nil {
		return target, ErrNotFound
	}
	return out, nil
}

func (s *ServiceFinder[T]) FindAll(target T) ([]T, error) {
	var out []T
	data, err := s.Driver.ReadAll(target.GetCollection())
	if err != nil {
		return nil, fmt.Errorf("error finding all: %w", err)
	}
	for _, raw := range data {
		var item T
		if err = json.Unmarshal([]byte(raw), &item); err != nil {
			return nil, fmt.Errorf("error unmarshaling item: %w", err)
		}
		if !target.IsEqual(item) {
			continue
		}
		out = append(out, item)
	}
	maxRange := 10
	if len(out) > maxRange {
		out = out[:maxRange]
	}
	return out, nil
}
