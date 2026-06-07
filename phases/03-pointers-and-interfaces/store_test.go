package main

import (
	"testing"
)

func TestTodoStoreInterface(t *testing.T) {
	// Compile-time check: does *InMemoryTodoStore implement TodoStore?
	var store TodoStore
	store = NewInMemoryTodoStore()

	// Verify initial empty list
	todos := store.List()
	if len(todos) != 0 {
		t.Errorf("Expected empty list, got %d items", len(todos))
	}

	// Add items
	t1 := store.Add("Task 1")
	if t1.ID != 1 || t1.Title != "Task 1" || t1.Completed {
		t.Errorf("Add failed: got %+v", t1)
	}

	t2 := store.Add("Task 2")
	if t2.ID != 2 {
		t.Errorf("Add failed: second ID is %d, expected 2", t2.ID)
	}

	// List items
	todos = store.List()
	if len(todos) != 2 {
		t.Errorf("Expected list of size 2, got %d", len(todos))
	}

	// Toggle completion status
	ok := store.ToggleComplete(1)
	if !ok {
		t.Errorf("ToggleComplete(1) returned false, expected true")
	}

	todos = store.List()
	if !todos[0].Completed {
		t.Errorf("Expected Task 1 to be completed, got false")
	}

	// Delete item
	ok = store.Delete(1)
	if !ok {
		t.Errorf("Delete(1) returned false, expected true")
	}

	todos = store.List()
	if len(todos) != 1 {
		t.Errorf("Expected list of size 1 after deletion, got %d", len(todos))
	}
	if todos[0].ID != 2 {
		t.Errorf("Expected remaining item to be ID 2, got %d", todos[0].ID)
	}
}

func TestStoreIsolation(t *testing.T) {
	// Verify that state is isolated inside each instance (no globals used)
	s1 := NewInMemoryTodoStore()
	s2 := NewInMemoryTodoStore()

	s1.Add("S1 Task")
	
	if len(s1.List()) != 1 {
		t.Errorf("Expected s1 to have 1 task")
	}
	if len(s2.List()) != 0 {
		t.Errorf("Expected s2 to be empty, but it has %d tasks. Are you still using global package-level variables?", len(s2.List()))
	}
}
