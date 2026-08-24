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

	operationResult := make(chan error, 1)
	go func() {
		store.Add("module", "pipe")
		_, err := store.Get("module")
		operationResult <- err
	}()

	select {
	case err := <-operationResult:
		if err != nil {
			t.Fatalf("Get failed after deleting a missing pipe: %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("pipe store lock remained held after deleting a missing pipe")
	}
}
