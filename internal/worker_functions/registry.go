package workerfunctions

import (
	"fmt"

	sdk "github.com/dibbla-agents/sdk-go"
	"github.com/dibbla-agents/go-worker-starter-template/internal/state"
)

// WorkerFunction defines the interface that all worker functions must implement
type WorkerFunction interface {
	GetName() string
	GetVersion() string
	GetDescription() string
	GetTags() []string
	Register(server *sdk.Server, ags *state.AsyncGlobalState) error
}

// Registry manages all registered worker functions
type Registry struct {
	functions []WorkerFunction
}

// NewRegistry creates a new function registry
func NewRegistry() *Registry {
	return &Registry{
		functions: make([]WorkerFunction, 0),
	}
}

// Register adds a new worker function to the registry
func (r *Registry) Register(fn WorkerFunction) {
	r.functions = append(r.functions, fn)
}

// RegisterAll registers all functions with the SDK server
func (r *Registry) RegisterAll(server *sdk.Server, ags *state.AsyncGlobalState) error {
	for _, fn := range r.functions {
		if err := fn.Register(server, ags); err != nil {
			return fmt.Errorf("failed to register function %s: %w", fn.GetName(), err)
		}
		fmt.Printf("✅ Registered function: %s (v%s)\n", fn.GetName(), fn.GetVersion())
	}
	return nil
}

// GetFunctions returns all registered functions
func (r *Registry) GetFunctions() []WorkerFunction {
	return r.functions
}

