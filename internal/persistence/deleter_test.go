//go:build unit

package persistence

import (
	"testing"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestServiceDeleter(t *testing.T) {
	t.Run("new successfully", newServiceDeleterSuccessfully)
	t.Run("delete", func(t *testing.T) {
		t.Run("successfully", deleteSuccessfully)
		t.Run("with error", deleteWithError)
	})
}

func newServiceDeleterSuccessfully(t *testing.T) {
	db, _ := newTestDB(t)
	deleter := NewServiceDeleter[*domain.Task](db)
	assert.NotNil(t, deleter)
}

func deleteSuccessfully(t *testing.T) {
	db, _ := newTestDB(t)
	deleter := NewServiceDeleter[*domain.Task](db)
	task := newTestTask("01KR8EEX18E7GY9FHC6FFA8D9G", "Delete Me")
	err := db.Write(task.GetCollection(), task.GetID(), task)
	assert.NoError(t, err)
	err = deleter.Delete(t.Context(), task)
	assert.NoError(t, err)
}

func deleteWithError(t *testing.T) {
	db, _ := newTestDB(t)
	deleter := NewServiceDeleter[*domain.Task](db)
	task := newTestTask("missing-id", "Not Existing")
	err := deleter.Delete(t.Context(), task)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "failed to delete the entity")
}
