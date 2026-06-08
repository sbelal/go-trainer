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

#### `if` statements
In Go, `if` statements do not require parentheses `()` around the condition, but curly braces `{}` are **always** mandatory for the code blocks, even if there is only a single statement.

```go
// Example of if-else structure
priority := 7

if priority >= 5 {
    // executes if condition is true
    fmt.Println("High priority")
} else if priority >= 3 {
    // executes if condition is false but this else-if is true
    fmt.Println("Medium priority")
} else {
    // executes if all conditions are false
    fmt.Println("Low priority")
}
```

#### `for` loops
The `for` loop is Go's **only** loop construct. There is no `while` or `do-while` loop in Go! You can represent a standard C-style three-component loop, or a while-style conditional loop:

```go
// Standard three-component for loop: init; condition; post
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// "While" loop style: condition only
count := 0
for count < 3 {
    fmt.Println(count)
    count++
}
```

### 5. Functions and String Formatting

#### Functions and Return Values
In Go, functions are declared with the `func` keyword. You must specify the types of all parameters and the return type(s) explicitly.

Here is an example function that takes two numbers, adds them, and returns the result:
```go
// func Name(parameterName parameterType) returnType
func AddNumbers(a int, b int) int {
    sum := a + b
    return sum // Uses the return keyword to send a value back
}
```

If a function does not return anything, you simply omit the return type:
```go
func SayHello(name string) {
    fmt.Println("Hello, " + name)
}
```

#### String Formatting with `fmt.Sprintf`
To construct formatted strings, Go provides the `fmt.Sprintf` function (part of the standard `"fmt"` package). It works similarly to template literals in JS or f-strings in Python, but uses verb specifiers (like `%s` for strings, `%d` for integers, `%f` or `%v` for general values).

Unlike `fmt.Printf` (which prints directly to the console), `fmt.Sprintf` returns the formatted string so you can use it in variables or return values.

To use functions from the `"fmt"` package, you must import it at the top of your file:
```go
package main

import "fmt" // Import the fmt package

func FormatGreeting(name string, age int) string {
    // %s is for string, %d is for integer
    message := fmt.Sprintf("Hello %s, you are %d years old.", name, age)
    return message
}
```

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
