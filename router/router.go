package router

import (
	"database/sql"
	"go-Beitler-api/handler/mdsHandler"
	"go-Beitler-api/repository/mdsRepository"
	"go-Beitler-api/service/mdsService"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// Create repository, service and handler
	repo := mdsRepository.NewMdsRepository(db)
	service := mdsService.NewMdsService(repo)
	handler := mdsHandler.NewMdsHandler(service)

	// API routes
	apiRouter := router.PathPrefix("/api").Subrouter()
	mdsRouter := apiRouter.PathPrefix("/mds").Subrouter()

	// MDS endpoints
	mdsRouter.HandleFunc("", handler.Create).Methods("POST")
	mdsRouter.HandleFunc("", handler.GetAll).Methods("GET")
	mdsRouter.HandleFunc("/{id:[0-9]+}", handler.Delete).Methods("DELETE")

	// Serve static files for frontend
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	return router
}
