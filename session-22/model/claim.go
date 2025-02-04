package model

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserName string `json:"username"`
	jwt.RegisteredClaims
}
