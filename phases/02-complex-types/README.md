# Phase 2: Complex Types (Slices, Maps, and Structs)

In this phase, you will learn about Go's structural data types and collections: Structs, Slices (dynamic arrays), and Maps (hash maps/dictionaries).

---

## Conceptual Overview: Go vs. JavaScript & Python

### 1. Structs (Custom Data Blueprints)
Go is not a class-based object-oriented language. It does not have classes or inheritance. Instead, it uses **Structs** to define user-defined custom schemas, representing structured data.

- **JS/Python:** You use classes or raw object/dictionary literals.
- **Go:** You define a strict layout.
```go
type Todo struct {
    ID        int
    Title     string
    Completed bool
}
```

#### Initializing Struct Variables
To use a struct, you can instantiate and initialize it in a few different ways:

1. **Inline Initialization (Struct Literals):**
   You specify the field names and their values inside curly braces `{}`. This is the most common way to initialize a struct with values.
   ```go
   todo := Todo{
       ID:        1,
       Title:     "Buy Milk",
       Completed: false,
   }
   ```

2. **Declaration and Field Assignment:**
   You declare the variable using `var` (which automatically initializes fields to their zero values) and then assign values to the fields individually using the dot (`.`) operator.
   ```go
   var todo Todo
   todo.ID = 1
   todo.Title = "Buy Milk"
   todo.Completed = false
   ```

### 2. Slices (Dynamic Arrays)
Go has **Arrays** (fixed size, immutable length) and **Slices** (dynamically-sized views into arrays). In practice, you will use **Slices** 99% of the time.

- **JS:** `let list = [1, 2, 3]; list.push(4);`
- **Python:** `list = [1, 2, 3]; list.append(4)`
- **Go:** 
  ```go
  var list []int = []int{1, 2, 3}
  list = append(list, 4) // Note: append returns a new slice!
  ```

### 3. Deleting from a Slice
In JS or Python, deleting an item from an array is simple (e.g. `list.splice(i, 1)` or `del list[i]`).
Go has **no built-in delete function** for slices! To delete an item at index `i` from slice `s`, you must concatenate the parts before and after index `i` using the `...` (variadic unpacking) operator:
```go
s = append(s[:i], s[i+1:]...)
```
> 💡 **How this works:** `s[:i]` takes everything up to (but not including) index `i`. `s[i+1:]...` takes everything from index `i+1` to the end, unpacking the elements into `append()`.

### 4. Maps (Hash Tables / Dictionaries)
Maps associate keys with values.
- **JS:** `const m = {}; m["key"] = "val"; delete m["key"];`
- **Python:** `m = {}; m["key"] = "val"; del m["key"]`
- **Go:**
  ```go
  // Initializing a map
  m := make(map[string]string)
  m["key"] = "val"
  
  // Deleting a key
  delete(m, "key")
  ```

### 5. Iterating over Slices and Maps (`for range`)
To loop through a slice or map, Go uses the `for range` construct. This returns both the **index** and the **value** (or key and value for maps).

```go
names := []string{"Alice", "Bob", "Charlie"}

// Loop with index and value
for index, name := range names {
    fmt.Printf("Index: %d, Value: %s\n", index, name)
}

// If you don't need the index, use the blank identifier (_)
for _, name := range names {
    fmt.Println(name)
}
```

> [!WARNING]
> **Modifying Structs in a Loop (Copy-by-Value Pitfall):**
> When you iterate over a slice of structs (e.g., `[]Todo`) using `for _, todo := range Todos`, the `todo` variable is a **copy** of the element in the slice.
> Modifying `todo.Completed = true` inside the loop will **not** modify the actual element in the `Todos` slice!
> To modify the element in the slice, you must access it using its index:
> ```go
> for i, todo := range Todos {
>     if todo.ID == targetID {
>         // Access the slice element directly by index to modify it
>         Todos[i].Completed = !Todos[i].Completed
>     }
> }
> ```

---

## Hands-On Exercise

Create a file named `app/todo.go` with the package declaration `package main`. Implement the following:

### 1. The `Todo` Struct
Define a struct named `Todo` with three fields:
- `ID` (int)
- `Title` (string)
- `Completed` (bool)

### 2. Global Storage
Define two global variables:
- `var Todos []Todo` (a slice of `Todo` structs, initialized as `nil` or empty)
- `var nextID = 1` (an integer to keep track of the auto-incrementing IDs)

### 3. Add Todo Item
Create a function named `AddTodoItem` that accepts a `title` (string) and returns the created `Todo` struct:
- It should instantiate a new `Todo` with `ID: nextID`, `Title: title`, and `Completed: false`.
- Increment `nextID` by `1`.
- Append the new `Todo` to the global `Todos` slice.
- Return the created `Todo`.

```go
func AddTodoItem(title string) Todo {
    // Implement logic here
}
```

### 4. List Todos
Create a function named `ListTodos` that takes no arguments and returns a slice of `Todo` structs:
- Simply return the global `Todos` slice.

### 5. Toggle Todo Completion Status
Create a function named `ToggleTodoComplete` that accepts an `id` (int) and returns a boolean (`bool`):
- Loop through the global `Todos` slice.
- If a todo's `ID` matches the input `id`, flip its `Completed` status (i.e. `true` becomes `false`, `false` becomes `true`), save it back to the slice, and return `true`.
- If no matching todo is found after the loop, return `false`.

### 6. Delete Todo Item
Create a function named `DeleteTodoItem` that accepts an `id` (int) and returns a boolean (`bool`):
- Loop through the global `Todos` slice to find the item's index.
- If found, remove it from the slice using the slice-append pattern:
  `Todos = append(Todos[:i], Todos[i+1:]...)`
- Return `true`.
- If no matching todo is found, return `false`.

---

## Verifying Your Progress

To run the verifier for this phase:
```bash
node verify-phase.js 2
```

The verifier will compile your code and run unit tests to check that struct fields match, IDs increment correctly, completion toggling works, and slice deletion operates without errors.
