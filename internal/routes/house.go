package routes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ourhouz/houz/internal/db"
	"github.com/ourhouz/houz/internal/schemas"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
)

// House is the router for the /house endpoint
func House(r chi.Router) {
	coll := db.Database.Collection(schemas.HouseCollection)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		house := q.Get("house_id")

		houseId, err := primitive.ObjectIDFromHex(house)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var h schemas.House
		err = coll.FindOne(db.Ctx, bson.M{"_id": houseId}).Decode(&h)
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
		type Body struct {
			Name   string `json:"name"`
			UserId string `json:"user_id"`
		}

		raw, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var b Body
		if err = json.Unmarshal(raw, &b); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		h, err := schemas.NewHouse(b.Name, b.UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newHouse, err := coll.InsertOne(db.Ctx, h)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := json.Marshal(newHouse)
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
