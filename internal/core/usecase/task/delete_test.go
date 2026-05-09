//go:build unit

package task

import (
	"fmt"
	"testing"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
)

func TestDeleteUseCase(t *testing.T) {
	t.Run("successfully", deleteSuccessfully)
	t.Run("with error", func(t *testing.T) {
		t.Run("when task is nil", deleteWithNilTask)
		t.Run("when task ID is empty", deleteWithEmptyID)
		t.Run("when persistence returns an error", deleteWithPersistenceError)
	})
}

func deleteSuccessfully(t *testing.T) {
	deleter := NewDelete(&mockPersistenceDeleter{})
	task := &domain.Task{ID: "1"}
	err := deleter.Delete(task)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func deleteWithNilTask(t *testing.T) {
	deleter := NewDelete(&mockPersistenceDeleter{})
	err := deleter.Delete(nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func deleteWithEmptyID(t *testing.T) {
	deleter := NewDelete(&mockPersistenceDeleter{})
	task := &domain.Task{ID: ""}
	err := deleter.Delete(task)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func deleteWithPersistenceError(t *testing.T) {
	deleter := NewDelete(&mockPersistenceDeleterWithError{})
	task := &domain.Task{ID: "1"}
	err := deleter.Delete(task)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

/*
Mocks below
*/

type (
	mockPersistenceDeleter          struct{}
	mockPersistenceDeleterWithError struct{}
)

func (m *mockPersistenceDeleter) Delete(_ *domain.Task) error {
	return nil
}

func (m *mockPersistenceDeleterWithError) Delete(_ *domain.Task) error {
	return fmt.Errorf("persistence error")
}
