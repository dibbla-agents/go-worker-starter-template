# How to Add a Frontend (Optional)

This guide shows how to add an embedded web UI to your worker using Vite + React.

## Overview

The frontend is **optional** and lives in the `_optional/` directory. When enabled:
- Frontend assets are embedded in the Go binary
- HTTP API endpoints for frontend communication
- No separate web server needed
- Single deployable artifact

---

## Quick Setup

### Step 1: Copy Frontend Files

```bash
# From your project root
cp -r _optional/frontend ./frontend
cp -r _optional/internal_frontend ./internal/frontend
cp -r _optional/http_handlers ./internal/http_handlers
```

### Step 2: Install Dependencies

```bash
cd frontend
npm install
```

### Step 3: Development Mode

Start the Go worker first (with HTTP server), then run the frontend dev server:

```bash
# Terminal 1: Run the worker
go run ./cmd/worker

# Terminal 2: Run frontend with hot reload
cd frontend
npm run dev
```

The frontend opens at `http://localhost:5173` with hot reload.
API calls are proxied to the Go worker at `:8080`.

### Step 4: Build for Production

```bash
cd frontend
npm run build
```

Outputs to `internal/frontend/dist/` for Go embedding.

---

## Enable in Go Worker

Update `cmd/worker/main.go`:

```go
package main

import (
	"log"
	"net/http"
	"os"

	sdk "github.com/dibbla-agents/sdk-go"
	"github.com/joho/godotenv"

	// Import frontend router and handlers
	"github.com/dibbla-agents/go-worker-starter-template/internal/frontend"
	"github.com/dibbla-agents/go-worker-starter-template/internal/http_handlers/greeting"
)

func main() {
	// ... existing setup ...

	// Create HTTP router (serves API + embedded frontend)
	router := frontend.NewRouter()

	// Register HTTP API handlers on the router's mux
	greeting.Register(router.Mux())
	log.Println("   ✅ HTTP: POST /api/greeting")

	// Start HTTP server on port 8080
	go func() {
		log.Println("🌐 Starting HTTP server on :8080")
		if err := http.ListenAndServe(":8080", router.Handler()); err != nil {
			log.Printf("HTTP server error: %v", err)
		}
	}()

	// ... start gRPC server ...
}
```

---

## Project Structure After Setup

```
your-project/
├── frontend/                    # Frontend source (Vite + React)
│   ├── src/
│   │   ├── main.tsx
│   │   ├── App.tsx
│   │   └── index.css
│   ├── index.html
│   ├── package.json
│   └── vite.config.ts
├── internal/
│   ├── frontend/               # Embedded assets + router (Go)
│   │   ├── router.go           # HTTP router with frontend fallback
│   │   └── dist/               # Built frontend (after npm run build)
│   └── http_handlers/          # HTTP API endpoints
│       └── greeting/
│           └── handler.go
└── cmd/worker/
    └── main.go
```

---

## Adding HTTP Endpoints

### Create a New Endpoint

1. Create `internal/http_handlers/your_handler/handler.go`:

```go
package yourhandler

import (
	"encoding/json"
	"net/http"
)

type Input struct {
	Field string `json:"field"`
}

type Output struct {
	Result string `json:"result"`
}

func Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/your_endpoint", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input Input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Your logic here
	output := Output{Result: "processed: " + input.Field}

	json.NewEncoder(w).Encode(output)
}
```

2. Register in `main.go`:

```go
import "github.com/your-org/your-project/internal/http_handlers/yourhandler"

// In main():
yourhandler.Register(router)
```

### HTTP Method Patterns (Go 1.22+)

Use method-specific patterns:
- `"GET /api/items"` - GET only
- `"POST /api/items"` - POST only
- `"PUT /api/items/{id}"` - with path parameter
- `"DELETE /api/items/{id}"` - with path parameter

Access path parameters: `r.PathValue("id")`

---

## Development Workflow

### Frontend Development

1. Run the Go worker with HTTP server
2. Run `npm run dev` in `frontend/`
3. Edit components in `src/`
4. API calls are proxied to Go automatically

### Full Stack Testing

1. Build frontend: `cd frontend && npm run build`
2. Build worker: `go build -o worker ./cmd/worker`
3. Run worker: `./worker`
4. Open `http://localhost:8080`

---

## Calling API from Frontend

The example `App.tsx` shows how to call API endpoints:

```tsx
const callGreeting = async () => {
  const response = await fetch('/api/greeting', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ name: 'World' }),
  })
  const data = await response.json()
  console.log(data.message) // "Hello, World!"
}
```

In development, Vite proxies `/api/*` to `http://localhost:8080`.
In production, everything is served from the same binary.

---

## Customization

### Adding Components

Create new components in `frontend/src/components/`:

```tsx
// frontend/src/components/StatusCard.tsx
export function StatusCard({ title, value }: { title: string; value: string }) {
  return (
    <div className="card">
      <h3>{title}</h3>
      <p>{value}</p>
    </div>
  )
}
```

### Styling

Edit `frontend/src/index.css` for global styles. The default theme uses:
- Dark background with blue accents
- Clean, minimal UI
- Responsive layout

---

## Build Notes

- **Always rebuild frontend before Go build** if frontend changed
- Frontend is embedded at compile time - changes require recompilation
- Built assets are ~200KB gzipped for the default template

---

## Troubleshooting

| Issue | Solution |
|-------|----------|
| `embed: no matching files` | Run `npm run build` first |
| Frontend not updating | Rebuild both frontend and Go binary |
| 404 on routes | SPA routing is handled automatically |
| CORS errors in dev | Vite proxy handles this - ensure worker is running |
| API returns 404 | Check handler is registered and method matches |
