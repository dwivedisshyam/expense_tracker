package model

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
