package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/dwivedisshyam/expense_tracker/pkg"
	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

func exemptPath(path string) bool {
	paths := map[string]bool{
		"/login": true,
	}
	_, ok := paths[path]

	return ok
}

func Auth(app *pkg.App) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if exemptPath(r.URL.Path) {
				h.ServeHTTP(w, r)
				return
			}

			authToken := r.Header.Get("Authorization")
			if authToken == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			authToken = strings.TrimPrefix(authToken, "Bearer ")

			key := []byte(app.GetEnv("JWT_KEY"))
			if len(key) == 0 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			claims := &model.Claims{}

			t, err := jwt.ParseWithClaims(authToken, claims, func(t *jwt.Token) (interface{}, error) {
				return key, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if !t.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			v := mux.Vars(r)

			id := strconv.Itoa(int(claims.ID))

			if id != v["user_id"] {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}
