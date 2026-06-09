# Phase 3: Pointers and Interfaces

In this phase, you will learn two of Go's most important architectural features: **Pointers** (managing memory references) and **Interfaces** (implicit decoupling). You will transition your Todo store from global functions to a struct-based interface implementation.

---

## Conceptual Overview: Go vs. JavaScript & Python

### 1. Where are the Classes? (Building Services the Go Way)
In JS and Python, if you want to build a service that maintains state (like a todo store), you would define a **Class**:

```javascript
// JavaScript / TypeScript Class
class InMemoryTodoStore {
    constructor() {
        this.todos = [];
        this.nextID = 1;
    }

    add(title) {
        const todo = { id: this.nextID++, title, completed: false };
        this.todos.push(todo);
        return todo;
    }
}
```

In Go, **there are no classes, inheritance, or class constructors.** Instead, Go separates **Data** (using structs) from **Behavior** (using receiver methods) and ties them together:

```go
// Go Way: 1. Define the Data Schema (Struct)
type InMemoryTodoStore struct {
    todos  []Todo
    nextID int
}

// Go Way: 2. Define the Constructor Function
func NewInMemoryTodoStore() *InMemoryTodoStore {
    return &InMemoryTodoStore{
        todos:  []Todo{},
        nextID: 1,
    }
}

// Go Way: 3. Define the Behavior (Receiver Method)
// (s *InMemoryTodoStore) binds this function to the struct as a method
func (s *InMemoryTodoStore) Add(title string) Todo {
    todo := Todo{ID: s.nextID, Title: title, Completed: false}
    s.todos = append(s.todos, todo)
    s.nextID++
    return todo
}
```
This separation ensures that data structures remain lightweight and easily serializable, while behavior can be composed and decoupled using interfaces.

### 2. Pointers (Value vs. Reference Semantics)
- **JS/Python:** Objects and lists are passed by reference. If you pass an object/dictionary to a function and mutate its fields, the changes affect the caller's object.
- **Go:** Everything is **passed by value** (copied). If you pass a struct to a method or function, Go copies the entire struct. Changes made inside the function are lost when it returns!

To mutate a struct or pass it efficiently without copying, you must use **Pointers**:
- `*Type`: Declares a pointer (a variable that stores the memory address of a value).
- `&value`: The "address-of" operator. It returns the memory address of `value`.

```go
type Counter struct {
    count int
}

// VALUE RECEIVER: Copies the struct. Increment fails!
func (c Counter) IncrementValue() {
    c.count++ 
}

// POINTER RECEIVER: Receives the memory address. Increment succeeds!
func (c *Counter) IncrementPointer() {
    c.count++
}
```

#### Initializing and Returning Struct Pointers
To create a struct and get a pointer to it, use the `&` (address-of) operator on a struct literal:

```go
type User struct {
    Name string
}

// Returns a pointer to a User (*User)
func NewUser(name string) *User {
    return &User{
        Name: name,
    }
}
```

### 3. Interfaces (Implicit Duck Typing)
In many languages (like Java or TypeScript), a class must explicitly state that it implements an interface (e.g. `class MyStore implements TodoStore`).

Go has **implicit interfaces**. If a struct implements all the methods defined in an interface, Go automatically considers it to implement that interface. There is no `implements` keyword!

This is similar to Python's "duck typing" ("if it walks like a duck and quacks like a duck, it's a duck"), but Go validates this **at compile time**, ensuring static type safety!

---

## Hands-On Exercise

Create a file named `app/store.go` with the package declaration `package main`. Implement the following:

### 1. The `TodoStore` Interface
Define an interface named `TodoStore` that declares the following four method signatures:
```go
type TodoStore interface {
    Add(title string) Todo
    List() []Todo
    ToggleComplete(id int) bool
    Delete(id int) bool
}
```

### 2. The `InMemoryTodoStore` Struct
Define a struct named `InMemoryTodoStore` that will hold the state. The fields should be private (lowercase, to encapsulate them):
```go
type InMemoryTodoStore struct {
    todos  []Todo
    nextID int
}
```

### 3. Constructor function
Create a helper constructor function named `NewInMemoryTodoStore` that takes no arguments and returns a **pointer** to an initialized `InMemoryTodoStore` (`*InMemoryTodoStore`):
```go
func NewInMemoryTodoStore() *InMemoryTodoStore {
    // Return a pointer using &
}
```

### 4. Implement the Methods
Write the methods on `*InMemoryTodoStore` to satisfy the `TodoStore` interface. You must use **pointer receivers** (`(s *InMemoryTodoStore)`) because these methods mutate the store's internal state:

- **`Add(title string) Todo`**:
  - Creates a new `Todo` with `ID: s.nextID`, `Title: title`, `Completed: false`.
  - Appends it to `s.todos`.
  - Increments `s.nextID`.
  - Returns the created `Todo`.
- **`List() []Todo`**:
  - Returns `s.todos`.
- **`ToggleComplete(id int) bool`**:
  - Loops over `s.todos`. If matching `ID` is found, flips `Completed`, saves, and returns `true`.
  - Otherwise returns `false`.
- **`Delete(id int) bool`**:
  - Loops over `s.todos`. If matching `ID` is found, deletes it using the slice-append pattern:
    `s.todos = append(s.todos[:i], s.todos[i+1:]...)`
    and returns `true`.
  - Otherwise returns `false`.

---

## Verifying Your Progress

To run the verifier:
```bash
node verify-phase.js 3
```

The verifier will compile your code and run tests that:
1. Confirm `InMemoryTodoStore` implements the `TodoStore` interface.
2. Confirm state mutations persist across multiple method calls (validating correct pointer receiver usage).
