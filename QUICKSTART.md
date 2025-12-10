# Starter Template - Quick Start

Get a structured worker project running in 5 minutes.

## Prerequisites

- [Go 1.23+](https://go.dev/dl/) installed
- Dibbla account with API token

## Step 1: Create Your Project

```bash
# Install gonew (one-time)
go install golang.org/x/tools/cmd/gonew@latest

# Create your project
gonew github.com/dibbla-agents/go-worker-starter-template@latest github.com/your-org/my-worker
cd my-worker
```

## Step 2: Configure Environment

Create a `.env` file:

```env
SERVER_API_TOKEN=ak_your_token_here
SERVER_NAME=my-worker
```

> Get your token from [app.dibbla.com/dashboard](https://app.dibbla.com/dashboard) → Settings → API Keys

> **Docker users:** Add `HTTP_HOST=0.0.0.0` to your `.env` file.

## Step 3: Run the Worker

```bash
go mod tidy
go run ./cmd/worker
```

### Stopping the Worker

Press **`Ctrl + C`** to stop the worker.

> **Note:** Use a regular terminal (PowerShell, Terminal, CMD). Some IDE integrated terminals may not handle `Ctrl+C` properly.

## Success!

You should see:

```
🚀 Starting Worker...
🔧 Creating SDK server...
📝 Registering worker functions...
   ✅ Registered: greeting
🎯 Starting worker server 'my-worker'...
✅ gRPC client successfully connected to workflow server
```

The template includes a `greeting` function out of the box.

---

## Step 4: Test with Frontend (Optional)

Open a **second terminal** (keep the worker running in the first):

```bash
cd frontend
npm install
npm run dev
```

Open **http://localhost:5173** and test the greeting function directly.

> The frontend proxies `/api/*` to the worker on port 8080.

To stop the frontend: Press **`Ctrl + C`** in the frontend terminal.

---

## Add Your Own Function

Create `internal/worker_functions/hello/hello.go`:

```go
package hello

import (
    "fmt"
    sdk "github.com/dibbla-agents/sdk-go"
)

type Input struct {
    Name string `json:"name"`
}

type Output struct {
    Message string `json:"message"`
}

func Register(server *sdk.Server) {
    fn := sdk.NewSimpleFunction[Input, Output](
        "hello", "1.0.0", "Say hello",
    ).WithHandler(func(input Input) (Output, error) {
        return Output{Message: fmt.Sprintf("Hello, %s!", input.Name)}, nil
    })
    server.RegisterFunction(fn)
}
```

Register it in `cmd/worker/main.go`:

```go
import "github.com/your-org/my-worker/internal/worker_functions/hello"

// In main(), after other Register calls:
hello.Register(server)
log.Println("   ✅ Registered: hello")
```

Stop the worker (`Ctrl + C`) and run again:

```bash
go run ./cmd/worker
```

---

## Project Structure

```
my-worker/
├── cmd/worker/main.go           # Entry point
├── internal/
│   └── worker_functions/        # Your functions go here
│       └── greeting/            # Example function
├── .env                         # Your config (create this)
└── env.example                  # Config template
```

---

## Next Steps

- [Add Worker Functions](docs/how_to/add_worker_function.md)
- [Create Jobs](docs/how_to/create_jobs.md) - Multi-step workflows
- [Docker Deployment](docs/how_to/docker_deployment.md)

---

## Troubleshooting

| Error | Solution |
|-------|----------|
| `SERVER_API_TOKEN environment variable is required` | Create `.env` file with your token |
| `invalid or expired API token` | Generate new token from dashboard |
| `cannot find module` | Run `go mod tidy` |
| Port 8080 already in use | Stop the other worker first, or check Task Manager for `worker.exe` |

---

## Advanced: Registry Pattern

The examples above use the **simple pattern** (one `Register()` function per file). This is recommended for most use cases.

For advanced scenarios where functions need **shared state** (database connections, caches, etc.), see the registry pattern in `internal/worker_functions/registry.go`.

