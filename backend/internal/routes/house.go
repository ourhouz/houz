package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// House is the router for the /house endpoint
func House(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO
	})
}
