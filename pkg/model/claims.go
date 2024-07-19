package model

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	ProjectID string `json:"project_id"`
	ID        string `json:"id"`
	Email     string `json:"email"`
	jwt.RegisteredClaims
}
