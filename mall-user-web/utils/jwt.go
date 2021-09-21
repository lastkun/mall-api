package utils

import (
	"github.com/dgrijalva/jwt-go"
)

//jwt playload
type CustomClaims struct {
	ID       int32
	NickName string
	Role     int32
	jwt.StandardClaims
}
