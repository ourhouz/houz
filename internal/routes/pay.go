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

func Pay(r chi.Router) {
	{
		const route = "/entry"

		// TODO: get detailed payEntry info (includes populated payItem refs)
		r.Get(route, func(w http.ResponseWriter, r *http.Request) {
			q := r.URL.Query()

			id, err := primitive.ObjectIDFromHex(q.Get("entry_id"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			coll := db.Database.Collection()
			var p schemas.PayEntry

			w.Write([]byte("Hello, World!"))
		})

		// TODO: create new payEntry
		r.Post(route, func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, World!"))
		})
	}

	{
		const route = "/period"
		coll := db.Database.Collection(schemas.HouseCollection)

		r.Post(route, func(w http.ResponseWriter, r *http.Request) {
			type Body struct {
				HouseId schemas.Id `json:"house_id"`
				Start   int64      `json:"start"`
			}

			b, err := parseBody[Body](r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			p := schemas.NewPayPeriod(b.Start)

			_, err = coll.UpdateOne(
				r.Context(),
				bson.M{"_id": b.HouseId},
				bson.M{"$push": bson.M{"pay_periods": p}},
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := json.Marshal(p)
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
}
