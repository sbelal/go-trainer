# Phase 4: Errors and JSON Manipulation

In this phase, you will master two fundamental topics in Go back-end development:
1. **Explicit Error Handling:** Why Go does not use exceptions (`try-catch`).
2. **JSON Serialization:** Using struct tags and the `encoding/json` package to work with JSON payloads.

---

## Conceptual Overview: Go vs. JavaScript & Python

### 1. No Exceptions (Explicit Error Handling)
- **JS/Python:** Errors are thrown as exceptions. If they aren't caught with `try-catch` / `try-except`, the program crashes.
  ```javascript
  try {
      let data = JSON.parse(raw);
  } catch (err) {
      console.error("Failed to parse JSON:", err);
  }
  ```
- **Go:** Errors are treated as normal return values. Functions that can fail return an `error` interface as their last return value. You must check it explicitly.
  ```go
  var todo Todo
  err := json.Unmarshal(raw, &todo)
  if err != nil {
      log.Println("Failed to parse JSON:", err)
      return err
  }
  ```
  This forces developers to deal with failures at the point of origin, leading to much more robust and predictable code.

### 2. Returning Multiple Values and Creating Errors
Go functions can return multiple values. This is commonly used to return a result alongside an error.

To return an error, you import the `"errors"` package and use `errors.New("message")`. If a function completes successfully, you return `nil` for the error.

```go
import "errors"

// Returns two values: a string and an error
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        // Return zero value for the result, and a new error
        return 0, errors.New("cannot divide by zero")
    }
    // Return the result and nil (no error)
    return a / b, nil
}
```

To call this function and handle the error:
```go
result, err := Divide(10, 0)
if err != nil {
    // Handle error
    fmt.Println("Error:", err)
    return
}
fmt.Println("Result:", result)
```

### 3. Struct Tags & JSON Mapping
By default, Go fields must start with an uppercase letter to be "exported" (publicly visible outside the package). Since standard JSON APIs use lowercase/camelCase (e.g., `title`, `id`), Go uses **Struct Tags** in backticks to map Go fields to JSON keys.

```go
type Todo struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}
```

#### JSON Serialization and Deserialization (`json.Marshal` & `json.Unmarshal`)
To work with JSON payloads, Go provides the `"encoding/json"` standard library package.

* **`json.Marshal` (Serialization):** Converts a Go value (like a struct or slice of structs) into a JSON byte slice (`[]byte`).
* **`json.Unmarshal` (Deserialization):** Parses JSON data in a byte slice (`[]byte`) and writes the result into a Go variable. **Crucial:** You must pass a pointer (using `&`) to the variable so that `json.Unmarshal` can modify it directly.

Here is an example demonstrating both:
```go
package main

import (
    "encoding/json" // Note: import path is "encoding/json", but we use prefix `json` in code
    "fmt"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func ExampleJSON() {
    u := User{Name: "Alice", Age: 30}

    // 1. Serialize (Marshal)
    jsonData, err := json.Marshal(u)
    if err != nil {
        fmt.Println("Error serializing:", err)
        return
    }
    fmt.Println(string(jsonData)) // Output: {"name":"Alice","age":30}

    // 2. Deserialize (Unmarshal)
    var parsedUser User
    // We pass &parsedUser (a pointer) so json.Unmarshal can write data into it
    err = json.Unmarshal(jsonData, &parsedUser)
    if err != nil {
        fmt.Println("Error deserializing:", err)
        return
    }
    fmt.Printf("Parsed: %+v\n", parsedUser)
}
```

---

## Hands-On Exercise

### Step 1: Add JSON Tags to `Todo` Struct
Modify the `Todo` struct inside `app/todo.go`. Add tags so that `ID` maps to `"id"`, `Title` maps to `"title"`, and `Completed` maps to `"completed"`.

### Step 2: Refactor `TodoStore` to return Errors
Modify `app/store.go` to add validation error handling:
1. Import `"errors"`.
2. Update the `TodoStore` interface's `Add` method signature to return both a `Todo` and an `error`:
   ```go
   type TodoStore interface {
       Add(title string) (Todo, error)
       List() []Todo
       ToggleComplete(id int) bool
       Delete(id int) bool
   }
   ```
3. Update the `InMemoryTodoStore` struct's `Add` method implementation:
   - Check if `title` is empty (`""`).
   - If empty, return an empty `Todo{}` and an error created using `errors.New("title cannot be empty")`.
   - If valid, add the todo to the slice and return `todo, nil`.

### Step 3: Implement JSON Serialization Helpers
Create a new file named `app/json.go` under `package main`. Implement two helper functions:

1. **`SerializeTodos`**:
   - Signature: `func SerializeTodos(todos []Todo) ([]byte, error)`
   - Converts a slice of `Todo` structs to a JSON byte array using `json.Marshal(todos)`.
   - Return the serialized byte array and the error.

2. **`DeserializeTodo`**:
   - Signature: `func DeserializeTodo(data []byte) (Todo, error)`
   - Parses a JSON byte array into a single `Todo` struct using `json.Unmarshal(data, &todo)`.
   - Return the parsed `Todo` and the error.

---

## Verifying Your Progress

To run the verifier:
```bash
node verify-phase.js 4
```

The verifier will compile your code and run unit tests to check:
1. The JSON tag mapping on `Todo` struct.
2. The empty title error validation on `InMemoryTodoStore.Add`.
3. The JSON serialization and deserialization functions (including handling invalid JSON payloads).
