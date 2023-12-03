package server

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ourhouz/houz/internal/db"
	"golang.org/x/crypto/bcrypt"
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

		if len(body.HouseName) == 0 {
			err = errors.New("house name cannot be empty")
			return
		}
		if len(body.HouseName) > 100 {
			err = errors.New("house name cannot be longer than 100 characters")
			return
		}

		house := db.House{
			Name: body.HouseName,
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		user := db.User{
			Name:         body.Username,
			PasswordHash: hash,
			HouseID:      house.ID,
		}

		house.Owner = user
		house.Housemates = append(house.Housemates, user)

		db.Database.Create(&house)
		db.Database.Create(&user)

		writeUserJWT(w, user)

		writeJson(w, house)
	})
}
