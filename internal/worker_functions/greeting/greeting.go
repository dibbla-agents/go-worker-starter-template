// Package greeting provides a simple worker function example.
// This is a minimal example - see example_function_advanced for a full-featured template.
package greeting

import (
	"fmt"

	sdk "github.com/dibbla-agents/sdk-go"
)

// Input defines what the function receives
type GreetingInput struct {
	Name string `json:"name"`
}

// Output defines what the function returns
type GreetingOutput struct {
	Message string `json:"message"`
}

// Register registers the greeting function with the SDK server
func Register(server *sdk.Server) {
	fn := sdk.NewSimpleFunction[GreetingInput, GreetingOutput](
		"greeting",
		"1.0.0",
		"Generate a greeting message",
	).WithHandler(func(input GreetingInput) (GreetingOutput, error) {
		if input.Name == "" {
			return GreetingOutput{}, fmt.Errorf("name is required")
		}
		return GreetingOutput{
			Message: fmt.Sprintf("Hello, %s!", input.Name),
		}, nil
	})

	server.RegisterFunction(fn)
}

