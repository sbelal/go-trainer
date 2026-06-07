# 🐹 Go-Trainer

**Master Go by building a production-ready REST API.** A structured, step-by-step roadmap to learning Go by doing, with built-in verification to ensure you've nailed each concept before moving on.

By the end of this course, you will have built a complete, database-backed REST API for a TODO list application using Go, SQLite, and clean architecture.

---

## Who Is This For?

This trainer is designed for software engineers who:

- ✅ Know programming fundamentals in a dynamic language like **JavaScript (Node.js)** or **Python**
- ✅ Have a command line environment (Windows Command Prompt/PowerShell, macOS, or Linux)
- 🌱 Are complete beginners or intermediate-level developers looking to adopt Go

We will use your existing knowledge of JavaScript and Python as mental models, contrasting Go's static typing, compilation, memory layout, interfaces, and error handling philosophies with their Python and JS equivalents.

---

## Prerequisites

Ensure you have the following installed on your system:

| Tool | Recommended Version | Purpose | Verify Command |
|---|---|---|---|
| **Go** | `1.22+` | The Go compiler and toolchain | `go version` |
| **Node.js** | `18+` | Runs the verification engine | `node --version` |
| **Git** | `Any` | Version control | `git --version` |
| **VSCode** | `Latest` | Recommended editor | [Download](https://code.visualstudio.com/) |

### Recommended VSCode Extensions
- **Go** (`golang.go`): Official extension for Go (provides autocomplete, formatting on save, linting, debugging, and test execution).
- **Markdown Preview Enhanced** (`shd101wyy.markdown-preview-enhanced`): For viewing lesson files in a rich rendered view.

---

## Getting Started

### 1. Clone the Repository
Clone this repository locally and open it in your preferred IDE:
```bash
git clone <this-repo-url> go-trainer
cd go-trainer
code .
```

### 2. Verify Your Environment
Before starting, ensure Go and Node.js are correctly installed on your system by running:
```bash
go version
node --version
```
If both commands return their respective versions, you are ready to start! Head over to Phase 0 to set up your project workspace.

---

## How It Works

1. **Read the Lesson:** Go to the phase's folder (e.g. `phases/00-environment-setup/`) and read the `README.md`.
2. **Do the Hands-On Exercise:** Write/modify the Go source code inside the `app/` folder according to the instructions.
3. **Verify Your Progress:** Run the verifier tool for that phase:
   ```bash
   node verify-phase.js <phase_number>
   ```
4. **Iterate:** If a check fails (❌), read the feedback and the suggested lesson section, fix the code, and re-run the verifier.
5. **Next Phase:** Once all checks pass (✅), move to the next phase!

---

## The Go-Trainer Roadmap

You will build a single TODO REST API inside the `app/` folder, adding layers of functionality phase-by-phase.

| Phase | Topic | What You Will Learn & Do |
|:---:|---|---|
| **[Phase 0: Environment Setup](phases/00-environment-setup/README.md)** | Go Toolchain & Hello World | Configure your workspace, initialize Go modules, compile, and run your first Go code. |
| **[Phase 1: Basics & Syntax](phases/01-basics-and-syntax/README.md)** | Go Language Fundamentals | Variables, types, Go's strict static compilation, functions with multi-returns, control flow, and zero values. |
| **[Phase 2: Complex Types](phases/02-complex-types/README.md)** | Slices, Maps & Structs | Manipulate dynamic lists (Slices), key-value lookups (Maps), and define data layouts with Structs. |
| **[Phase 3: Pointers & Interfaces](phases/03-pointers-and-interfaces/README.md)** | Reference Semantics & Duck Typing | Understand pointers (`*` vs `&`), value vs pointer receivers, and decoupling code using implicit Go interfaces. |
| **[Phase 4: Errors & JSON](phases/04-errors-and-json/README.md)** | Error Handling & Marshalling | Handle errors explicitly (no `try-catch`!), add JSON struct tags, and marshal/unmarshal data payload requests. |
| **[Phase 5: HTTP REST API](phases/05-rest-api/README.md)** | Server Endpoints & Routing | Build an HTTP server using `net/http` and use Go 1.22's path-matching router for RESTful requests. |
| **[Phase 6: SQLite Integration](phases/06-sqlite-integration/README.md)** | SQL Database Persistence | Connect to SQLite database using a CGO-free SQL driver, run table schemas, and perform CRUD queries. |
| **[Phase 7: Clean Architecture](phases/07-clean-architecture/README.md)** | Package Structuring & Configs | Reorganize files into standard layouts (`cmd/`, `internal/`), use env variables, and write robust testing. |
| **[Phase 8: Concurrency & Mutexes](phases/08-concurrency-mutexes/README.md)** | Thread Safety & Mutexes | Understand Go's multi-threaded model, discover data races, and use `sync.RWMutex` to secure in-memory data structures. |
| **[Phase 9: HTTP Middleware](phases/09-http-middleware/README.md)** | Request Interception | Implement logging and API Key authentication middleware using Go's handler decorator pattern. |
| **[Phase 10: Context & Timeout](phases/10-context-cancellations/README.md)** | Lifecycle Cancellations | propagate `context.Context` to HTTP pipelines and SQL queries to auto-cancel slow operations. |
| **[Phase 11: Testing & Mocking](phases/11-testing-mocking/README.md)** | Idiomatic Unit Testing | Write table-driven test configurations and implement mocks to test controllers in complete database isolation. |

---

## License

MIT
