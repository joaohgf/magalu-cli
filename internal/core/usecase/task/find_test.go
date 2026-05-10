//go:build unit

package task

import (
	"context"
	"fmt"
	"testing"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestFindUseCase(t *testing.T) {
	t.Run("successfully", findSuccessfully)
	t.Run("with errors", func(t *testing.T) {
		t.Run("when task ID is empty", findWithEmptyTaskID)
		t.Run("when finding task", findWithErrorFindingTask)
	})
}

func findSuccessfully(t *testing.T) {
	find := NewFind(new(mockPersistenceFinder))
	task := &domain.Task{ID: "1"}
	foundTask, err := find.Find(t.Context(), task)
	assert.NoError(t, err)
	assert.Equal(t, task.ID, foundTask.ID)
}

func findWithEmptyTaskID(t *testing.T) {
	find := NewFind(new(mockPersistenceFinder))
	task := &domain.Task{ID: ""}
	_, err := find.Find(t.Context(), task)
	assert.Error(t, err)
}

func findWithErrorFindingTask(t *testing.T) {
	find := NewFind(new(mockPersistenceFinderWithError))
	task := &domain.Task{ID: "1"}
	_, err := find.Find(t.Context(), task)
	assert.Error(t, err)
}

/*
Mocks below
*/

type (
	mockPersistenceFinder          struct{}
	mockPersistenceFinderWithError struct{}
)

func (m *mockPersistenceFinder) Find(_ context.Context, target *domain.Task) (*domain.Task, error) {
	return &domain.Task{ID: target.ID}, nil
}
func (m *mockPersistenceFinder) FindAll(_ context.Context, _ *domain.Task) (*domain.Task, error) {
	return nil, nil
}
func (m *mockPersistenceFinderWithError) Find(_ context.Context, _ *domain.Task) (*domain.Task, error) {
	return nil, fmt.Errorf("errors finding task")
}
