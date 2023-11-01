package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ourhouz/houz/internal/db"
	"github.com/ourhouz/houz/internal/schemas"
	"go.mongodb.org/mongo-driver/bson"
)

func Pay(r chi.Router) {
	payEntry(r)
	payPeriod(r)
}

// payEntry is the router for the /pay/entry endpoint
func payEntry(r chi.Router) {
	const route = "/entry"

	r.Get(route, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	r.Post(route, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
}

// payPeriod is the router for the /pay/item endpoint
func payPeriod(r chi.Router) {
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

		p, err := schemas.NewPayPeriod(b.Start)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

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
