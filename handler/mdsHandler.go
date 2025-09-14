```go
package handler

import (
	"encoding/json"
	"go-Beitler-api/repository"
	"go-Beitler-api/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type MdsHandler struct {
	service service.MdsService
}

type createMdsRequest struct {
	Name          string `json:"name"`
	Comments      string `json:"comments"`
	EffectiveFrom string `json:"effectiveFrom"`
	EffectiveTo   string `json:"effectiveTo"`
	IsPPAgreed    bool   `json:"isPPAgreed"`
	DocumentPath  string `json:"documentPath"`
}

func NewMdsHandler(service service.MdsService) *MdsHandler {
	return &MdsHandler{service}
}

func (h *MdsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Deleted successfully"})
}

func (h *MdsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createMdsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	effectiveFrom, err := time.Parse("2006-01-02", req.EffectiveFrom)
	if err != nil {
		http.Error(w, "Invalid effective from date format", http.StatusBadRequest)
		return
	}

	effectiveTo, err := time.Parse("2006-01-02", req.EffectiveTo)
	if err != nil {
		http.Error(w, "Invalid effective to date format", http.StatusBadRequest)
		return
	}

	entry := repository.MdsEntry{
		Name:          req.Name,
		Comments:      req.Comments,
		EffectiveFrom: effectiveFrom,
		EffectiveTo:   effectiveTo,
		IsPPAgreed:    req.IsPPAgreed,
		DocumentPath:  req.DocumentPath,
	}

	id, err := h.service.Create(entry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "MDS entry has been saved successfully.",
		"id":      id,
	})
}

func (h *MdsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	entries, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}
```
