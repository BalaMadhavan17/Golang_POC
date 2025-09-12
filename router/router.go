```
package router

import (
	"go-Beitler-api/handler"

	"github.com/gorilla/mux"
)

func InitRouter(h *handler.MdsHandler) *mux.Router {
	r := mux.NewRouter()
	
	// MDS API endpoints
	r.HandleFunc("/api/mds", h.GetAll).Methods("GET")
	r.HandleFunc("/api/mds", h.Create).Methods("POST")
	r.HandleFunc("/api/mds/{id}", h.GetById).Methods("GET")
	r.HandleFunc("/api/mds/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/api/mds/{id}", h.Delete).Methods("DELETE")

	return r
}
```
