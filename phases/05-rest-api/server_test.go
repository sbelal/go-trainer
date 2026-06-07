package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRESTAPI(t *testing.T) {
	store := NewInMemoryTodoStore()
	router := NewRouter(store)

	// 1. GET /todos (initially empty)
	req := httptest.NewRequest("GET", "/todos", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("GET /todos expected status 200, got %d", rr.Code)
	}
	if cType := rr.Header().Get("Content-Type"); !strings.Contains(cType, "application/json") {
		t.Errorf("Expected Content-Type application/json, got %q", cType)
	}
	if strings.TrimSpace(rr.Body.String()) != "[]" && strings.TrimSpace(rr.Body.String()) != "null" {
		t.Errorf("Expected empty JSON list, got %q", rr.Body.String())
	}

	// 2. POST /todos (create task)
	payload := []byte(`{"title":"Buy Milk"}`)
	req = httptest.NewRequest("POST", "/todos", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("POST /todos expected status 201, got %d", rr.Code)
	}

	var created Todo
	err := json.Unmarshal(rr.Body.Bytes(), &created)
	if err != nil {
		t.Fatalf("Failed to parse POST response: %v", err)
	}
	if created.ID != 1 || created.Title != "Buy Milk" || created.Completed {
		t.Errorf("Created todo values mismatch: %+v", created)
	}

	// 3. POST /todos (validation error: empty title)
	invalidPayload := []byte(`{"title":""}`)
	req = httptest.NewRequest("POST", "/todos", bytes.NewBuffer(invalidPayload))
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("POST /todos empty title expected status 400, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "title cannot be empty") {
		t.Errorf("Expected error message in response, got %q", rr.Body.String())
	}

	// 4. POST /todos/{id}/toggle (toggle complete)
	req = httptest.NewRequest("POST", "/todos/1/toggle", nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("POST /todos/1/toggle expected status 200, got %d", rr.Code)
	}

	// Verify completion status
	list := store.List()
	if len(list) != 1 || !list[0].Completed {
		t.Errorf("Expected todo to be completed, got %+v", list)
	}

	// 5. POST /todos/{id}/toggle (non-existent ID)
	req = httptest.NewRequest("POST", "/todos/999/toggle", nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("POST /todos/999/toggle expected status 404, got %d", rr.Code)
	}

	// 6. DELETE /todos/{id} (delete item)
	req = httptest.NewRequest("DELETE", "/todos/1", nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK && rr.Code != http.StatusNoContent {
		t.Errorf("DELETE /todos/1 expected status 200 or 204, got %d", rr.Code)
	}

	list = store.List()
	if len(list) != 0 {
		t.Errorf("Expected store to be empty after deletion, but has %d elements", len(list))
	}

	// 7. DELETE /todos/{id} (already deleted / non-existent)
	req = httptest.NewRequest("DELETE", "/todos/1", nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("DELETE /todos/1 (non-existent) expected status 404, got %d", rr.Code)
	}
}
