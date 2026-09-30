// An MCP server on Dibbla.
//
// This process is a Dibbla tool server: it connects out to the platform,
// registers the functions in tools/, and serves calls to them. Deployed as an
// app whose dibbla.yaml carries `mcp: <name>`, every function becomes a tool
// at https://mcp.<your-dibbla-domain>/platform/servers/<name>.
//
// There is no MCP code here. The platform speaks MCP, handles login and
// decides who may call; you write functions.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	sdk "github.com/dibbla-agents/sdk-go"

	"github.com/dibbla-agents/go-worker-starter-template/_optional/mcp/tools"
)

func main() {
	// The name this server registers under. It is the last part of the MCP
	// address, and it must be the same as `mcp:` in dibbla.yaml.
	name := os.Getenv("SERVER_NAME")
	if name == "" {
		log.Fatal("SERVER_NAME is not set. dibbla.yaml sets it for a deployed app; for a local run, copy env.example to .env.")
	}

	// Everything else comes from the environment (see env.example). On the
	// platform the credential is the app's own workload identity, so a
	// deployed app needs no token.
	server, err := sdk.New(sdk.WithServerName(name))
	if err != nil {
		log.Fatalf("creating the tool server: %v", err)
	}

	// Each function registered here is one tool on the MCP address.
	tools.Register(server)

	go serveStatus(name)

	log.Printf("starting tool server %q", name)
	if err := server.Start(); err != nil {
		log.Fatalf("tool server stopped: %v", err)
	}
}

// serveStatus answers on the app's URL. The tools are not served here: they
// are on the platform's MCP address. This page is what makes the app show up
// with a URL in the console, and what the platform's health check reaches.
func serveStatus(name string) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s is an MCP server on Dibbla.\n\nConnect a client to it with:\n\n    dibbla mcp server %s\n", name, name)
	})
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Printf("status page stopped: %v", err)
	}
}
