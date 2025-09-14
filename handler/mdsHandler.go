```go
package handler

import (
	"encoding/json"
	"go-Beitler-api/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type MdsEntryRequest struct {
	Name          string    `json:"name"`
	Comments      string    `json:"comments"`
	EffectiveFrom string    `json:"effectiveFrom"`
	EffectiveTo   string    `json:"effectiveTo"`
	IsPPAgreed    bool      `json:"isPPAgreed"`
	DocumentPath  string    `json:"documentPath"`
}

type MdsHandler struct {
	service service.MdsService
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
	w.Header().Set("Content-Type", "application/json")

	var req MdsEntryRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Parse dates
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

	id, err := h.service.Create(req.Name, req.Comments, effectiveFrom, effectiveTo, req.IsPPAgreed, req.DocumentPath)
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
	w.Header().Set("Content-Type", "application/json")

	entries, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(entries)
}
```
