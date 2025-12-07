package health

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Handler struct {
	db *sql.DB
}

func New(db *sql.DB) (*Handler, error) {
	if db == nil {
		return nil, fmt.Errorf("health service requires *sql.DB, it is nil")
	}

	return &Handler{
		db: db,
	}, nil
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"status":    "healthy",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// Make it also Kubernetes ready
func (h *Handler) Ready(w http.ResponseWriter, r *http.Request) {
	if err := h.db.Ping(); err != nil {
		http.Error(w, `{"status":"not ready","reason":"database"}`, http.StatusServiceUnavailable)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"status":    "ready",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}
