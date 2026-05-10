//go:build unit

package task

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/enum"
)

func TestAddUseCase(t *testing.T) {
	t.Run("successfully", addSuccessfully)
	t.Run("with errors", func(t *testing.T) {
		t.Run("validation nil task", addWithNilTask)
		t.Run("validation empty ID", addWithEmptyID)
		t.Run("validation empty title", addWithEmptyTitle)
		t.Run("validation unknown priority", addWithUnknownPriority)
		t.Run("validation unknown status", addWithUnknownStatus)
		t.Run("persistence errors", addWithPersistenceError)
	})
}

func addSuccessfully(t *testing.T) {
	add := NewAdd(&mockPersistenceSaver{})
	task := domain.NewTask("1", "Test Task", "Description", enum.PriorityHigh, time.Now(), "tag1", "tag2")
	saved, err := add.Save(t.Context(), task)
	if err != nil {
		t.Fatalf("expected no errors, got %v", err)
	}
	if saved.ID != task.ID {
		t.Errorf("expected ID %s, got %s", task.ID, saved.ID)
	}
	if saved.Title != task.Title {
		t.Errorf("expected Title %s, got %s", task.Title, saved.Title)
	}
	if saved.Description != task.Description {
		t.Errorf("expected Description %s, got %s", task.Description, saved.Description)
	}
	if saved.Priority != task.Priority {
		t.Errorf("expected Priority %s, got %s", task.Priority.String(), saved.Priority.String())
	}
	if saved.Status != task.Status {
		t.Errorf("expected Status %s, got %s", task.Status.String(), saved.Status.String())
	}
	for i, tag := range task.Tags {
		if saved.Tags[i] != tag {
			t.Errorf("expected Tag %s, got %s", tag, saved.Tags[i])
		}
	}
}

func addWithNilTask(t *testing.T) {
	add := NewAdd(new(mockPersistenceSaver))
	_, err := add.Save(t.Context(), nil)
	if err == nil {
		t.Fatal("expected errors, got nil")
	}
}

func addWithEmptyID(t *testing.T) {
	add := NewAdd(new(mockPersistenceSaver))
	task := domain.NewTask("", "Test Task", "Description", enum.PriorityHigh, time.Now())
	_, err := add.Save(t.Context(), task)
	if err == nil {
		t.Fatal("expected errors, got nil")
	}
}

func addWithEmptyTitle(t *testing.T) {
	add := NewAdd(new(mockPersistenceSaver))
	task := domain.NewTask("1", "", "Description", enum.PriorityHigh, time.Now())
	_, err := add.Save(t.Context(), task)
	if err == nil {
		t.Fatal("expected errors, got nil")
	}
}

func addWithUnknownPriority(t *testing.T) {
	add := NewAdd(new(mockPersistenceSaver))
	task := domain.NewTask("1", "Test Task", "Description", enum.PriorityUnknown, time.Now())
	_, err := add.Save(t.Context(), task)
	if err == nil {
		t.Fatal("expected errors, got nil")
	}
}

func addWithUnknownStatus(t *testing.T) {
	add := NewAdd(new(mockPersistenceSaver))
	task := domain.NewTask("1", "Test Task", "Description", enum.PriorityHigh, time.Now(), "tag1", "tag2")
	task.Status = enum.StatusUnknown
	_, err := add.Save(t.Context(), task)
	if err == nil {
		t.Fatal("expected errors, got nil")
	}
}

func addWithPersistenceError(t *testing.T) {
	add := NewAdd(&mockPersistenceSaverWithError{})
	task := domain.NewTask("1", "Test Task", "Description", enum.PriorityHigh, time.Now(), "tag1", "tag2")
	_, err := add.Save(t.Context(), task)
	if err == nil {
		t.Fatal("expected errors, got nil")
	}
}

/*
Mocks below
*/

type (
	mockPersistenceSaver          struct{}
	mockPersistenceSaverWithError struct{}
)

func (m *mockPersistenceSaver) Save(_ context.Context, task *domain.Task) (*domain.Task, error) {
	return task, nil
}

func (m *mockPersistenceSaverWithError) Save(_ context.Context, _ *domain.Task) (*domain.Task, error) {
	return nil, fmt.Errorf("failed to save task")
}
