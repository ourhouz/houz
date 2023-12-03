package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ourhouz/houz/internal/db"
)

func payRouter(r chi.Router) {
	{
		const route = "/entry"

		r.Get(route, func(w http.ResponseWriter, r *http.Request) {
			// TODO
		})

		r.Post(route, func(w http.ResponseWriter, r *http.Request) {
			// TODO
		})
	}

	{
		const route = "/period"

		r.Post(route, func(w http.ResponseWriter, r *http.Request) {
			if r.Context().Value("auth") == false {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if r.Context().Value("userRouter") == nil || r.Context().Value("houseRouter") == nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			body, err := parseBody[struct {
				StartTime time.Time `json:"startTime"`
			}](r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			house := r.Context().Value("houseRouter").(db.House)
			pp := db.PayPeriod{
				HouseID:   house.ID,
				StartTime: body.StartTime,
			}
			house.PayPeriods = append(house.PayPeriods, pp)

			db.Database.Create(&pp)
			db.Database.Save(&house)

			writeJson(w, pp)
		})
	}
}
