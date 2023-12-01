package routes

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/ourhouz/houz/internal/auth"
	"github.com/ourhouz/houz/internal/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func User(r chi.Router) {
	r.Post("/create", func(w http.ResponseWriter, r *http.Request) {
		type RequestBody struct {
			HouseId  uint   `json:"house_id"`
			Name     string `json:"name"`
			Password string `json:"password"`
		}

		body, err := parseBody[RequestBody](r)
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

		result := db.Database.Where("id = ?", body.HouseId).Take(&db.House{})
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			err = errors.New("house with id " + strconv.Itoa(int(body.HouseId)) + " doesn't exist")
			return
		}

		var user db.User
		result = db.Database.Where(&db.User{
			HouseID: body.HouseId,
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
			HouseID:      body.HouseId,
		}
		db.Database.Create(&user)

		t, err := auth.CreateUserJWT(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Authorization", "Bearer "+t)
		w.WriteHeader(http.StatusCreated)
	})

	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		type RequestBody struct {
			HouseId  uint   `json:"house_id"`
			Name     string `json:"name"`
			Password string `json:"password"`
		}

		body, err := parseBody[RequestBody](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var user db.User
		result := db.Database.Where(&db.User{
			HouseID: body.HouseId,
			Name:    body.Name,
		}).First(&user)
		if result.RowsAffected == 0 {
			http.Error(w, "User doesn't exist", http.StatusBadRequest)
			return
		}

		err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(body.Password))
		if err != nil {
			http.Error(w, "Incorrect password", http.StatusBadRequest)
			return
		}

		t, err := auth.CreateUserJWT(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Authorization", "Bearer "+t)
		w.WriteHeader(http.StatusCreated)
	})
}
