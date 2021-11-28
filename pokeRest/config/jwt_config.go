package config

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func NewJwtCustomClaims(id uint, name string) JwtCustomClaims {
	standardClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	}

	return JwtCustomClaims{
		id,
		name,
		standardClaims,
	}
}
