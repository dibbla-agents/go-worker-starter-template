// Package tools holds the functions this MCP server offers. Each function is
// one tool: its name, description and input fields are what an MCP client
// shows to the model, so write them for a reader who has nothing else.
package tools

import sdk "github.com/dibbla-agents/sdk-go"

// Register adds every function to the server. Add your own here.
func Register(server *sdk.Server) {
	server.RegisterFunction(greeting())
}
