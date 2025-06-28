package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting TokenOne Backend Server...")
	"github.com/tokenone/tokenone-backend/internal/api"
)

func main() {
	fmt.Println("Starting TokenOne Backend Server...")

	mux := http.NewServeMux()

	// Health check handler
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	// Basic handler
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "TokenOne Backend is running!")
	})

	// Configure API routes
	api.ConfigureRoutes(mux)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
