# Phase 11: Table-Driven Tests and Mocking

In this final phase, you will learn how to write idiomatic unit tests in Go. You will learn the **Table-Driven Testing** pattern (a Go community standard) and how to create a **Mock** implementation of your `TodoStore` interface to test HTTP handlers in isolation without needing a database connection.

---

## Conceptual Overview: Go vs. JavaScript & Python

### 1. Table-Driven Tests
- **JS/Python:** In Jest or PyTest, you write multiple separate test blocks for different scenarios (e.g., `it("returns 400 when title is empty")`, `it("returns 401 when key is missing")`). This results in repetitive setup code.
- **Go:** Go developers use **Table-Driven Tests**. You define a slice of test cases (the "table"), where each case specifies its name, inputs, mock behavior, and expected outputs. You then run a single loop that executes `t.Run()` on each test case.
  ```go
  tests := []struct {
      name             string
      inputTitle       string
      expectedCode     int
  }{
      {"valid title", "Buy Groceries", http.StatusCreated},
      {"empty title", "", http.StatusBadRequest},
  }
  
  for _, tt := range tests {
      t.Run(tt.name, func(t *testing.T) {
          // run test using tt.inputTitle and check tt.expectedCode
      })
  }
  ```

### 2. Mocking Interfaces (Behavior Injection)
- **JS/Python:** Testing libraries dynamically mock methods using runtime magic (e.g., `jest.fn()` or `unittest.mock.patch`).
- **Go:** Because Go is compiled and statically typed, you mock by **defining a new struct that implements the target interface**, delegating its methods to function fields that you can swap out dynamically in each test case:

```go
type MockTodoStore struct {
    AddFunc func(ctx context.Context, title string) (todo.Todo, error)
}

// Satisfying the interface
func (m *MockTodoStore) Add(ctx context.Context, title string) (todo.Todo, error) {
    return m.AddFunc(ctx, title)
}
```

---

## Hands-On Exercise

You will write a test file named `app/internal/handler/router_test.go` under `package handler`.

### Step 1: Create the Mock Store
At the top of `app/internal/handler/router_test.go`, define a struct `MockTodoStore` that satisfies the `todo.TodoStore` interface:
```go
package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo-api/internal/todo"
)

type MockTodoStore struct {
	AddFunc            func(ctx context.Context, title string) (todo.Todo, error)
	ListFunc           func(ctx context.Context) ([]todo.Todo, error)
	ToggleCompleteFunc func(ctx context.Context, id int) (bool, error)
	DeleteFunc         func(ctx context.Context, id int) (bool, error)
}

func (m *MockTodoStore) Add(ctx context.Context, title string) (todo.Todo, error) {
	return m.AddFunc(ctx, title)
}

func (m *MockTodoStore) List(ctx context.Context) ([]todo.Todo, error) {
	return m.ListFunc(ctx)
}

func (m *MockTodoStore) ToggleComplete(ctx context.Context, id int) (bool, error) {
	return m.ToggleCompleteFunc(ctx, id)
}

func (m *MockTodoStore) Delete(ctx context.Context, id int) (bool, error) {
	return m.DeleteFunc(ctx, id)
}
```

---

### Step 2: Implement Table-Driven Test for `POST /todos`
Write a function `TestRouter_AddTodo(t *testing.T)` that tests the `POST /todos` endpoint.
Define a table struct representing the test cases:
```go
func TestRouter_AddTodo(t *testing.T) {
	tests := []struct {
		name           string
		apiKey         string
		payload        string
		mockAdd        func(ctx context.Context, title string) (todo.Todo, error)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:   "Missing API Key Header",
			apiKey: "",
			payload: `{"title": "Valid Todo"}`,
			mockAdd: nil, // shouldn't be called
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `"error":"unauthorized"`,
		},
		{
			name:   "Invalid API Key",
			apiKey: "wrong_key",
			payload: `{"title": "Valid Todo"}`,
			mockAdd: nil, // shouldn't be called
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `"error":"unauthorized"`,
		},
		{
			name:   "Success Case",
			apiKey: "supersecret",
			payload: `{"title": "Buy Milk"}`,
			mockAdd: func(ctx context.Context, title string) (todo.Todo, error) {
				return todo.Todo{ID: 10, Title: title, Completed: false}, nil
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   `"title":"Buy Milk"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 1. Initialize mock store
			mockStore := &MockTodoStore{
				AddFunc: tt.mockAdd,
			}

			// 2. Build router using the mock store
			router := NewRouter(mockStore)

			// 3. Create HTTP Request
			req := httptest.NewRequest("POST", "/todos", strings.NewReader(tt.payload))
			req.Header.Set("Content-Type", "application/json")
			if tt.apiKey != "" {
				req.Header.Set("X-API-Key", tt.apiKey)
			}

			// 4. Record response
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			// 5. Assert status code
			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, rr.Code)
			}

			// 6. Assert response body contains expected output
			if !strings.Contains(rr.Body.String(), tt.expectedBody) {
				t.Errorf("Expected body to contain %q, got %q", tt.expectedBody, rr.Body.String())
			}
		})
	}
}
```

---

## Verifying Your Progress

To run the verifier:
```bash
node verify-phase.js 11
```

The verifier will compile your code and check that `app/internal/handler/router_test.go` exists, and that running the tests inside your `handler` package executes and passes successfully:
```bash
go test -v ./app/internal/handler/...
```
Congratulations! Once this phase passes, you have finished all core and advanced phases of the Go Trainer!
