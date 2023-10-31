package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func PayRouter(r chi.Router) {
	payEntryRouter(r)
}

func payEntryRouter(r chi.Router) {
	const route = "/entry"

	r.Get(route, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	r.Post(route, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
}
