```
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

type MdsEntry struct {
	Name          string    `json:"name"`
	Comments      string    `json:"comments"`
	EffectiveFrom string    `json:"effectiveFrom"`
	EffectiveTo   string    `json:"effectiveTo"`
	IsPPAgreed    bool      `json:"isPPAgreed"`
}

type MdsListEntry struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Comments      string    `json:"comments"`
	EffectiveFrom string    `json:"effectiveFrom"`
	EffectiveTo   string    `json:"effectiveTo"`
	IsPPAgreed    bool      `json:"isPPAgreed"`
	CreatedAt     string    `json:"createdAt"`
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
	var entry MdsEntry
	
	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	effectiveFrom, err := time.Parse("2006-01-02", entry.EffectiveFrom)
	if err != nil {
		http.Error(w, "Invalid effective from date format", http.StatusBadRequest)
		return
	}
	
	effectiveTo, err := time.Parse("2006-01-02", entry.EffectiveTo)
	if err != nil {
		http.Error(w, "Invalid effective to date format", http.StatusBadRequest)
		return
	}
	
	id, err := h.service.Create(entry.Name, entry.Comments, effectiveFrom, effectiveTo, entry.IsPPAgreed)
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
	
	responseEntries := []MdsListEntry{}
	for _, entry := range entries {
		responseEntries = append(responseEntries, MdsListEntry{
			ID:            entry.ID,
			Name:          entry.Name,
			Comments:      entry.Comments,
			EffectiveFrom: entry.EffectiveFrom.Format("2006-01-02"),
			EffectiveTo:   entry.EffectiveTo.Format("2006-01-02"),
			IsPPAgreed:    entry.IsPPAgreed,
			CreatedAt:     entry.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	
	json.NewEncoder(w).Encode(responseEntries)
}
```
