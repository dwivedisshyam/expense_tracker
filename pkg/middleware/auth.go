package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	gofrHTTP "gofr.dev/pkg/gofr/http"
)

func exemptPath(path string) bool {
	paths := map[string]bool{
		"POST /login": true,
		"POST /users": true,
	}
	_, ok := paths[path]

	return ok
}

func Authentication(jwtKey string) gofrHTTP.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if exemptPath(r.Method + " " + r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}

			authToken := r.Header.Get("Authorization")
			if authToken == "" {
				errorResponse(w, errors.Unauthenticated("invalid auth token"))
				return
			}

			authToken = strings.TrimPrefix(authToken, "Bearer ")

			key := []byte(jwtKey)
			claims := &model.Claims{}

			t, err := jwt.ParseWithClaims(authToken, claims, func(t *jwt.Token) (interface{}, error) {
				return key, nil
			})
			if err != nil {
				errorResponse(w, errors.Unauthenticated("invalid auth token"))
				return
			}

			if !t.Valid {
				errorResponse(w, errors.Unauthenticated("invalid auth token"))
				return
			}

			if claims.ID != mux.Vars(r)["user_id"] {
				errorResponse(w, errors.Unauthorized("un-authorized request"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func errorResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	json.NewEncoder(w).Encode(map[string]error{"error": err})
}
