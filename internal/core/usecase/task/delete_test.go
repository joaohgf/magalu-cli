//go:build unit

package task

import (
	"context"
	"fmt"
	"testing"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
)

func TestDeleteUseCase(t *testing.T) {
	t.Run("successfully", deleteSuccessfully)
	t.Run("with errors", func(t *testing.T) {
		t.Run("when task is nil", deleteWithNilTask)
		t.Run("when task ID is empty", deleteWithEmptyID)
		t.Run("when persistence returns an errors", deleteWithPersistenceError)
	})
}

func deleteSuccessfully(t *testing.T) {
	deleter := NewDelete(&mockPersistenceDeleter{})
	task := &domain.Task{ID: "1"}
	err := deleter.Delete(t.Context(), task)
	if err != nil {
		t.Fatalf("expected no errors, got %v", err)
	}
}

func deleteWithNilTask(t *testing.T) {
	deleter := NewDelete(&mockPersistenceDeleter{})
	err := deleter.Delete(t.Context(), nil)
	if err == nil {
		t.Fatal("expected errors, got nil")
	}
}

func deleteWithEmptyID(t *testing.T) {
	deleter := NewDelete(&mockPersistenceDeleter{})
	task := &domain.Task{ID: ""}
	err := deleter.Delete(t.Context(), task)
	if err == nil {
		t.Fatal("expected errors, got nil")
	}
}

func deleteWithPersistenceError(t *testing.T) {
	deleter := NewDelete(&mockPersistenceDeleterWithError{})
	task := &domain.Task{ID: "1"}
	err := deleter.Delete(t.Context(), task)
	if err == nil {
		t.Fatal("expected errors, got nil")
	}
}

/*
Mocks below
*/

type (
	mockPersistenceDeleter          struct{}
	mockPersistenceDeleterWithError struct{}
)

func (m *mockPersistenceDeleter) Delete(_ context.Context, _ *domain.Task) error {
	return nil
}

func (m *mockPersistenceDeleterWithError) Delete(_ context.Context, _ *domain.Task) error {
	return fmt.Errorf("persistence errors")
}
