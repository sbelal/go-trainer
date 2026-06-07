package todo

import (
	"fmt"
	"sync"
	"testing"
)

func TestConcurrentStoreAccess(t *testing.T) {
	store := NewInMemoryTodoStore()
	var wg sync.WaitGroup

	numGoroutines := 100

	// Concurrent writes
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			_, err := store.Add(fmt.Sprintf("Concurrent Task %d", id))
			if err != nil {
				t.Errorf("Unexpected error adding task concurrently: %v", err)
			}
		}(i)
	}

	// Concurrent reads & updates
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			_ = store.List()
			_ = store.ToggleComplete(id)
		}(i)
	}

	// Concurrent deletions
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			_ = store.Delete(id)
		}(i)
	}

	wg.Wait()
}
