# Getting Started

Quick guide to set up and run this worker template.

---

## Prerequisites

- Go 1.23+
- Go SDK repository cloned (required)
- PostgreSQL (optional, for database features)

---

## Project Structure

```
worker_starter_template/
├── cmd/
│   └── worker/              # Application entry point (main.go)
├── internal/
│   ├── config/              # Configuration management
│   ├── embeddings/          # OpenAI embeddings client
│   ├── external_api/        # External API integrations
│   ├── frontend/            # Embedded frontend (Go embed)
│   │   ├── dist/            # Built frontend assets
│   │   └── embed.go         # Go embed configuration
│   ├── jobs/                # Job definitions (multi-step workflows)
│   │   └── tasks/           # Reusable task definitions
│   ├── models/              # Database models (GORM)
│   ├── state/               # Global state management
│   └── worker_functions/    # Your worker function implementations
│       ├── example_function/
│       └── registry.go      # Function registration
├── frontend/                # Frontend development project
│   ├── src/
│   │   ├── components/      # UI components
│   │   ├── pages/           # Page views
│   │   ├── assets/          # Static assets
│   │   └── ...
│   ├── index.html           # HTML entry point
│   ├── vite.config.js       # Vite configuration
│   └── dist/                # Build output (copied to internal/frontend/dist)
├── docs/                    # Documentation and guides
├── docker-compose.yml       # Docker Compose configuration
├── Dockerfile.worker        # Worker Docker image
├── go.mod                   # Go module dependencies
└── env.example              # Environment variable template
```

**Key Directories:**
- **`cmd/worker/`** - Start here: main application entry point
- **`internal/worker_functions/`** - Add your custom worker functions here
- **`internal/jobs/`** - Define multi-step workflows
- **`internal/models/`** - Database models (if using PostgreSQL)
- **`frontend/`** - Frontend development project (Vite + components + pages)
- **`internal/frontend/`** - Embedded frontend build (served by Go worker)
- **`docs/how_to/`** - Detailed guides for common tasks

**Key Files:**
- **`docker-compose.yml`** - Multi-container deployment setup
- **`Dockerfile.worker`** - Worker containerization

---

## Setup

### 1. Customize Your Project

After creating from the template, update the module name in `go.mod`:

```go
module github.com/YOUR_USERNAME/YOUR_PROJECT_NAME

go 1.23
```

**If you use an external SDK:**

Clone your SDK and add replace directives in `go.mod`:

```bash
git clone https://github.com/your-org/your-sdk.git
```

In `go.mod`:
```go
replace github.com/your-org/your-sdk/pkg => /path/to/your-sdk/pkg
```

### 2. Build and Run

**This is a Go project** - download dependencies and build before running:

```bash
# Download dependencies (required)
go mod download

# Build the binary (required)
go build -o worker ./cmd/worker

# Run
./worker
```

### 3. Optional: Configure Environment

Copy and configure the environment file:

```bash
# Copy the example file
cp env.example .env

# Edit .env with your actual values
# See env.example for all available options
```

Key common environment variables:
- `GRPC_SERVER_ADDRESS` - gRPC server address
- `SERVER_NAME` - Worker identifier
- `SERVER_API_TOKEN` - API authentication token
- `OPENAI_API_KEY` - OpenAI API key (if using AI features)
- `DATABASE_URL` - PostgreSQL connection string
- `LOG_LEVEL` - Logging level (info, debug, warn, error)

---

## Add Your First Worker Function

1. **Create function directory:**
   ```bash
   mkdir internal/worker_functions/my_function
   cp internal/worker_functions/example_function/functions.go internal/worker_functions/my_function/
   ```

2. **Edit the function** in `my_function/functions.go`

3. **Register it** in `internal/worker_functions/registry.go`:
   ```go
   my_function.RegisterMyFunction(w, ags)
   ```

4. **Rebuild and run:**
   ```bash
   go build -o worker ./cmd/worker && ./worker
   ```

---

## Enable Database (Optional)

1. Uncomment `DB` field in `internal/state/async_global_state.go`
2. Define models in `internal/models/models.go`
3. Uncomment migration in `cmd/worker/main.go`
4. Rebuild: `go build -o worker ./cmd/worker`

---

## Common Commands

```bash
# Build and run
go mod download
go build -o worker ./cmd/worker
./worker

# Development
go fmt ./...        # Format code
go test ./...       # Run tests
go vet ./...        # Check for issues
```

---

## Troubleshooting

| Issue | Solution |
|-------|----------|
| "cannot find module" | Run `go mod download`, check SDK paths in `go.mod` |
| "command not found: worker" | Build first: `go build -o worker ./cmd/worker` |
| Worker doesn't respond | Check `GRPC_SERVER_ADDRESS`, verify function is registered |
| Database errors | Check PostgreSQL is running, verify `DB` is uncommented |

---

## Next Steps

- [Add Worker Functions](docs/how_to/add_worker_function.md) - Detailed guide
- [Create Jobs](docs/how_to/create_jobs.md) - Multi-step workflows  
- [Create Tasks](docs/how_to/create_tasks.md) - Reusable operations
- [Docker Deployment](docs/how_to/docker_deployment.md) - Production setup
- [Documentation Index](docs/README.md) - All guides
