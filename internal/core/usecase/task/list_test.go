//go:build unit

package task

import (
	"context"
	"testing"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/errors"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	t.Run("successfully", listSuccessfully)
	t.Run("with error", listWithError)
}

func listSuccessfully(t *testing.T) {
	useCase := NewList(new(mockPersistenceList))
	target := domain.NewTaskFilter(nil)
	result, err := useCase.All(t.Context(), target)
	assert.NoError(t, err)
	assert.Equal(t, target, result)
}

func listWithError(t *testing.T) {
	useCase := NewList(new(mockPersistenceListWithError))
	target := domain.NewTaskFilter(nil)
	_, err := useCase.All(t.Context(), target)
	assert.Error(t, err)
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
