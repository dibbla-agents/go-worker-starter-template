// Package frontend provides embedded frontend assets and HTTP routing for the worker.
//
// To use this package:
// 1. Build the frontend: cd frontend && npm run build
// 2. Import this package in your main.go
// 3. Create a router and register your API handlers
// 4. Serve with: http.ListenAndServe(":8080", router.Handler())
package frontend

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"
)

//go:embed dist/*
var distFS embed.FS

// Router provides HTTP routing with embedded frontend support.
// API routes take priority, with frontend serving as fallback.
type Router struct {
	mux *http.ServeMux
}

// NewRouter creates a new HTTP router.
// Register your API handlers, then call Handler() to get the final http.Handler.
func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

// HandleFunc registers a handler for the given pattern.
// Use Go 1.22+ patterns like "POST /api/greeting" for method-specific routes.
func (r *Router) HandleFunc(pattern string, handler http.HandlerFunc) {
	r.mux.HandleFunc(pattern, handler)
}

// Handle registers an http.Handler for the given pattern.
func (r *Router) Handle(pattern string, handler http.Handler) {
	r.mux.Handle(pattern, handler)
}

// Mux returns the underlying *http.ServeMux for registering handlers.
// Use this to pass to handler registration functions that expect *http.ServeMux.
// Example: greeting.Register(router.Mux())
func (r *Router) Mux() *http.ServeMux {
	return r.mux
}

// Handler returns the final http.Handler that serves both API routes and the frontend.
// API routes registered via HandleFunc/Handle take priority.
// All other requests fall back to the embedded frontend (SPA-friendly).
func (r *Router) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Try API routes first (anything starting with /api)
		if strings.HasPrefix(req.URL.Path, "/api") {
			r.mux.ServeHTTP(w, req)
			return
		}

		// For non-API routes, try to serve static files from embedded frontend
		dist, err := fs.Sub(distFS, "dist")
		if err != nil {
			http.Error(w, "Frontend not available", http.StatusInternalServerError)
			return
		}

		// Try to serve the exact file
		fileServer := http.FileServer(http.FS(dist))

		// Check if file exists
		path := strings.TrimPrefix(req.URL.Path, "/")
		if path == "" {
			path = "index.html"
		}

		if _, err := fs.Stat(dist, path); err == nil {
			fileServer.ServeHTTP(w, req)
			return
		}

		// For SPA routing: serve index.html for any non-existent path
		req.URL.Path = "/"
		fileServer.ServeHTTP(w, req)
	})
}

// Handler returns an http.Handler that serves the embedded frontend files.
// Use this for simple static file serving without API routes.
// For API + frontend, use NewRouter() instead.
func Handler() http.Handler {
	dist, err := fs.Sub(distFS, "dist")
	if err != nil {
		panic("failed to create sub filesystem: " + err.Error())
	}
	return http.FileServer(http.FS(dist))
}

// FS returns the embedded filesystem for custom handling.
func FS() fs.FS {
	dist, _ := fs.Sub(distFS, "dist")
	return dist
}

