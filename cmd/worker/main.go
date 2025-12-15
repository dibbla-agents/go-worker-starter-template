package main

import (
	"bytes"
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

// loadEnvFile loads environment variables from .env file, handling Windows UTF-8 BOM
func loadEnvFile() error {
	content, err := os.ReadFile(".env")
	if err != nil {
		return err
	}

	// Strip UTF-8 BOM if present (common on Windows)
	content = bytes.TrimPrefix(content, []byte{0xEF, 0xBB, 0xBF})

	// Parse the content
	envMap, err := godotenv.Unmarshal(string(content))
	if err != nil {
		return err
	}

	// Set environment variables
	for k, v := range envMap {
		os.Setenv(k, v)
	}
	return nil
}

func main() {
	log.Println("🚀 Starting Worker...")

	// Load environment variables from .env file
	if err := loadEnvFile(); err != nil {
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

	// gRPC server config (for self-hosted servers)
	grpcServerAddress := os.Getenv("GRPC_SERVER_ADDRESS")
	if grpcServerAddress == "" {
		grpcServerAddress = "grpc.dibbla.com:443" // default Dibbla cloud
	}

	// TLS defaults to true (required for grpc.dibbla.com:443)
	// Set GRPC_USE_TLS=false for local/dev servers without TLS
	grpcUseTLS := os.Getenv("GRPC_USE_TLS") != "false"

	// HTTP server config (localhost only by default - no firewall prompt)
	httpHost := os.Getenv("HTTP_HOST")
	if httpHost == "" {
		httpHost = "127.0.0.1"
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	httpAddr := httpHost + ":" + httpPort

	// Create SDK server
	log.Println("🔧 Creating SDK server...")
	log.Printf("📡 Connecting to gRPC server: %s (TLS: %v)", grpcServerAddress, grpcUseTLS)
	server, err := sdk.New(
		sdk.WithServerName(serverName),
		sdk.WithServerApiToken(serverApiToken),
		sdk.WithGrpcServerAddress(grpcServerAddress),
		sdk.WithGrpcTLS(grpcUseTLS),
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
		log.Printf("🌐 Starting HTTP server on %s", httpAddr)
		if err := http.ListenAndServe(httpAddr, router.Handler()); err != nil {
			log.Printf("HTTP server error: %v", err)
		}
	}()

	// Start the server (blocks forever)
	log.Printf("🎯 Starting worker server '%s'...", serverName)
	if err := server.Start(); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
