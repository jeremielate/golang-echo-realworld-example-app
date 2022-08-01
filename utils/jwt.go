package utils

import (
	_ "embed"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//go:embed jwt-secret.bin
var JWTSecret []byte

func GenerateJWT(id uint) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString(JWTSecret)
	return t
}
