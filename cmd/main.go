package main

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewMux()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("recieved request from ", "path:", r.URL.Path)
		w.Write([]byte("welcome to website"))
	})
	http.ListenAndServe(":3000", router)
}
