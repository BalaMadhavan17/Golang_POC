package main

import (
	"go-Beitler-api/config"
	"go-Beitler-api/handler"
	"go-Beitler-api/repository"
	"go-Beitler-api/router"
	"go-Beitler-api/service"
	"log"
	"net/http"
	"os"
)

func main() {
	// Connect DB
	config.ConnectDB()

	// Dependency Injection
	repo := repository.NewMdsRepository(config.DB)
	svc := service.NewMdsService(repo)
	h := handler.NewMdsHandler(svc)

	// Init Router
	r := router.InitRouter(h)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("🚀 Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
