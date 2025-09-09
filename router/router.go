package router

import (
	"go-Beitler-api/handler"

	"github.com/gorilla/mux"
)

func InitRouter(h *handler.MdsHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/mds", h.GetAll).Methods("GET")
	r.HandleFunc("/api/mds/{id}", h.GetByID).Methods("GET")
	r.HandleFunc("/api/mds", h.Create).Methods("POST")
	r.HandleFunc("/api/mds/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/api/mds/{id}", h.Delete).Methods("DELETE")
	return r
}
