package tools

import (
	"context"
	"testing"

	sdk "github.com/dibbla-agents/sdk-go"
)

func TestGreet(t *testing.T) {
	out, err := greet(context.Background(), GreetingInput{Name: " Ada "})
	if err != nil {
		t.Fatalf("greet: %v", err)
	}
	if out.Message != "Hello, Ada!" || out.CalledBy != "" {
		t.Fatalf("greet = %+v", out)
	}

	if _, err := greet(context.Background(), GreetingInput{}); err == nil {
		t.Fatal("greet without a name: want an error")
	}
}

func TestGreetNamesTheCaller(t *testing.T) {
	ctx := sdk.ContextWithCaller(context.Background(), sdk.Caller{
		UserID:   "u1",
		Email:    "ada@example.com",
		Identity: sdk.IdentityUserAuthenticated,
	})
	out, err := greet(ctx, GreetingInput{Name: "Ada"})
	if err != nil {
		t.Fatalf("greet: %v", err)
	}
	if out.CalledBy != "ada@example.com" {
		t.Fatalf("CalledBy = %q", out.CalledBy)
	}
}
