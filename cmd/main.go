package main

import (
	"net/http"

	views "github.com/ImranZahoor/bit2word/views/home"
	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewMux()
	home := views.Index("hi")
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		home.Render(r.Context(), w)
	})
	http.ListenAndServe(":3000", router)
}
