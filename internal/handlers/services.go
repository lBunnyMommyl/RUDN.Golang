package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"portfolio-go/internal/storage"
)

func ServicesHandler(store *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}

		writeJSON(w, http.StatusOK, store.GetServices())
	}
}

func ServiceByIDHandler(store *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}

		idText := strings.TrimPrefix(r.URL.Path, "/api/services/")
		id, err := strconv.Atoi(idText)

		if err != nil {
			writeError(w, http.StatusBadRequest, "invalid service id")
			return
		}

		service, err := store.GetServiceByID(id)
		if err != nil {
			writeError(w, http.StatusNotFound, "service not found")
			return
		}

		writeJSON(w, http.StatusOK, service)
	}
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{
		"error": message,
	})
}