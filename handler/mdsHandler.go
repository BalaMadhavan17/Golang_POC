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

type createMdsEntryRequest struct {
	Name          string `json:"name"`
	Comments      string `json:"comments"`
	EffectiveFrom string `json:"effectiveFrom"`
	EffectiveTo   string `json:"effectiveTo"`
	IsPPAgreed    bool   `json:"isPPAgreed"`
	DocumentPath  string `json:"documentPath"`
}

type createMdsEntryResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
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

func (h *MdsHandler) CreateEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req createMdsEntryRequest
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

	// Create entry object
	entry := &repository.MdsEntry{
		Name:          req.Name,
		Comments:      req.Comments,
		EffectiveFrom: effectiveFrom,
		EffectiveTo:   effectiveTo,
		IsPPAgreed:    req.IsPPAgreed,
		DocumentPath:  req.DocumentPath,
	}

	// Call service to create entry
	id, err := h.service.CreateEntry(entry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createMdsEntryResponse{
		ID:      id,
		Message: "MDS entry has been saved successfully.",
	})
}
```
