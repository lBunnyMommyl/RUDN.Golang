package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"portfolio-go/internal/models"
	"portfolio-go/internal/storage"
)

func MessagesHandler(store *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, http.StatusOK, store.GetMessages())

		case http.MethodPost:
			createMessage(w, r, store)

		default:
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		}
	}
}

func createMessage(w http.ResponseWriter, r *http.Request, store *storage.Storage) {
	var message models.Message

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	message.Name = strings.TrimSpace(message.Name)
	message.Email = strings.TrimSpace(message.Email)
	message.Message = strings.TrimSpace(message.Message)

	if message.Name == "" || message.Email == "" || message.Message == "" {
		writeError(w, http.StatusBadRequest, "all fields are required")
		return
	}

	created := store.AddMessage(message)

	writeJSON(w, http.StatusCreated, created)
}