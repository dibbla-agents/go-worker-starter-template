# How to Add a Worker Function

This guide shows two approaches: **Simple** (for most use cases) and **Advanced** (for functions needing shared resources).

---

## Option 1: Simple Function (Recommended)

Best for stateless functions that don't need database or shared resources.

### Step 1: Create Function File

Create `internal/worker_functions/your_function/your_function.go`:

```go
package yourfunction

import (
	"fmt"

	sdk "github.com/dibbla-agents/sdk-go"
)

// Input defines what the function receives
type YourInput struct {
	Name string `json:"name"`
}

// Output defines what the function returns
type YourOutput struct {
	Message string `json:"message"`
}

// Register registers the function with the SDK server
func Register(server *sdk.Server) {
	fn := sdk.NewSimpleFunction[YourInput, YourOutput](
		"your_function",  // Unique name
		"1.0.0",          // Version
		"Description of what it does",
	).WithHandler(func(input YourInput) (YourOutput, error) {
		if input.Name == "" {
			return YourOutput{}, fmt.Errorf("name is required")
		}
		return YourOutput{
			Message: fmt.Sprintf("Hello, %s!", input.Name),
		}, nil
	})

	server.RegisterFunction(fn)
}
```

### Step 2: Register in main.go

Add the import and register call:

```go
import (
	// ... existing imports ...
	yourfunction "github.com/dibbla-agents/go-worker-starter-template/internal/worker_functions/your_function"
)

func main() {
	// ... server setup ...

	// Register your function
	yourfunction.Register(server)
	log.Println("   ✅ Registered: your_function")

	// ... start server ...
}
```

### Step 3: Build and Test

**Linux/Mac:**
```bash
go build -o worker ./cmd/worker && ./worker
```

**Windows (PowerShell):**
```powershell
go build -o worker.exe .\cmd\worker; .\worker.exe
```

---

## Option 2: Advanced Function (With Shared State)

Use this when your function needs database access, cache, or other shared resources.

### Step 1: Create Function File

Create `internal/worker_functions/your_function/functions.go`:

```go
package yourfunction

import (
	"fmt"
	"log"

	sdk "github.com/dibbla-agents/sdk-go"
	"github.com/dibbla-agents/go-worker-starter-template/internal/state"
)

// Input/Output structs
type YourInput struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type YourOutput struct {
	Result string `json:"result"`
}

// YourFunction implements the WorkerFunction interface
type YourFunction struct{}

func NewYourFunction() *YourFunction {
	return &YourFunction{}
}

// Required interface methods
func (f *YourFunction) GetName() string        { return "your_function" }
func (f *YourFunction) GetVersion() string     { return "1.0.0" }
func (f *YourFunction) GetDescription() string { return "Describe what it does" }
func (f *YourFunction) GetTags() []string      { return []string{"tag1", "tag2"} }

// Register with SDK (receives AsyncGlobalState for shared resources)
func (f *YourFunction) Register(server *sdk.Server, ags *state.AsyncGlobalState) error {
	fn := sdk.NewSimpleFunction[YourInput, YourOutput](
		f.GetName(),
		f.GetVersion(),
		f.GetDescription(),
	).WithHandler(func(input YourInput) (YourOutput, error) {
		return f.handler(input, ags)
	}).WithTags(f.GetTags()...)

	server.RegisterFunction(fn)
	return nil
}

// Handler with access to shared state
func (f *YourFunction) handler(input YourInput, ags *state.AsyncGlobalState) (YourOutput, error) {
	log.Printf("Processing: %s", input.Name)

	// Validate
	if input.Name == "" {
		return YourOutput{}, fmt.Errorf("name is required")
	}

	// Example: Database access (if AGS is configured)
	// if ags != nil && ags.DB != nil {
	// 	var records []models.YourModel
	// 	ags.DB.Where("name = ?", input.Name).Find(&records)
	// }

	result := fmt.Sprintf("Hello %s, you are %d years old", input.Name, input.Age)
	return YourOutput{Result: result}, nil
}
```

### Step 2: Register with Registry

In `cmd/worker/main.go`, use the registry pattern:

```go
import (
	"github.com/dibbla-agents/go-worker-starter-template/internal/state"
	workerfunctions "github.com/dibbla-agents/go-worker-starter-template/internal/worker_functions"
	yourfunction "github.com/dibbla-agents/go-worker-starter-template/internal/worker_functions/your_function"
)

func main() {
	// ... server setup ...

	// Initialize shared state
	ags, err := state.NewAsyncGlobalState()
	if err != nil {
		log.Fatalf("Failed to initialize state: %v", err)
	}
	defer ags.Close()

	// Use registry for advanced functions
	registry := workerfunctions.NewRegistry()
	registry.Register(yourfunction.NewYourFunction())
	
	if err := registry.RegisterAll(server, ags); err != nil {
		log.Fatalf("Failed to register functions: %v", err)
	}

	// ... start server ...
}
```

---

## Key Points

| Aspect | Simple | Advanced |
|--------|--------|----------|
| **Use when** | Stateless, no DB | Needs DB, cache, shared resources |
| **Boilerplate** | Minimal (~30 lines) | More structure (~60 lines) |
| **Registration** | Direct: `fn.Register(server)` | Via registry with AGS |
| **State access** | None | Via `AsyncGlobalState` |

- **Unique name**: Function names must be unique across all workers
- **JSON tags**: Required for all Input/Output struct fields
- **Error handling**: Return descriptive errors for validation failures
- **Versioning**: Update version when changing input/output contracts

---

## Example References

- **Simple**: `internal/worker_functions/greeting/greeting.go`
- **Advanced**: `internal/worker_functions/example_function/functions.go`
