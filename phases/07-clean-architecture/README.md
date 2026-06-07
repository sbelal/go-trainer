# Phase 7: Clean Architecture and Package Refactoring

In this final phase, you will refactor your TODO API from a flat file structure into a production-grade directory layout. You will learn about standard Go project conventions (`cmd/` and `internal/`), packaging, resolving imports, injecting configurations via environment variables, and writing end-to-end integration tests.

---

## Conceptual Overview: Go vs. JavaScript & Python

### 1. The Standard Go Directory Layout
In Node.js or Python, you can name directories whatever you like (`controllers/`, `views/`, `src/`).
Go projects follow a community standard directory structure to separate concerns and enforce access boundaries:
- **`cmd/`**: Entrypoints. Each folder here represents an executable program (e.g. `cmd/todo-api/main.go` compiles into `todo-api`). Code here should only do configuration, dependency wiring, and startup.
- **`internal/`**: Private application code. Go **compiler enforces** that packages inside an `internal/` folder can only be imported by code inside the parent directory. This keeps your internal implementations encapsulated.

### 2. Resolving Package Import Paths
Unlike Node's relative paths (`../../controllers/todo`) or Python's package paths, Go imports are **always absolute, starting with the module name defined in `go.mod`**.
If your `go.mod` is:
```go
module todo-api
```
And you have a package `handler` inside `app/internal/handler`, you import it like this from anywhere else in your project:
```go
import "todo-api/internal/handler"
```

### 3. Preventing Circular Imports
In JS or Python, file A can import file B, and file B can import file A (circular dependency).
In Go, **circular imports are a compile-time error!**
```
Package A imports Package B
Package B imports Package A  --> Compile error!
```
To avoid circular imports, Go developers:
- Define interfaces in domain packages (e.g. `internal/todo`) that have no dependencies.
- Make implementation packages (e.g. `internal/sqlite`) import the domain package, not the other way around.
- Let the entrypoint (`cmd/`) instantiate the database implementation and inject it into the handlers (Dependency Injection).

---

## Hands-On Exercise

You will refactor your existing flat files inside the `app/` folder into structured directories.

### Step 1: Create the Project Folders
Create the following directories inside `app/`:
```bash
mkdir -p app/cmd/todo-api
mkdir -p app/internal/todo
mkdir -p app/internal/sqlite
mkdir -p app/internal/handler
```

### Step 2: Refactor Code into Packages

#### 1. Package `todo` (Domain Logic)
Move `app/todo.go` and `app/json.go` into `app/internal/todo/`.
- Change their package declaration at the top to: `package todo`.
- In `todo.go`, define the `Todo` struct, `TodoStore` interface, `InMemoryTodoStore` struct, and `NewInMemoryTodoStore()`.
- In `json.go`, place the `SerializeTodos` and `DeserializeTodo` functions.

#### 2. Package `sqlite` (Database Persistence)
Move `app/sqlite_store.go` into `app/internal/sqlite/`.
- Change package declaration to: `package sqlite`.
- Import the `todo` package: `import "todo-api/internal/todo"`.
- Update the store fields and method signatures to use the `todo` prefix (e.g., `todo.Todo`, `todo.TodoStore`).
  Example:
  ```go
  func (s *SqliteTodoStore) Add(title string) (todo.Todo, error) { ... }
  ```

#### 3. Package `handler` (REST API Routing)
Move `app/server.go` into `app/internal/handler/`.
- Change package declaration to: `package handler`.
- Import the `todo` package: `import "todo-api/internal/todo"`.
- Update `NewRouter` to accept `todo.TodoStore` and map the endpoints. Make sure to update the type of `Todo` reference inside your HTTP encoder/decoder calls to `todo.Todo`.

#### 4. The main entrypoint
Move `app/main.go` into `app/cmd/todo-api/`.
- Package declaration remains: `package main`.
- Import packages `sqlite` and `handler`:
  ```go
  import (
      "todo-api/internal/handler"
      "todo-api/internal/sqlite"
  )
  ```
- Modify the `main` function to check environment variables `PORT` (defaulting to `"8080"`) and `DB_PATH` (defaulting to `"todos.db"`) using `os.Getenv()`. Connect to SQLite using the configured path, build the router, and run the server.

---

### Step 3: Cleanup Root Folder
To avoid compilation errors (due to duplicate declarations), delete the old files in the `app/` root folder:
- Delete `app/basics.go`
- Delete `app/todo.go`
- Delete `app/store.go`
- Delete `app/json.go`
- Delete `app/sqlite_store.go`
- Delete `app/server.go`
- Delete `app/main.go`

Only `go.mod`, `go.sum`, `cmd/`, and `internal/` should remain in the `app/` folder.

---

## Verifying Your Progress

To run the final verifier:
```bash
node verify-phase.js 7
```

The verifier will check that you have:
1. Cleaned up all flat source files in `app/`.
2. Created the structured folder layout.
3. Enabled env variables `PORT` and `DB_PATH` correctly.
4. Compiled and executed a comprehensive E2E integration test suite simulating full API actions.
