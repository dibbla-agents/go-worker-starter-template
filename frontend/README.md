# Optional Frontend Module

This directory contains an optional Vite + React frontend that can be embedded in your Go worker.

## Quick Start

### 1. Copy All Frontend Files

```bash
# From your project root
cp -r _optional/frontend ./frontend
cp -r _optional/internal_frontend ./internal/frontend
cp -r _optional/http_handlers ./internal/http_handlers
```

### 2. Install Dependencies

```bash
cd frontend
npm install
```

### 3. Development Mode

**Terminal 1** - Start the Go worker (with HTTP server enabled):
```bash
go run ./cmd/worker
```

**Terminal 2** - Start frontend dev server:
```bash
cd frontend
npm run dev
```

Opens at `http://localhost:5173` with hot reload.
API calls to `/api/*` are automatically proxied to the Go worker.

### 4. Build for Production

```bash
cd frontend
npm run build
```

This outputs to `internal/frontend/dist/` which is embedded in the Go binary.

### 5. Enable in Go Worker

Update `cmd/worker/main.go`:

```go
import (
	"net/http"
	"github.com/your-org/your-project/internal/frontend"
	"github.com/your-org/your-project/internal/http_handlers/greeting"
)

func main() {
	// ... existing setup ...

	// Create HTTP router
	router := frontend.NewRouter()

	// Register API handlers
	greeting.Register(router.Mux())

	// Start HTTP server (serves API + frontend)
	go func() {
		log.Println("🌐 HTTP server on :8080")
		http.ListenAndServe(":8080", router.Handler())
	}()

	// ... start gRPC server ...
}
```

## Project Structure

```
frontend/                      # Frontend source (copy to ./frontend)
├── src/
│   ├── main.tsx              # Entry point
│   ├── App.tsx               # Main component (calls /api/greeting)
│   └── index.css             # Styles
├── index.html
├── package.json
└── vite.config.ts            # Build config + dev proxy

internal_frontend/             # Go embedding (copy to ./internal/frontend)
├── router.go                 # HTTP router + static file server
└── dist/                     # Built assets (after npm run build)

http_handlers/                 # API endpoints (copy to ./internal/http_handlers)
└── greeting/
    └── handler.go            # POST /api/greeting
```

## How It Works

1. **Development**: Vite dev server (`npm run dev`) proxies `/api/*` to Go worker at `:8080`
2. **Production**: Go binary serves both API routes and static frontend from embedded files
3. **API Pattern**: Create handlers in `http_handlers/`, register with router in `main.go`

## Customization

- **Theme**: Edit `src/App.tsx` and `src/index.css`
- **Components**: Add to `src/components/`
- **API Endpoints**: Add handlers to `internal/http_handlers/`

## Adding a New API Endpoint

1. Create `internal/http_handlers/your_handler/handler.go`:

```go
package yourhandler

import (
	"encoding/json"
	"net/http"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/your_endpoint", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Your logic here
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
```

2. Register in `main.go`:

```go
yourhandler.Register(router)
```

3. Call from frontend:

```tsx
const response = await fetch('/api/your_endpoint', {
	method: 'POST',
	headers: { 'Content-Type': 'application/json' },
	body: JSON.stringify({ /* your data */ }),
})
```

## Notes

- Frontend is served from the same binary - no separate server needed
- Assets are embedded at compile time using Go's `embed` package
- Run `npm run build` before `go build` to include latest frontend
- SPA routing is handled automatically (non-API routes serve `index.html`)
