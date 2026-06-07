# Phase 0: Environment Setup

Welcome to Go-Trainer! In this phase, you will configure your Go workspace, initialize a Go module, and write your first compilation-ready Go program.

---

## Conceptual Overview: Go vs. JavaScript & Python

Before writing code, let's understand how Go's design differs from languages you might already know:

### 1. Compiled vs. Interpreted
- **Node.js/Python:** These are interpreted/JIT-compiled languages. You feed source files directly to the runtime (`node app.js` or `python app.py`).
- **Go:** Go is a **statically compiled** language. The Go compiler compiles your code directly into a single, self-contained binary executable for the target operating system (e.g., `.exe` on Windows, or an ELF binary on Linux). You do not need Go installed on a target server to run a compiled Go binary.

### 2. Dependency Management
- **Node.js:** Uses `npm` and a `package.json` file.
- **Python:** Uses `pip` and a `requirements.txt` file (or `pyproject.toml`/`Pipfile`).
- **Go:** Uses Go Modules. You initialize your module with `go mod init <module-name>`, which creates a `go.mod` file. This file tracks your dependencies and Go toolchain version.

### 3. Entrypoint
- **Node.js/Python:** Any script file can be run, starting execution from line 1 of that file.
- **Go:** Executables always start execution from a package named `main` in a function named `main()`.

---

## Tools of the Trade

Here are the primary Go commands you'll use throughout your career:

- **`go mod init <name>`**: Initializes a new Go module in the current directory.
- **`go run <file.go>`**: Compiles your code in memory and executes it immediately. Great for development.
- **`go build`**: Compiles your package and writes the binary output to the disk.
- **`go test ./...`**: Runs tests in the current directory and all subdirectories.
- **`go fmt`**: Automatically formats your Go code following official conventions (tabs, spacing, bracket alignment). Go developers have zero arguments about code style because the compiler toolchain enforces one standard style!

---

## Hands-On Exercise: Create "Hello World"

Your objective in this phase is to set up a new module named `todo-api` inside the `app/` folder and write a program that prints `Hello, World!`.

### Step 1: Create the App Directory and Initialize the Module
Create a folder named `app` in the root of this repository. In your terminal, navigate into `app` and initialize a module named `todo-api`:

```bash
mkdir app
cd app
go mod init todo-api
```
This will create a `go.mod` file. Open it to see its contents. It should specify the module name and your Go compiler version.

### Step 2: Create `main.go`
Inside the `app` folder, create a file named `main.go` and write the following code:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

### Step 3: Run Your Code
Verify that it works by running this command in your terminal from inside the `app` folder:
```bash
go run main.go
```
It should print:
```
Hello, World!
```

---

## Verifying Your Progress
Once you have created `app/go.mod` and `app/main.go`, return to the root of the `go-trainer` directory and run:
```bash
node verify-phase.js 0
```
If everything is set up correctly, all checks will pass, and you will unlock Phase 1!
