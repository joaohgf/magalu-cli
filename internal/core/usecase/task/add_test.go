//go:build unit

package task

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)
	assert.Equal(t, task.ID, saved.ID)
	assert.Equal(t, task.Title, saved.Title)
	assert.Equal(t, task.Description, saved.Description)
	assert.Equal(t, task.Priority, saved.Priority)
	assert.Equal(t, task.Status, saved.Status)
	assert.Equal(t, task.Tags, saved.Tags)
}

func addWithNilTask(t *testing.T) {
	add := NewAdd(new(mockPersistenceSaver))
	_, err := add.Save(t.Context(), nil)
	assert.Error(t, err)
}

func addWithEmptyID(t *testing.T) {
	add := NewAdd(new(mockPersistenceSaver))
	task := domain.NewTask("", "Test Task", "Description", enum.PriorityHigh, time.Now())
	_, err := add.Save(t.Context(), task)
	assert.Error(t, err)
}

func addWithEmptyTitle(t *testing.T) {
	add := NewAdd(new(mockPersistenceSaver))
	task := domain.NewTask("1", "", "Description", enum.PriorityHigh, time.Now())
	_, err := add.Save(t.Context(), task)
	assert.Error(t, err)
}

func addWithUnknownPriority(t *testing.T) {
	add := NewAdd(new(mockPersistenceSaver))
	task := domain.NewTask("1", "Test Task", "Description", enum.PriorityUnknown, time.Now())
	_, err := add.Save(t.Context(), task)
	assert.Error(t, err)
}

func addWithUnknownStatus(t *testing.T) {
	add := NewAdd(new(mockPersistenceSaver))
	task := domain.NewTask("1", "Test Task", "Description", enum.PriorityHigh, time.Now(), "tag1", "tag2")
	task.Status = enum.StatusUnknown
	_, err := add.Save(t.Context(), task)
	assert.Error(t, err)
}

func addWithPersistenceError(t *testing.T) {
	add := NewAdd(&mockPersistenceSaverWithError{})
	task := domain.NewTask("1", "Test Task", "Description", enum.PriorityHigh, time.Now(), "tag1", "tag2")
	_, err := add.Save(t.Context(), task)
	assert.Error(t, err)
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
