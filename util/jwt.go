package util

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const SecretKey = "mysecretkey"

func GenerateJWT(issuer string) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": issuer,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	return claims.SignedString([]byte(SecretKey))

}

func ParseJWT(cookie string) (string, error) {

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return "", err

	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	return claims.Issuer, nil

}
