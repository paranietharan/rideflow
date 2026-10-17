package health

import (
	"encoding/json"
	"net/http"
	"time"

	"rideflow/internal/platform/mongodb"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	mongo *mongodb.Client
}

func NewHandler(mongo *mongodb.Client) *Handler {
	return &Handler{mongo: mongo}
}
func (h *Handler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/healthz", h.Liveness)
	r.Get("/readyz", h.Readiness)
	return r
}

func (h *Handler) Liveness(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"status":    "ok",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

func (h *Handler) Readiness(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := h.mongo.Ping(ctx); err != nil {
		writeJSON(w, http.StatusServiceUnavailable, map[string]any{
			"status":    "unavailable",
			"database":  "disconnected",
			"error":     err.Error(),
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"status":    "ready",
		"database":  "connected",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
