//go:build unit

package persistence

import (
	"testing"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestServiceFinder(t *testing.T) {
	t.Run("new successfully", newServiceFinderSuccessfully)
	t.Run("find", func(t *testing.T) {
		t.Run("successfully", findSuccessfully)
		t.Run("with error", findWithError)
	})
}

func newServiceFinderSuccessfully(t *testing.T) {
	db, _ := newTestDB(t)
	finder := NewServiceFinder[*domain.Task](db)
	assert.NotNil(t, finder)
}

func findSuccessfully(t *testing.T) {
	db, _ := newTestDB(t)
	finder := NewServiceFinder[*domain.Task](db)
	task := newTestTask("01KR801MXH2NG3XD5NTY3F2MRS", "Find Me")
	err := db.Write(task.GetCollection(), task.GetID(), task)
	assert.NoError(t, err)
	found, err := finder.Find(t.Context(), task)
	assert.NoError(t, err)
	assert.Equal(t, task.ID, found.ID)
	assert.Equal(t, task.Title, found.Title)
}

func findWithError(t *testing.T) {
	db, _ := newTestDB(t)
	finder := NewServiceFinder[*domain.Task](db)
	task := newTestTask("non-existent-id", "Not Found")
	_, err := finder.Find(t.Context(), task)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "entity not found")
}
