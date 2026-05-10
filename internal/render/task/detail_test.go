//go:build unit

package task

import (
	"testing"
	"time"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/stretchr/testify/assert"
)

func TestDetailView(t *testing.T) {
	t.Run("new successfully", newDetailViewSuccessfully)
	t.Run("table", func(t *testing.T) {
		t.Run("successfully", detailViewTableSuccessfully)
		t.Run("with error", func(t *testing.T) {
			t.Run("append", detailViewTableErrorAppend)
			t.Run("render", detailViewTableErrorRender)
		})
	})
	t.Run("json successfully", detailViewJSONSuccessfully)
	t.Run("yaml successfully", detailViewYAMLSuccessfully)
}

func newDetailViewSuccessfully(t *testing.T) {
	view := NewDetailView(new(mockTable))
	assert.NotNil(t, view)
}

func detailViewTableSuccessfully(t *testing.T) {
	view := NewDetailView(new(mockTable))
	task := &domain.Task{
		ID:          "01KR8021T5FZE79CSANG5FACK0",
		Title:       "Test Task",
		Description: "This is a test task",
		Priority:    enum.PriorityMedium,
		Status:      enum.StatusInProgress,
		CreatedAt:   time.Now(),
		Tags:        []string{"work", "urgent"},
	}
	err := view.Table(task)
	assert.NoError(t, err)
}

func detailViewTableErrorAppend(t *testing.T) {
	view := NewDetailView(new(mockTableWithAppendError))
	task := &domain.Task{
		ID:          "01KR8021T5FZE79CSANG5FACK0",
		Title:       "Test Task",
		Description: "This is a test task",
		Priority:    enum.PriorityMedium,
		Status:      enum.StatusInProgress,
		CreatedAt:   time.Now(),
		Tags:        []string{"work", "urgent"},
	}
	err := view.Table(task)
	assert.Error(t, err)
}

func detailViewTableErrorRender(t *testing.T) {
	view := NewDetailView(new(mockTableWithRenderError))
	task := &domain.Task{
		ID:          "01KR8021T5FZE79CSANG5FACK0",
		Title:       "Test Task",
		Description: "This is a test task",
		Priority:    enum.PriorityMedium,
		Status:      enum.StatusInProgress,
		CreatedAt:   time.Now(),
		Tags:        []string{"work", "urgent"},
	}
	err := view.Table(task)
	assert.Error(t, err)
}

func detailViewJSONSuccessfully(t *testing.T) {
	task := &domain.Task{
		ID:          "01KR8021T5FZE79CSANG5FACK0",
		Title:       "Test Task",
		Description: "This is a test task",
		Priority:    enum.PriorityMedium,
		Status:      enum.StatusInProgress,
		CreatedAt:   time.Now(),
		Tags:        []string{"work"},
	}
	view := NewDetailView(new(mockTable))
	err := view.JSON(task)
	assert.NoError(t, err)
}

func detailViewYAMLSuccessfully(t *testing.T) {
	view := NewDetailView(new(mockTable))
	task := &domain.Task{
		ID:          "01KR8021T5FZE79CSANG5FACK0",
		Title:       "Test Task",
		Description: "This is a test task",
		Priority:    enum.PriorityMedium,
		Status:      enum.StatusInProgress,
		CreatedAt:   time.Now(),
		Tags:        []string{"work"},
	}
	err := view.YAML(task)
	assert.NoError(t, err)
}
