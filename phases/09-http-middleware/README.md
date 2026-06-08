# Phase 9: HTTP Middleware (Request Interceptors)

In this phase, you will learn to implement the **Middleware** pattern in Go using the standard library. You will write a logging middleware to record server response times, and an API Key Authentication middleware to secure modifications (`POST` and `DELETE` requests) to your Todo app.

---

## Conceptual Overview: Go vs. JavaScript & Python

Middleware functions intercept and process requests before they reach the final handler (e.g. for authentication, logging, CORS).

### 1. Express (JS) vs. Go Middleware Model
- **JS (Express):** Uses the `next()` callback pattern:
  ```javascript
  const logger = (req, res, next) => {
      console.log(`${req.method} ${req.url}`);
      next();
  };
  app.use(logger);
  ```
- **Go:** Uses a **Decorator pattern**. A middleware is simply a function that takes an `http.Handler` as an argument, wraps it inside a new `http.HandlerFunc`, performs operations, and calls the original handler's `ServeHTTP` method:
  ```go
  func Logger(next http.Handler) http.Handler {
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
          log.Printf("Request: %s %s", r.Method, r.URL.Path)
          next.ServeHTTP(w, r) // pass control to next handler
      })
  }
  ```

### 2. Wrapping the Entire Mux vs. Wrapping Specific Routes
- **Wrapping Everything:** You can apply middleware globally by wrapping the final router/mux:
  ```go
  router := NewRouter(store)
  loggedRouter := Logger(router)
  http.ListenAndServe(":8080", loggedRouter)
  ```
- **Wrapping Specific Routes:** You can apply middleware selectively by wrapping individual handler registrations:
  ```go
  mux.Handle("POST /todos", Auth(http.HandlerFunc(postHandler)))
  ```

### 3. Measuring Duration and Reading Request Headers

#### Timing Operations
To measure how long an operation takes, Go provides the `"time"` package:
* `time.Now()` returns the current local time.
* `time.Since(startTime)` returns the elapsed time since `startTime` as a duration.

#### Logging
In Go, you use the standard `"log"` package to output timestamped logging information to the console:
```go
import (
    "log"
    "time"
)

func ProcessRequest() {
    start := time.Now()
    
    // ... do some work ...
    
    duration := time.Since(start)
    log.Printf("Process took %s", duration)
}
```

#### Reading Request Headers
To retrieve HTTP headers sent by the client, use the `Header` field of the `*http.Request` object and call its `.Get("Header-Name")` method. Note that Go normalizes header keys, but it is idiomatic to use standard capitalization:
```go
func HandleAuth(w http.ResponseWriter, r *http.Request) {
    // Retrieves value of the header "X-API-Key"
    apiKey := r.Header.Get("X-API-Key")
    
    if apiKey == "" {
        log.Println("X-API-Key header is missing!")
    }
}
```

---

## Hands-On Exercise

You will create a file named `app/internal/handler/middleware.go` and update `app/internal/handler/handler.go` to add logging and API key auth headers.

### Step 1: Write the Middlewares
Create `app/internal/handler/middleware.go` with the package `handler` and implement:

1. **`RequestLogger`**:
   - Signature: `func RequestLogger(next http.Handler) http.Handler`
   - Logs the HTTP Method, URL Path, and the time taken to process the request.
   - **Hint:** Record the start time using `time.Now()` before calling `next.ServeHTTP(w, r)`, then compute duration using `time.Since(start)`. Log it using the `"log"` package.

2. **`APIKeyAuth`**:
   - Signature: `func APIKeyAuth(next http.Handler) http.Handler`
   - Retrieves the header `"X-API-Key"` from the request.
   - If the header is missing or does not equal `"supersecret"`, it should:
     - Set the header `Content-Type` to `application/json`.
     - Write the status code `401 Unauthorized`.
     - Return the JSON body `{"error":"unauthorized"}`.
     - **Do not** call `next.ServeHTTP`.
   - If the header matches, call `next.ServeHTTP(w, r)`.

---

### Step 2: Apply the Middleware in `handler.go`
Open `app/internal/handler/handler.go`. 

1. Apply `APIKeyAuth` to all state-mutating endpoints:
   - `POST /todos`
   - `POST /todos/{id}/toggle`
   - `DELETE /todos/{id}`
   *Leave `GET /todos` public (unauthenticated).*

   Example:
   ```go
   // Instead of: mux.HandleFunc("POST /todos", handleAddTodo)
   // Use Handle instead of HandleFunc because APIKeyAuth returns an http.Handler:
   mux.Handle("POST /todos", APIKeyAuth(http.HandlerFunc(handleAddTodo)))
   ```

2. Apply `RequestLogger` globally in `NewRouter`:
   - Before returning from `NewRouter`, wrap the complete multiplexer `mux` inside the logger:
   ```go
   func NewRouter(store todo.TodoStore) http.Handler {
       mux := http.NewServeMux()
       
       // ... register routes ...
       
       // Wrap the mux with Logger middleware
       return RequestLogger(mux)
   }
   ```
   *Note:* The return type of `NewRouter` should change from `*http.ServeMux` to `http.Handler` because `RequestLogger` returns an `http.Handler`. You will also need to update the signature in `app/cmd/todo-api/main.go` if it refers to `*http.ServeMux` (changing it to `http.Handler` or using type inference).

---

## Verifying Your Progress

To run the verifier:
```bash
node verify-phase.js 9
```

The verifier will copy a middleware test file into `app/internal/handler/` and run tests that:
1. Verify that `GET /todos` works without any credentials.
2. Verify that write actions (`POST` and `DELETE`) return `401 Unauthorized` and `{"error":"unauthorized"}` if `X-API-Key` is missing or incorrect.
3. Verify that write actions succeed when the valid `X-API-Key: supersecret` header is provided.
4. Verify that the logger prints output to the stdout.
