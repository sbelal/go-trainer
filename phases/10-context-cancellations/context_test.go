package todo

import (
	"context"
	"testing"
)

func TestCancelledContext(t *testing.T) {
	store := NewInMemoryTodoStore()

	// Create a cancelled context
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	// Test Add with cancelled context
	_, err := store.Add(ctx, "Test Task")
	if err == nil {
		t.Error("Expected Add to return error when called with a cancelled context, got nil")
	}

	// Test List with cancelled context
	_, err = store.List(ctx)
	if err == nil {
		t.Error("Expected List to return error when called with a cancelled context, got nil")
	}

	// Test ToggleComplete with cancelled context
	_, err = store.ToggleComplete(ctx, 1)
	if err == nil {
		t.Error("Expected ToggleComplete to return error when called with a cancelled context, got nil")
	}

	// Test Delete with cancelled context
	_, err = store.Delete(ctx, 1)
	if err == nil {
		t.Error("Expected Delete to return error when called with a cancelled context, got nil")
	}
}
