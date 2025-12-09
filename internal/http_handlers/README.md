# HTTP Handlers (Optional)

This directory contains HTTP API handlers that can be used with the optional frontend.

## Structure

```
http_handlers/
├── greeting/           # Example handler
│   └── handler.go
└── your_handler/       # Add your handlers here
    └── handler.go
```

## Usage

### 1. Copy to Your Project

```bash
# From your project root
cp -r _optional/http_handlers ./internal/http_handlers
```

### 2. Register Handlers in main.go

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
    
    // Register API handlers on the router's mux
    greeting.Register(router.Mux())
    
    // Start HTTP server (serves API + frontend)
    go func() {
        log.Println("🌐 HTTP server on :8080")
        http.ListenAndServe(":8080", router.Handler())
    }()

    // ... start gRPC server ...
}
```

## Creating a New Handler

1. Create a new directory: `internal/http_handlers/your_handler/`
2. Create `handler.go`:

```go
package yourhandler

import (
    "encoding/json"
    "net/http"
)

type Input struct {
    // Your input fields
}

type Output struct {
    // Your output fields
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
    output := Output{}
    
    json.NewEncoder(w).Encode(output)
}
```

3. Register in `main.go`

## HTTP Method Patterns (Go 1.22+)

Use method-specific patterns:
- `"GET /api/items"` - GET requests only
- `"POST /api/items"` - POST requests only  
- `"PUT /api/items/{id}"` - PUT with path parameter
- `"DELETE /api/items/{id}"` - DELETE with path parameter

Access path parameters with `r.PathValue("id")`.

