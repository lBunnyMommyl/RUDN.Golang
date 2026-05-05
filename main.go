package main

import (
	"log"
	"net/http"

	"portfolio-go/internal/handlers"
	"portfolio-go/internal/storage"
)

func main() {
	store := storage.NewStorage()

	mux := http.NewServeMux()

	mux.HandleFunc("/api/services", handlers.ServicesHandler(store))
	mux.HandleFunc("/api/services/", handlers.ServiceByIDHandler(store))
	mux.HandleFunc("/api/messages", handlers.MessagesHandler(store))

	fileServer := http.FileServer(http.Dir("./web"))
	mux.Handle("/", fileServer)

	log.Println("Сервер запущен: http://localhost:8080")

	err := http.ListenAndServe(":8080", corsMiddleware(mux))
	if err != nil {
		log.Fatal(err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}