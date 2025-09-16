package main

import (
	"go-Beitler-api/router"
	"log"
	"net/http"
)

func main() {
	handler := router.SetupRoutes()
	
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Could not start server: %s", err.Error())
	}
}
