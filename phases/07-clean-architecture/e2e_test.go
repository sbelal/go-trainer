package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"todo-api/internal/handler"
	"todo-api/internal/sqlite"
	"todo-api/internal/todo"
)

func TestCleanArchitectureE2E(t *testing.T) {
	const dbPath = "e2e_todos.db"
	os.Remove(dbPath)
	defer os.Remove(dbPath)

	// Test SQLite initialization in the new clean architecture path
	store, err := sqlite.NewSqliteTodoStore(dbPath)
	if err != nil {
		t.Fatalf("Failed to initialize sqlite store from clean path: %v", err)
	}

	// Verify interface compatibility
	var _ todo.TodoStore = store

	router := handler.NewRouter(store)

	// Add an item using handler
	payload, _ := json.Marshal(map[string]string{"title": "E2E Task"})
	req := httptest.NewRequest("POST", "/todos", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status Created (201), got %d", rr.Code)
	}

	var created todo.Todo
	err = json.Unmarshal(rr.Body.Bytes(), &created)
	if err != nil {
		t.Fatalf("Failed to parse todo: %v", err)
	}
	if created.Title != "E2E Task" || created.Completed {
		t.Errorf("Unexpected created todo values: %+v", created)
	}

	// Verify DB file exists
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		t.Errorf("Expected database file %s to be created, but it does not exist", dbPath)
	}
}
