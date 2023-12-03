package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ourhouz/houz/internal/config"
)

func Init() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(3 * time.Second))

	r.Use(extractUserJWT)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("pong")); err != nil {
			return
		}
	})
	r.Get("/auth", func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value("auth") == false {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	})

	r.Route("/user", userRouter)
	r.Route("/house", houseRouter)
	r.Route("/pay", payRouter)

	err := http.ListenAndServe(":"+config.Env.Port, r)
	if err != nil {
		panic(err)
	}
}
