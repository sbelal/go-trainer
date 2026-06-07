package main

import (
	"database/sql"
	"os"
	"testing"

	_ "modernc.org/sqlite"
)

const testDBPath = "test_todos.db"

func cleanDB(t *testing.T) {
	os.Remove(testDBPath)
}

func TestSqliteTodoStore(t *testing.T) {
	cleanDB(t)
	t.Cleanup(func() {
		cleanDB(t)
	})

	// Test store instantiation
	store, err := NewSqliteTodoStore(testDBPath)
	if err != nil {
		t.Fatalf("Failed to create SqliteTodoStore: %v", err)
	}

	// Verify that table exists by querying it directly
	db, err := sql.Open("sqlite", testDBPath)
	if err != nil {
		t.Fatalf("Failed to open test database directly: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatalf("Failed to ping direct database connection: %v", err)
	}

	// Verify interface implementation
	var _ TodoStore = store

	// 1. Verify list is empty
	todos := store.List()
	if len(todos) != 0 {
		t.Errorf("Expected empty list, got %d items", len(todos))
	}

	// 2. Test Add
	t1, err := store.Add("DB Task 1")
	if err != nil {
		t.Fatalf("Add returned error: %v", err)
	}
	if t1.ID == 0 {
		t.Error("Expected auto-incremented ID to be non-zero")
	}
	if t1.Title != "DB Task 1" || t1.Completed {
		t.Errorf("Add returned incorrect values: %+v", t1)
	}

	// 3. Test Add empty validation
	_, err = store.Add("")
	if err == nil {
		t.Error("Expected Add(\"\") to fail with error, but it succeeded")
	}

	// 4. Test List
	todos = store.List()
	if len(todos) != 1 {
		t.Errorf("Expected list size 1, got %d", len(todos))
	}
	if todos[0].ID != t1.ID || todos[0].Title != "DB Task 1" {
		t.Errorf("List returned incorrect item details: %+v", todos[0])
	}

	// 5. Test ToggleComplete
	ok := store.ToggleComplete(t1.ID)
	if !ok {
		t.Errorf("ToggleComplete(%d) returned false, expected true", t1.ID)
	}

	// Retrieve directly from DB to verify persistence
	var completed bool
	err = db.QueryRow("SELECT completed FROM todos WHERE id = ?", t1.ID).Scan(&completed)
	if err != nil {
		t.Fatalf("Failed to query database directly: %v", err)
	}
	if !completed {
		t.Error("Expected todo completed status to be true in the database")
	}

	// Toggle back
	ok = store.ToggleComplete(t1.ID)
	if !ok {
		t.Errorf("ToggleComplete(%d) back returned false", t1.ID)
	}
	err = db.QueryRow("SELECT completed FROM todos WHERE id = ?", t1.ID).Scan(&completed)
	if err != nil {
		t.Fatalf("Failed to query database directly: %v", err)
	}
	if completed {
		t.Error("Expected todo completed status to be false in the database after second toggle")
	}

	// Toggle non-existent
	ok = store.ToggleComplete(999)
	if ok {
		t.Error("ToggleComplete(999) returned true, expected false")
	}

	// 6. Test Delete
	ok = store.Delete(t1.ID)
	if !ok {
		t.Errorf("Delete(%d) returned false, expected true", t1.ID)
	}

	todos = store.List()
	if len(todos) != 0 {
		t.Errorf("Expected empty list after deleting the only item, got %d items", len(todos))
	}

	// Delete non-existent
	ok = store.Delete(t1.ID)
	if ok {
		t.Errorf("Delete(%d) returned true on already deleted item", t1.ID)
	}
}
