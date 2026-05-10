//go:build unit

package task

import (
	"context"
	"testing"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/errors"
)

func TestList(t *testing.T) {
	t.Run("successfully", listSuccessfully)
	t.Run("with error", listWithError)
}

func listSuccessfully(t *testing.T) {
	useCase := NewList(new(mockPersistenceList))
	target := domain.NewTaskFilter(nil)
	result, err := useCase.All(t.Context(), target)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected a result, got nil")
	}
}

func listWithError(t *testing.T) {
	useCase := NewList(new(mockPersistenceListWithError))
	target := domain.NewTaskFilter(nil)
	_, err := useCase.All(t.Context(), target)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

/*
Mocks below
*/

type (
	mockPersistenceList          struct{}
	mockPersistenceListWithError struct{}
)

func (m *mockPersistenceList) List(_ context.Context, filter *domain.TaskFilter) (*domain.TaskFilter, error) {
	return filter, nil
}

func (m *mockPersistenceListWithError) List(_ context.Context, _ *domain.TaskFilter) (*domain.TaskFilter, error) {
	return nil, errors.New("failed to list tasks")
}
