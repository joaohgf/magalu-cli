//go:build unit

package persistence

import (
	"testing"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestServiceSaver(t *testing.T) {
	t.Run("new successfully", newServiceSaverSuccessfully)
	t.Run("save", func(t *testing.T) {
		t.Run("successfully", saveSuccessfully)
		t.Run("with error", saveWithError)
	})
}

func newServiceSaverSuccessfully(t *testing.T) {
	db, _ := newTestDB(t)
	saver := NewServiceSaver[*domain.Task](db)
	assert.NotNil(t, saver)
}

func saveSuccessfully(t *testing.T) {
	db, _ := newTestDB(t)
	saver := NewServiceSaver[*domain.Task](db)
	task := newTestTask("01KR8021T5FZE79CSANG5FACK0", "Test Task")
	saved, err := saver.Save(t.Context(), task)
	assert.NoError(t, err)
	assert.Equal(t, task.ID, saved.ID)
	assert.Equal(t, task.Title, saved.Title)
}

func saveWithError(t *testing.T) {
	db, _ := newTestDB(t)
	saver := NewServiceSaver[*domain.Task](db)
	task := newTestTask("", "Invalid Task")
	_, err := saver.Save(t.Context(), task)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "failed to save the entity")
}
