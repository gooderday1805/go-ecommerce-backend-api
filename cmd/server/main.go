package main

import (
	"log"
	"github.com/gin-contrib/cors"
	"github.com/gooderday1805/go-ecommerce-backend-api/internal/routers"
)

func main() {
	// Create router using our router configuration
	r := routers.NewRouters()

	// Add CORS middleware to avoid errors when calling API from browser/Postman
	r.Use(cors.Default())

	// Run server on port 8080, listening on all IP addresses
	port := ":8080"
	log.Printf("Server is running on http://localhost%s\n", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
