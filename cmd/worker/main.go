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
	// ADVANCED (rare): capability providers — replace a platform built-in
	// (custom tool_search ranking / memory selection). Most workers don't
	// need this; see internal/providers/providers.go before enabling.
	// "github.com/dibbla-agents/go-worker-starter-template/internal/providers"
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

	// Set environment variables. A variable that is already set wins: .env
	// is for local runs, and must never override what the platform gives a
	// deployed worker (e.g. its own installation's gRPC address).
	for k, v := range envMap {
		if _, set := os.LookupEnv(k); set {
			continue
		}
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

	// Authentication: an API token when you run the worker yourself. Deployed
	// on Dibbla, the platform mounts a workload identity token
	// (DIBBLA_IDENTITY_TOKEN_FILE) and the SDK uses it, so no token is needed.
	serverApiToken := os.Getenv("SERVER_API_TOKEN")
	if serverApiToken == "" && os.Getenv("DIBBLA_IDENTITY_TOKEN_FILE") == "" {
		log.Fatal("❌ SERVER_API_TOKEN environment variable is required")
	}

	// gRPC server address and TLS: set them only to choose a server yourself
	// (a self-hosted one, or localhost:50051 with GRPC_USE_TLS=false). Left
	// unset, the SDK decides: deployed on Dibbla it connects to the
	// installation the worker runs on (DIBBLA_GRPC_SERVER_ADDRESS, set by the
	// platform), run locally it defaults to grpc.dibbla.com:443 with TLS.
	// Passing a default here would override the platform's address.
	grpcServerAddress := os.Getenv("GRPC_SERVER_ADDRESS")
	grpcUseTLS := os.Getenv("GRPC_USE_TLS")

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
	opts := []sdk.Option{sdk.WithServerName(serverName)}
	if serverApiToken != "" {
		opts = append(opts, sdk.WithServerApiToken(serverApiToken))
	}
	if grpcServerAddress != "" {
		log.Printf("📡 Connecting to gRPC server: %s", grpcServerAddress)
		opts = append(opts, sdk.WithGrpcServerAddress(grpcServerAddress))
	} else if platform := os.Getenv("DIBBLA_GRPC_SERVER_ADDRESS"); platform != "" {
		log.Printf("📡 Connecting to this installation's gRPC server: %s", platform)
	} else {
		log.Println("📡 Connecting to the SDK's default gRPC server (grpc.dibbla.com:443)")
	}
	if grpcUseTLS != "" {
		opts = append(opts, sdk.WithGrpcTLS(grpcUseTLS != "false"))
	}
	server, err := sdk.New(opts...)
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

	// ADVANCED (rare): register the example capability providers. Leave this
	// commented unless you are building a custom tool_search or memory
	// implementation — built-ins cover normal workflows.
	// providers.MustRegister(server)

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
