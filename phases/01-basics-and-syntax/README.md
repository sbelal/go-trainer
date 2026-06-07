# Phase 1: Basics and Syntax

In this phase, you will learn the core building blocks of Go: variable declarations, static types, type conversion, basic control flow, and functions.

---

## Conceptual Overview: Go vs. JavaScript & Python

Go is statically typed and strongly typed. It does not perform implicit type conversions (coercion) for you.

### 1. Variables and Type Inference
In JS/Python, variables are dynamic:
```javascript
// JS
let name = "Buy Milk";
name = 42; // Allowed
```

In Go, variables are static. Once declared, they cannot change their type:
```go
// Go
var name string = "Buy Milk"
// or shorthand (inside functions only):
name := "Buy Milk" 

name = 42 // COMPILE ERROR: cannot use 42 (type int) as type string
```

### 2. No Implicit Type Coercion (Strict Type Safety)
In JS or Python, you can mix ints and floats easily:
```python
# Python
total = 10     # int
completed = 3.5 # float
rate = completed / total # Allowed: auto-converts total to float
```

In Go, this is a compilation error! You must convert types **explicitly**:
```go
// Go
var total int = 10
var completed float64 = 3.5

// COMPILE ERROR: invalid operation: completed / total (mismatched types float64 and int)
var rate float64 = completed / float64(total) // CORRECT
```

### 3. Zero Values
In JS, uninitialized variables are `undefined`. In Python, they don't exist until assigned, or they are `None`.
In Go, variables always have a default **Zero Value** if not explicitly initialized:
- `int`, `float64` $\rightarrow$ `0`
- `string` $\rightarrow$ `""` (empty string)
- `bool` $\rightarrow$ `false`
- Pointers, slices, maps, channels, interfaces $\rightarrow$ `nil`

### 4. Control Flow: If & For
- **`if` statements:** Do not require parentheses `()` around the condition, but curly braces `{}` are **always** mandatory.
- **`for` loops:** The `for` loop is Go's **only** loop construct. There is no `while` or `do-while` loop in Go! You represent a `while` loop simply using `for <condition>`.

---

## Hands-On Exercise

Create a file named `app/basics.go` with the package declaration `package main`. Implement the following:

### 1. Module-level Variable
Define a global variable named `AppName` with the value `"TodoApp"`.
```go
package main

var AppName = "TodoApp"
```

### 2. Formatter Function
Create a function named `FormatTodo` that accepts a `title` (string) and `completed` (bool), and returns a formatted string.
- If `completed` is `true`, return `"<title> [Completed]"` (e.g. `"Buy Milk [Completed]"`).
- If `completed` is `false`, return `"<title> [Pending]"` (e.g. `"Buy Milk [Pending]"`).

```go
func FormatTodo(title string, completed bool) string {
    // Hint: You can use standard string concatenation (+) or fmt.Sprintf
}
```

### 3. Priority Evaluator
Create a function named `IsPriorityHigh` that accepts `priority` (int) and returns a boolean (`bool`).
- Return `true` if `priority` is greater than or equal to `5`.
- Return `false` otherwise.

### 4. Rate Calculator
Create a function named `CalculateCompletionRate` that accepts `completedCount` (int) and `totalCount` (int) and returns a `float64` rate.
- If `totalCount` is `0`, return `0.0` (to avoid division-by-zero errors).
- Otherwise, divide `completedCount` by `totalCount` and return the result.
- **Crucial hint:** You must explicitly cast both integers to `float64` before division: `float64(a) / float64(b)`.

---

## Verifying Your Progress

To run the verifier for this phase:
```bash
node verify-phase.js 1
```

Behind the scenes, the verification engine will compile your code and run tests against your functions to verify they handle all conditions (including division by zero) correctly.
