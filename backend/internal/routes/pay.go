package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Pay(r chi.Router) {
	{
		const route = "/entry"

		r.Get(route, func(w http.ResponseWriter, r *http.Request) {
			// TODO
		})

		r.Post(route, func(w http.ResponseWriter, r *http.Request) {
			// TODO
		})
	}

	{
		const route = "/period"

		r.Post(route, func(w http.ResponseWriter, r *http.Request) {
			// TODO
		})
	}
}
