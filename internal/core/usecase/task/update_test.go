//go:build unit

package task

import (
	"testing"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/enum"
)

func TestUpdateUseCase(t *testing.T) {
	t.Run("successfully", updateSuccessfully)
	t.Run("with errors", func(t *testing.T) {
		t.Run("task nil", updateWithNilTask)
		t.Run("task without ID", updateWithoutID)
		t.Run("not found", updateNotFound)
		t.Run("invalid priority", updateWithInvalidPriority)
		t.Run("without any field", updateWithoutAnyField)
		t.Run("errors saving", updateErrorSaving)
	})
}

func updateSuccessfully(t *testing.T) {
	update := NewUpdate(new(mockPersistenceSaver), new(mockPersistenceFinder))
	task := &domain.Task{ID: "1", Title: "Test Task"}
	updatedTask, err := update.Save(t.Context(), task)
	if err != nil {
		t.Fatalf("expected no errors, got %v", err)
	}
	if updatedTask == nil {
		t.Fatal("expected updated task, got nil")
	}
	if updatedTask.ID != task.ID {
		t.Fatal("expected task ID to remain unchanged")
	}
	if updatedTask.Title != task.Title {
		t.Fatal("expected task title to be updated")
	}
}

func updateWithNilTask(t *testing.T) {
	update := NewUpdate(new(mockPersistenceSaver), new(mockPersistenceFinder))
	_, err := update.Save(t.Context(), nil)
	if err == nil {
		t.Fatal("expected errors, got nil")
	}
}

func updateWithoutID(t *testing.T) {
	update := NewUpdate(new(mockPersistenceSaver), new(mockPersistenceFinder))
	task := &domain.Task{Title: "Test Task"}
	_, err := update.Save(t.Context(), task)
	if err == nil {
		t.Fatal("expected errors, got nil")
	}
}

func updateNotFound(t *testing.T) {
	update := NewUpdate(new(mockPersistenceSaver), new(mockPersistenceFinderWithError))
	task := &domain.Task{ID: "1", Title: "Test Task"}
	_, err := update.Save(t.Context(), task)
	if err == nil {
		t.Fatal("expected errors, got nil")
	}
}

func updateWithInvalidPriority(t *testing.T) {
	update := NewUpdate(new(mockPersistenceSaver), new(mockPersistenceFinder))
	task := &domain.Task{ID: "1", Priority: enum.PriorityUnknown}
	_, err := update.Save(t.Context(), task)
	if err == nil {
		t.Fatal("expected errors, got nil")
	}
}

func updateWithoutAnyField(t *testing.T) {
	update := NewUpdate(new(mockPersistenceSaver), new(mockPersistenceFinder))
	task := &domain.Task{ID: "1"}
	_, err := update.Save(t.Context(), task)
	if err == nil {
		t.Fatal("expected errors, got nil")
	}
}

func updateErrorSaving(t *testing.T) {
	update := NewUpdate(new(mockPersistenceSaverWithError), new(mockPersistenceFinder))
	task := &domain.Task{ID: "1", Title: "Test Task"}
	_, err := update.Save(t.Context(), task)
	if err == nil {
		t.Fatal("expected errors, got nil")
	}
}
