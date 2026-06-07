package main

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestTodoStructTags(t *testing.T) {
	todo := Todo{}
	st := reflect.TypeOf(todo)

	// Check ID tag
	idField, ok := st.FieldByName("ID")
	if !ok {
		t.Fatal("Todo struct is missing field ID")
	}
	if tag := idField.Tag.Get("json"); tag != "id" {
		t.Errorf("Expected json tag for ID to be 'id', got '%s'", tag)
	}

	// Check Title tag
	titleField, ok := st.FieldByName("Title")
	if !ok {
		t.Fatal("Todo struct is missing field Title")
	}
	if tag := titleField.Tag.Get("json"); tag != "title" {
		t.Errorf("Expected json tag for Title to be 'title', got '%s'", tag)
	}

	// Check Completed tag
	completedField, ok := st.FieldByName("Completed")
	if !ok {
		t.Fatal("Todo struct is missing field Completed")
	}
	if tag := completedField.Tag.Get("json"); tag != "completed" {
		t.Errorf("Expected json tag for Completed to be 'completed', got '%s'", tag)
	}
}

func TestStoreAddValidationError(t *testing.T) {
	store := NewInMemoryTodoStore()

	// Try adding with empty title
	todo, err := store.Add("")
	if err == nil {
		t.Fatal("Expected Add(\"\") to return an error, got nil")
	}
	if err.Error() != "title cannot be empty" {
		t.Errorf("Expected error message 'title cannot be empty', got '%v'", err.Error())
	}
	if todo.ID != 0 || todo.Title != "" || todo.Completed {
		t.Errorf("Expected empty Todo on error, got %+v", todo)
	}

	// Try adding with valid title
	todo, err = store.Add("Buy Groceries")
	if err != nil {
		t.Fatalf("Expected Add(\"Buy Groceries\") to return no error, got %v", err)
	}
	if todo.Title != "Buy Groceries" {
		t.Errorf("Expected Title to be 'Buy Groceries', got '%s'", todo.Title)
	}
}

func TestSerializationHelpers(t *testing.T) {
	// 1. Test SerializeTodos
	todos := []Todo{
		{ID: 1, Title: "Task 1", Completed: true},
		{ID: 2, Title: "Task 2", Completed: false},
	}

	bytes, err := SerializeTodos(todos)
	if err != nil {
		t.Fatalf("SerializeTodos returned error: %v", err)
	}

	jsonStr := string(bytes)
	if !strings.Contains(jsonStr, `"id":1`) || !strings.Contains(jsonStr, `"title":"Task 1"`) || !strings.Contains(jsonStr, `"completed":true`) {
		t.Errorf("Serialized JSON does not contain expected keys or values: %s", jsonStr)
	}

	// 2. Test DeserializeTodo (Valid JSON)
	validJSON := []byte(`{"id": 42, "title": "Test Deserialize", "completed": true}`)
	parsed, err := DeserializeTodo(validJSON)
	if err != nil {
		t.Fatalf("DeserializeTodo returned error on valid JSON: %v", err)
	}
	if parsed.ID != 42 || parsed.Title != "Test Deserialize" || !parsed.Completed {
		t.Errorf("DeserializeTodo returned incorrect struct values: %+v", parsed)
	}

	// 3. Test DeserializeTodo (Invalid JSON)
	invalidJSON := []byte(`{"id": 42, "title": `) // malformed
	_, err = DeserializeTodo(invalidJSON)
	if err == nil {
		t.Error("Expected DeserializeTodo to fail on invalid JSON, but it returned nil error")
	}
}
