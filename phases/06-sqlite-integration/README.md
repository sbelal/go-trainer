# Phase 6: SQLite Integration

In this phase, you will replace your temporary in-memory store with a persistent SQLite database. You will learn to use Go's standard library `database/sql` package, use a CGO-free driver, run database schema migrations, and execute SQL CRUD queries.

Thanks to the `TodoStore` interface you created in Phase 3, you can swap the entire database implementation without changing a single line of your HTTP handler code!

---

## Conceptual Overview: Go vs. JavaScript & Python

Working with databases in Go centers around the standard library's `database/sql` package, which defines a unified database connection pool manager.

### 1. Zero C-Dependencies (Avoiding CGO friction)
- **JS/Python:** Node's `sqlite3` or Python's native `sqlite3` compile binary wrappers automatically.
- **Go:** Standard SQLite wrappers (like `github.com/mattn/go-sqlite3`) require a C compiler (GCC) installed on the host machine.
To avoid compiler toolchain issues on Windows and Linux, we use a **pure-Go SQLite driver** (`modernc.org/sqlite`), which is a transpiled version of SQLite in pure Go. It implements the exact `database/sql` interface, requiring zero system compiler dependencies.

### 2. Importing Drivers and Connection Pools
In Go, you import the driver anonymously (`_`) to execute its package initialization (registering it with `database/sql`), and then use the generic `sql.Open`:
```go
import (
    "database/sql"
    _ "modernc.org/sqlite" // registers driver as "sqlite"
)

db, err := sql.Open("sqlite", "todos.db")
```
> 💡 `sql.Open` does **not** actually establish a connection to the database. It merely sets up the connection pool manager. To verify the connection, you must call `db.Ping()`.

### 3. Executing Statements and Scanning Rows
- **`db.Exec`**: Used for queries that modify data (DDL creation, INSERT, UPDATE, DELETE). It returns an `sql.Result` with info like `LastInsertId()` and `RowsAffected()`.
- **`db.QueryRow`**: Used for queries expected to return at most one row. You scan directly from it:
  ```go
  var title string
  err := db.QueryRow("SELECT title FROM todos WHERE id = ?", id).Scan(&title)
  ```
- **`db.Query`**: Used for queries returning multiple rows. You iterate using a loop:
  ```go
  rows, err := db.Query("SELECT id, title, completed FROM todos")
  defer rows.Close() // ALWAYS close rows to avoid connection leaks!
  
  for rows.Next() {
      var t Todo
      err := rows.Scan(&t.ID, &t.Title, &t.Completed)
      // append to slice
  }
  ```

---

## Hands-On Exercise

### Step 1: Install the SQLite Driver
Open your terminal inside the `app/` directory and add the pure-Go SQLite package:
```bash
go get modernc.org/sqlite
```
This will download the library and update your `go.mod` and `go.sum` files automatically.

### Step 2: Implement the `SqliteTodoStore`
Create a new file named `app/sqlite_store.go` with package `package main`. Implement the following:

1. **Struct Definition:**
   ```go
   type SqliteTodoStore struct {
       db *sql.DB
   }
   ```
2. **Constructor Function:**
   Write `NewSqliteTodoStore` which initializes the database, pings it, runs the migration DDL, and returns the store:
   ```go
   func NewSqliteTodoStore(dbPath string) (*SqliteTodoStore, error) {
       db, err := sql.Open("sqlite", dbPath)
       if err != nil {
           return nil, err
       }
       
       // Ping the database
       if err = db.Ping(); err != nil {
           return nil, err
       }
       
       // Create tables DDL statement
       query := `
       CREATE TABLE IF NOT EXISTS todos (
           id INTEGER PRIMARY KEY AUTOINCREMENT,
           title TEXT NOT NULL,
           completed BOOLEAN NOT NULL DEFAULT 0
       );`
       
       _, err = db.Exec(query)
       if err != nil {
           return nil, err
       }
       
       return &SqliteTodoStore{db: db}, nil
   }
   ```
3. **Implement the `TodoStore` Methods:**
   Implement the four methods on `*SqliteTodoStore` to satisfy the `TodoStore` interface:
   - **`Add(title string) (Todo, error)`**: Check if `title` is empty. If empty, return validation error. Otherwise, execute an insert SQL command, retrieve the `LastInsertId()`, and return the populated `Todo` struct.
   - **`List() []Todo`**: Retrieve all rows from the database. Iterate over rows, scan them into `Todo` structs, and return the slice. Ensure rows are closed properly with `defer rows.Close()`.
   - **`ToggleComplete(id int) bool`**: Query the todo's current `completed` value using `QueryRow().Scan()`. If not found (returns `sql.ErrNoRows`), return `false`. Update the `completed` value to its opposite (`!currentCompleted`), and return `true`.
   - **`Delete(id int) bool`**: Execute a delete SQL statement. Check `result.RowsAffected()`. If rows affected is `0` (meaning no todo had that ID), return `false`. Otherwise, return `true`.

---

### Step 3: Wire it Up in `main.go`
Update `app/main.go` to use the database instead of the in-memory store:
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Initialize SQLite Store (will create todos.db in the directory)
	store, err := NewSqliteTodoStore("todos.db")
	if err != nil {
		fmt.Printf("Failed to connect to SQLite: %v\n", err)
		return
	}

	router := NewRouter(store)

	fmt.Println("Server is running on http://localhost:8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
```

---

## Verifying Your Progress

To run the verifier:
```bash
node verify-phase.js 6
```

The verifier will compile your code and run unit tests that check if a SQLite DB is initialized, tables are created, and CRUD operations successfully persist items to the database across store instances.
