# Phase 10: Context and Request Cancellations

In this phase, you will master Go's most important pattern for managing request life cycles and system deadlines: **`context.Context`**. You will refactor your stores and handlers to accept and propagate contexts, write a request timeout middleware, and utilize context-aware database queries.

---

## Conceptual Overview: Go vs. JavaScript & Python

### 1. What is Context?
In Go, a `Context` is a standard object that carries deadline schedules, cancellation signals, and request-scoped values across API boundaries and nested function calls.

- **JS / Python:** Cancelling operations mid-flight is often difficult or handled via custom event bindings (like JS `AbortController`).
- **Go:** Passing `ctx context.Context` as the **first argument** of functions is a universal idiom. If a user closes their browser tab or a query exceeds its timeout, the context is cancelled, signaling all downstream functions (like SQL queries or remote HTTP calls) to immediately stop working and release system resources.

### 2. Creating Contexts with Deadlines
You can create a child context that automatically cancels itself after a specific duration:
```go
import (
    "context"
    "time"
)

// parent context, duration
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel() // ALWAYS call cancel to release resources when done!
```

### 3. Context-Aware SQL Queries
Go's `database/sql` package provides context-aware alternatives for every database operation:
- `db.ExecContext(ctx, query, args...)` instead of `db.Exec(query, args...)`
- `db.QueryRowContext(ctx, query, args...)` instead of `db.QueryRow(query, args...)`
- `db.QueryContext(ctx, query, args...)` instead of `db.Query(query, args...)`

If `ctx` is cancelled before the query finishes, SQLite will immediately halt execution and return a `context.Canceled` or `context.DeadlineExceeded` error, preventing wasted database cycles.

---

## Hands-On Exercise

### Step 1: Update the `TodoStore` Interface
Open `app/internal/todo/store.go` and update the `TodoStore` interface:
1. Import `"context"`.
2. Add `ctx context.Context` as the first argument in all methods.
3. Add an `error` as the last return type for all methods (since database queries can fail):
```go
type TodoStore interface {
    Add(ctx context.Context, title string) (Todo, error)
    List(ctx context.Context) ([]Todo, error)
    ToggleComplete(ctx context.Context, id int) (bool, error)
    Delete(ctx context.Context, id int) (bool, error)
}
```

### Step 2: Update `InMemoryTodoStore` Methods
Update the method signatures and bodies in `app/internal/todo/store.go` to match the interface:
- Add `ctx context.Context` arguments.
- Return `nil` or errors (e.g. `return todos, nil`).
- In all methods, check if the context is already cancelled before executing:
  ```go
  if err := ctx.Err(); err != nil {
      return nil, err // returns context.Canceled or context.DeadlineExceeded
  }
  ```

### Step 3: Update `SqliteTodoStore` Methods
Open `app/internal/sqlite/store.go`.
1. Update all methods to accept `ctx context.Context` and match the new interface signatures.
2. Refactor all database calls to use their `Context` equivalents:
   - Change `s.db.Exec(...)` $\rightarrow$ `s.db.ExecContext(ctx, ...)`
   - Change `s.db.Query(...)` $\rightarrow$ `s.db.QueryContext(ctx, ...)`
   - Change `s.db.QueryRow(...)` $\rightarrow$ `s.db.QueryRowContext(ctx, ...)`
3. Bubble up any SQLite errors returned by these queries.

### Step 4: Write Timeout Middleware
Open `app/internal/handler/middleware.go`. Implement a timeout decorator:
- Signature: `func Timeout(timeout time.Duration) func(http.Handler) http.Handler`
- Returns a middleware function that sets a deadline timeout context on the request:
```go
func Timeout(timeout time.Duration) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx, cancel := context.WithTimeout(r.Context(), timeout)
            defer cancel()
            
            // Pass the request with the new child context
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
```

### Step 5: Update Router Handlers to Pass Context
Open `app/internal/handler/handler.go`. 
1. Pass the request context `r.Context()` to all store method calls.
   Example:
   ```go
   todos, err := h.store.List(r.Context())
   ```
2. Handle any store errors appropriately (e.g. returning `500 Internal Server Error` with `{"error": err.Error()}`).
3. Apply the `Timeout` middleware in `NewRouter` so that every request is capped at **1 second**:
   ```go
   // Inside NewRouter:
   // Wrap with Timeout first, then RequestLogger
   handler := Timeout(1 * time.Second)(mux)
   return RequestLogger(handler)
   ```

---

## Verifying Your Progress

To run the verifier:
```bash
node verify-phase.js 10
```

The verifier copies test files into your packages to assert:
1. All signatures in `TodoStore` match.
2. Context cancellations are correctly handled in `InMemoryTodoStore` and `SqliteTodoStore`.
3. Slow requests are automatically aborted by the timeout middleware (returning context cancelled errors).
