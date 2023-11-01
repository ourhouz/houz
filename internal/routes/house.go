package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ourhouz/houz/internal/db"
	"github.com/ourhouz/houz/internal/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// House is the router for the /house endpoint
func House(r chi.Router) {
	coll := db.Database.Collection(schemas.HouseCollection)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		id, err := primitive.ObjectIDFromHex(q.Get("house_id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var h schemas.House
		err = coll.FindOne(r.Context(), bson.M{"_id": id}).Decode(&h)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := json.Marshal(h)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if _, err = w.Write(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		type RequestBody struct {
			Name   string     `json:"name"`
			UserId schemas.Id `json:"user_id"`
		}

		b, err := parseBody[RequestBody](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		h, err := schemas.NewHouse(b.Name, b.UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = coll.InsertOne(r.Context(), h)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res, err := json.Marshal(h)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err = w.Write(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
	})
}
