package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/vijayraghavareddy/tax-calculation/handlers"
)

func main() {
	router := mux.NewRouter()

	// API routes
	router.HandleFunc("/api/v1/calculate-tax", corsMiddleware(handlers.CalculateTax)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/health", corsMiddleware(handlers.HealthCheck)).Methods("GET", "OPTIONS")

	// Serve static files
	staticDir := "./static"
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	// Serve index.html for root path
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, staticDir+"/index.html")
	}).Methods("GET")

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Printf("Web UI available at http://localhost:%s", port)
	log.Printf("API endpoints at http://localhost:%s/api/v1/", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}

// corsMiddleware adds CORS headers to responses
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
