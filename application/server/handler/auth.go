package handler

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
)

type jwtCustomClaims struct {
	UID   uint   `json:"uid"`
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

var signingKey = []byte(os.Getenv("SECRET_KEY"))

var Config = middleware.JWTConfig{
	Claims:     &jwtCustomClaims{},
	SigningKey: signingKey,
}
