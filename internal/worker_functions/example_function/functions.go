// Package examplefunction provides a template for creating worker functions.
//
// To create your own function:
// 1. Copy this folder and rename it
// 2. Update package name to match folder
// 3. Customize Input/Output structs
// 4. Implement your logic in handler()
// 5. Import and register in cmd/worker/main.go
package examplefunction

import (
	"fmt"
	"log"

	"github.com/FatsharkStudiosAB/codex/workflows/workers/go/sdk"
	"worker_starter_template/internal/state"
)

// Input defines what the function receives
type ExampleInput struct {
	Name  string `json:"name"`  // Required field
	Count int    `json:"count"` // Optional field (set default in handler)
}

// Output defines what the function returns
type ExampleOutput struct {
	Result  string `json:"result"`
	Success bool   `json:"success"`
}

// ExampleFunction struct
type ExampleFunction struct{}

// NewExampleFunction creates a new instance
func NewExampleFunction() *ExampleFunction {
	return &ExampleFunction{}
}

// GetName returns the unique function identifier
func (f *ExampleFunction) GetName() string {
	return "example_function"
}

// GetVersion returns the function version
func (f *ExampleFunction) GetVersion() string {
	return "1.0.0"
}

// GetDescription returns what the function does
func (f *ExampleFunction) GetDescription() string {
	return "Example function that processes input and returns a result"
}

// GetTags returns searchable tags
func (f *ExampleFunction) GetTags() []string {
	return []string{"example", "template"}
}

// Register registers the function with the SDK server
func (f *ExampleFunction) Register(server *sdk.Server, ags *state.AsyncGlobalState) error {
	fn := sdk.NewSimpleFunction[ExampleInput, ExampleOutput](
		f.GetName(),
		f.GetVersion(),
		f.GetDescription(),
	).WithHandler(func(input ExampleInput) (ExampleOutput, error) {
		return f.handler(input, ags)
	}).WithTags(f.GetTags()...)

	server.RegisterFunction(fn)
	return nil
}

// handler implements the function logic
func (f *ExampleFunction) handler(input ExampleInput, ags *state.AsyncGlobalState) (ExampleOutput, error) {
	log.Printf("🚀 Processing: %s", input.Name)

	// Validate input
	if input.Name == "" {
		return ExampleOutput{}, fmt.Errorf("name is required")
	}

	// Set defaults
	count := input.Count
	if count <= 0 {
		count = 1
	}

	// Your logic here
	result := fmt.Sprintf("Processed %s %d times", input.Name, count)

	// Example: Database access (if AsyncGlobalState is initialized)
	// if ags != nil && ags.DB != nil {
	// 	var records []models.MyModel
	// 	ags.DB.Where("name = ?", input.Name).Find(&records)
	// }

	log.Printf("✅ Completed successfully")

	return ExampleOutput{
		Result:  result,
		Success: true,
	}, nil
}

