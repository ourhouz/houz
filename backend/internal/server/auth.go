package server

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ourhouz/houz/internal/config"
	"github.com/ourhouz/houz/internal/db"
)

// extractToken is a middleware that parses then injects a JWT into r.Context()
func extractToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "auth", false)

		token, found := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
		if !found {
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		user, house, err := parseUserJWT(token)
		if err != nil {
			w.Header().Add("Authorization", "")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(r.Context(), "auth", true)
		ctx = context.WithValue(ctx, "userRouter", user)
		ctx = context.WithValue(ctx, "houseRouter", house)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type claims struct {
	HouseId uint   `json:"house_id"`
	Name    string `json:"name"`
	jwt.RegisteredClaims
}

// writeUserJWT creates then writes a JWT to the response header
func writeUserJWT(w http.ResponseWriter, user db.User) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		user.HouseID,
		user.Name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "houz",
		},
	})
	s, err := t.SignedString(config.Env.JWTKey)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Authorization", "Bearer "+s)
}

// parseUserJWT verifies then parses a JWT and returns the userRouter and houseRouter associated with it
func parseUserJWT(s string) (user db.User, house db.House, err error) {
	var c claims
	_, err = jwt.ParseWithClaims(s, &c, func(t *jwt.Token) (interface{}, error) {
		return config.Env.JWTKey, nil
	})
	if err != nil {
		err = errors.New("invalid token")
		return
	}

	// redundant, but might check for deleted houseRouter
	result := db.Database.Where("id = ?", c.HouseId).First(&house)
	if result.RowsAffected == 0 {
		err = errors.New("houseRouter not found")
		return
	}

	result = db.Database.Where(&db.User{
		HouseID: c.HouseId,
		Name:    c.Name,
	}).First(&user)
	if result.RowsAffected == 0 {
		err = errors.New("userRouter not found")
		return
	}

	return
}
