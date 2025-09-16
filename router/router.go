package router

import (
	"go-Beitler-api/config"
	mdsHandlerPkg "go-Beitler-api/handler"
	"go-Beitler-api/repository"
	mdsServicePkg "go-Beitler-api/service"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func SetupRoutes() http.Handler {
	r := mux.NewRouter()

	// Get database connection
	db := config.GetDB()

	// Initialize repository
	mdsRepo := repository.NewMdsRepository(db)

	// Initialize service
	mdsService := mdsServicePkg.NewMdsService(mdsRepo)

	// Initialize handler
	mdsHandler := mdsHandlerPkg.NewMdsHandler(mdsService)

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	// Health check route
	api.HandleFunc("/health", healthCheckHandler).Methods("GET")
	// MDS routes
	api.HandleFunc("/mds", mdsHandler.Create).Methods("POST")
	api.HandleFunc("/mds", mdsHandler.GetAll).Methods("GET")
	api.HandleFunc("/mds/{id:[0-9]+}", mdsHandler.Delete).Methods("DELETE")

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	return c.Handler(r)
}
