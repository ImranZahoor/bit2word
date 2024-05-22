package main

import (
	"net/http"

	views "github.com/ImranZahoor/bit2word/views/pages"
	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewMux()
	fileServer := http.FileServer(http.Dir("./public"))
	router.Handle("/public/*", http.StripPrefix("/public/", fileServer))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		views.Home().Render(r.Context(), w)
	})
	router.Get("/services", func(w http.ResponseWriter, r *http.Request) {
		views.Services().Render(r.Context(), w)
	})
	router.Get("/about-us", func(w http.ResponseWriter, r *http.Request) {
		views.AboutUs().Render(r.Context(), w)
	})
	router.Get("/contact-us", func(w http.ResponseWriter, r *http.Request) {
		views.ContactUs().Render(r.Context(), w)
	})

	http.ListenAndServe(":3000", router)
}
