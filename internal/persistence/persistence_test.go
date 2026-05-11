//go:build unit

package persistence

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/joaohgf/magalu-cli/internal/core/domain"
	"github.com/joaohgf/magalu-cli/internal/enum"
	"github.com/nanobox-io/scribble"
)

func newTestDB(t *testing.T) (*scribble.Driver, string) {
	dir := t.TempDir()
	db, err := scribble.New(dir, nil)
	if err != nil {
		t.Fatalf("failed to create test db: %v", err)
	}
	return db, dir
}

func newTestTask(id, title string) *domain.Task {
	return &domain.Task{
		ID:        id,
		Title:     title,
		Status:    enum.StatusInProgress,
		Priority:  enum.PriorityMedium,
		CreatedAt: time.Now(),
	}
}

func writeRawTaskFile(t *testing.T, baseDir, fileName, raw string) {
	tasksDir := filepath.Join(baseDir, "tasks")
	if err := os.MkdirAll(tasksDir, os.ModePerm); err != nil {
		t.Fatalf("failed to create tasks directory: %v", err)
	}
	if err := os.WriteFile(filepath.Join(tasksDir, fileName), []byte(raw), 0o644); err != nil {
		t.Fatalf("failed to write raw task file: %v", err)
	}
}
