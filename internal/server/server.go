package server

import (
	"net/http"

	"github.com/ourhouz/houz/internal/config"
	"github.com/ourhouz/houz/internal/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Init() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/house", routes.House)

	r.Route("/pay", routes.Pay)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("pong")); err != nil {
			return
		}
	})

	err := http.ListenAndServe(":"+config.Env.Port, r)
	if err != nil {
		panic(err)
	}
}
