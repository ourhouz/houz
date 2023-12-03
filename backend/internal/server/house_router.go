package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ourhouz/houz/internal/db"
)

// houseRouter is the router for the /house endpoint
func houseRouter(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value("auth") == false {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if r.Context().Value("user") == nil || r.Context().Value("house") == nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		house := r.Context().Value("house").(db.House)
		err := db.Database.Select("id, name").Model(&house).Association("Owner").Find(&house.Owner)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = db.Database.Select("id, name").Model(&house).Association("Housemates").Find(&house.Housemates)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = db.Database.Model(&house).Association("PayPeriods").Find(&house.PayPeriods)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		writeJson(w, house)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := parseBody[struct {
			HouseName string `json:"houseName"`
			Username  string `json:"username"`
			Password  string `json:"password"`
		}](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		house, err := db.CreateHouse(body.HouseName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := db.CreateUser(house.ID, body.Username, body.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		house.Owner = user
		house.Housemates = append(house.Housemates, user)

		db.Database.Create(&house)
		db.Database.Create(&user)

		writeUserJWT(w, user)

		writeJson(w, house)
	})
}
