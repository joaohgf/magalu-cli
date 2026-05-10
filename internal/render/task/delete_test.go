//go:build unit

package task

import (
	"testing"
	"time"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/joaohgf/magalu-cli/internal/errors"
	"github.com/stretchr/testify/assert"
)

func TestDeleteView(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		t.Run("successfully", newDeleteViewSuccessfully)
	})
	t.Run("table", func(t *testing.T) {
		t.Run("successfully", deleteViewTableSuccessfully)
		t.Run("with error", func(t *testing.T) {
			t.Run("append", deleteViewTableErrorAppend)
			t.Run("render", deleteViewTableErrorRender)
		})
	})
}

func newDeleteViewSuccessfully(t *testing.T) {
	view := NewDeleteView(new(mockTable))
	assert.NotNil(t, view)
	assert.NotNil(t, view.DetailView)
}

func deleteViewTableSuccessfully(t *testing.T) {
	view := NewDeleteView(new(mockTable))
	task := &domain.Task{
		ID:        "01KR8021T5FZE79CSANG5FACK0",
		Title:     "Test Task",
		Priority:  enum.PriorityMedium,
		Status:    enum.StatusInProgress,
		CreatedAt: time.Now(),
	}
	err := view.Table(task)
	assert.NoError(t, err)
}

func deleteViewTableErrorAppend(t *testing.T) {
	view := NewDeleteView(new(mockTableWithAppendError))
	task := &domain.Task{
		ID:        "01KR8021T5FZE79CSANG5FACK0",
		Title:     "Test Task",
		Priority:  enum.PriorityMedium,
		Status:    enum.StatusInProgress,
		CreatedAt: time.Now(),
	}
	err := view.Table(task)
	assert.Error(t, err)
}

func deleteViewTableErrorRender(t *testing.T) {
	view := NewDeleteView(new(mockTableWithRenderError))
	task := &domain.Task{
		ID:        "01KR8021T5FZE79CSANG5FACK0",
		Title:     "Test Task",
		Priority:  enum.PriorityMedium,
		Status:    enum.StatusInProgress,
		CreatedAt: time.Now(),
	}
	err := view.Table(task)
	assert.Error(t, err)
}

/*
Mocks below
*/

type (
	mockTable                struct{}
	mockTableWithAppendError struct {
		mockTable
	}
	mockTableWithRenderError struct {
		mockTable
	}
)

func (m *mockTable) Header(_ ...any) {}

func (m *mockTable) Render() error {
	return nil
}

func (m *mockTable) Append(_ ...any) error {
	return nil
}

func (m *mockTable) Footer(_ ...any) {}

func (m *mockTable) Close() error {
	return nil
}

func (m *mockTableWithAppendError) Append(_ ...any) error {
	return errors.New("mock append error")
}
func (m *mockTableWithRenderError) Render() error {
	return errors.New("mock render error")
}
