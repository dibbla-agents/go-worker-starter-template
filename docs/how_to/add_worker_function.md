# How to Add a Worker Function

Quick guide to creating a new worker function in 3 steps.

## Step 1: Create Function Folder & File

Create a new folder in `internal/worker_functions/your_function/` with a `functions.go` file:

```go
package yourfunction

import (
	"fmt"
	"github.com/FatsharkStudiosAB/codex/workflows/workers/go/sdk"
	"worker_starter_template/internal/state"
)

// Define Input/Output
type YourFunctionInput struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type YourFunctionOutput struct {
	Result string `json:"result"`
}

// Function struct and constructor
type YourFunction struct{}

func NewYourFunction() *YourFunction {
	return &YourFunction{}
}

// Required methods
func (f *YourFunction) GetName() string {
	return "your_function_name" // Must be unique
}

func (f *YourFunction) GetVersion() string {
	return "1.0.0"
}

func (f *YourFunction) GetDescription() string {
	return "Describe what your function does"
}

func (f *YourFunction) GetTags() []string {
	return []string{"tag1", "tag2"}
}

// Register with SDK
func (f *YourFunction) Register(server *sdk.Server, ags *state.AsyncGlobalState) error {
	fn := sdk.NewSimpleFunction[YourFunctionInput, YourFunctionOutput](
		f.GetName(),
		f.GetVersion(),
		f.GetDescription(),
	).WithHandler(func(input YourFunctionInput) (YourFunctionOutput, error) {
		return f.handler(input, ags)
	}).WithTags(f.GetTags()...)

	server.RegisterFunction(fn)
	return nil
}

// Your logic goes here
func (f *YourFunction) handler(input YourFunctionInput, ags *state.AsyncGlobalState) (YourFunctionOutput, error) {
	// Validate
	if input.Name == "" {
		return YourFunctionOutput{}, fmt.Errorf("name is required")
	}

	// Process
	result := fmt.Sprintf("Hello %s, you are %d years old", input.Name, input.Age)

	return YourFunctionOutput{
		Result: result,
	}, nil
}
```

## Step 2: Register in `cmd/worker/main.go`

Add your import:
```go
yourfunction "worker_starter_template/internal/worker_functions/your_function"
```

Register the function:
```go
registry.Register(yourfunction.NewYourFunction())
```

## Step 3: Test

```bash
go build ./cmd/worker/
./worker
```

## Key Points

- **Unique name**: `GetName()` must be unique across all functions
- **JSON tags**: Required for all Input/Output fields
- **Error handling**: Return errors from handler when validation fails
- **Database access**: Use `ags` (AsyncGlobalState) for database/shared resources (see `example_function`)

## Example Reference

See `internal/worker_functions/example_function/functions.go` for a complete working example.

