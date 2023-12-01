package auth

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

func ExtractToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "auth", false)

		token, found := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
		if !found {
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		user, house, err := VerifyUserJWT(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(r.Context(), "auth", true)
		ctx = context.WithValue(ctx, "user", user)
		ctx = context.WithValue(ctx, "house", house)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type claims struct {
	HouseId uint   `json:"house_id"`
	Name    string `json:"name"`
	jwt.RegisteredClaims
}

func CreateUserJWT(user db.User) (s string, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		user.HouseId,
		user.Name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "houz",
		},
	})
	s, err = t.SignedString(config.Env.JWTKey)

	return
}

func VerifyUserJWT(s string) (user db.User, house db.House, err error) {
	var c claims
	_, err = jwt.ParseWithClaims(s, &c, func(t *jwt.Token) (interface{}, error) {
		return config.Env.JWTKey, nil
	})
	if err != nil {
		err = errors.New("invalid token")
		return
	}

	// redundant, but might check for deleted house
	result := db.Database.Where("id = ?", c.HouseId).First(&house)
	if result.RowsAffected == 0 {
		err = errors.New("house not found")
		return
	}

	result = db.Database.Where(&db.User{
		HouseId: c.HouseId,
		Name:    c.Name,
	}).First(&user)
	if result.RowsAffected == 0 {
		err = errors.New("user not found")
		return
	}

	return
}
