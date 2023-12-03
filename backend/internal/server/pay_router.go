package server

import (
	"net/http"
	"strconv"
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

		r.Get(route, func(w http.ResponseWriter, r *http.Request) {
			if r.Context().Value("auth") == false {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if r.Context().Value("user") == nil || r.Context().Value("house") == nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			q := r.URL.Query()
			payPeriodID, err := strconv.Atoi(q.Get("payPeriodID"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			var payPeriod db.PayPeriod
			result := db.Database.
				Where("id = ?", payPeriodID).
				Preload("PayEntries").
				Preload("PayEntries.PayItems").
				Preload("PayEntries.PayItems.PayItemDues").
				Take(&payPeriod)
			if result.RowsAffected == 0 {
				http.Error(w, "pay period with id "+q.Get("payPeriodID")+" doesn't exist", http.StatusBadRequest)
				return
			}

			writeJson(w, payPeriod)
		})

		r.Post(route, func(w http.ResponseWriter, r *http.Request) {
			if r.Context().Value("auth") == false {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if r.Context().Value("user") == nil || r.Context().Value("house") == nil {
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
			if body.StartTime.IsZero() {
				http.Error(w, "start time cannot be empty", http.StatusBadRequest)
				return
			}

			house := r.Context().Value("house").(db.House)
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
