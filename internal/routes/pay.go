package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Pay(r chi.Router) {
	payEntry(r)
}

func payEntry(r chi.Router) {
	const route = "/entry"

	r.Get(route, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	r.Post(route, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
}
