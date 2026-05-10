//go:build unit

package task

import (
	"testing"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/stretchr/testify/assert"
)

func TestDoneUseCase(t *testing.T) {
	t.Run("successfully", doneSuccessfully)
	t.Run("with errors", func(t *testing.T) {
		t.Run("when task is nil", doneWithNilTask)
		t.Run("when task ID is empty", doneWithEmptyTaskID)
		t.Run("when finding task", doneWithErrorFindingTask)
		t.Run("when saving task", doneWithErrorSavingTask)
	})
}

func doneSuccessfully(t *testing.T) {
	done := NewDone(new(mockPersistenceSaver), new(mockPersistenceFinder))
	task := &domain.Task{ID: "1"}
	updatedTask, err := done.Save(t.Context(), task)
	assert.NoError(t, err)
	assert.Equal(t, task.ID, updatedTask.ID)
	assert.Equal(t, enum.StatusDone, updatedTask.Status)
	assert.NotNil(t, updatedTask.DoneAt)
}

func doneWithNilTask(t *testing.T) {
	done := NewDone(new(mockPersistenceSaver), new(mockPersistenceFinder))
	_, err := done.Save(t.Context(), nil)
	assert.Error(t, err)
}

func doneWithEmptyTaskID(t *testing.T) {
	done := NewDone(new(mockPersistenceSaver), new(mockPersistenceFinder))
	task := &domain.Task{ID: ""}
	_, err := done.Save(t.Context(), task)
	assert.Error(t, err)
}

func doneWithErrorFindingTask(t *testing.T) {
	done := NewDone(new(mockPersistenceSaver), new(mockPersistenceFinderWithError))
	task := &domain.Task{ID: "1"}
	_, err := done.Save(t.Context(), task)
	assert.Error(t, err)
}

func doneWithErrorSavingTask(t *testing.T) {
	done := NewDone(new(mockPersistenceSaverWithError), new(mockPersistenceFinder))
	task := &domain.Task{ID: "1"}
	_, err := done.Save(t.Context(), task)
	assert.Error(t, err)
}
