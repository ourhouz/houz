package server

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/ourhouz/houz/internal/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func userRouter(r chi.Router) {
	r.Post("/create", func(w http.ResponseWriter, r *http.Request) {
		body, err := parseBody[struct {
			HouseID  uint   `json:"houseID"`
			Name     string `json:"name"`
			Password string `json:"password"`
		}](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if len(body.Name) == 0 {
			err = errors.New("name cannot be empty")
			return
		}
		if len(body.Name) > 72 {
			// bcrypt limit
			err = errors.New("name cannot be longer than 72 characters")
			return
		}

		result := db.Database.Where("id = ?", body.HouseID).Take(&db.House{})
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			err = errors.New("house with id " + strconv.Itoa(int(body.HouseID)) + " doesn't exist")
			return
		}

		var user db.User
		result = db.Database.Where(&db.User{
			HouseID: body.HouseID,
			Name:    body.Name,
		}).Take(&user)
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			err = errors.New("user with name " + body.Name + " already exists")
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		user = db.User{
			Name:         body.Name,
			PasswordHash: hash,
			HouseID:      body.HouseID,
		}
		db.Database.Create(&user)

		writeUserJWT(w, user)

		w.WriteHeader(http.StatusCreated)
	})

	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		body, err := parseBody[struct {
			HouseID  uint   `json:"houseID"`
			Name     string `json:"name"`
			Password string `json:"password"`
		}](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var user db.User
		result := db.Database.Where(&db.User{
			HouseID: body.HouseID,
			Name:    body.Name,
		}).Take(&user)
		if result.RowsAffected == 0 {
			http.Error(w, "user doesn't exist", http.StatusBadRequest)
			return
		}

		err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(body.Password))
		if err != nil {
			http.Error(w, "Incorrect password", http.StatusBadRequest)
			return
		}

		writeUserJWT(w, user)

		w.WriteHeader(http.StatusCreated)
	})
}
