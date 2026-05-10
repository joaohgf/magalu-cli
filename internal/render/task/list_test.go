//go:build unit

package task

import (
	"testing"
	"time"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		t.Run("successfully", newListSuccessfully)
	})
	t.Run("table", func(t *testing.T) {
		t.Run("successfully", listTableSuccessfully)
		t.Run("with no tasks", listTableNoTasks)
		t.Run("with error", func(t *testing.T) {
			t.Run("append", listTableErrorAppend)
			t.Run("render", listTableErrorRender)
		})
	})
	t.Run("json successfully", listJSONSuccessfully)
	t.Run("yaml successfully", listYAMLSuccessfully)
}

func newListSuccessfully(t *testing.T) {
	list := NewList(new(mockTable))
	assert.NotNil(t, list)
}

func listTableSuccessfully(t *testing.T) {
	list := NewList(new(mockTable))
	filter := &domain.TaskFilter{
		Page:  1,
		Size:  10,
		Total: 2,
		Data: []*domain.Task{
			{
				ID:        "01KR8021T5FZE79CSANG5FACK0",
				Title:     "Task 1",
				Priority:  enum.PriorityHigh,
				Status:    enum.StatusInProgress,
				CreatedAt: time.Now(),
			},
			{
				ID:        "01KR801MXH2NG3XD5NTY3F2MRS",
				Title:     "Task 2",
				Priority:  enum.PriorityMedium,
				Status:    enum.StatusDone,
				CreatedAt: time.Now(),
			},
		},
	}
	err := list.Table(filter)
	assert.NoError(t, err)
}

func listTableNoTasks(t *testing.T) {
	list := NewList(new(mockTable))
	filter := &domain.TaskFilter{
		Page:  1,
		Size:  10,
		Total: 0,
		Data:  []*domain.Task{},
	}
	err := list.Table(filter)
	assert.NoError(t, err)
}

func listTableErrorAppend(t *testing.T) {
	list := NewList(new(mockTableWithAppendError))
	filter := &domain.TaskFilter{
		Page:  1,
		Size:  10,
		Total: 1,
		Data: []*domain.Task{
			{
				ID:        "01KR8021T5FZE79CSANG5FACK0",
				Title:     "Task 1",
				Priority:  enum.PriorityHigh,
				Status:    enum.StatusInProgress,
				CreatedAt: time.Now(),
			},
		},
	}
	err := list.Table(filter)
	assert.Error(t, err)
}

func listTableErrorRender(t *testing.T) {
	list := NewList(new(mockTableWithRenderError))
	filter := &domain.TaskFilter{
		Page:  1,
		Size:  10,
		Total: 1,
		Data: []*domain.Task{
			{
				ID:        "01KR8021T5FZE79CSANG5FACK0",
				Title:     "Task 1",
				Priority:  enum.PriorityHigh,
				Status:    enum.StatusInProgress,
				CreatedAt: time.Now(),
			},
		},
	}
	err := list.Table(filter)
	assert.Error(t, err)
}

func listJSONSuccessfully(t *testing.T) {
	list := NewList(new(mockTable))
	filter := &domain.TaskFilter{
		Page:  1,
		Size:  10,
		Total: 2,
		Data: []*domain.Task{
			{
				ID:        "01KR8021T5FZE79CSANG5FACK0",
				Title:     "Task 1",
				Priority:  enum.PriorityHigh,
				Status:    enum.StatusInProgress,
				CreatedAt: time.Now(),
			},
			{
				ID:        "01KR801MXH2NG3XD5NTY3F2MRS",
				Title:     "Task 2",
				Priority:  enum.PriorityMedium,
				Status:    enum.StatusDone,
				CreatedAt: time.Now(),
			},
		},
	}
	err := list.JSON(filter)
	assert.NoError(t, err)
}

func listYAMLSuccessfully(t *testing.T) {
	list := NewList(new(mockTable))
	filter := &domain.TaskFilter{
		Page:  1,
		Size:  10,
		Total: 2,
		Data: []*domain.Task{
			{
				ID:        "01KR8021T5FZE79CSANG5FACK0",
				Title:     "Task 1",
				Priority:  enum.PriorityHigh,
				Status:    enum.StatusInProgress,
				CreatedAt: time.Now(),
			},
			{
				ID:        "01KR801MXH2NG3XD5NTY3F2MRS",
				Title:     "Task 2",
				Priority:  enum.PriorityMedium,
				Status:    enum.StatusDone,
				CreatedAt: time.Now(),
			},
		},
	}
	err := list.YAML(filter)
	assert.NoError(t, err)
}
