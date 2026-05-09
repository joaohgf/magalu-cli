//go:build unit

package task

import (
	"fmt"
	"testing"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/enum"
)

func TestDoneUseCase(t *testing.T) {
	t.Run("successfully", doneSuccessfully)
	t.Run("with error", func(t *testing.T) {
		t.Run("when task is nil", doneWithNilTask)
		t.Run("when task ID is empty", doneWithEmptyTaskID)
		t.Run("when finding task", doneWithErrorFindingTask)
		t.Run("when saving task", doneWithErrorSavingTask)
	})
}

func doneSuccessfully(t *testing.T) {
	done := NewDone(new(mockPersistenceSaver), new(mockPersistenceFinder))
	task := &domain.Task{ID: "1"}
	updatedTask, err := done.Save(task)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if updatedTask == nil {
		t.Fatal("expected updated task, got nil")
	}
	if updatedTask.Status != enum.StatusDone {
		t.Fatal("expected task to be marked as done")
	}
	if updatedTask.DoneAt == nil || updatedTask.DoneAt.IsZero() {
		t.Fatal("expected DoneAt to be set")
	}
}

func doneWithNilTask(t *testing.T) {
	done := NewDone(new(mockPersistenceSaver), new(mockPersistenceFinder))
	_, err := done.Save(nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func doneWithEmptyTaskID(t *testing.T) {
	done := NewDone(new(mockPersistenceSaver), new(mockPersistenceFinder))
	task := &domain.Task{ID: ""}
	_, err := done.Save(task)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func doneWithErrorFindingTask(t *testing.T) {
	done := NewDone(new(mockPersistenceSaver), new(mockPersistenceFinderWithError))
	task := &domain.Task{ID: "1"}
	_, err := done.Save(task)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func doneWithErrorSavingTask(t *testing.T) {
	done := NewDone(new(mockPersistenceSaverWithError), new(mockPersistenceFinder))
	task := &domain.Task{ID: "1"}
	_, err := done.Save(task)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

/*
Mocks below
*/

type (
	mockPersistenceSaver           struct{}
	mockPersistenceSaverWithError  struct{}
	mockPersistenceFinder          struct{}
	mockPersistenceFinderWithError struct {
		mockPersistenceFinder
	}
)

func (m *mockPersistenceSaver) Save(task *domain.Task) (*domain.Task, error) {
	return task, nil
}

func (m *mockPersistenceSaverWithError) Save(_ *domain.Task) (*domain.Task, error) {
	return nil, fmt.Errorf("failed to save task")
}

func (m *mockPersistenceFinder) Find(task *domain.Task) (*domain.Task, error) {
	return task, nil
}

func (m *mockPersistenceFinder) FindAll(_ *domain.Task) ([]*domain.Task, error) {
	return nil, nil
}

func (m *mockPersistenceFinderWithError) Find(_ *domain.Task) (*domain.Task, error) {
	return nil, fmt.Errorf("failed to find task")
}
