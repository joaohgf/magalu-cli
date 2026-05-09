//go:build unit

package task

import (
	"fmt"
	"testing"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
)

func TestFindUseCase(t *testing.T) {
	t.Run("successfully", findSuccessfully)
	t.Run("with error", func(t *testing.T) {
		t.Run("when task ID is empty", findWithEmptyTaskID)
		t.Run("when finding task", findWithErrorFindingTask)
	})
}

func findSuccessfully(t *testing.T) {
	find := NewFind(new(mockPersistenceFinder))
	task := &domain.Task{ID: "1"}
	foundTask, err := find.Find(task)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if foundTask == nil {
		t.Fatal("expected found task, got nil")
	}
	if foundTask.ID != task.ID {
		t.Fatalf("expected found task ID to be %s, got %s", task.ID, foundTask.ID)
	}
}

func findWithEmptyTaskID(t *testing.T) {
	find := NewFind(new(mockPersistenceFinder))
	task := &domain.Task{ID: ""}
	_, err := find.Find(task)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func findWithErrorFindingTask(t *testing.T) {
	find := NewFind(new(mockPersistenceFinderWithError))
	task := &domain.Task{ID: "1"}
	_, err := find.Find(task)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

/*
Mocks below
*/

type (
	mockPersistenceFinder          struct{}
	mockPersistenceFinderWithError struct{}
)

func (m *mockPersistenceFinder) Find(target *domain.Task) (*domain.Task, error) {
	return &domain.Task{ID: target.ID}, nil
}
func (m *mockPersistenceFinder) FindAll(_ *domain.Task) (*domain.Task, error) {
	return nil, nil
}
func (m *mockPersistenceFinderWithError) Find(_ *domain.Task) (*domain.Task, error) {
	return nil, fmt.Errorf("error finding task")
}
