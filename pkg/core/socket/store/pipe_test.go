package store

import (
	"testing"
	"time"
)

func TestDeleteMissingPipeReleasesLock(t *testing.T) {
	store := NewPipeStore()
	deleted := make(chan struct{})

	go func() {
		store.Delete("missing")
		close(deleted)
	}()

	select {
	case <-deleted:
	case <-time.After(time.Second):
		t.Fatal("Delete blocked while handling a missing pipe")
	}

	store.Add("module", "pipe")
	if _, err := store.Get("module"); err != nil {
		t.Fatalf("Get failed after deleting a missing pipe: %v", err)
	}
}
