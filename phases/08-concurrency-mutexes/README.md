# Phase 8: Concurrency, Mutexes, and Data Races

In this phase, you will learn about Go's multi-threaded execution model, the danger of data races when sharing memory, and how to synchronize access to in-memory state using Mutual Exclusion locks (`sync.Mutex` and `sync.RWMutex`).

---

## Conceptual Overview: Go vs. JavaScript & Python

### 1. The Multi-Threaded Reality
- **JS/Node.js:** Runs on a single-threaded event loop. If you have a global object `const cache = {}`, multiple incoming HTTP requests can read/write to `cache` without crashing the process, because JS executes callbacks sequentially (non-concurrently).
- **Python:** Uses the Global Interpreter Lock (GIL), which prevents multiple native threads from executing Python bytecodes at once.
- **Go:** Go is built for high-performance multi-threading. The Go runtime runs your code on a scheduler that maps lightweight threads called **Goroutines** onto multiple physical CPU cores. 

**Every incoming HTTP request in a Go server is handled in a separate, concurrent Goroutine.**

### 2. What is a Data Race?
If two or more Goroutines access the same memory location concurrently, and at least one of the accesses is a write, you have a **Data Race**.
In Go, a data race on an in-memory collection (like a Slice or a Map) will corrupt memory and cause your application to panic and crash unpredictably.

### 3. Synchronization with Mutexes
To make code "thread-safe", we use locks to ensure only one Goroutine can access a critical section of code at a time:
- **`sync.Mutex` (Mutual Exclusion):** A simple lock. Only one goroutine can read or write at a time.
- **`sync.RWMutex` (Reader/Writer Mutex):** An optimized lock. Multiple goroutines can read concurrently (using `RLock()`), but only a single goroutine can write (using `Lock()`). This is perfect for data stores where reads are much more frequent than writes.

```go
import "sync"

type SafeCounter struct {
    mu    sync.Mutex
    value int
}

func (c *SafeCounter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock() // defer unlocks automatically when the function exits
    c.value++
}
```

---

## Hands-On Exercise

You will refactor the in-memory Todo store in your clean architecture path (`app/internal/todo/store.go`) to make it thread-safe.

### Step 1: Update `InMemoryTodoStore`
Open `app/internal/todo/store.go`. 
1. Import `"sync"`.
2. Add a `mu sync.RWMutex` field to the `InMemoryTodoStore` struct:
   ```go
   type InMemoryTodoStore struct {
       mu     sync.RWMutex
       todos  []Todo
       nextID int
   }
   ```

### Step 2: Implement Thread-Safe Methods
Update the four methods of `*InMemoryTodoStore` using the reader/writer mutex:

- **`Add`**: Needs a write lock.
  ```go
  func (s *InMemoryTodoStore) Add(title string) (Todo, error) {
      s.mu.Lock()
      defer s.mu.Unlock()
      // ... existing add logic ...
  }
  ```
- **`List`**: Needs a read lock. Additionally, to avoid data races if the caller tries to modify the returned slice, you should **copy** the slice before returning it:
  ```go
  func (s *InMemoryTodoStore) List() []Todo {
      s.mu.RLock()
      defer s.mu.RUnlock()
      
      // Make a copy of s.todos
      copied := make([]Todo, len(s.todos))
      copy(copied, s.todos)
      return copied
  }
  ```
- **`ToggleComplete`**: Needs a write lock. Acquire `s.mu.Lock()` at the start, and unlock it.
- **`Delete`**: Needs a write lock. Acquire `s.mu.Lock()` at the start, and unlock it.

---

## Verifying Your Progress

To run the verifier:
```bash
node verify-phase.js 8
```

The verifier will copy a concurrency test file into `app/internal/todo/`. This test spawns hundreds of concurrent Goroutines reading, writing, and deleting from the same `InMemoryTodoStore` simultaneously. 

It executes this test with Go's built-in **Race Detector**:
```bash
go test -race ./app/internal/todo/...
```
If your mutex implementation is missing or incorrect, the race detector will catch it, print a data race report, and fail the verification.
