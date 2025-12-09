package main

import (
	"log"
	"net/http"
	"os"

	sdk "github.com/dibbla-agents/sdk-go"
	"github.com/joho/godotenv"

	// Built-in functions
	"github.com/dibbla-agents/go-worker-starter-template/internal/worker_functions/greeting"

	// Frontend and HTTP handlers (optional - remove if not using frontend)
	"github.com/dibbla-agents/go-worker-starter-template/internal/frontend"
	httpgreeting "github.com/dibbla-agents/go-worker-starter-template/internal/http_handlers/greeting"

	// TODO: Import your worker functions here
	// Example:
	// myfunction "github.com/dibbla-agents/go-worker-starter-template/internal/worker_functions/my_function"

	// Advanced: For functions needing shared state (database, cache, etc.)
	// "github.com/dibbla-agents/go-worker-starter-template/internal/state"
	// workerfunctions "github.com/dibbla-agents/go-worker-starter-template/internal/worker_functions"
)

func main() {
	log.Println("🚀 Starting Worker...")

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  Warning: .env file not found, using system environment variables")
	}

	// Get configuration from environment
	serverName := os.Getenv("SERVER_NAME")
	if serverName == "" {
		serverName = "worker-starter"
	}

	serverApiToken := os.Getenv("SERVER_API_TOKEN")
	if serverApiToken == "" {
		log.Fatal("❌ SERVER_API_TOKEN environment variable is required")
	}

	// Create SDK server
	log.Println("🔧 Creating SDK server...")
	server, err := sdk.New(
		sdk.WithServerName(serverName),
		sdk.WithServerApiToken(serverApiToken),
	)
	if err != nil {
		log.Fatalf("❌ Failed to create SDK server: %v", err)
	}

	// Register worker functions
	log.Println("📝 Registering worker functions...")

	// Register the greeting function (simple example)
	greeting.Register(server)
	log.Println("   ✅ Registered: greeting")

	// TODO: Register your functions here
	// myfunction.Register(server)

	// Advanced: For functions needing shared state (database, etc.)
	// Uncomment the imports above and use:
	// ags, err := state.NewAsyncGlobalState()
	// registry := workerfunctions.NewRegistry()
	// registry.Register(examplefunction.NewExampleFunction())
	// registry.RegisterAll(server, ags)

	// Start HTTP server with frontend (optional - remove if not using frontend)
	router := frontend.NewRouter()
	httpgreeting.Register(router.Mux())
	log.Println("   ✅ HTTP: POST /api/greeting")

	go func() {
		log.Println("🌐 Starting HTTP server on :8080")
		if err := http.ListenAndServe(":8080", router.Handler()); err != nil {
			log.Printf("HTTP server error: %v", err)
		}
	}()

	// Start the server (blocks forever)
	log.Printf("🎯 Starting worker server '%s'...", serverName)
	if err := server.Start(); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
