package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo-api/internal/todo"
)

func TestAuthMiddleware(t *testing.T) {
	store := todo.NewInMemoryTodoStore()
	router := NewRouter(store) // Returns http.Handler wrapped in RequestLogger

	// 1. GET /todos should not require auth
	req := httptest.NewRequest("GET", "/todos", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("GET /todos expected status 200, got %d", rr.Code)
	}

	// 2. POST /todos without auth header should fail (401)
	payload := []byte(`{"title":"Auth Task"}`)
	req = httptest.NewRequest("POST", "/todos", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("POST /todos without key expected status 401, got %d", rr.Code)
	}
	if cType := rr.Header().Get("Content-Type"); !strings.Contains(cType, "application/json") {
		t.Errorf("Expected Content-Type application/json, got %q", cType)
	}
	if !strings.Contains(rr.Body.String(), `"error":"unauthorized"`) {
		t.Errorf("Expected unauthorized error body, got %q", rr.Body.String())
	}

	// 3. POST /todos with wrong auth header should fail (401)
	req = httptest.NewRequest("POST", "/todos", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", "wrongkey")
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("POST /todos with wrong key expected status 401, got %d", rr.Code)
	}

	// 4. POST /todos with correct auth header should succeed (201)
	req = httptest.NewRequest("POST", "/todos", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", "supersecret")
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("POST /todos with correct key expected status 201, got %d", rr.Code)
	}

	// 5. DELETE /todos/{id} without auth header should fail (401)
	req = httptest.NewRequest("DELETE", "/todos/1", nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("DELETE /todos/1 without key expected status 401, got %d", rr.Code)
	}

	// 6. DELETE /todos/{id} with correct auth header should pass (200 since ID 1 exists now)
	req = httptest.NewRequest("DELETE", "/todos/1", nil)
	req.Header.Set("X-API-Key", "supersecret")
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK && rr.Code != http.StatusNoContent {
		t.Errorf("DELETE /todos/1 with correct key expected status 200 or 204, got %d", rr.Code)
	}
}
