package main

import (
	"go-Beitler-api/config"
	"log"
	"os"
)

func main() {
	// Connect DB
	config.ConnectDB()

	// Dependency Injection

	// Init Router

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("🚀 Server running on port", port)
}
