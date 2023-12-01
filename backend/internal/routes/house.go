package routes

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ourhouz/houz/internal/auth"
	"github.com/ourhouz/houz/internal/db"
	"golang.org/x/crypto/bcrypt"
)

// House is the router for the /house endpoint
func House(r chi.Router) {
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

		writeJson(w, house)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		type RequestBody struct {
			HouseName string `json:"house_name"`
			Username  string `json:"username"`
			Password  string `json:"password"`
		}

		body, err := parseBody[RequestBody](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, user, err := createHouse(body.HouseName, body.Username, body.Password)

		t, err := auth.CreateUserJWT(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Authorization", "Bearer "+t)
		w.WriteHeader(http.StatusCreated)
	})
}

func createHouse(houseName, username, password string) (house db.House, user db.User, err error) {
	if len(houseName) == 0 {
		err = errors.New("house name cannot be empty")
		return
	}
	if len(houseName) > 100 {
		err = errors.New("house name cannot be longer than 100 characters")
		return
	}

	house = db.House{
		Name: houseName,
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user = db.User{
		Name:         username,
		PasswordHash: hash,
		HouseId:      house.ID,
	}

	house.Owner = user
	house.Housemates = append(house.Housemates, user)

	db.Database.Create(&house)
	db.Database.Create(&user)

	return
}
