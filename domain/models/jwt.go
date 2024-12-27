package models

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	UserId uint   `json:"id_user"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}