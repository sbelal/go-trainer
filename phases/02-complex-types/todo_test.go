package main

import (
	"testing"
)

func resetState() {
	Todos = nil
	nextID = 1
}

func TestAddTodoItem(t *testing.T) {
	resetState()

	t1 := AddTodoItem("Learn Go")
	if t1.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", t1.ID)
	}
	if t1.Title != "Learn Go" {
		t.Errorf("Expected Title to be 'Learn Go', got '%s'", t1.Title)
	}
	if t1.Completed {
		t.Errorf("Expected Completed to be false, got true")
	}

	t2 := AddTodoItem("Build API")
	if t2.ID != 2 {
		t.Errorf("Expected ID to be 2, got %d", t2.ID)
	}

	todos := ListTodos()
	if len(todos) != 2 {
		t.Errorf("Expected 2 todos, got %d", len(todos))
	}
}

func TestToggleTodoComplete(t *testing.T) {
	resetState()

	AddTodoItem("Task 1") // ID 1
	AddTodoItem("Task 2") // ID 2

	// Toggle existing
	ok := ToggleTodoComplete(1)
	if !ok {
		t.Errorf("ToggleTodoComplete(1) returned false, expected true")
	}

	todos := ListTodos()
	if !todos[0].Completed {
		t.Errorf("Expected todos[0].Completed to be true, got false")
	}

	// Toggle back
	ok = ToggleTodoComplete(1)
	if !ok {
		t.Errorf("ToggleTodoComplete(1) returned false, expected true")
	}
	if todos[0].Completed {
		t.Errorf("Expected todos[0].Completed to be false, got true")
	}

	// Toggle non-existent
	ok = ToggleTodoComplete(999)
	if ok {
		t.Errorf("ToggleTodoComplete(999) returned true, expected false")
	}
}

func TestDeleteTodoItem(t *testing.T) {
	resetState()

	AddTodoItem("Task A") // ID 1
	AddTodoItem("Task B") // ID 2
	AddTodoItem("Task C") // ID 3

	// Delete in the middle (ID 2)
	ok := DeleteTodoItem(2)
	if !ok {
		t.Errorf("DeleteTodoItem(2) returned false, expected true")
	}

	todos := ListTodos()
	if len(todos) != 2 {
		t.Errorf("Expected 2 todos remaining, got %d", len(todos))
	}

	if todos[0].ID != 1 || todos[1].ID != 3 {
		t.Errorf("Expected remaining IDs to be 1 and 3, got %d and %d", todos[0].ID, todos[1].ID)
	}

	// Delete non-existent
	ok = DeleteTodoItem(999)
	if ok {
		t.Errorf("DeleteTodoItem(999) returned true, expected false")
	}
}
