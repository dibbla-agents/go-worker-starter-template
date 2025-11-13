package main

import (
	"log"
	"os"

	"github.com/FatsharkStudiosAB/codex/workflows/workers/go/sdk"
	"github.com/joho/go-dotenv/godotenv"
	
	"worker_starter_template/internal/state"
	workerfunctions "worker_starter_template/internal/worker_functions"
	
	// TODO: Import your worker functions here
	// Example:
	// myfunction "worker_starter_template/internal/worker_functions/my_function"
)

func main() {
	log.Println("🚀 Starting Worker...")

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  Warning: .env file not found, using system environment variables")
	}

	// TODO: Uncomment if you need database/shared resources
	// log.Println("📊 Initializing async global state...")
	// ags, err := state.NewAsyncGlobalState()
	// if err != nil {
	// 	log.Fatalf("❌ Failed to initialize async global state: %v", err)
	// }
	// defer func() {
	// 	if err := ags.Close(); err != nil {
	// 		log.Printf("⚠️  Warning: Failed to close async global state: %v", err)
	// 	}
	// }()
	// log.Println("✅ Async global state initialized")

	var ags *state.AsyncGlobalState = nil

	// Create SDK server
	log.Println("🔧 Creating SDK server...")

	serverName := os.Getenv("SERVER_NAME")
	if serverName == "" {
		serverName = "worker-starter"
	}

	grpcServerAddress := os.Getenv("GRPC_SERVER_ADDRESS")
	if grpcServerAddress == "" {
		grpcServerAddress = "localhost:50051"
	}

	serverApiToken := os.Getenv("SERVER_API_TOKEN")

	server, err := sdk.New(
		sdk.WithServerName(serverName),
		sdk.WithGrpcServerAddress(grpcServerAddress),
		sdk.WithServerApiToken(serverApiToken),
	)
	if err != nil {
		log.Fatalf("❌ Failed to create SDK server: %v", err)
	}

	// Register worker functions
	log.Println("📝 Registering worker functions...")
	registry := workerfunctions.NewRegistry()

	// TODO: Register your worker functions here
	// Example:
	// registry.Register(myfunction.NewMyFunction())

	if err := registry.RegisterAll(server, ags); err != nil {
		log.Fatalf("❌ Failed to register functions: %v", err)
	}

	// Start the server (blocks forever)
	log.Printf("🎯 Starting worker server '%s' connecting to workflow server at %s...", serverName, grpcServerAddress)
	if err := server.Start(); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
