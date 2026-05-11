//go:build unit

package persistence

import (
	"context"
	"testing"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestServiceLister(t *testing.T) {
	t.Run("new successfully", newServiceListerSuccessfully)
	t.Run("list", func(t *testing.T) {
		t.Run("successfully", listSuccessfully)
		t.Run("empty collection", listWithEmptyCollection)
		t.Run("with filter", listWithFilter)
		t.Run("with invalid json", listWithInvalidJSON)
	})
}

func newServiceListerSuccessfully(t *testing.T) {
	db, _ := newTestDB(t)
	lister := NewServiceLister[*domain.Task, *domain.TaskFilter](db)
	assert.NotNil(t, lister)
}

func listSuccessfully(t *testing.T) {
	db, _ := newTestDB(t)
	lister := NewServiceLister[*domain.Task, *domain.TaskFilter](db)
	task1 := newTestTask("01KR8EEVQ0T1QJ2EDN5B0767SH", "Task 1")
	task2 := newTestTask("01KR8EEWBQKXB8F2YMVQKBR1TC", "Task 2")
	assert.NoError(t, db.Write(task1.GetCollection(), task1.GetID(), task1))
	assert.NoError(t, db.Write(task2.GetCollection(), task2.GetID(), task2))

	filter := domain.NewTaskFilter(&domain.Task{})
	filter.SetPage(1)
	filter.SetSize(10)

	out, err := lister.List(context.Background(), filter)

	assert.NoError(t, err)
	assert.Len(t, out.GetContent(), 2)
	assert.Equal(t, 2, out.GetTotal())
}

func listWithEmptyCollection(t *testing.T) {
	db, _ := newTestDB(t)
	lister := NewServiceLister[*domain.Task, *domain.TaskFilter](db)
	filter := domain.NewTaskFilter(&domain.Task{})

	out, err := lister.List(context.Background(), filter)

	assert.NoError(t, err)
	assert.Empty(t, out.GetContent())
	assert.Equal(t, 0, out.GetTotal())
}

func listWithFilter(t *testing.T) {
	db, _ := newTestDB(t)
	lister := NewServiceLister[*domain.Task, *domain.TaskFilter](db)
	task1 := newTestTask("01KR8EEXR0TFN7RPKAJG198A67", "Work Task")
	task2 := newTestTask("01KR8EEYBHNZPETYTB62WY3M1E", "Home Task")
	assert.NoError(t, db.Write(task1.GetCollection(), task1.GetID(), task1))
	assert.NoError(t, db.Write(task2.GetCollection(), task2.GetID(), task2))

	filter := domain.NewTaskFilter(&domain.Task{Title: "work"})
	filter.SetPage(1)
	filter.SetSize(10)

	out, err := lister.List(context.Background(), filter)

	assert.NoError(t, err)
	assert.Len(t, out.GetContent(), 1)
	assert.Equal(t, "Work Task", out.GetContent()[0].Title)
	assert.Equal(t, 1, out.GetTotal())
}

func listWithInvalidJSON(t *testing.T) {
	db, dir := newTestDB(t)
	lister := NewServiceLister[*domain.Task, *domain.TaskFilter](db)
	writeRawTaskFile(t, dir, "broken.json", "{not-a-valid-json")

	filter := domain.NewTaskFilter(&domain.Task{})

	_, err := lister.List(context.Background(), filter)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "invalid entity")
}
