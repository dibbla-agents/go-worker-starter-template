package tools

import (
	"context"
	"errors"
	"fmt"
	"strings"

	sdk "github.com/dibbla-agents/sdk-go"
)

// GreetingInput is the tool's input. The json tags are the argument names
// the client sees.
type GreetingInput struct {
	Name string `json:"name"`
}

// GreetingOutput is what the tool returns.
type GreetingOutput struct {
	Message  string `json:"message"`
	CalledBy string `json:"called_by,omitempty"`
}

// greeting is the example tool. Copy it, rename it, and replace the handler.
func greeting() *sdk.SimpleFunction[GreetingInput, GreetingOutput] {
	return sdk.NewSimpleFunction[GreetingInput, GreetingOutput](
		"greeting",
		"1.0.0",
		"Greet a person by name. Returns the greeting and the email of the signed-in user who asked for it.",
	).WithContextHandler(greet)
}

func greet(ctx context.Context, in GreetingInput) (GreetingOutput, error) {
	name := strings.TrimSpace(in.Name)
	if name == "" {
		return GreetingOutput{}, errors.New("name is required")
	}
	out := GreetingOutput{Message: fmt.Sprintf("Hello, %s!", name)}

	// The platform tells the function who is calling. The identity comes from
	// the caller's Dibbla login, not from the tool's arguments, so it can be
	// trusted when deciding what a person may read or change.
	if caller, ok := sdk.CallerFromContext(ctx); ok && caller.IsUser() {
		out.CalledBy = caller.Email
	}
	return out, nil
}
