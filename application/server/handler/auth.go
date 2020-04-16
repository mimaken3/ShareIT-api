package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
)

type jwtCustomClaims struct {
	UID  uint   `json:"uid"`
	Name string `json:"name"`
	jwt.StandardClaims
}

var signingKey = []byte("secret")

var Config = middleware.JWTConfig{
	Claims:     &jwtCustomClaims{},
	SigningKey: signingKey,
}
