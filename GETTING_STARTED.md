# Getting Started

Quick guide to set up and run this worker template.

---

## Prerequisites

- Go 1.23+
- PostgreSQL (optional, for database features)

---

## Project Structure

```
your-project/
├── cmd/
│   └── worker/              # Application entry point (main.go)
├── internal/
│   ├── config/              # Configuration management
│   ├── embeddings/          # OpenAI embeddings client (optional)
│   ├── jobs/                # Job definitions (multi-step workflows)
│   │   └── tasks/           # Reusable task definitions
│   ├── models/              # Database models (GORM)
│   ├── state/               # Global state management
│   └── worker_functions/    # Your worker function implementations
│       ├── example_function/
│       └── registry.go      # Function registration
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
- **`docs/how_to/`** - Detailed guides for common tasks

**Key Files:**
- **`docker-compose.yml`** - Multi-container deployment setup
- **`Dockerfile.worker`** - Worker containerization

---

## Setup


### 1. Verify Module Name

Open your `go.mod` file and check the first line.

*   ✅ **If you used `gonew`:** It should already match your project (e.g., `module github.com/your-org/your-project`). **Skip to Step 2.**
*   📝 **If you used `git clone`:** You must manually update it now:
    ```go
    // Change this line to your new repository path
    module github.com/YOUR_USERNAME/YOUR_PROJECT_NAME
    ```

### 2. Build and Run

**Linux/Mac:**
```bash
# Download dependencies
go mod download

# Build the binary
go build -o worker ./cmd/worker

# Run
./worker
```

**Windows (PowerShell):**
```powershell
# Download dependencies
go mod download

# Build the binary
go build -o worker.exe .\cmd\worker

# Run
.\worker.exe
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
   
   **Linux/Mac:**
   ```bash
   go build -o worker ./cmd/worker && ./worker
   ```
   
   **Windows (PowerShell):**
   ```powershell
   go build -o worker.exe .\cmd\worker; .\worker.exe
   ```

---

## Enable Database (Optional)

1. Uncomment `DB` field in `internal/state/async_global_state.go`
2. Define models in `internal/models/models.go`
3. Uncomment migration in `cmd/worker/main.go`
4. Rebuild: `go build -o worker ./cmd/worker`

---

## Common Commands

**Linux/Mac:**
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

**Windows (PowerShell):**
```powershell
# Build and run
go mod download
go build -o worker.exe .\cmd\worker
.\worker.exe

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
