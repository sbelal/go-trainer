# Phase 5: HTTP REST API

In this phase, you will build a functional REST API using Go's standard library `net/http` package. You will learn how to create a router, define endpoints with HTTP methods and path parameters (using Go 1.22+ syntax), and handle JSON request and response bodies.

---

## Conceptual Overview: Go vs. JavaScript & Python

If you are used to Express (Node.js) or Flask/FastAPI (Python), Go's standard library `net/http` is incredibly powerful and requires no external framework to build robust production-grade APIs.

### 1. Defining Routes and Routers
- **JS (Express):**
  ```javascript
  const express = require('express');
  const app = express();
  app.get('/todos', (req, res) => { res.json(todos); });
  ```
- **Go (1.22+ Standard library):**
  Using the built-in HTTP request multiplexer `http.ServeMux`:
  ```go
  mux := http.NewServeMux()
  mux.HandleFunc("GET /todos", func(w http.ResponseWriter, r *http.Request) {
      // Logic here
  })
  ```
  Note the route pattern syntax: `"METHOD /path"`. Go automatically matches the HTTP verb and the route!

### 2. Path Parameters
- **JS (Express):** `:id` parsed via `req.params.id`.
- **Python (Flask):** `<int:id>` parsed via handler argument.
- **Go (1.22+):** `{id}` parsed via `r.PathValue("id")`.
  ```go
  mux.HandleFunc("DELETE /todos/{id}", func(w http.ResponseWriter, r *http.Request) {
      idStr := r.PathValue("id")
      id, err := strconv.Atoi(idStr) // Parse string to int
      // ...
  })
  ```

### 3. Parsing Request Bodies and Writing JSON Responses
In Go, reading request bodies requires decoding the raw byte stream from `r.Body`. Go's `json.NewDecoder` is highly optimized for this.

- **Writing JSON:** Set the `Content-Type` header, set the status code, and encode the struct:
  ```go
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(data)
  ```
- **Reading JSON:** Decode `r.Body` directly into a struct:
  ```go
  var req AddRequest
  err := json.NewDecoder(r.Body).Decode(&req)
  ```

---

## Hands-On Exercise

Create a file named `app/server.go` with package declaration `package main`. Implement the following:

### Step 1: Create the Router
Implement a function named `NewRouter` that accepts a `TodoStore` interface and returns a pointer to an `http.ServeMux`:
```go
func NewRouter(store TodoStore) *http.ServeMux {
    mux := http.NewServeMux()
    
    // Register routes here
    
    return mux
}
```

### Step 2: Register and Implement the REST Endpoints
Implement handlers inside `NewRouter` (using closure or helper structs) for the following endpoints:

#### 1. `GET /todos` (List Todos)
- Status code: `200 OK`
- Response header: `Content-Type: application/json`
- Body: JSON array of all Todo items.

#### 2. `POST /todos` (Add Todo)
- Parse JSON body: `{"title": "..."}`. You can define a helper struct:
  ```go
  type AddTodoRequest struct {
      Title string `json:"title"`
  }
  ```
- Call `store.Add(req.Title)`.
- If JSON decoding fails or the store returns an error (e.g., empty title):
  - Status code: `400 Bad Request`
  - Response body: `{"error": "<error message>"}`
- If successful:
  - Status code: `201 Created`
  - Response body: JSON representation of the created `Todo`.

#### 3. `DELETE /todos/{id}` (Delete Todo)
- Retrieve `id` parameter from route: `r.PathValue("id")`.
- Convert the string `id` to an integer using `strconv.Atoi`. If conversion fails, return `400 Bad Request` with `{"error": "invalid id"}`.
- Call `store.Delete(id)`.
- If the todo was found and deleted:
  - Status code: `200 OK` (you can return `{"success": true}` or empty response)
- If not found:
  - Status code: `404 Not Found`
  - Response body: `{"error": "todo not found"}`

#### 4. `POST /todos/{id}/toggle` (Toggle Todo Completion Status)
- Retrieve `id` parameter from route, convert to integer. If conversion fails, return `400 Bad Request`.
- Call `store.ToggleComplete(id)`.
- If found:
  - Status code: `200 OK`
  - Response body: `{"success": true}`
- If not found:
  - Status code: `404 Not Found`
  - Response body: `{"error": "todo not found"}`

---

### Step 3: Run the Server in `main.go`
Update your `app/main.go` to wire everything together and run the web server:
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	store := NewInMemoryTodoStore()
	router := NewRouter(store)

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
```

---

## Verifying Your Progress

To run the verifier:
```bash
node verify-phase.js 5
```

The verifier will compile your code and execute a Go HTTP integration test suite that tests each route, status codes, headers, and validation errors under simulated requests.
