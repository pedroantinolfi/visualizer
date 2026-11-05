package main

import (
	"log"
	"net/http"
	"time"
)

// loggingMiddleware intercepts incoming requests to add headers and log performance.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Disable caching during development so the browser always loads the latest main.wasm
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		// Pass the request to the next handler (the file server or API endpoints)
		next.ServeHTTP(w, r)

		// Log the HTTP method, requested path, and execution time
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	// ServeMux is the standard HTTP request multiplexer (router) in Go.
	mux := http.NewServeMux()

	// 1. Static files route (HTML, CSS, JS, WASM)
	fileServer := http.FileServer(http.Dir("./web"))
	mux.Handle("/", fileServer)

	// 2. Future expansion area: register new API routes here
	// mux.HandleFunc("GET /api/status", statusHandler)
	// mux.HandleFunc("POST /api/save", saveTreeHandler)

	// Wrap the entire router with our middleware
	handler := loggingMiddleware(mux)

	// Configure the HTTP server with strict timeouts to prevent resource leaks
	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Println("Server listening on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server crashed:", err)
	}
}
